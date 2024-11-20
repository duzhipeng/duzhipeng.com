import { createApp } from "vue";
import Router from "./router";
import App from "./App.vue";
import AxiosHttp from "./utils/axios";
// import { createPinia } from "pinia";
import "./assets/theme.less";

const app = createApp(App);
app.config.productionTip = false;
app.use(Router);
app.use(AxiosHttp);
// app.use(createPinia());
app.mount("#app");
