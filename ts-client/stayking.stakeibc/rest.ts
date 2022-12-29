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

export enum ValidatorValidatorStatus {
    ACTIVE = 'ACTIVE',
    INACTIVE = 'INACTIVE',
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

export interface StakeibcDelegation {
    delegate_acct_address?: string
    validator?: StakeibcValidator
    amt?: string
}

export interface StakeibcEpochTracker {
    epoch_identifier?: string

    /** @format uint64 */
    epoch_number?: string

    /** @format uint64 */
    next_epoch_start_time?: string

    /** @format uint64 */
    duration?: string
}

export interface StakeibcHostZone {
    chain_id?: string
    connection_id?: string
    bech32prefix?: string
    transfer_channel_id?: string
    validators?: StakeibcValidator[]
    blacklisted_validators?: StakeibcValidator[]
    withdrawal_account?: StakeibcICAAccount
    fee_account?: StakeibcICAAccount
    delegation_account?: StakeibcICAAccount
    redemption_account?: StakeibcICAAccount

    /** ibc denom on stride */
    ibc_denom?: string

    /** native denom on host zone */
    host_denom?: string

    /**
     * TODO(TEST-68): Should we make this an array and store the last n redemption
     * rates then calculate a TWARR?
     */
    last_redemption_rate?: string
    redemption_rate?: string

    /**
     * stores how many days we should wait before issuing unbondings
     * @format uint64
     */
    unbonding_frequency?: string

    /** TODO(TEST-101) int to dec */
    staked_bal?: string
    address?: string
}

export interface StakeibcICAAccount {
    address?: string
    delegations?: StakeibcDelegation[]
    target?: StakeibcICAAccountType
}

export enum StakeibcICAAccountType {
    DELEGATION = 'DELEGATION',
    FEE = 'FEE',
    WITHDRAWAL = 'WITHDRAWAL',
    REDEMPTION = 'REDEMPTION',
}

export type StakeibcMsgAddValidatorResponse = object

export type StakeibcMsgChangeValidatorWeightResponse = object

export type StakeibcMsgClaimUndelegatedTokensResponse = object

export type StakeibcMsgClearBalanceResponse = object

export type StakeibcMsgDeleteValidatorResponse = object

export type StakeibcMsgLiquidStakeResponse = object

export type StakeibcMsgRebalanceValidatorsResponse = object

export type StakeibcMsgRedeemStakeResponse = object

export type StakeibcMsgRegisterHostZoneResponse = object

export type StakeibcMsgRestoreInterchainAccountResponse = object

export type StakeibcMsgUpdateValidatorSharesExchRateResponse = object

export interface StakeibcParams {
    /**
     * define epoch lengths, in stride_epochs
     * @format uint64
     */
    rewards_interval?: string

    /** @format uint64 */
    delegate_interval?: string

    /** @format uint64 */
    deposit_interval?: string

    /** @format uint64 */
    redemption_rate_interval?: string

    /** @format uint64 */
    stride_commission?: string

    /**
     * zone_com_address stores which addresses to
     * send the Stride commission too, as well as what portion
     * of the fee each address is entitled to
     * TODO implement this
     */
    zone_com_address?: Record<string, string>

    /** @format uint64 */
    reinvest_interval?: string

    /** @format uint64 */
    validator_rebalancing_threshold?: string

    /** @format uint64 */
    ica_timeout_nanos?: string

    /** @format uint64 */
    buffer_size?: string

    /** @format uint64 */
    ibc_timeout_blocks?: string

    /** @format uint64 */
    fee_transfer_timeout_nanos?: string

    /** @format uint64 */
    max_stake_ica_calls_per_epoch?: string

    /** @format uint64 */
    safety_min_redemption_rate_threshold?: string

    /** @format uint64 */
    safety_max_redemption_rate_threshold?: string

    /** @format uint64 */
    ibc_transfer_timeout_nanos?: string

    /** @format uint64 */
    safety_num_validators?: string
}

export interface StakeibcQueryAllEpochTrackerResponse {
    epoch_tracker?: StakeibcEpochTracker[]

    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse
}

export interface StakeibcQueryAllHostZoneResponse {
    host_zone?: StakeibcHostZone[]

    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse
}

export interface StakeibcQueryGetEpochTrackerResponse {
    epoch_tracker?: StakeibcEpochTracker
}

export interface StakeibcQueryGetHostZoneResponse {
    host_zone?: StakeibcHostZone
}

export interface StakeibcQueryGetICAAccountResponse {
    ica_account?: StakeibcICAAccount
}

export interface StakeibcQueryGetValidatorsResponse {
    validators?: StakeibcValidator[]
}

export interface StakeibcQueryInterchainAccountFromAddressResponse {
    interchain_account_address?: string
}

export interface StakeibcQueryModuleAddressResponse {
    addr?: string
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface StakeibcQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: StakeibcParams
}

export interface StakeibcValidator {
    name?: string
    address?: string
    status?: ValidatorValidatorStatus

    /** @format uint64 */
    commission_rate?: string
    delegation_amt?: string

    /** @format uint64 */
    weight?: string
    internal_exchange_rate?: StakeibcValidatorExchangeRate
}

export interface StakeibcValidatorExchangeRate {
    internal_tokens_to_shares_rate?: string

    /** @format uint64 */
    epoch_number?: string
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string

    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string

    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string

    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    count_total?: boolean

    /** reverse is set to true if results are to be returned in the descending order. */
    reverse?: boolean
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
    /**
     * next_key is the key to be passed to PageRequest.key to
     * query the next page most efficiently
     * @format byte
     */
    next_key?: string

    /**
     * total is total number of results available if PageRequest.count_total
     * was set, its value is undefined otherwise
     * @format uint64
     */
    total?: string
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
 * @title stayking/stakeibc/callbacks.proto
 * @version version not set
 */
export class Api<
    SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryEpochTrackerAll
     * @summary Queries a list of EpochTracker items.
     * @request GET:/stayking/stakeibc/epoch_tracker
     */
    queryEpochTrackerAll = (
        query?: {
            'pagination.key'?: string
            'pagination.offset'?: string
            'pagination.limit'?: string
            'pagination.count_total'?: boolean
            'pagination.reverse'?: boolean
        },
        params: RequestParams = {}
    ) =>
        this.request<StakeibcQueryAllEpochTrackerResponse, RpcStatus>({
            path: `/stayking/stakeibc/epoch_tracker`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryEpochTracker
     * @summary Queries a EpochTracker by index.
     * @request GET:/stayking/stakeibc/epoch_tracker/{epoch_identifier}
     */
    queryEpochTracker = (epochIdentifier: string, params: RequestParams = {}) =>
        this.request<StakeibcQueryGetEpochTrackerResponse, RpcStatus>({
            path: `/stayking/stakeibc/epoch_tracker/${epochIdentifier}`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryHostZoneAll
     * @summary Queries a list of HostZone items.
     * @request GET:/stayking/stakeibc/host_zone
     */
    queryHostZoneAll = (
        query?: {
            'pagination.key'?: string
            'pagination.offset'?: string
            'pagination.limit'?: string
            'pagination.count_total'?: boolean
            'pagination.reverse'?: boolean
        },
        params: RequestParams = {}
    ) =>
        this.request<StakeibcQueryAllHostZoneResponse, RpcStatus>({
            path: `/stayking/stakeibc/host_zone`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryHostZone
     * @summary Queries a HostZone by id.
     * @request GET:/stayking/stakeibc/host_zone/{chain_id}
     */
    queryHostZone = (chainId: string, params: RequestParams = {}) =>
        this.request<StakeibcQueryGetHostZoneResponse, RpcStatus>({
            path: `/stayking/stakeibc/host_zone/${chainId}`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryIcaAccount
     * @summary Queries a ICAAccount by index.
     * @request GET:/stayking/stakeibc/ica_account
     */
    queryIcaAccount = (params: RequestParams = {}) =>
        this.request<StakeibcQueryGetICAAccountResponse, RpcStatus>({
            path: `/stayking/stakeibc/ica_account`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryModuleAddress
     * @summary Queries a list of ModuleAddress items.
     * @request GET:/stayking/stakeibc/module_address/{name}
     */
    queryModuleAddress = (name: string, params: RequestParams = {}) =>
        this.request<StakeibcQueryModuleAddressResponse, RpcStatus>({
            path: `/stayking/stakeibc/module_address/${name}`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/stayking/stakeibc/params
     */
    queryParams = (params: RequestParams = {}) =>
        this.request<StakeibcQueryParamsResponse, RpcStatus>({
            path: `/stayking/stakeibc/params`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryValidators
     * @summary Queries a Validator by host zone.
     * @request GET:/stayking/stakeibc/validators/{chain_id}
     */
    queryValidators = (chainId: string, params: RequestParams = {}) =>
        this.request<StakeibcQueryGetValidatorsResponse, RpcStatus>({
            path: `/stayking/stakeibc/validators/${chainId}`,
            method: 'GET',
            format: 'json',
            ...params,
        })
}
