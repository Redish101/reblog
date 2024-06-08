import { message } from "antd";
import { history } from "umi";

const useApi = async (url: string | URL | Request, options?: RequestInit) => {
  const token = localStorage.getItem("token");

  if (token) {
    const headers =
      options && options.headers ? new Headers(options.headers) : new Headers();

    headers.append("Authorization", token);

    options = {
      ...options,
      headers,
    };
  }

  const res = await fetch(url, options);

  if (res.status == 401) {
    localStorage.removeItem("token");
    message.open({
      type: "info",
      content: "登录信息已过期，请重新登录",
    });
    history.push("/login");
  }

  return res;
};

export default useApi;
