<template>
  <div class="wrapper">
    <a-row :gutter="20" justify="center">
      <a-col :xs="24" :lg="15" class="left">
        <a-avatar :size="300" shape="square" class="logo">
          <img src="../assets/logo.png" alt="avatar" />
        </a-avatar>
        <a-space direction="vertical" fill>
          <a-typography-text code>电子邮箱</a-typography-text>
          <a-typography-text copyable>DUZHIPENG@ME.COM</a-typography-text>
          <a-link
            href="https://keys.openpgp.org/vks/v1/by-fingerprint/1C5E38AEB5BEBAFEB6BCC9D0F3C82540B55E0FEA"
            target="_blank"
          >
            <a-typography-text type="primary" class="text-size-small"
              >( PGP PUBLIC KEY )</a-typography-text
            ></a-link
          >
          <a-typography-text code>呼号</a-typography-text>
          <a-typography-text bold type="success">BH4DGU</a-typography-text>
          <a-typography-text code>QQ号</a-typography-text>
          <a-typography-text copyable>7171891</a-typography-text>
          <a-typography-text code>微信号</a-typography-text>
          <a-typography-text copyable>DUZHIPENGSH</a-typography-text>
          <a-typography-text code>常住</a-typography-text>
          <a-typography-text>中国 · 上海</a-typography-text>
        </a-space>
      </a-col>
      <a-col :xs="24" :lg="9">
        <a-divider orientation="center">研究</a-divider>
        <a-link href="/paper/server_config" icon>
          配置服务器的操作的备忘（v2023.2)
        </a-link>
      </a-col>
    </a-row>
  </div>
</template>

<script>
export default {
  name: "Index",
  components: {},
  mounted() {},
  data() {
    return {};
  },
  methods: {},
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.wrapper {
  margin: 30px auto;
  padding: 10px;
  max-width: 980px;
}
.left {
  text-align: center;
}
.logo {
  background-color: #fff;
  margin-bottom: 30px;
}
</style>
