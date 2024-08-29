export type ServerConfig = {
  url: string;
};

export type Config = {
  server: ServerConfig;
  cache: RequestCache;
};
