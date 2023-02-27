<template>
    <el-container>
      <!--Layout Header-->
      <el-header class="layout-index-header"  style="--wails-draggable:drag">
        <LayoutHeader />
      </el-header>

      <!--Main-->
      <!--Layout Main-->
      <el-main class="layout-index-main" :style="defaultHeight">
        <LayoutMain />
      </el-main>

      <!--Layout Footer-->
      <el-footer class="layout-index-footer" style="--wails-draggable:drag">
        <LayoutFooter />
      </el-footer>
    </el-container>


</template>

<script>
import LayoutHeader from "./header.vue";
import LayoutFooter from "./footer.vue";
import LayoutMain from "./main.vue";

export default {
  name: "Layout",
  components: { LayoutHeader, LayoutMain, LayoutFooter},
  data() {
    return {
      isCollapse: false,
      defaultHeight: {
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 70 + "px";
      console.log(this.defaultHeight.height);
    },
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },
};
</script>

<style lang="scss" scoped>
.layout-index-header{
  padding: 0;
  background: #2B2D30;
  height: 40px;
}

.layout-index-footer {
  background: #2B2D30;
  height: 30px;
  border-bottom: none;
}

.layout-index-main {
  padding: 2px 0;
  background: #1E1F22;
  color: #c8c9cc;
}

</style>
