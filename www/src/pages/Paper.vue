<template>
  <div class="menu">
    <a-breadcrumb>
      <a-breadcrumb-item>
        <a-link href="/"> 首页 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item> {{ meta.title }} </a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <a-row :gutter="20">
    <!--    侧栏-->
    <a-col flex="200px">
      <a-affix :offsetTop="80">
        <a-anchor class="anchor">
          <a-anchor-link v-for="(v, k) in meta.sections" :href="'#' + v.id">
            {{ v.name }}
            <template #sublist v-if="v.sub">
              <a-anchor-link :href="'#' + v.id" v-for="(v, k) in v.sub">
                {{ v.name }}
              </a-anchor-link>
            </template>
          </a-anchor-link>
        </a-anchor>
      </a-affix>
    </a-col>
    <a-col flex="auto">
      <div class="content">
        <router-view @syncMeta="syncMeta"> </router-view>
      </div>
    </a-col>
  </a-row>
  <div class="menu">
    <a-breadcrumb>
      <a-breadcrumb-item>
        <a-link href="/"> 首页 </a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item> {{ meta.title }} </a-breadcrumb-item>
    </a-breadcrumb>
  </div>
</template>

<script>
import "../assets/prism";
import { IconHome } from "@arco-design/web-vue/es/icon";

export default {
  name: "PaperContent",
  components: { IconHome },
  mounted() {},
  data() {
    return {
      meta: {},
    };
  },
  methods: {
    syncMeta(payload) {
      console.log("info", payload);
      this.meta = payload;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
@import "../assets/prism.css";
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
