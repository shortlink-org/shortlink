const PROFILING_ENV_FLAG = "DD_PROFILING_ENABLED";
const DEFAULT_SERVICE_NAME = "proxy";
const DEFAULT_PYROSCOPE_SERVER = "http://pyroscope:4040";

/**
 * Initialize Pyroscope profiling if enabled via environment variable.
 *
 * Requires DD_PROFILING_ENABLED=true to activate.
 * Configures service name and server address from environment variables.
 *
 * Side effects: Starts Pyroscope profiler if enabled.
 */
export async function initializeProfiling(): Promise<void> {
  if (process.env[PROFILING_ENV_FLAG] !== "true") {
    return;
  }

  try {
    // Dynamic import to avoid loading Pyroscope when not needed
    const Pyroscope = await import("@pyroscope/nodejs");

    const serviceName = process.env.SERVICE_NAME || DEFAULT_SERVICE_NAME;
    const serverAddress =
      process.env.PYROSCOPE_SERVER_ADDRESS || DEFAULT_PYROSCOPE_SERVER;

    Pyroscope.init({
      appName: serviceName,
      serverAddress: serverAddress,
    });

    Pyroscope.start();
    console.log(`[Pyroscope] enabled for "${serviceName}"`);
  } catch (err) {
    console.error("[Pyroscope] failed to initialize:", err);
  }
}

