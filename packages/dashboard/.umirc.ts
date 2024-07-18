import { defineConfig } from "umi";

export default defineConfig({
  history: { type: "hash" },
  routes: [
    {
      path: "/",
      component: "index",
    },
    {
      path: "/login",
      component: "login",
      layout: false,
    },
    {
      path: "/init",
      component: "init",
      layout: false,
    },
    {
      path: "/article",
      component: "article",
    },
    {
      path: "/article/create",
      component: "article/create",
    },
    {
      path: "/article/edit/:slug",
      component: "article/edit",
    },
    {
      path: "/site",
      component: "site",
    },
    {
      path: "/friend",
      component: "friend",
    },
    {
      path: "/*",
      component: "404",
    },
  ],
  npmClient: "pnpm",
  title: "reblog dashboard",
  icons: {},
  esbuildMinifyIIFE: true,
  proxy: {
    "/api": {
      target: "http://localhost:3000",
    },
  },
});
