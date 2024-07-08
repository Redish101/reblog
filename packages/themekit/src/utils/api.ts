export type Resp<T> = {
  success: boolean;
  msg: string;
  data?: T | null;
};

export class Api {
  constructor(private baseURL: string) {}

  async get<T>(url: string, options?: RequestInit) {
    return await useApi<T>(this.baseURL + url, options);
  }
}

async function useApi<T>(url: string, options?: RequestInit) {
  const response = await fetch(url, options);
  const json = await response.json();
  if (!response.ok) {
    const text = await response.text();
    throw new Error(`reblog api error: ${text}`);
  }
  const data = json as Resp<T>;

  if (!data.success) throw new Error(`reblog api error: ${data.msg}`);

  return data;
}
