import { createRouter } from '@axis-backstage/plugin-readme-backend';
import { Router } from 'express';
import { PluginEnvironment } from '../types';

export default async function createPlugin(
  env: PluginEnvironment,
): Promise<Router> {
  return await createRouter({
    logger: env.logger,
    config: env.config,
    reader: env.reader,
    discovery: env.discovery,
    tokenManager: env.tokenManager,
  });
}
