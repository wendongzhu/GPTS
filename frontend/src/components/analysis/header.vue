<template>
  <div>
    <el-row >
<!--    <div class="collection-index-header-txt" :style="defaultHeaderWidth">-->
<!--      &lt;!&ndash;中间预留&ndash;&gt;-->
<!--    </div>-->
    <div class="history-header-button">
        <el-button-group >
          <el-button icon="Connection"  color="#2B2D30">输入振荡数据</el-button>
          <el-button icon="Postcard" bg color="#2B2D30">输出融合结果</el-button>
        </el-button-group>

      <font size="3" font-weight="bold" style="padding-left: 700px">测试人员：XXX</font>
      <font size="3" font-weight="bold" style="padding-left: 100px">测试地点：XXX</font>
      <font size="3" font-weight="bold" style="padding-left: 100px">测试单位：XXX</font>

    </div>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "header",
  data() {
    return {
      defaultHeaderWidth: {
        width: "",
      },
    };
  },


  methods: {
    getFooterWidth() {
      this.defaultHeaderWidth.width = window.innerWidth - 40 - 2 - 40 - 2 - 450 + "px";
    },
  },
  created() {
    window.addEventListener("resize", this.getFooterWidth);
    this.getFooterWidth();
  },
}
</script>

<style scoped>
.history-header-button{
  padding: 15px 0 15px 0;

}
</style>