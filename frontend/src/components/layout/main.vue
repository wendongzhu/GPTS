<template>
  <el-container>
    <el-aside class="layout-index-aside" :style="defaultHeight" style="--wails-draggable:drag">
      <LayoutLeftAside/>
    </el-aside>

    <!--Layout Main-->
    <el-main class="layout-index-home" :style="defaultHeight">
      <router-view></router-view>
    </el-main>

    <el-aside class="layout-index-aside" :style="defaultHeight" style="--wails-draggable:drag">
      <LayoutRightAside/>
    </el-aside>
  </el-container>
</template>

<script>
import LayoutLeftAside from "./left_aside.vue";
import LayoutRightAside from "./right_aside.vue"
import LayoutHome from "../home/index.vue";

export default {
  name: "main",
  components: { LayoutLeftAside, LayoutHome, LayoutRightAside },
  data() {
    return {
      defaultHeight: {
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 70 - 4 + "px";
      console.log(this.defaultHeight.height);
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
.layout-index-aside {
  width: 40px;
  background: #2B2D30;
}

.layout-index-home {
  padding: 0 0 0 2px;
  background: #1E1F22;
}
</style>