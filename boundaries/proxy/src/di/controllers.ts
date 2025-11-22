import { ProxyController } from "../infrastructure/http/fastify/controllers/ProxyController.js";
import { MetricsController } from "../infrastructure/http/fastify/controllers/MetricsController.js";

export const CONTROLLERS = {
  proxyController: ProxyController,
  metricsController: MetricsController,
};

