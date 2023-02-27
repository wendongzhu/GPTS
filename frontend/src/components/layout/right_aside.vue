<template>
  <div id="aside">
    <el-row style="padding: 5px 5px 0 5px">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Bell"
          size="large"
          @click="collectionFunc" />
    </el-row>
    <el-row style="padding: 5px 5px 0 5px">
      <el-button
          style="width: 30px; height: 30px; margin: 0; padding: 0;border-radius: 5px"
          color="#2B2D30"
          icon="Coin"
          size="large"
          @click="collectionFunc" />
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
    // settingFunc(){
    //   this.$router.push("/setting");
    // },
    // collectionFunc(){
    //   this.$router.push("/collection");
    // },
    // analysisFunc(){
    //   this.$router.push("/analysis");
    // },
    // historyFunc(){
    //   this.$router.push("/history");
    // },
    // helpFunc(){
    //   this.$router.push("/help");
    // },
  }
}

</script>


<style scoped>
.el-button{
  /*width: 90px;*/
  /*height: 60px;*/
  /*padding: 0;*/
  /*/deep/ i {*/
  /*  icon-size: 20px;*/
  /*}*/

}

</style>
