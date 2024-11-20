import { createRouter, createWebHistory } from "vue-router";

const Index = () => import("./pages/Index.vue");
const Paper = () => import("./pages/Paper.vue");
const ServerConfig = () => import("./pages/papers/ServerConfig.vue");
// const Login = () => import("./pages/Login.vue");
// const Garden = () => import("./pages/garden/Index.vue");
// const Shop = () => import("./pages/Items.vue");
// const Item = () => import("./pages/Item.vue");
// const Order = () => import("./pages/Order.vue");
// const Checkout = () => import("./pages/Checkout.vue");
// const Pay = () => import("./pages/Pay.vue");

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
      title: "文库",
    },
  },
  {
    path: "/paper/server_config",
    name: "ServerConfig",
    component: ServerConfig,
    meta: {
      title: "配置服务器的操作的备忘（v2023.01)",
    },
  },
  // {
  //   path: "/server_config",
  //   name: "ServerConfig",
  //   component: ServerConfig,
  //   meta: {
  //     title: "配置服务器的操作的备忘（v2023.01)",
  //   },
  // },
  // {
  //   path: "/garden",
  //   name: "garden",
  //   component: Garden,
  //   meta: {
  //     title: "菜园",
  //   },
  // },
  // {
  //   path: "/shop",
  //   name: "shop",
  //   component: Shop,
  //   meta: {
  //     title: "小店",
  //   },
  // },
  // {
  //   path: "/item/:id",
  //   name: "Item",
  //   component: Item,
  //   meta: { title: "小店" },
  // },
  // {
  //   path: "/order",
  //   name: "order",
  //   component: Order,
  //   meta: { title: "订单" },
  // },
  // {
  //   path: "/checkout",
  //   name: "Checkout",
  //   component: Checkout,
  //   meta: {
  //     title: "收银台",
  //   },
  // },
  // {
  //   path: "/pay",
  //   name: "Pay",
  //   component: Pay,
  //   meta: {
  //     title: "付款",
  //   },
  // },
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
