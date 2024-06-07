import { defineConfig } from "umi";

export default defineConfig({
  routes: [
    {
      path: "/",
      component: "index",
    },
    {
      path: "/init",
      component: "init",
      layout: false,
    },
    {
      path: "/*",
      component: "404",
    },
  ],
  npmClient: "pnpm",
  title: "reblog dashboard",
  icons: {},
});
