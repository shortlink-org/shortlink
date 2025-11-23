/**
 * Initialize profiling (currently disabled - Pyroscope not compatible with Node.js 25+).
 *
 * This function is kept for API compatibility but does nothing.
 * Pyroscope will be re-enabled when compatibility is restored.
 */
export async function initializeProfiling(): Promise<void> {
  // Pyroscope is disabled due to Node.js 25+ compatibility issues
  // Will be re-enabled when @pyroscope/nodejs supports Node.js 25+
  return Promise.resolve();
}
