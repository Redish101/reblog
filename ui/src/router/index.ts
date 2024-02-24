import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "home",
    component: () => import("@/view/Home.vue"),
    meta: {
      title: "首页",
    },
  },
  {
    path: "/login",
    name: "login",
    component: () => import("@/view/Login.vue"),
    meta: {
      title: "登录",
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach(async (to, _) => {
  if (to.name != "login" && localStorage.getItem("token") == null) {
    router.push({ name: "login" });
  }
  document.title = `${to.meta["title"]} | reblog admin`;
});

export default router;
