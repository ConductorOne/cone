//go:build darwin && cgo

package managedconfig

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <stdlib.h>
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

import "unsafe"

// readManagedConfig reads the managed device configuration from the macOS
// managed-preferences store using CoreFoundation directly. It calls
// CFPreferencesCopyAppValue, which resolves the managed-preferences layer
// (administrator-pushed policy) for the "ai.c1" application domain, exactly as
// the `defaults read` fallback does — but without shelling out.
//
// A missing key, a value that is not a string, or any lookup failure yields a
// nil map. This function never returns an error and never panics.
func readManagedConfig() map[string]string {
	v := copyManagedString(Namespace, KeyTenantDomain)
	if v == "" {
		return nil
	}
	return map[string]string{KeyTenantDomain: v}
}

// copyManagedString reads a single string value from the managed-preferences
// layer for the given application domain, returning "" when the key is absent,
// is not a string, or cannot be read.
//
// Memory ownership follows the CoreFoundation "Create/Copy" rule mirrored from
// crypto/x509/internal/macos: every ref obtained from a Create*/Copy* call is
// owned (+1 retain) and must be released; refs obtained from Get* calls are not
// owned and must not be released.
func copyManagedString(appID, key string) string {
	// C.CString allocates C memory that must be freed explicitly.
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cApp := C.CString(appID)
	defer C.free(unsafe.Pointer(cApp))

	// CFStringCreateWithCString returns an owned (+1) CFString, or NULL if the
	// bytes are not valid for the encoding. Release both when done.
	keyRef := C.CFStringCreateWithCString(C.kCFAllocatorDefault, cKey, C.kCFStringEncodingUTF8)
	if keyRef == nil {
		return ""
	}
	defer C.CFRelease(C.CFTypeRef(keyRef))

	appRef := C.CFStringCreateWithCString(C.kCFAllocatorDefault, cApp, C.kCFStringEncodingUTF8)
	if appRef == nil {
		return ""
	}
	defer C.CFRelease(C.CFTypeRef(appRef))

	// CFPreferencesCopyAppValue returns an owned (+1) value, or NULL when the
	// key is absent. Release it whether or not it is the type we want.
	val := C.CFPreferencesCopyAppValue(keyRef, appRef)
	if val == nil {
		return ""
	}
	defer C.CFRelease(val)

	// The value may be any property-list type. Confirm it is actually a
	// CFString before extracting; a non-string value yields "" rather than a
	// crash.
	if C.CFGetTypeID(val) != C.CFStringGetTypeID() {
		return ""
	}

	return cfStringToGoString(C.CFStringRef(val))
}

// cfStringToGoString extracts a Go string from a CFString without assuming a
// fixed buffer size, mirroring the robust extraction in
// crypto/x509/internal/macos: try the zero-copy fast path first, then fall back
// to CFStringGetCString with a buffer sized by CFStringGetMaximumSizeForEncoding.
// The caller retains ownership of s; this function releases nothing it does not
// create.
func cfStringToGoString(s C.CFStringRef) string {
	// Fast path: CFStringGetCStringPtr may return a pointer to the CFString's
	// internal UTF-8 storage. It is owned by the CFString (not by us), so it is
	// copied by C.GoString and never released. It returns NULL when no such
	// direct representation exists, in which case we fall back below.
	if p := C.CFStringGetCStringPtr(s, C.kCFStringEncodingUTF8); p != nil {
		return C.GoString(p)
	}

	// Fallback: size a buffer large enough for the worst-case UTF-8 encoding of
	// the string's length, plus one for the NUL terminator, then copy into it.
	length := C.CFStringGetLength(s)
	maxSize := C.CFStringGetMaximumSizeForEncoding(length, C.kCFStringEncodingUTF8) + 1
	buf := (*C.char)(C.malloc(C.size_t(maxSize)))
	defer C.free(unsafe.Pointer(buf))

	// CFStringGetCString returns a Boolean (0 on failure).
	if C.CFStringGetCString(s, buf, maxSize, C.kCFStringEncodingUTF8) == 0 {
		return ""
	}
	return C.GoString(buf)
}
