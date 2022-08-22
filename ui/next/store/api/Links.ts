/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

import {
  V1AddRequest,
  V1AddResponse,
  V1GetResponse,
  V1ListResponse,
  V1UpdateRequest,
  V1UpdateResponse,
} from './data-contracts'
import { ContentType, HttpClient, RequestParams } from './http-client'

export class Links<
  SecurityDataType = unknown,
> extends HttpClient<SecurityDataType> {
  /**
   * @description List links
   *
   * @name ListLinks
   * @summary List links
   * @request GET:/links
   * @response `200` `V1ListResponse` OK
   */
  listLinks = (params: RequestParams = {}) =>
    this.request<V1ListResponse, any>({
      path: `/links`,
      method: 'GET',
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Update link
   *
   * @name UpdateLink
   * @summary Update link
   * @request PUT:/links
   * @response `200` `V1UpdateResponse` OK
   */
  updateLink = (link: V1UpdateRequest, params: RequestParams = {}) =>
    this.request<V1UpdateResponse, any>({
      path: `/links`,
      method: 'PUT',
      body: link,
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Add link
   *
   * @name AddLink
   * @summary Add link
   * @request POST:/links
   * @response `200` `V1AddResponse` OK
   */
  addLink = (link: V1AddRequest, params: RequestParams = {}) =>
    this.request<V1AddResponse, any>({
      path: `/links`,
      method: 'POST',
      body: link,
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Get link
   *
   * @name GetLink
   * @summary Get link
   * @request GET:/links/{hash}
   * @response `200` `V1GetResponse` OK
   */
  getLink = (hash: string, params: RequestParams = {}) =>
    this.request<V1GetResponse, any>({
      path: `/links/${hash}`,
      method: 'GET',
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Delete link
   *
   * @name DeleteLink
   * @summary Delete link
   * @request DELETE:/links/{hash}
   * @response `200` `void`
   */
  deleteLink = (hash: string, params: RequestParams = {}) =>
    this.request<void, any>({
      path: `/links/${hash}`,
      method: 'DELETE',
      type: ContentType.Json,
      ...params,
    })
}
