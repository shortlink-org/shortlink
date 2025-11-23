/**
 * Telemetry initialization module
 * MUST be imported FIRST before any other modules that use:
 * - amqplib
 * - ioredis
 * - winston
 *
 * This ensures OpenTelemetry instrumentation is set up before these modules are loaded.
 */

import * as dotenv from "dotenv";

// Load environment variables FIRST
// Note: dotenv v17+ doesn't have 'silent' option, message is informational only
dotenv.config();

// Initialize OpenTelemetry BEFORE any other imports
import { initializeTelemetry } from "./infrastructure/telemetry.js";
initializeTelemetry();
