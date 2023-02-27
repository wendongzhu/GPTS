<template>
<div class="history-main" >
  <el-table class="history-main-table" :data="tableData" border  :style="historyTable" >
    <el-table-column prop="序号" label="序号" width="60" />
    <el-table-column prop="实验日期" label="实验日期" width="110" />
    <el-table-column prop="实验时间" label="实验时间" width="110" />
    <el-table-column prop="采样时长" label="采样时长" width="110" />
    <el-table-column prop="通道数据" label="通道数据" width="110" />
    <el-table-column prop="采样文件" label="采样文件" width="180" />
    <el-table-column prop="是否融合" label="是否融合" width="110" />
    <el-table-column prop="振荡文件" label="振荡文件" width="180" />
    <el-table-column prop="融合文件" label="融合文件" width="180" />
    <el-table-column prop="测试人员" label="测试人员" width="110" />
    <el-table-column prop="测试单位" label="测试单位" />
    <el-table-column prop="测试地点" label="测试地点" />
    <el-table-column prop="测试概要" label="测试概要" />

  </el-table>
</div>
</template>


<script>
export default {
  name: "main",
  data() {
    return {
      historyTable: {
        width: "",
        height: "",
        background: "#1E1F22",
        color: "#c8c9cc",
      },
      tableData: [
        {
          序号: '1',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xlsx',
          是否融合: '是',
          振荡文件: 'yyy_20220226.xlsx',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
        {
          序号: '2',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xlsx',
          是否融合: '是',
          振荡文件: 'yyy_20220226.xlsx',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
        {
          序号: '3',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xlsx',
          是否融合: '是',
          振荡文件: 'yyy_20220226.xlsx',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
        {
          序号: '4',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xlsx',
          是否融合: '是',
          振荡文件: 'yyy_20220226.xlsx',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
        {
          序号: '5',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xlsx',
          是否融合: '是',
          振荡文件: 'yyy_20220226.xlsx',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
        {
          序号: '6',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xlsx',
          是否融合: '是',
          振荡文件: 'yyy_20220226.xlsx',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
        {
          序号: '7',
          实验日期: '2016-05-03',
          实验时间: '09:10:50',
          采样时长: '30',
          通道数据: '8',
          采样文件: 'xxx_20220226.xal',
          是否融合: '否',
          振荡文件: 'yyy_20220226.xal',
          融合文件: 'xxx_xxx',
          测试人员: 'gpt',
          测试地点: '北京',
          测试单位: '北京',
          测试概要: 'No. 189, ',
        },
      ]
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getHeight() {
      this.historyTable.width = window.innerWidth - 40 - 2 + "px"
      this.historyTable.height = window.innerHeight - 70 - 2 -60 - 250 - 2 - 2 + "px"
    },

    //根据情况显示背景色
    cellStyle(row){
      if(row.column.是否融合==="否"){
        return 'background:#2D5AB9'
      }else if(row.column.是否融合==="是"){
        return 'background:#974356'
      }
    }
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },

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