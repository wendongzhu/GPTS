<template>
  <el-container>
    <!--Collection Header-->
    <el-header class="analysis-index-header" :style="analysisHeaderSize">
      <AnalysisHeader />
    </el-header>
    <el-container>
      <el-container>
        <!--Collection Main-->
        <el-main class="analysis-index-main" :style="analysisMainSize">
          <AnalysisMain />
        </el-main>
        <!--Collection Footer-->
        <el-footer class="analysis-index-footer" :style="analysisFooterSize">
          <AnalysisFooter />
        </el-footer>
      </el-container>
      <!--Collection Right Aside-->
      <el-aside class="analysis-index-right-aside" :style="analysisRightAsideSize">
        <AnalysisRightAside />
      </el-aside>
    </el-container>

  </el-container>
</template>

<script>
import AnalysisHeader from "./header.vue"
import AnalysisMain from  "./main.vue"
import AnalysisFooter from "./footer.vue"
import AnalysisRightAside from "./right_aside.vue"

export default {
  name: "index",
  components: { AnalysisHeader, AnalysisMain, AnalysisFooter, AnalysisRightAside },
  data() {
    return {
      analysisHeaderSize: {
        width: "",
      },
      analysisFooterSize: {
        width: "",
      },
      analysisMainSize: {
        width: "",
        height: "",
      },
      analysisRightAsideSize: {
        width: "",
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.analysisHeaderSize.width = window.innerWidth - 40 - 2 -40 -2+ "px"
      this.analysisFooterSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2+ "px"
      this.analysisMainSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 -2 + "px"
      this.analysisMainSize.height = window.innerHeight - 70 - 2 -60 - 250 - 2 + "px"
      this.analysisRightAsideSize.height = window.innerHeight - 70 - 2 -60  - 2  -2 + "px"
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
.analysis-index-header{
  padding: 0;
  background: #2B2D30;
  height: 60px;
}

.analysis-index-footer{
  padding: 0;
  background: #2B2D30;
  height: 250px;
}
.analysis-index-main{
  padding: 1px 1px 1px 1px;
  background: #1E1F22;
}
.analysis-index-right-aside{
  width: 180px;
  margin-top: 2px;
  margin-right: 2px;
  background: #2B2D30;

}


</style>