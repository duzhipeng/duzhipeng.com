<template>
  <div class="content">
    <a-row justify="center" :gutter="[15, 15]">
      <!-- Logo -->
      <a-col flex="140px">
        <div class="card card-type-A">
          <a-avatar :size="120" shape="square">
            <img src="../assets/logo.png" alt="avatar" class="logo" />
          </a-avatar>
        </div>
      </a-col>
      <!-- Email -->
      <a-col flex="240px">
        <div class="card card-type-A">
          <a-space direction="vertical" :size="3">
            <a-typography-text class="title"> ✉️ </a-typography-text>
            <a-typography-text copyable>MAIL@DUZHIPENG.COM</a-typography-text>
            <a-link
              href="https://keys.openpgp.org/vks/v1/by-fingerprint/1C5E38AEB5BEBAFEB6BCC9D0F3C82540B55E0FEA"
              target="_blank"
            >
              <a-typography-text type="primary" class="text-size-small">
                ⏬ PGP PUBLIC KEY
              </a-typography-text>
            </a-link>
          </a-space>
        </div>
      </a-col>
      <!-- QQ -->
      <a-col flex="240px">
        <div class="card card-type-A">
          <a-space direction="vertical" :size="3">
            <a-typography-text class="title"> QQ </a-typography-text>
            <a-typography-text copyable>7171891</a-typography-text>
            <a-typography-text class="text-size-small">
              验证问题请输写 来自博客
            </a-typography-text>
          </a-space>
        </div>
      </a-col>
      <!-- WeChat -->
      <a-col flex="240px">
        <div class="card card-type-A">
          <a-space direction="vertical" :size="3">
            <a-typography-text class="title"> 💬 </a-typography-text>
            <a-typography-text copyable copy-text="BH4GDU">
              微信号：BH4GDU
            </a-typography-text>
            <a-typography-text class="text-size-small">
              验证问题请输写 来自博客
            </a-typography-text>
          </a-space>
        </div>
      </a-col>
      <!-- HAM -->
      <a-col flex="640px">
        <a-space direction="vertical" :size="15" fill>
          <div class="card card-type-A">
            <a-space :size="3" class="card-type-A">
              <a-avatar :size="120" shape="square">
                <img class="logo" src="../assets/CRAC-LOGO.svg" alt="avatar" />
              </a-avatar>
              <a-space direction="vertical" :size="3">
                <a-typography-text class="subtitle">HAM</a-typography-text>
                <a-typography-text>
                  🆔
                  <a-typography-text copyable>BH4GDU</a-typography-text>
                  <a-typography-text> | </a-typography-text>
                  <a-link href="https://www.qrz.com/db/bh4gdu" target="_blank"
                    >QRZ Home Page</a-link
                  >
                </a-typography-text>
                <a-typography-text class="text-size-small">
                  YAESU FT-818ND & 泉盛 UV-K6 ｜ 持八木打 🛰️｜ 常在7～10频道 |
                  <a-popover position="right">
                    <a-typography-text class="text-size-small"
                      >B类操作证 📃</a-typography-text
                    >
                    <template #content>
                      <img
                        class="image-large"
                        src="../assets/certificateImage.png"
                        alt="avatar"
                      />
                    </template>
                  </a-popover>
                </a-typography-text>
              </a-space>
            </a-space>
          </div>
          <div class="card">
            <a-typography-text class="title"> 🔖 </a-typography-text>
            <a-space direction="vertical" fill>
              <template #split>
                <a-divider type="dashed" :margin="3" />
              </template>
              <a-typography-text>
                <a-link href="/paper/hukou" icon> 上海落户记（2023.7) </a-link>
                <!--                - 成果-->
              </a-typography-text>
              <a-typography-text>
                <a-link href="/paper/server_config" icon>
                  配置服务器的操作的备忘（2023.2)
                </a-link>
                <!--                - 成果-->
              </a-typography-text>
            </a-space>
          </div>
        </a-space>
      </a-col>
      <!-- Schools -->
      <a-col flex="240px">
        <div class="card">
          <a-space direction="vertical" :size="20">
            <a-typography-text class="title"> 🏫 </a-typography-text>
            <a-popover>
              <img
                class="image-small logo"
                src="../assets/JIAOTONG-LOGO.png"
                alt="avatar"
              />

              <template #content>
                <img
                  class="image-big logo"
                  src="../assets/JIAOTONG-LOGO.png"
                  alt="avatar"
                />
              </template>
            </a-popover>
            <a-popover>
              <img
                class="image-small logo"
                src="../assets/EDUCATION-logo.jpg"
                alt="avatar"
              />

              <template #content>
                <img
                  class="image-big logo"
                  src="../assets/EDUCATION-logo.jpg"
                  alt="avatar"
                />
              </template>
            </a-popover>
            <a-popover>
              <img
                class="image-small logo"
                src="../assets/chebu_logo.jpg"
                alt="avatar"
              />
              <template #content>
                <img
                  class="image-big logo"
                  src="../assets/chebu_logo.jpg"
                  alt="avatar"
                />
              </template>
            </a-popover>
          </a-space>
        </div>
      </a-col>
    </a-row>
  </div>

  <div class="wrapper">
    <a-row :gutter="20" justify="center">
      <a-col :xs="24" :lg="15" class="left"> </a-col>
      <a-col :xs="24" :lg="9"> </a-col>
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
.content {
  margin: 30px auto;
  padding: 10px;
  max-width: 980px;
}
.card {
  background-color: #fff;
  border-radius: 8px;
  padding: 10px;
}
.card-type-A {
  height: 120px;
}
.card-type-B {
  height: 275px;
}
.logo {
  background-color: #fff;
}
.title {
  font-size: 45px;
}
.subtitle {
  font-size: 35px;
}
.image-small {
  max-width: 150px;
  max-height: 50px;
}
.image-big {
  max-width: 350px;
  max-height: 150px;
}
.image-large {
  max-width: 550px;
}
</style>
