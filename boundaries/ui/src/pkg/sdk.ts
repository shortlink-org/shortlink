import { Configuration, FrontendApi } from '@ory/client'

const NEXT_PUBLIC_API_URI = `${process.env.NEXT_PUBLIC_API_URI || ''}/api/auth`

const ory = new FrontendApi(
  new Configuration({
    basePath: NEXT_PUBLIC_API_URI,
    baseOptions: {
      withCredentials: true,
    },
  }),
)

export default ory
