<template>
  <div>
    <el-row class="setting-header-button">
      <el-button icon="RefreshRight"  color="#2B2D30">默认恢复</el-button>
      <el-button icon="Check" bg color="#2B2D30">保存参数</el-button>
      <el-button icon="Close" color="#2B2D30">放弃退出</el-button>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "header"
}
</script>

<style scoped>
.setting-header-button{
  padding: 15px 0px 15px 0px;

}
</style>