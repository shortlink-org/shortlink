/**
 * Application entry point.
 *
 * This file serves as the main entry point for the compiled application.
 * It delegates to the bootstrap module which follows Clean Architecture principles.
 *
 * @module app
 */

import "reflect-metadata";
import server from "./application/bootstrap.js";

export default server;
