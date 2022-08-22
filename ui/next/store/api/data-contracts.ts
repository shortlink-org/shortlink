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
