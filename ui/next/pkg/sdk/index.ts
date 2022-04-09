import { Configuration, V0alpha2Api } from '@ory/client'
const KRATOS_PUBLIC_API = process.env.KRATOS_PUBLIC_API || 'http://127.0.0.1:4433'

export default new V0alpha2Api(new Configuration({
  basePath: KRATOS_PUBLIC_API,
  baseOptions: {
    withCredentials: true,
  }
}))
