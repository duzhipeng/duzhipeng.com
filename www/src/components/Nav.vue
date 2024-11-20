<template>
  <div class="header">
    <a-space :size="30">
      <!--Logo-->
      <router-link to="/">
        <img class="logo" src="../assets/logo.png" alt="阿杜" />
      </router-link>
      <a-divider direction="vertical" />
      <!--Menu-->
      <a-link href="/garden">菜园</a-link>
      <a-link href="/shop">小店</a-link>
      <a-link href="/order">订单</a-link>
      <!--Status-->
      <a-divider direction="vertical" />
      <a-badge
        :status="system.status[system.data.monitor]"
        :text="
          'GATLING ' +
          system.data.monitor +
          ' · ' +
          dateTime(system.sync_time, 'YYYY-MM-DD HH:mm:ss')
        "
      />
    </a-space>
  </div>
</template>

<script>
import api from "../api";
import moment from "moment";

export default {
  name: "Nav",
  components: {},
  created() {
    this.checkStatus(api.status);
  },
  mounted() {
    this.monitor();
  },
  computed() {},
  data() {
    return {
      system: {
        status: {
          CONNECTING: "warning",
          OFFLINE: "normal",
          ONLINE: "success",
          ERROR: "danger",
        },
        data: {
          monitor: "CONNECTING",
        },
        sync_time: "",
      },
    };
  },
  methods: {
    dateTime(value, format) {
      if (value === "") {
        return "";
      }
      return moment(value).format(format);
    },
    monitor() {
      self.setInterval(() => {
        self.setTimeout(() => {
          this.checkStatus(api.status);
        }, 0);
      }, 10000);
    },
    checkStatus(api) {
      this.$http
        .get(api)
        .then((response) => {
          // handle success
          if (response.status === 200) {
            this.$data.system.data = response.data["data"];
            this.$data.system.sync_time = new Date();
          } else {
            this.$message.error("系统服务异常。 ${response.data}");
          }
        })
        .catch((error) => {
          // handle error
          this.$message.error("系统服务中断。本地网络不通。");
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
.header {
  background-color: #fff;
  padding: 5px 10px;
}

.logo {
  width: 30px;
}
</style>
