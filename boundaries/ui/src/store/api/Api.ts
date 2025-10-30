/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link {
  /** Create at */
  created_at?: TimestamppbTimestamp

  /** Describe of link */
  describe?: string

  /** Hash by URL + salt */
  hash?: string

  /** Update at */
  updated_at?: TimestamppbTimestamp

  /** URL */
  url?: string
}

export interface GithubComBatazorShortlinkInternalServicesLinkInfrastructureRpcLinkV1Link {
  LinkServiceServer?: any
}

export interface TimestamppbTimestamp {
  /**
   * Non-negative fractions of a second at nanosecond resolution. Negative
   * second values with fractions must still have non-negative nanos values
   * that count forward in time. Must be from 0 to 999,999,999
   * inclusive.
   */
  nanos?: number

  /**
   * Represents seconds of UTC time since Unix epoch
   * 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
   * 9999-12-31T23:59:59Z inclusive.
   */
  seconds?: number
}

export interface V1AddRequest {
  link?: GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link
}

export interface V1AddResponse {
  link?: GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link
}

export interface V1GetResponse {
  link?: GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link
}

export interface V1Links {
  link?: GithubComBatazorShortlinkInternalServicesLinkInfrastructureRpcLinkV1Link[]
}

export interface V1ListResponse {
  links?: V1Links
}

export interface V1UpdateRequest {
  link?: GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link
}

export interface V1UpdateResponse {
  link?: GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link
}

export type QueryParamsType = Record<string | number, any>
export type ResponseFormat = keyof Omit<Body, 'body' | 'bodyUsed'>

export interface FullRequestParams extends Omit<RequestInit, 'body'> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean
  /** request path */
  path: string
  /** content type of request body */
  type?: ContentType
  /** query params */
  query?: QueryParamsType
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat
  /** request body */
  body?: unknown
  /** base url */
  baseUrl?: string
  /** request cancellation token */
  cancelToken?: CancelToken
}

export type RequestParams = Omit<FullRequestParams, 'body' | 'method' | 'query' | 'path'>

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string
  baseApiParams?: Omit<RequestParams, 'baseUrl' | 'cancelToken' | 'signal'>
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void
  customFetch?: typeof fetch
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D
  error: E
}

type CancelToken = Symbol | string | number

export enum ContentType {
  Json = 'application/json',
  FormData = 'multipart/form-data',
  UrlEncoded = 'application/x-www-form-urlencoded',
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = 'http://localhost:7070/api'
  private securityData: SecurityDataType | null = null
  private securityWorker?: ApiConfig<SecurityDataType>['securityWorker']
  private abortControllers = new Map<CancelToken, AbortController>()
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams)

  private baseApiParams: RequestParams = {
    credentials: 'same-origin',
    headers: {},
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
  }

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig)
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data
  }

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key)
    return `${encodedKey}=${encodeURIComponent(typeof value === 'number' ? value : `${value}`)}`
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key])
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key]
    return value.map((v: any) => this.encodeQueryParam(key, v)).join('&')
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {}
    const keys = Object.keys(query).filter((key) => 'undefined' !== typeof query[key])
    return keys.map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key))).join('&')
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery)
    return queryString ? `?${queryString}` : ''
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === 'object' || typeof input === 'string') ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key]
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === 'object' && property !== null
              ? JSON.stringify(property)
              : `${property}`,
        )
        return formData
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  }

  protected mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    }
  }

  protected createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken)
      if (abortController) {
        return abortController.signal
      }
      return void 0
    }

    const abortController = new AbortController()
    this.abortControllers.set(cancelToken, abortController)
    return abortController.signal
  }

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken)

    if (abortController) {
      abortController.abort()
      this.abortControllers.delete(cancelToken)
    }
  }

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === 'boolean' ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {}
    const requestParams = this.mergeRequestParams(params, secureParams)
    const queryString = query && this.toQueryString(query)
    const payloadFormatter = this.contentFormatters[type || ContentType.Json]
    const responseFormat = format || requestParams.format

    return this.customFetch(`${baseUrl || this.baseUrl || ''}${path}${queryString ? `?${queryString}` : ''}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { 'Content-Type': type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : requestParams.signal,
      body: typeof body === 'undefined' || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>
      r.data = null as unknown as T
      r.error = null as unknown as E

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data
              } else {
                r.error = data
              }
              return r
            })
            .catch((e) => {
              r.error = e
              return r
            })

      if (cancelToken) {
        this.abortControllers.delete(cancelToken)
      }

      if (!response.ok) throw data
      return data
    })
  }
}

/**
 * @title Shortlink API
 * @version 1.0
 * @license MIT (http://www.opensource.org/licenses/MIT)
 * @termsOfService http://swagger.io/terms/
 * @baseUrl http://localhost:7070/api
 * @contact Shortlink repository <support@swagger.io> (https://github.com/shortlink-org/shortlink/issues)
 *
 * Shortlink API
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  links = {
    /**
     * @description List links
     *
     * @name ListLinks
     * @summary List links
     * @request GET:/links
     */
    listLinks: (params: RequestParams = {}) =>
      this.request<V1ListResponse, any>({
        path: `/links`,
        method: 'GET',
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Add link
     *
     * @name AddLink
     * @summary Add link
     * @request POST:/links
     */
    addLink: (link: V1AddRequest, params: RequestParams = {}) =>
      this.request<V1AddResponse, any>({
        path: `/links`,
        method: 'POST',
        body: link,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Update link
     *
     * @name UpdateLink
     * @summary Update link
     * @request PUT:/links/:hash
     */
    updateLink: (hash: string, link: V1UpdateRequest, params: RequestParams = {}) =>
      this.request<V1UpdateResponse, any>({
        path: `/links/${hash}`,
        method: 'PUT',
        body: link,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Get link
     *
     * @name GetLink
     * @summary Get link
     * @request GET:/links/{hash}
     */
    getLink: (hash: string, params: RequestParams = {}) =>
      this.request<V1GetResponse, any>({
        path: `/links/${hash}`,
        method: 'GET',
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Delete link
     *
     * @name DeleteLink
     * @summary Delete link
     * @request DELETE:/links/{hash}
     */
    deleteLink: (hash: string, params: RequestParams = {}) =>
      this.request<void, any>({
        path: `/links/${hash}`,
        method: 'DELETE',
        type: ContentType.Json,
        ...params,
      }),
  }
}
