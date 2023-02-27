<template>
  <div>
    <el-row >
      <div class="collection-index-header-txt" :style="defaultHeaderWidth">
        <!--中间预留-->
      </div>

      <el-button-group class="history-header-button">
        <el-button icon="Search"  color="#2B2D30">筛选数据</el-button>
        <el-button icon="View" bg color="#2B2D30">查看数据</el-button>
        <el-button icon="Delete" color="#2B2D30">删除数据</el-button>
        <el-button icon="Collection" color="#2B2D30">导出数据</el-button>
      </el-button-group>

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