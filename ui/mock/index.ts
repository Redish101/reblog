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
  "GET /api/admin/tokenState": {
    success: true,
    msg: "success",
    data: true,
  },
  "GET /api/article/list": {
    success: true,
    msg: "success",
    data: {
      count: 2,
      articles: [
        {
          id: 1,
          slug: "hello-world",
          title: "Hello, World!",
          content: "This is the first article.",
          created_at: "2024-01-01 00:00:00",
          updated_at: "2024-01-01 00:00:00",
        },
        {
          id: 2,
          slug: "about-me",
          title: "About Me",
          content: "This is a about me article.",
          created_at: "2021-01-02 00:00:00",
          updated_at: "2021-01-02 00:00:00",
        },
      ],
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
        data: null,
      });
    }
  },
  "GET /api/site": {
    success: true,
    msg: "success",
    data: {
      name: "伟大的站点",
      url: "https://the-greatest-site.com",
      desc: "这是世界上最伟大的站点",
    },
  },
});
