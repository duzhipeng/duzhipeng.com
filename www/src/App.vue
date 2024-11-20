<template>
  <Notice />
  <div class="wrapper"><router-view /></div>
  <Footer />
</template>

<script>
// @ is an alias to /src
import Notice from "../src/components/Notice.vue";
import Footer from "../src/components/Footer.vue";

export default {
  components: {
    Notice,
    Footer,
  },
};
</script>

<style scoped>
.wrapper {
  padding: 10px;
}
</style>
