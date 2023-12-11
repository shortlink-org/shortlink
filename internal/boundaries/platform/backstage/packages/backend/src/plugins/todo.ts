import { Router } from 'express';
import { CatalogClient } from '@backstage/catalog-client';
import {
  createRouter,
  TodoReaderService,
  TodoScmReader,
} from '@backstage/plugin-todo-backend';
import { PluginEnvironment } from '../types';

export default async function createPlugin(
  env: PluginEnvironment,
): Promise<Router> {
  const todoReader = TodoScmReader.fromConfig(env.config, {
    logger: env.logger,
    reader: env.reader,
  });

  const catalogClient = new CatalogClient({
    discoveryApi: env.discovery,
  });

  const todoService = new TodoReaderService({
    todoReader,
    catalogClient,
  });

  return await createRouter({ todoService });
}
