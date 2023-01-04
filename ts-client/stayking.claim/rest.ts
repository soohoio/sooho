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

export enum ClaimAction {
    ACTION_FREE = 'ACTION_FREE',
    ACTION_LIQUID_STAKE = 'ACTION_LIQUID_STAKE',
    ACTION_DELEGATE_STAKE = 'ACTION_DELEGATE_STAKE',
}

export interface ClaimAirdrop {
    airdrop_identifier?: string

    /**
     * seconds
     * @format date-time
     */
    airdrop_start_time?: string

    /** seconds */
    airdrop_duration?: string

    /** denom of claimable asset */
    claim_denom?: string

    /** airdrop distribution account */
    distributor_address?: string

    /** ustrd tokens claimed so far in the current period */
    claimed_so_far?: string
}

export interface ClaimClaimRecord {
    /** airdrop identifier */
    airdrop_identifier?: string

    /** address of claim user */
    address?: string

    /** weight that represent the portion from total allocation */
    weight?: string

    /**
     * true if action is completed
     * index of bool in array refers to action enum #
     */
    action_completed?: boolean[]
}

export interface ClaimMsgClaimFreeAmountResponse {
    claimed_amount?: V1Beta1Coin[]
}

export type ClaimMsgCreateAirdropResponse = object

export type ClaimMsgDeleteAirdropResponse = object

export type ClaimMsgSetAirdropAllocationsResponse = object

export interface ClaimQueryClaimRecordResponse {
    claim_record?: ClaimClaimRecord
}

export interface ClaimQueryClaimableForActionResponse {
    coins?: V1Beta1Coin[]
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 */
export interface ClaimQueryDistributorAccountBalanceResponse {
    /** params defines the parameters of the module. */
    distributor_account_balance?: V1Beta1Coin[]
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 */
export interface ClaimQueryParamsResponse {
    /** params defines the parameters of the module. */
    params?: StaykingclaimParams
}

export interface ClaimQueryTotalClaimableResponse {
    coins?: V1Beta1Coin[]
}

export interface ClaimQueryUserVestingsResponse {
    spendable_coins?: V1Beta1Coin[]
    periods?: VestingPeriod[]
}

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface ProtobufAny {
    /**
     * A URL/resource name that uniquely identifies the type of the serialized
     * protocol buffer message. This string must contain at least
     * one "/" character. The last segment of the URL's path must represent
     * the fully qualified name of the type (as in
     * `path/google.protobuf.Duration`). The name should be in a canonical form
     * (e.g., leading "." is not accepted).
     *
     * In practice, teams usually precompile into the binary all types that they
     * expect it to use in the context of Any. However, for URLs which use the
     * scheme `http`, `https`, or no scheme, one can optionally set up a type
     * server that maps type URLs to message definitions as follows:
     * * If no scheme is provided, `https` is assumed.
     * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
     *   value in binary format, or produce an error.
     * * Applications are allowed to cache lookup results based on the
     *   URL, or have them precompiled into a binary to avoid any
     *   lookup. Therefore, binary compatibility needs to be preserved
     *   on changes to types. (Use versioned type names to manage
     *   breaking changes.)
     * Note: this functionality is not currently available in the official
     * protobuf release, and it is not used for type URLs beginning with
     * type.googleapis.com.
     * Schemes other than `http`, `https` (or the empty scheme) might be
     * used with implementation specific semantics.
     */
    '@type'?: string
}

export interface RpcStatus {
    /** @format int32 */
    code?: number
    message?: string
    details?: ProtobufAny[]
}

/**
 * Params defines the claim module's parameters.
 */
export interface StaykingclaimParams {
    airdrops?: ClaimAirdrop[]
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
    denom?: string
    amount?: string
}

/**
 * Period defines a length of time and amount of coins that will vest.
 */
export interface VestingPeriod {
    /** @format int64 */
    start_time?: string

    /** @format int64 */
    length?: string
    amount?: V1Beta1Coin[]

    /** @format int32 */
    action_type?: number
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
 * @title stayking/claim/claim.proto
 * @version version not set
 */
export class Api<
    SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryClaimRecord
     * @request GET:/claim/claim_record/{address}
     */
    queryClaimRecord = (
        address: string,
        query?: { airdrop_identifier?: string },
        params: RequestParams = {}
    ) =>
        this.request<ClaimQueryClaimRecordResponse, RpcStatus>({
            path: `/claim/claim_record/${address}`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryClaimableForAction
     * @request GET:/claim/claimable_for_action/{address}/{action}
     */
    queryClaimableForAction = (
        address: string,
        action: 'ACTION_FREE' | 'ACTION_LIQUID_STAKE' | 'ACTION_DELEGATE_STAKE',
        query?: { airdrop_identifier?: string },
        params: RequestParams = {}
    ) =>
        this.request<ClaimQueryClaimableForActionResponse, RpcStatus>({
            path: `/claim/claimable_for_action/${address}/${action}`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryDistributorAccountBalance
     * @request GET:/claim/module_account_balance
     */
    queryDistributorAccountBalance = (
        query?: { airdrop_identifier?: string },
        params: RequestParams = {}
    ) =>
        this.request<ClaimQueryDistributorAccountBalanceResponse, RpcStatus>({
            path: `/claim/module_account_balance`,
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
     * @request GET:/claim/params
     */
    queryParams = (params: RequestParams = {}) =>
        this.request<ClaimQueryParamsResponse, RpcStatus>({
            path: `/claim/params`,
            method: 'GET',
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryTotalClaimable
     * @request GET:/claim/total_claimable/{address}
     */
    queryTotalClaimable = (
        address: string,
        query?: { airdrop_identifier?: string; include_claimed?: boolean },
        params: RequestParams = {}
    ) =>
        this.request<ClaimQueryTotalClaimableResponse, RpcStatus>({
            path: `/claim/total_claimable/${address}`,
            method: 'GET',
            query: query,
            format: 'json',
            ...params,
        })

    /**
     * No description
     *
     * @tags Query
     * @name QueryUserVestings
     * @request GET:/claim/user_vestings/{address}
     */
    queryUserVestings = (address: string, params: RequestParams = {}) =>
        this.request<ClaimQueryUserVestingsResponse, RpcStatus>({
            path: `/claim/user_vestings/${address}`,
            method: 'GET',
            format: 'json',
            ...params,
        })
}