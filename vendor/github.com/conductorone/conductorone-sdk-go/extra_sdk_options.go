package conductoroneapi

import (
	"net/url"
	"strings"
)

const ClientIdGolangSDK = "2RCzHlak5q7CY14SdBc8HoZEJRf"

func WithTenant(input string) (SDKOption, error) {
	input = strings.ToLower(input)

	var err error
	u := &url.URL{}
	if !strings.Contains(input, "//") {
		if !strings.Contains(input, ".") {
			input += ".conductor.one"
		}
		u.Host = input
	} else {
		u, err = url.Parse(input)
		if err != nil {
			return nil, err
		}
	}

	parts := strings.Split(u.Host, ".")
	if len(parts) == 3 && parts[1] == "conductor" && parts[2] == "one" {
		return WithTenantDomain(parts[0]), nil
	}

	u.Scheme = "https"
	return WithServerURL(u.String()), nil
}
