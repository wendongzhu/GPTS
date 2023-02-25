import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [

  {
    path: "/setting",
    name: "setting",
    component: () => import("../components/setting/index.vue"),
  },
  {
    path: "/collection",
    name: "collection",
    component: () => import("../components/collection/index.vue"),
  },
  {
    path: "/analysis",
    name: "analysis",
    component: () => import("../components/analysis/index.vue"),
  },
  {
    path: "/history",
    name: "history",
    component: () => import("../components/history/index.vue"),
  },
  {
    path: "/help",
    name: "help",
    component: () => import("../components/help/index.vue"),
  },

];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
