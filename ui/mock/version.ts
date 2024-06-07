interface Response {
  success: boolean;
  msg: string;
  data?: any;
}

const mock: Record<string, Response> = {
  "GET /api/version": {
    success: true,
    msg: "success",
    data: {
      version: "dev",
      commit: "dev",
      app_name: "reblog-dev.dev(mock)",
      runtime: "go1.22(mock)",
    },
  },
  "POST /api/admin/login": {
    success: true,
    msg: "success",
    data: {
      token: "admin_token",
    },
  },
  "GET /api/admin/userInfo": {
    success: true,
    msg: "success",
    data: {
      username: "admin",
      nickname: "admin",
    },
  },
};

export default mock;
