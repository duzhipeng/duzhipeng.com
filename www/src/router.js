import { createRouter, createWebHistory } from "vue-router";

const Index = () => import("./pages/Index.vue");
const Paper = () => import("./pages/Paper.vue");
const ServerConfig = () => import("./pages/papers/ServerConfig.md");
const HukouOfShanghai = () => import("./pages/papers/HukouOfShanghai.md");

const routes = [
  {
    path: "/",
    name: "Index",
    component: Index,
    meta: {
      title: "",
    },
  },
  {
    path: "/paper",
    name: "Paper",
    component: Paper,
    meta: {
      title: "研究",
    },
    children: [
      {
        path: "server_config",
        name: "ServerConfig",
        component: ServerConfig,
        meta: {
          title: "配置服务器的操作的备忘（2023.01)",
        },
      },
      {
        path: "hukou",
        name: "HukouOfShanghai",
        component: HukouOfShanghai,
        meta: {
          title: "上海落户记（2023.7)",
        },
      },
    ],
  },
];

const Router = createRouter({
  history: createWebHistory(),
  routes,
});

// 前置拦截器
const defaultTitle = "阿杜的计算技术研究";
Router.beforeEach((to) => {
  // 更改页面Title
  document.title = to.meta["title"]
    ? to.meta["title"] + " - " + defaultTitle
    : defaultTitle;
});
export default Router;
