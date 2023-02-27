<template>
  <div>
<!--    <el-row>-->
<!--      <el-col :span="8"><div class="grid-content ep-bg-purple" />-->
<!--        <el-space >-->
<!--          <el-card  class="box-card" style="width: 250px">-->
<!--            <div v-for="o in 4" :key="o" class="text item">-->
<!--              {{ 'List item ' + o }}-->
<!--            </div>-->
<!--          </el-card>-->
<!--        </el-space>-->
<!--      </el-col>-->
<!--      <el-col :span="16"><div class="grid-content ep-bg-purple-light" />-->
<!--        <el-card shadow="never" style="width: 800px; height: 200px">-->
<!--          <el-table class="history-main-table" :data="tableData" border  style="width: 100%" >-->
<!--            <el-table-column prop="通道名称" label="通道名称" width="140" />-->
<!--            <el-table-column prop="通电状态" label="供电通电状态设置" width="140" />-->
<!--            <el-table-column prop="电压阈值" label="电压阈值设置(V)" width="140" />-->
<!--            <el-table-column prop="电流阈值" label="电流阈值设置(V)" width="140" />-->
<!--            <el-table-column prop="探管编号" label="探管编号" width="140" />-->

<!--          </el-table>-->
<!--        </el-card>-->

<!--      </el-col>-->
<!--    </el-row>-->
  </div>
</template>

<script>
export default {
  name: "header",
  data() {
    return {
      historyTable: {
        width: "",
        height: "",
        color: "#c8c9cc",
      },
      tableData: [
        {
          通道名称: '1',
          通电状态: '2016-05-03',
          电压阈值: '09:10:50',
          电流阈值: '30',
          探管编号: '8',
        },
        {
          通道名称: '1',
          通电状态: '2016-05-03',
          电压阈值: '09:10:50',
          电流阈值: '30',
          探管编号: '8',
        },
        {
          通道名称: '1',
          通电状态: '2016-05-03',
          电压阈值: '09:10:50',
          电流阈值: '30',
          探管编号: '8',
        },
        {
          通道名称: '1',
          通电状态: '2016-05-03',
          电压阈值: '09:10:50',
          电流阈值: '30',
          探管编号: '8',
        },
        {
          通道名称: '1',
          通电状态: '2016-05-03',
          电压阈值: '09:10:50',
          电流阈值: '30',
          探管编号: '8',
        },

      ]
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.historyTable.width = window.innerWidth - 40 - 2 + "px"
      this.historyTable.height = window.innerHeight - 70 - 2 - 60 - 250 - 2 - 2 + "px"
    },

    created() {
      //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
      window.addEventListener("resize", this.getHeight);
      this.getHeight();
    },
  }
}
</script>

<style scoped>
/*背景透明*/
::v-deep .el-table__expanded-cell {
  background-color: transparent !important;
}

::v-deep .el-table th,
::v-deep .el-table tr,
::v-deep .el-table td {
  background-color: transparent;
}

</style>