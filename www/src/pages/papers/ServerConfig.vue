<template>
  <div class="menu">
    <a-breadcrumb>
      <a-breadcrumb-item>
        <a-link href="/"> 首页 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item>
        <a-link href="/paper"> 研究 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item>配置服务器的操作的备忘（v2023.2)</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <a-row :gutter="20">
    <a-col flex="200px">
      <a-affix :offsetTop="80">
        <a-anchor class="anchor">
          <a-anchor-link href="#system">适用系统</a-anchor-link>
          <a-anchor-link href="#environment">更新环境</a-anchor-link>
          <a-anchor-link href="#disk">挂载数据盘</a-anchor-link>
          <a-anchor-link href="#adduser">增加用户</a-anchor-link>
          <a-anchor-link href="#install">安装必要软件</a-anchor-link>
          <a-anchor-link href="#nginx">
            配置 Nginx
            <template #sublist>
              <a-anchor-link href="#single">单前端配置参考</a-anchor-link>
              <a-anchor-link href="#multiple">前后端配置参考</a-anchor-link>
              <a-anchor-link href="#enable">启用配置</a-anchor-link>
            </template>
          </a-anchor-link>
          <a-anchor-link href="#supervisor">
            配置 Supervisor
            <template #sublist>
              <a-anchor-link href="#extension">文件后缀 .conf</a-anchor-link>
              <a-anchor-link href="#common">常用命令</a-anchor-link>
            </template>
          </a-anchor-link>
        </a-anchor>
      </a-affix>
    </a-col>
    <a-col flex="auto">
      <div class="content"><Content /></div>
    </a-col>
  </a-row>
  <div class="menu">
    <a-breadcrumb>
      <a-breadcrumb-item>
        <a-link href="/"> 首页 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item>
        <a-link href="/paper"> 研究 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item>配置服务器的操作的备忘（v2023.01)</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
</template>

<script>
import "../../assets/prism";
import Content from "./server_config.md";
import { IconHome } from "@arco-design/web-vue/es/icon";

export default {
  name: "Index",
  components: { Content, IconHome },
  mounted() {},
  data() {
    return {};
  },
  methods: {},
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
@import "../../assets/prism.css";
</style>
<style scoped>
.menu {
  margin-top: 30px;
}
.anchor {
  margin: 30px 0 0 20px;
  background-color: #fff;
}
.content {
  max-width: 680px;
  color: var(--color-text-1);
  line-height: 1.5715;
}
</style>
