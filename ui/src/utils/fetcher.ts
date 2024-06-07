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

  return fetch(url, options);
};

export default useApi;
