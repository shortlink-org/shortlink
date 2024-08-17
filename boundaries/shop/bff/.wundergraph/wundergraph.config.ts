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
    filePath: 'https://raw.githubusercontent.com/shortlink-org/shortlink/main/boundaries/shop/admin/docs/public/Shop%20Admin%20API.yaml',
  },
  baseURL: 'https://shop.shortlink.org/api/goods',
})

// configureWunderGraph emits the configuration
configureWunderGraphApplication({
	apis: [countries, goods],
	server,
	operations,
	generate: {
		codeGenerators: [],
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