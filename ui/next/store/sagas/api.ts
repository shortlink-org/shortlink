import { Api, ApiConfig } from "../api/Api";

const config: ApiConfig = {
  baseUrl: 'http://localhost:3000/api',
}

export default new Api(config)
