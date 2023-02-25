<template>
  <div id="aside">
    <el-row style="padding: 5px 5px 0 5px">
        <el-button
            style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
            color="#2B2D30"
            icon="Tools"
            size="large"
            @click="settingFunc" />
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Menu"
          size="large"
          @click="collectionFunc" />
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="TrendCharts"
          size="large"
          @click="analysisFunc" />
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Histogram"
          size="large"
          @click="historyFunc" />
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Help"
          size="large"
          @click="helpFunc" />

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
