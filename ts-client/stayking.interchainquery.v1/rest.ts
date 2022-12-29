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

export interface CryptoProofOp {
    type?: string

    /** @format byte */
    key?: string

    /** @format byte */
    data?: string
}

export interface CryptoProofOps {
    ops?: CryptoProofOp[]
}

export interface ProtobufAny {
    '@type'?: string
}

export interface RpcStatus {
    /** @format int32 */
    code?: number
    message?: string
    details?: ProtobufAny[]
}

/**
 * MsgSubmitQueryResponse represents a message type to fulfil a query request.
 */
export interface V1MsgSubmitQueryResponse {
    chain_id?: string
    query_id?: string

    /** @format byte */
    result?: string
    proof_ops?: CryptoProofOps

    /** @format int64 */
    height?: string
    from_address?: string
}

/**
* MsgSubmitQueryResponseResponse defines the MsgSubmitQueryResponse response
type.
*/
export type V1MsgSubmitQueryResponseResponse = object

export interface V1Query {
    id?: string
    connection_id?: string
    chain_id?: string
    query_type?: string

    /** @format byte */
    request?: string
    callback_id?: string

    /** @format uint64 */
    ttl?: string
    request_sent?: boolean
}

export interface V1QueryPendingQueriesResponse {
    pending_queries?: V1Query[]
}

import axios, {
    AxiosInstance,
    AxiosRequestConfig,
    AxiosResponse,
    ResponseType,
} from 'axios'

export type QueryParamsType = Record<string | number, any>

export interface FullRequestParams
    extends Omit<
        AxiosRequestConfig,
        'data' | 'params' | 'url' | 'responseType'
    > {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean
    /** request path */
    path: string
    /** content type of request body */
    type?: ContentType
    /** query params */
    query?: QueryParamsType
    /** format of response (i.e. response.json() -> format: "json") */
    format?: ResponseType
    /** request body */
    body?: unknown
}

export type RequestParams = Omit<
    FullRequestParams,
    'body' | 'method' | 'query' | 'path'
>

export interface ApiConfig<SecurityDataType = unknown>
    extends Omit<AxiosRequestConfig, 'data' | 'cancelToken'> {
    securityWorker?: (
        securityData: SecurityDataType | null
    ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void
    secure?: boolean
    format?: ResponseType
}

export enum ContentType {
    Json = 'application/json',
    FormData = 'multipart/form-data',
    UrlEncoded = 'application/x-www-form-urlencoded',
}

export class HttpClient<SecurityDataType = unknown> {
    public instance: AxiosInstance
    private securityData: SecurityDataType | null = null
    private securityWorker?: ApiConfig<SecurityDataType>['securityWorker']
    private secure?: boolean
    private format?: ResponseType

    constructor({
        securityWorker,
        secure,
        format,
        ...axiosConfig
    }: ApiConfig<SecurityDataType> = {}) {
        this.instance = axios.create({
            ...axiosConfig,
            baseURL: axiosConfig.baseURL || '',
        })
        this.secure = secure
        this.format = format
        this.securityWorker = securityWorker
    }

    public setSecurityData = (data: SecurityDataType | null) => {
        this.securityData = data
    }

    private mergeRequestParams(
        params1: AxiosRequestConfig,
        params2?: AxiosRequestConfig
    ): AxiosRequestConfig {
        return {
            ...this.instance.defaults,
            ...params1,
            ...(params2 || {}),
            headers: {
                ...(this.instance.defaults.headers || {}),
                ...(params1.headers || {}),
                ...((params2 && params2.headers) || {}),
            },
        }
    }

    private createFormData(input: Record<string, unknown>): FormData {
        return Object.keys(input || {}).reduce((formData, key) => {
            const property = input[key]
            formData.append(
                key,
                property instanceof Blob
                    ? property
                    : typeof property === 'object' && property !== null
                    ? JSON.stringify(property)
                    : `${property}`
            )
            return formData
        }, new FormData())
    }

    public request = async <T = any, _E = any>({
        secure,
        path,
        type,
        query,
        format,
        body,
        ...params
    }: FullRequestParams): Promise<AxiosResponse<T>> => {
        const secureParams =
            ((typeof secure === 'boolean' ? secure : this.secure) &&
                this.securityWorker &&
                (await this.securityWorker(this.securityData))) ||
            {}
        const requestParams = this.mergeRequestParams(params, secureParams)
        const responseFormat = (format && this.format) || void 0

        if (
            type === ContentType.FormData &&
            body &&
            body !== null &&
            typeof body === 'object'
        ) {
            requestParams.headers.common = { Accept: '*/*' }
            requestParams.headers.post = {}
            requestParams.headers.put = {}

            body = this.createFormData(body as Record<string, unknown>)
        }

        return this.instance.request({
            ...requestParams,
            headers: {
                ...(type && type !== ContentType.FormData
                    ? { 'Content-Type': type }
                    : {}),
                ...(requestParams.headers || {}),
            },
            params: query,
            responseType: responseFormat,
            data: body,
            url: path,
        })
    }
}

/**
 * @title stayking/interchainquery/v1/genesis.proto
 * @version version not set
 */
export class Api<
    SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags QueryService
     * @name QueryServicePendingQueries
     * @request GET:/Stride-Labs/stayking/interchainquery/pending_queries
     */
    queryServicePendingQueries = (params: RequestParams = {}) =>
        this.request<V1QueryPendingQueriesResponse, RpcStatus>({
            path: `/Stride-Labs/stayking/interchainquery/pending_queries`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Msg
     * @name MsgSubmitQueryResponse
     * @summary SubmitQueryResponse defines a method for submit query responses.
     * @request POST:/interchainquery/tx/v1beta1/submitquery
     */
    msgSubmitQueryResponse = (
        body: V1MsgSubmitQueryResponse,
        params: RequestParams = {}
    ) =>
        this.request<V1MsgSubmitQueryResponseResponse, RpcStatus>({
            path: `/interchainquery/tx/v1beta1/submitquery`,
            method: 'POST',
            body: body,
            type: ContentType.Json,
            format: 'json',
            ...params,
        })
}
