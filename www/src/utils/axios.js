import axios from "axios";
import router from "../router";

const AxiosHttp = {
  install: (app) => {
    // Axios
    axios.defaults.timeout = 3000;
    axios.defaults.baseURL = import.meta.env.VITE_APP_CORE_URL;
    axios.defaults.headers.post["Content-Type"] =
      "application/json;charset=UTF-8";
    axios.defaults.withCredentials = true;

    // 请求拦截器
    axios.interceptors.request.use(
      function (config) {
        let method = config.method.toLowerCase();
        if (
          method === "post" ||
          method === "put" ||
          method === "patch" ||
          method === "delete"
        ) {
        }
        return config;
      },
      function (error) {
        return Promise.reject(error);
      }
    );
    // 响应拦截器
    axios.interceptors.response.use(
      function (response) {
        return response;
      },
      function (error) {
        if (error.response) {
          switch (error.response.status) {
            // 401 身份认证信息有误（用户名密码错误）
            // 403 越权操作
            case 401:
            case 403:
              if (router.currentRoute.value.path !== "/login") {
                router
                  .replace({
                    path: "/login",
                    query: {
                      redirect: router.currentRoute.value.fullPath,
                    },
                  })
                  .catch((err) => err);
              }
              break;
          }
        }
        return Promise.reject(error);
      }
    );
    app.config.globalProperties.$http = axios;
  },
};
export default AxiosHttp;
