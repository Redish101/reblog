import { message } from "antd";
import { history } from "umi";

const useApi = async (
  url: string | URL | Request,
  opts?: {
    method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH" | "OPTIONS" | "HEAD";
    data?: any;
    options?: RequestInit;
  },
) => {
  const token = localStorage.getItem("token");

  let headers =
    opts?.options?.headers
      ? new Headers(opts.options.headers)
      : new Headers();

  headers.append("Authorization", token!);
  headers.append("Content-Type", "application/json");

  const requestOptions: RequestInit = {
    method: opts?.method || "GET",
    headers,
    body: opts?.data ? JSON.stringify(opts.data) : undefined,
    ...opts?.options,
  };

  const res = await fetch(url, requestOptions);

  if (res.status === 401) {
    localStorage.removeItem("token");

    message.open({
      type: "warning",
      content: "请登录"
    })

    history.push("/login")
  }

  return res;
};

export default useApi;
