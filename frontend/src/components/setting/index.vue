<template>
  <el-container>
    <!--Setting Header-->
    <el-header class="setting-index-header" :style="settingHeader">
      <SettingHeader />
    </el-header>
    <!--Setting Main-->
    <el-main class="setting-index-main" :style="settingMain">
      <SettingMain />
    </el-main>
  </el-container>
</template>

<script>
import SettingHeader from "../setting/header.vue"
import SettingMain from  "../setting/main.vue"

export default {
  name: "index",
  components: { SettingHeader, SettingMain },
  data() {
    return {
      settingHeader: {
        width: "",
      },
      settingMain:{
        height: "",
        width: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.settingHeader.width = window.innerWidth - 40 - 2 + "px"
      this.settingHeader.width = window.innerWidth - 40 - 2 + "px"
      this.settingMain.height = window.innerHeight - 70 -60 - 4- 4 + "px";
    },
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },

}
</script>

<style scoped>
.setting-index-header{
  padding: 0;
  background: #2B2D30;
  height: 60px;
}
</style>