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

export interface ProtobufAny {
    '@type'?: string
}

export interface RpcStatus {
    /** @format int32 */
    code?: number
    message?: string
    details?: ProtobufAny[]
}

export interface V1Beta1DistributionProportions {
    /**
     * staking defines the proportion of the minted minted_denom that is to be
     * allocated as staking rewards.
     */
    staking?: string

    /**
     * community_pool defines the proportion of the minted mint_denom that is
     * to be allocated to the community pool: growth.
     */
    community_pool_growth?: string

    /**
     * community_pool defines the proportion of the minted mint_denom that is
     * to be allocated to the community pool: security budget.
     */
    community_pool_security_budget?: string

    /**
     * strategic_reserve defines the proportion of the minted mint_denom that is
     * to be allocated to the pool: strategic reserve.
     */
    strategic_reserve?: string
}

/**
 * Params holds parameters for the mint module.
 */
export interface V1Beta1Params {
    /** type of coin to mint */
    mint_denom?: string

    /** epoch provisions from the first epoch */
    genesis_epoch_provisions?: string

    /** mint epoch identifier */
    epoch_identifier?: string

    /**
     * number of epochs take to reduce rewards
     * @format int64
     */
    reduction_period_in_epochs?: string

    /** reduction multiplier to execute on each period */
    reduction_factor?: string

    /** distribution_proportions defines the proportion of the minted denom */
    distribution_proportions?: V1Beta1DistributionProportions

    /**
     * start epoch to distribute minting rewards
     * @format int64
     */
    minting_rewards_distribution_start_epoch?: string
}

/**
* QueryEpochProvisionsResponse is the response type for the
Query/EpochProvisions RPC method.
*/
export interface V1Beta1QueryEpochProvisionsResponse {
    /**
     * epoch_provisions is the current minting per epoch provisions value.
     * @format byte
     */
    epoch_provisions?: string
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 */
export interface V1Beta1QueryParamsResponse {
    /** params defines the parameters of the module. */
    params?: V1Beta1Params
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
 * @title stayking/mint/v1beta1/genesis.proto
 * @version version not set
 */
export class Api<
    SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryEpochProvisions
     * @summary EpochProvisions current minting epoch provisions value.
     * @request GET:/mint/v1beta1/epoch_provisions
     */
    queryEpochProvisions = (params: RequestParams = {}) =>
        this.request<V1Beta1QueryEpochProvisionsResponse, RpcStatus>({
            path: `/mint/v1beta1/epoch_provisions`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Params returns the total set of minting parameters.
     * @request GET:/mint/v1beta1/params
     */
    queryParams = (params: RequestParams = {}) =>
        this.request<V1Beta1QueryParamsResponse, RpcStatus>({
            path: `/mint/v1beta1/params`,
            method: 'GET',
            format: 'json',
            ...params,
        })
}
