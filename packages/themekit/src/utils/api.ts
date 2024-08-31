export type Resp<T> = {
  success: boolean;
  msg: string;
  data?: T | null;
};

export class Api {
  constructor(
    private baseURL: string,
    private requestCache: RequestCache = "no-store",
  ) {}

  async useApi<T>(url: string, options?: RequestInit) {
    const response = await fetch(url, {
      cache: this.requestCache,
      ...options,
    });
    const json = await response.json();
    const data = json as Resp<T>;

    return data;
  }

  async get<T>(url: string, options?: RequestInit) {
    return await this.useApi<T>(this.baseURL + url, options);
  }

  async post<T>(url: string, body: any, options?: RequestInit) {
    return await this.useApi<T>(this.baseURL + url, {
      method: "POST",
      body,
      ...options,
    });
  }
}
