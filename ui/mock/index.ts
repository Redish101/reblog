import { defineMock } from "umi";

export default defineMock({
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
  "POST /api/admin/login": (req, res) => {
    if (req.body.username === "admin" && req.body.password === "admin") {
      res.status(200).json({
        success: true,
        msg: "success",
        data: {
          token: "admin-token",
        },
      });
    } else {
      res.status(401).json({
        success: false,
        msg: "用户名或密码错误",
        data: {},
      });
    }
  },
  "GET /api/admin/userInfo": {
    success: true,
    msg: "success",
    data: {
      username: "admin",
      nickname: "admin",
    },
  },
  "POST /api/article/:slug": (req, res) => {
    if (req.headers.authorization != "admin-token") {
      res.status(401).json({
        success: false,
        msg: "请先登录",
        data: {},
      });
      return;
    }
    if (req.body.title && req.body.content) {
      res.status(200).json({
        success: true,
        msg: "success",
        data: {},
      });
    }
  },
});
