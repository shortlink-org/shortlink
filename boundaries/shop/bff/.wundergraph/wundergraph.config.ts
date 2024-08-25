import { configureWunderGraphApplication, cors, EnvironmentVariable, introspect, templates } from '@wundergraph/sdk';
import server from './wundergraph.server';
import operations from './wundergraph.operations';

const countries = introspect.graphql({
	apiNamespace: 'countries',
	url: 'https://countries.trevorblades.com/',
});

const goods = introspect.openApiV2({
  id: 'goods',
  apiNamespace: 'goods',
  source: {
    kind: 'file',
    filePath: './schema/swagger.yaml',
  },
  baseURL: 'http://127.0.0.1:8000/goods/',
})

// configureWunderGraph emits the configuration
configureWunderGraphApplication({
	apis: [goods, countries],
	server,
	operations,
	generate: {
		codeGenerators: [
      {
        templates: [
          // use all the typescript react templates to generate a client
          ...templates.typescript.all,
          templates.typescript.operations,
          templates.typescript.linkBuilder,
          templates.typescript.client,
        ],
      }
    ],
	},
  openApi: {
    title: "ShortLink Shop API",
  },
	cors: {
		...cors.allowAll,
		allowedOrigins:
			process.env.NODE_ENV === 'production'
				? [
						// change this before deploying to production to the actual domain where you're deploying your app
						'http://localhost:3000',
				  ]
				: ['http://localhost:3000', new EnvironmentVariable('WG_ALLOWED_ORIGIN')],
	},
	security: {
		enableGraphQLEndpoint: process.env.NODE_ENV !== 'production' || process.env.GITPOD_WORKSPACE_ID !== undefined,
	},
});
