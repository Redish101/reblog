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

export interface CommonResp {
  data?: any;
  msg?: string;
  success?: boolean;
}

export interface GormDeletedAt {
  time?: string;
  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface HandlerAdminLoginResp {
  token?: string;
}

export interface HandlerArticlesListResp {
  articles?: ModelArticle[];
  count?: number;
}

export interface ModelArticle {
  content?: string;
  created_at?: string;
  deleted_at?: GormDeletedAt;
  desc?: string;
  id?: number;
  slug?: string;
  title?: string;
  updated_at?: string;
}

export interface ModelSite {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  desc?: string;
  icon?: string;
  id?: number;
  name?: string;
  updated_at?: string;
  url?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<
  FullRequestParams,
  "body" | "method" | "query" | "path"
>;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown>
  extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "//localhost:3000";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) =>
    fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(
      typeof value === "number" ? value : `${value}`,
    )}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter(
      (key) => "undefined" !== typeof query[key],
    );
    return keys
      .map((key) =>
        Array.isArray(query[key])
          ? this.addArrayQueryParam(query, key)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string")
        ? JSON.stringify(input)
        : input,
    [ContentType.Text]: (input: any) =>
      input !== null && typeof input !== "string"
        ? JSON.stringify(input)
        : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(
    params1: RequestParams,
    params2?: RequestParams,
  ): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (
    cancelToken: CancelToken,
  ): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

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
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(
      `${baseUrl || this.baseUrl || ""}${path}${
        queryString ? `?${queryString}` : ""
      }`,
      {
        ...requestParams,
        headers: {
          ...(requestParams.headers || {}),
          ...(type && type !== ContentType.FormData
            ? { "Content-Type": type }
            : {}),
        },
        signal:
          (cancelToken
            ? this.createAbortSignal(cancelToken)
            : requestParams.signal) || null,
        body:
          typeof body === "undefined" || body === null
            ? null
            : payloadFormatter(body),
      },
    ).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title reblog api
 * @version 1.0
 * @license GPL-V3
 * @baseUrl //localhost:3000
 * @contact
 */
export class Api<
  SecurityDataType extends unknown,
> extends HttpClient<SecurityDataType> {
  admin = {
    /**
     * @description 管理员使用用户名和密码进行登录，若登录成功，返回token
     *
     * @tags 站点管理
     * @name LoginCreate
     * @summary 登录
     * @request POST:/admin/login
     */
    loginCreate: (
      data: {
        /** 用户名或邮箱 */
        username: string;
        /** 密码 */
        password: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<
        CommonResp & {
          data?: HandlerAdminLoginResp;
        },
        CommonResp
      >({
        path: `/admin/login`,
        method: "POST",
        body: data,
        type: ContentType.UrlEncoded,
        format: "json",
        ...params,
      }),

    /**
     * @description 更新站点的名称、URL、描述和图标
     *
     * @tags 站点管理
     * @name SiteUpdate
     * @summary 更新站点信息
     * @request PUT:/admin/site
     * @secure
     */
    siteUpdate: (
      data: {
        /** 站点名称 */
        name: string;
        /** 站点URL */
        url: string;
        /** 站点描述 */
        desc: string;
        /** 站点图标(base64格式) */
        icon: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<CommonResp, CommonResp>({
        path: `/admin/site`,
        method: "PUT",
        body: data,
        secure: true,
        type: ContentType.UrlEncoded,
        format: "json",
        ...params,
      }),

    /**
     * @description 获取当前token的状态
     *
     * @tags 站点管理
     * @name TokenStateList
     * @summary 获取token状态
     * @request GET:/admin/tokenState
     * @secure
     */
    tokenStateList: (params: RequestParams = {}) =>
      this.request<
        CommonResp & {
          data?: boolean;
        },
        any
      >({
        path: `/admin/tokenState`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),
  };
  article = {
    /**
     * @description 分页获取文章列表
     *
     * @tags 文章
     * @name ListList
     * @summary 分页获取文章列表
     * @request GET:/article/list
     */
    listList: (
      query?: {
        /** 页码, 默认为1 */
        pageIndex?: number;
        /** 每页数量, 默认为10 */
        pageSize?: number;
      },
      params: RequestParams = {},
    ) =>
      this.request<
        CommonResp & {
          data?: HandlerArticlesListResp;
        },
        CommonResp
      >({
        path: `/article/list`,
        method: "GET",
        query: query,
        format: "json",
        ...params,
      }),

    /**
     * @description 根据slug获取文章详情
     *
     * @tags 文章
     * @name ArticleDetail
     * @summary 获取文章详情
     * @request GET:/article/{slug}
     */
    articleDetail: (slug: string, params: RequestParams = {}) =>
      this.request<
        CommonResp & {
          data?: ModelArticle;
        },
        CommonResp
      >({
        path: `/article/${slug}`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * @description 根据slug更新文章的标题、描述和内容
     *
     * @tags 文章
     * @name ArticleUpdate
     * @summary 更新文章
     * @request PUT:/article/{slug}
     * @secure
     */
    articleUpdate: (
      slug: string,
      data: {
        /** 文章的标题 */
        title: string;
        /** 文章的描述 */
        desc: string;
        /** 文章的内容 */
        content: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<CommonResp, CommonResp>({
        path: `/article/${slug}`,
        method: "PUT",
        body: data,
        secure: true,
        type: ContentType.FormData,
        format: "json",
        ...params,
      }),

    /**
     * @description 添加一篇新的文章
     *
     * @tags 文章
     * @name ArticleCreate
     * @summary 添加文章
     * @request POST:/article/{slug}
     * @secure
     */
    articleCreate: (
      slug: string,
      data: {
        /** 文章标题 */
        title: string;
        /** 文章描述 */
        desc: string;
        /** 文章内容 */
        content: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<CommonResp, CommonResp>({
        path: `/article/${slug}`,
        method: "POST",
        body: data,
        secure: true,
        type: ContentType.FormData,
        format: "json",
        ...params,
      }),

    /**
     * @description 根据slug删除文章
     *
     * @tags 文章
     * @name ArticleDelete
     * @summary 删除文章
     * @request DELETE:/article/{slug}
     * @secure
     */
    articleDelete: (slug: string, params: RequestParams = {}) =>
      this.request<CommonResp, CommonResp>({
        path: `/article/${slug}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),
  };
  init = {
    /**
     * @description 使用给定的参数初始化站点
     *
     * @tags 站点管理
     * @name InitCreate
     * @summary 初始化站点
     * @request POST:/init
     */
    initCreate: (
      data: {
        /** 用户名 */
        username: string;
        /** 昵称 */
        nickname: string;
        /** 邮箱 */
        email: string;
        /** 密码 */
        password: string;
        /** 站点名称 */
        name: string;
        /** 站点URL */
        url: string;
        /** 站点描述 */
        desc?: string;
        /** 站点图标(base64格式) */
        icon?: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<CommonResp, CommonResp>({
        path: `/init`,
        method: "POST",
        body: data,
        type: ContentType.UrlEncoded,
        format: "json",
        ...params,
      }),
  };
  site = {
    /**
     * @description 获取站点信息
     *
     * @tags 站点
     * @name SiteList
     * @summary 获取站点信息
     * @request GET:/site
     */
    siteList: (params: RequestParams = {}) =>
      this.request<
        CommonResp & {
          data?: ModelSite;
        },
        any
      >({
        path: `/site`,
        method: "GET",
        format: "json",
        ...params,
      }),
  };
  user = {
    /**
     * @description 管理员更新用户信息
     *
     * @tags 站点管理
     * @name UserUpdate
     * @summary 更新用户信息
     * @request PUT:/user/{username}
     */
    userUpdate: (
      username: string,
      data: {
        /** 昵称 */
        nickname: string;
        /** 邮箱 */
        email: string;
        /** 密码 */
        password: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<CommonResp, CommonResp>({
        path: `/user/${username}`,
        method: "PUT",
        body: data,
        type: ContentType.FormData,
        format: "json",
        ...params,
      }),
  };
}
