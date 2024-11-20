<template>
  <div class="menu">
    <a-breadcrumb>
      <a-breadcrumb-item>
        <a-link href="/"> 首页 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item> 研究 </a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <div class="content">
    <a-link href="/paper/server_config" icon>
      配置服务器的操作的备忘（v2023.01)
    </a-link>
  </div>
  <div class="menu">
    <a-breadcrumb>
      <a-breadcrumb-item>
        <a-link href="/"> 首页 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item> 研究 </a-breadcrumb-item>
    </a-breadcrumb>
  </div>
</template>

<script>
import { IconArrowLeft } from "@arco-design/web-vue/es/icon";

export default {
  name: "Index",
  components: { IconArrowLeft },
  mounted() {},
  data() {
    return {};
  },
  methods: {},
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.menu {
  margin-top: 30px;
}
.content {
  margin: 30px 0 0 20px;
  max-width: 980px;
}
</style>
