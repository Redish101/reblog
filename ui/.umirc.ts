import { defineConfig } from "umi";

export default defineConfig({
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
      path: "/*",
      component: "404",
    },
  ],
  npmClient: "pnpm",
  title: "reblog dashboard",
  icons: {},
  clientLoader: {},
});
