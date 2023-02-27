<template>
  <el-container>
    <!--Setting Header-->
    <el-header class="history-index-header" :style="historyWidth">
      <HistoryHeader />
    </el-header>
    <!--Setting Main-->
    <el-main class="history-index-main" :style="historyMain">
      <HistoryMain />
    </el-main>
    <!--Setting Main-->
    <el-footer class="history-index-footer" :style="historyWidth">
      <HistoryFooter />
    </el-footer>
  </el-container>
</template>

<script>
import HistoryHeader from "../history/header.vue"
import HistoryMain from  "../history/main.vue"
import HistoryFooter from "../history/footer.vue"

export default {
  name: "index",
  components: { HistoryHeader, HistoryMain, HistoryFooter },
  data() {
    return {
      historyWidth: {
        width: "",
      },
      historyMain: {
        width: "",
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.historyWidth.width = window.innerWidth - 40 - 2 - 40 -2 + "px"
      this.historyMain.width = window.innerWidth - 40 - 2 - 40 -2 + "px"
      this.historyMain.height = window.innerHeight - 70 - 2 -60 - 250 - 2 + "px"
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
.history-index-header{
  padding: 0;
  background: #2B2D30;
  height: 60px;
}

.history-index-footer{
  padding: 0;
  background: #2B2D30;
  height: 250px;
}
.history-index-main{
  padding: 1px 1px 1px 1px;
  background: #1E1F22;
}

</style>