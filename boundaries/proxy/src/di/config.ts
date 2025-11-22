import { asClass, asValue } from "awilix";
import { AppConfig } from "../../infrastructure/config/index.js";
import { ExternalServicesConfig } from "../../infrastructure/config/ExternalServicesConfig.js";
import { CacheConfig } from "../../infrastructure/config/CacheConfig.js";

export const CONFIG = {
  appConfig: asValue(new AppConfig()),
  externalServicesConfig: asClass(ExternalServicesConfig).singleton(),
  cacheConfig: asClass(CacheConfig).singleton(),
};

