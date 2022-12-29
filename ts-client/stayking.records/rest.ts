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

export enum DepositRecordSource {
    STRIDE = 'STRIDE',
    WITHDRAWAL_ICA = 'WITHDRAWAL_ICA',
}

export interface GooglerpcStatus {
    /** @format int32 */
    code?: number
    message?: string
    details?: ProtobufAny[]
}

export interface ProtobufAny {
    '@type'?: string
}

export interface RecordsDepositRecord {
    /** @format uint64 */
    id?: string
    amount?: string
    denom?: string
    host_zone_id?: string
    status?: RecordsDepositRecordStatus

    /** @format uint64 */
    deposit_epoch_number?: string
    source?: DepositRecordSource
}

export enum RecordsDepositRecordStatus {
    TRANSFER_QUEUE = 'TRANSFER_QUEUE',
    TRANSFER_IN_PROGRESS = 'TRANSFER_IN_PROGRESS',
    DELEGATION_QUEUE = 'DELEGATION_QUEUE',
    DELEGATION_IN_PROGRESS = 'DELEGATION_IN_PROGRESS',
}

export interface RecordsEpochUnbondingRecord {
    /** @format uint64 */
    epoch_number?: string
    host_zone_unbondings?: RecordsHostZoneUnbonding[]
}

export interface RecordsHostZoneUnbonding {
    st_token_amount?: string
    native_token_amount?: string
    denom?: string
    host_zone_id?: string

    /** @format uint64 */
    unbonding_time?: string
    status?: RecordsHostZoneUnbondingStatus
    user_redemption_records?: string[]
}

export enum RecordsHostZoneUnbondingStatus {
    UNBONDING_QUEUE = 'UNBONDING_QUEUE',
    UNBONDING_IN_PROGRESS = 'UNBONDING_IN_PROGRESS',
    EXIT_TRANSFER_QUEUE = 'EXIT_TRANSFER_QUEUE',
    EXIT_TRANSFER_IN_PROGRESS = 'EXIT_TRANSFER_IN_PROGRESS',
    CLAIMABLE = 'CLAIMABLE',
}

/**
 * Params defines the parameters for the module.
 */
export type RecordsParams = object

export interface RecordsQueryAllDepositRecordResponse {
    deposit_record?: RecordsDepositRecord[]

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

export interface RecordsQueryAllEpochUnbondingRecordResponse {
    epoch_unbonding_record?: RecordsEpochUnbondingRecord[]

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

export interface RecordsQueryAllUserRedemptionRecordForUserResponse {
    user_redemption_record?: RecordsUserRedemptionRecord[]

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

export interface RecordsQueryAllUserRedemptionRecordResponse {
    user_redemption_record?: RecordsUserRedemptionRecord[]

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

export interface RecordsQueryGetDepositRecordResponse {
    deposit_record?: RecordsDepositRecord
}

export interface RecordsQueryGetEpochUnbondingRecordResponse {
    epoch_unbonding_record?: RecordsEpochUnbondingRecord
}

export interface RecordsQueryGetUserRedemptionRecordResponse {
    user_redemption_record?: RecordsUserRedemptionRecord
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RecordsQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: RecordsParams
}

export interface RecordsUserRedemptionRecord {
    /** {chain_id}.{epoch}.{sender} */
    id?: string
    sender?: string
    receiver?: string
    amount?: string
    denom?: string
    host_zone_id?: string

    /** @format uint64 */
    epoch_number?: string
    claim_is_pending?: boolean
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
 * @title stayking/records/callbacks.proto
 * @version version not set
 */
export class Api<
    SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryDepositRecordAll
     * @summary Queries a list of DepositRecord items.
     * @request GET:/Stride-Labs/stayking/records/deposit_record
     */
    queryDepositRecordAll = (
        query?: {
            'pagination.key'?: string
            'pagination.offset'?: string
            'pagination.limit'?: string
            'pagination.count_total'?: boolean
            'pagination.reverse'?: boolean
        },
        params: RequestParams = {}
    ) =>
        this.request<RecordsQueryAllDepositRecordResponse, GooglerpcStatus>({
            path: `/Stride-Labs/stayking/records/deposit_record`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryDepositRecord
     * @summary Queries a DepositRecord by id.
     * @request GET:/Stride-Labs/stayking/records/deposit_record/{id}
     */
    queryDepositRecord = (id: string, params: RequestParams = {}) =>
        this.request<RecordsQueryGetDepositRecordResponse, GooglerpcStatus>({
            path: `/Stride-Labs/stayking/records/deposit_record/${id}`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryEpochUnbondingRecordAll
     * @summary Queries a list of EpochUnbondingRecord items.
     * @request GET:/Stride-Labs/stayking/records/epoch_unbonding_record
     */
    queryEpochUnbondingRecordAll = (
        query?: {
            'pagination.key'?: string
            'pagination.offset'?: string
            'pagination.limit'?: string
            'pagination.count_total'?: boolean
            'pagination.reverse'?: boolean
        },
        params: RequestParams = {}
    ) =>
        this.request<
            RecordsQueryAllEpochUnbondingRecordResponse,
            GooglerpcStatus
        >({
            path: `/Stride-Labs/stayking/records/epoch_unbonding_record`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryEpochUnbondingRecord
     * @summary Queries a EpochUnbondingRecord by id.
     * @request GET:/Stride-Labs/stayking/records/epoch_unbonding_record/{epoch_number}
     */
    queryEpochUnbondingRecord = (
        epochNumber: string,
        params: RequestParams = {}
    ) =>
        this.request<
            RecordsQueryGetEpochUnbondingRecordResponse,
            GooglerpcStatus
        >({
            path: `/Stride-Labs/stayking/records/epoch_unbonding_record/${epochNumber}`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryUserRedemptionRecordAll
     * @summary Queries a list of UserRedemptionRecord items.
     * @request GET:/Stride-Labs/stayking/records/user_redemption_record
     */
    queryUserRedemptionRecordAll = (
        query?: {
            'pagination.key'?: string
            'pagination.offset'?: string
            'pagination.limit'?: string
            'pagination.count_total'?: boolean
            'pagination.reverse'?: boolean
        },
        params: RequestParams = {}
    ) =>
        this.request<
            RecordsQueryAllUserRedemptionRecordResponse,
            GooglerpcStatus
        >({
            path: `/Stride-Labs/stayking/records/user_redemption_record`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryUserRedemptionRecord
     * @summary Queries a UserRedemptionRecord by id.
     * @request GET:/Stride-Labs/stayking/records/user_redemption_record/{id}
     */
    queryUserRedemptionRecord = (id: string, params: RequestParams = {}) =>
        this.request<
            RecordsQueryGetUserRedemptionRecordResponse,
            GooglerpcStatus
        >({
            path: `/Stride-Labs/stayking/records/user_redemption_record/${id}`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryUserRedemptionRecordForUser
     * @summary Queries a list of UserRedemptionRecord items by chainId / userId pair.
     * @request GET:/Stride-Labs/stayking/records/user_redemption_record_for_user/{chain_id}/{day}/{address}/{limit}
     */
    queryUserRedemptionRecordForUser = (
        chainId: string,
        day: string,
        address: string,
        limit: string,
        query?: {
            'pagination.key'?: string
            'pagination.offset'?: string
            'pagination.limit'?: string
            'pagination.count_total'?: boolean
            'pagination.reverse'?: boolean
        },
        params: RequestParams = {}
    ) =>
        this.request<
            RecordsQueryAllUserRedemptionRecordForUserResponse,
            GooglerpcStatus
        >({
            path: `/Stride-Labs/stayking/records/user_redemption_record_for_user/${chainId}/${day}/${address}/${limit}`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/StrideLabs/stayking/records/params
     */
    queryParams = (params: RequestParams = {}) =>
        this.request<RecordsQueryParamsResponse, GooglerpcStatus>({
            path: `/StrideLabs/stayking/records/params`,
            method: 'GET',
            format: 'json',
            ...params,
        })
}
