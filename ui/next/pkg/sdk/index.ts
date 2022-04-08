import { Configuration, V0alpha2Api } from '@ory/client'
const KRATOS_PUBLIC_API = process.env.KRATOS_PUBLIC_API || 'https://shortlink.ddns.net/api/auth'

export default new V0alpha2Api(new Configuration({
  basePath: KRATOS_PUBLIC_API,
}))
