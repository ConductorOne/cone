// @c1/console - Structured logging wrapper for C1 Functions
// This wraps console methods to output structured JSON logs

export interface LogContext {
  functionId?: string;
  invocationId?: string;
}

let _context: LogContext = {};

export function setContext(ctx: LogContext) {
  _context = { ..._context, ...ctx };
}

export function getContext(): LogContext {
  return { ..._context };
}

interface StructuredLog {
  functionId: string;
  invocationId: string;
  logType: string;
  message: string;
  timestamp: string;
}

function formatMessage(...args: unknown[]): string {
  return args.map(arg => {
    if (typeof arg === 'string') return arg;
    try {
      return JSON.stringify(arg);
    } catch {
      return String(arg);
    }
  }).join(' ');
}

function createLogEntry(logType: string, ...args: unknown[]): StructuredLog {
  return {
    functionId: _context.functionId || "unknown",
    invocationId: _context.invocationId || "unknown",
    logType,
    message: formatMessage(...args),
    timestamp: new Date().toISOString()
  };
}

// Store original console methods
const originalConsole = {
  log: console.log.bind(console),
  error: console.error.bind(console),
  warn: console.warn.bind(console),
  info: console.info.bind(console),
  debug: console.debug.bind(console),
};

// Structured versions
export function log(...args: unknown[]) {
  const entry = createLogEntry("info", ...args);
  originalConsole.log(JSON.stringify(entry));
}

export function error(...args: unknown[]) {
  const entry = createLogEntry("error", ...args);
  originalConsole.log(JSON.stringify(entry));
}

export function warn(...args: unknown[]) {
  const entry = createLogEntry("warn", ...args);
  originalConsole.log(JSON.stringify(entry));
}

export function info(...args: unknown[]) {
  const entry = createLogEntry("info", ...args);
  originalConsole.log(JSON.stringify(entry));
}

export function debug(...args: unknown[]) {
  const entry = createLogEntry("info", ...args);
  originalConsole.log(JSON.stringify(entry));
}

// Install wrapper - replaces global console with structured logging
export function install() {
  globalThis.console.log = log;
  globalThis.console.error = error;
  globalThis.console.warn = warn;
  globalThis.console.info = info;
  globalThis.console.debug = debug;
}

// Restore original console
export function restore() {
  globalThis.console.log = originalConsole.log;
  globalThis.console.error = originalConsole.error;
  globalThis.console.warn = originalConsole.warn;
  globalThis.console.info = originalConsole.info;
  globalThis.console.debug = originalConsole.debug;
}

// Emit invocation end marker
export function endInvocation() {
  const id = _context.invocationId || "unknown";
  originalConsole.log(`###END:${id}###`);
}
