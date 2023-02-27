<template>
  <div id="aside">
    <el-row style="padding: 5px 5px 0 5px">

      <el-tooltip class="item" effect="dark" content="数据采集" placement="right-start">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Menu"
          size="large"
          @click="collectionFunc" />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-tooltip class="item" effect="dark" content="数据分析" placement="right-start">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="TrendCharts"
          size="large"
          @click="analysisFunc" />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-tooltip class="item" effect="dark" content="历史数据" placement="right-start">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Histogram"
          size="large"
          @click="historyFunc" />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-tooltip class="item" effect="dark" content="系统设置" placement="right-start">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Tools"
          size="large"
          @click="settingFunc" />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-tooltip class="item" effect="dark" content="帮助信息" placement="right-start">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="HelpFilled"
          size="large"
          @click="helpFunc" />
      </el-tooltip>
    </el-row>
  </div>

</template>

<script>
export default{
  name: "Aside",
  data() {
    return {

    };
  },
  methods: {
    settingFunc(){
      this.$router.push("/setting");
    },
    collectionFunc(){
      this.$router.push("/collection");
    },
    analysisFunc(){
      this.$router.push("/analysis");
    },
    historyFunc(){
      this.$router.push("/history");
    },
    helpFunc(){
      this.$router.push("/help");
    },
  }
}

</script>


<style scoped>


</style>
