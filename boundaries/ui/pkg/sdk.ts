import { FrontendApi, Configuration } from '@ory/client'

// Create a configured Ory client instance
const ory = new FrontendApi(
  new Configuration({
    basePath: process.env.NEXT_PUBLIC_ORY_SDK_URL || 'http://localhost:4433',
    baseOptions: {
      withCredentials: true,
    },
  })
)

export default ory
