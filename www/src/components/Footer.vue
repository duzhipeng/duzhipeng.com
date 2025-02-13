<template>
  <div class="content">
    <a-space direction="vertical" :size="10">
      <a-typography-text type="secondary">
        &copy; 2024 阿杜 &reg; 保留所有权利
      </a-typography-text>
      <span>
        <a-link href="https://beian.miit.gov.cn/" target="_blank">
          <a-typography-text type="secondary" class="text-size-small"
            >沪ICP备2023009730号-1
          </a-typography-text></a-link
        >
        <a-divider direction="vertical" />
        <a-link
          href="https://www.beian.gov.cn/portal/registerSystemInfo?recordcode=31011502019687"
          target="_blank"
        >
          <a-typography-text type="secondary" class="text-size-small">
            <img
              class="police-logo"
              src="/gaba.png"
              alt="沪公网安备 31011502019687号"
            />
            沪公网安备 31011502019687号
          </a-typography-text></a-link
        >
      </span>
    </a-space>
  </div>
</template>

<script>
export default {
  name: "Footer",
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
  margin: 60px auto;
  max-width: 980px;
  text-align: center;
}
.police-logo {
  width: 13px;
  vertical-align: text-top;
}
</style>
