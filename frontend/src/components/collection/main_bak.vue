<template>
  <div id="collection-main">
    <div class="collection-echarts" ref="collection_echarts_data" :style="collectionEchartsSize"> </div>
    <div>
      <el-row>
        ss
      </el-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, onUnmounted} from 'vue'
import * as echarts from 'echarts';


let collectionEchartsSize:any =  {width: "", height: "",}
let width:any = ""
let height:any = ""

const data = new Array(1, 4, 9);
const collection_echarts_data = ref()
let myChart: echarts.ECharts;
let colors = ['#c93756','#177cb0','#16a951','#8d4bbb','#ffb61e','#1baeae'];//自定义一个颜色数组，多系时会按照顺序使用自己定义的颜色数组，若不定义则使用默认的颜色数组

onMounted(() => {
  initEcharts(data)
  window.addEventListener("resize", ()=> myChart.resize());
  updateEcharts()
  getWindowSize()
})

const initEcharts = (es_data:any) => {
  myChart = echarts.init(collection_echarts_data.value);

  const option = {
    color : colors,
    tooltip: {
      trigger: "axis",
      axisPointer: {
        type: 'line',
        lineStyle: {
          color: '#16a951'
        }
      }
    },
    legend: {
      show: true,
      textStyle: {
        color: '#eeeeee',
      },
      data: ["温度", "供电电压", "供电电流", "计数", "计数率", "电源开关"],
    },
    grid: {
      left: "3%",
      right: "4%",
      bottom: "3%",
      containLabel: true,
    },
    xAxis: {
      type: "category",
      boundaryGap: false,
        axisTick: {
          show: false
        }
    },
    yAxis:[
      {
        //供电电流 第一个y轴位置在左侧
        position:'left',
        offset: 0,
        type : 'value',
        // 名字
        name: '供电电流(mA)',
        nameLocation: 'middle',
        nameGap: 23,
        min: 0, // 坐标Y轴刻度最小值
        max: 800, // 坐标Y轴刻度最大值
        interval: 50, // 坐标Y轴刻度间隔
        axisLine: { //坐标线
          show: true ,
          symbol: "pin",
          lineStyle: {
            color: colors[0],
            width: '0.5',
          }
        },
        axisTick: { // 坐标刻度
          show: true,
        },
        axisLabel: {
          show: true,
          fontFamily: "Arial",
          verticalAlign: "middle",
          formatter: '{value}'
        },
        splitLine: {
          show: false
        },
        axisPointer: {
          show: true,
          type: 'line',
          lineStyle: {
            color: '#16a951'
          }
        }
      },
      {
        //供电电压 第一个y轴位置在左侧
        position:'left',
        offset: 45,
        type : 'value',
        // 名字
        name: '供电电压(V)',
        nameLocation: 'middle',
        nameGap: 22,
        min: -10, // 坐标Y轴刻度最小值
        max: 30, // 坐标Y轴刻度最大值
        interval: 2, // 坐标Y轴刻度间隔
        axisLine: { //坐标线
          show: true ,
          symbol: "pin",
          lineStyle: {
            color: colors[1],
            width: '0.5',
          }
        },
        axisTick: { // 坐标刻度
          show: true,
        },
        axisLabel: {
          show: true,
          fontFamily: "Arial",
          verticalAlign: "middle",
          formatter: '{value}'
        },
        splitLine: {
          show: false
        },
        axisPointer: {
          show: true,
          type: 'line',
          lineStyle: {
            color: '#16a951'
          }
        }
      },
      {
        //计数率 第一个y轴位置在左侧
        position:'left',
        offset: 90,
        type : 'value',
        // 名字
        name: '采集盒温度(℃)',
        nameLocation: 'middle',
        nameGap: 23,
        min: 0, // 坐标Y轴刻度最小值
        max: 100, // 坐标Y轴刻度最大值
        interval: 5, // 坐标Y轴刻度间隔
        axisLine: { //坐标线
          show: true ,
          symbol: "pin",
          lineStyle: {
            color: colors[2],
            width: '0.5',
          }
        },
        axisTick: { // 坐标刻度
          show: true,
        },
        axisLabel: {
          show: true,
          fontFamily: "Arial",
          verticalAlign: "middle",
          formatter: '{value}'
        },
        splitLine: {
          show: false
        },
        axisPointer: {
          show: true,
          type: 'line',
          lineStyle: {
            color: '#16a951'
          }
        }
      },
      {
        //计数 第一个y轴位置在左侧
        position:'right',
        offset: 0,
        type : 'value',
        // 名字
        name: '计数(次)',
        nameLocation: 'middle',
        nameGap: 28,
        min: 0.9, // 坐标Y轴刻度最小值
        max: 1.1, // 坐标Y轴刻度最大值
        interval: 0.02, // 坐标Y轴刻度间隔
        axisLine: { //坐标线
          show: true ,
          symbol: "pin",
          lineStyle: {
            color: colors[3],
            width: '0.5',
          }
        },
        axisTick: { // 坐标刻度
          show: true,
        },
        axisLabel: {
          show: true,
          fontFamily: "Arial",
          verticalAlign: "middle",
          formatter: '{value}'
        },
        splitLine: {
          show: false
        },
        axisPointer: {
          show: true,
          type: 'line',
          lineStyle: {
            color: '#16a951'
          }
        }
      },
      {
        //计数率 第一个y轴位置在左侧
        position:'right',
        offset: 45,
        type : 'value',
        // 名字
        name: '计数率(%)',
        nameLocation: 'middle',
        nameGap: 28,
        min: 0, // 坐标Y轴刻度最小值
        max: 100, // 坐标Y轴刻度最大值
        interval: 5, // 坐标Y轴刻度间隔
        axisLine: { //坐标线
          show: true ,
          symbol: "pin",
          lineStyle: {
            color: colors[4],
            width: '0.5',
          }
        },
        axisTick: { // 坐标刻度
          show: true,
        },
        axisLabel: {
          show: true,
          fontFamily: "Arial",
          verticalAlign: "middle",
          formatter: '{value}'
        },
        splitLine: {
          show: false
        },
        axisPointer: {
          show: true,
          type: 'line',
          lineStyle: {
            color: '#16a951'
          }
        }
      },
      {
        //计数率 第一个y轴位置在左侧
        position:'right',
        offset: 90,
        type : 'value',
        name: '供电开关(%)',
        nameLocation: 'middle',
        nameGap: 32,
        min: -2, // 坐标Y轴刻度最小值
        max: 2, // 坐标Y轴刻度最大值
        interval: 0.2, // 坐标Y轴刻度间隔
        axisLine: { //坐标线
          show: true ,
          symbol: "pin",
          lineStyle: {
            color: colors[5],
            width: '0.5',
          }
        },
        axisTick: { // 坐标刻度
          show: true,
        },
        axisLabel: {
          show: true,
          fontFamily: "Arial",
          verticalAlign: "middle",
          formatter: '{value}'
        },
        splitLine: {
          show: false
        },
        axisPointer: {
          show: true,
          type: 'line',
          lineStyle: {
            color: '#16a951'
          }
        }
      },

    ],
    series: [
      {
        name: "供电电流",
        type: "line",
        stack: "Total",
        showSymbol: false,
        lineStyle: {
          width: 2, // 设置线宽
          color: colors[0],
        },
        yAxisIndex: 0,//使用第一个y轴，序号从0开始
        data: [631,633,435,539,631,733,634,433,336,631],
      },
      {
        name: "供电电压",
        type: "line",
        stack: "Total",
        showSymbol: false,
        lineStyle: {
          width: 2, // 设置线宽
          color: colors[1],
        },
        yAxisIndex: 1,//使用第一个y轴，序号从0开始
        data: [1,13,25,19,11,13,14,23,16,21],
      },
      {
        name: "采集盒温度",
        type: "line",
        stack: "Total",
        showSymbol: false,
        lineStyle: {
          width: 2, // 设置线宽
          color: colors[2],
        },
        yAxisIndex: 2,//使用第一个y轴，序号从0开始
        data: [21,33,45,39,51,33,44,53,66,31],
      },
      {
        name: "计数",
        type: "line",
        stack: "Total",
        showSymbol: false,
        lineStyle: {
          width: 2, // 设置线宽
          color: colors[3],
        },
        yAxisIndex: 3,//使用第一个y轴，序号从0开始
        data: [31,33,35,39,31,33,34,33,36,31],
      },
      {
        name: "计数率",
        type: "line",
        stack: "Total",
        showSymbol: false,
        lineStyle: {
          width: 2, // 设置线宽
          color: colors[4],
        },
        yAxisIndex: 4,//使用第一个y轴，序号从0开始
        data: [31,33,35,39,31,33,34,33,36,31],
      },
      {
        name: "电源开关",
        type: "line",
        stack: "Total",
        showSymbol: false,
        lineStyle: {
          width: 2, // 设置线宽
          color: colors[5],
        },
        yAxisIndex: 5,//使用第一个y轴，序号从0开始
        data: [1,1,1,0,0,0,1,1,0,0],
      },

    ],
  };


  myChart.setOption(option)
}
function updateEcharts() {
  myChart.resize()//重新渲染
}

//定义方法，获取高度减去头尾
function getWindowSize() {
  collectionEchartsSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2+ "px"
  collectionEchartsSize.height = window.innerHeight - 70 - 2 -60 - 250 - 2 - 40 + "px"
  width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2+ "px"
  height = window.innerHeight - 70 - 2 -60 - 250 - 2 - 40 + "px"
  console.log("echarts-size:", collectionEchartsSize)
}

// // ========== websocket 连接 ==========
// const ws = new WebSocket("ws://127.0.0.1:8090/ws/v1/index");
// //连接成功后触发
// ws.onopen = (e) => {
//   console.log('已连接');
// };
// //接收消息时触发
// ws.onmessage = (msg) => {
//   const temp = JSON.parse(msg.data) as number
//
//   if (data.length > 200 ){
//     data.splice(0,1)
//   }
//   data.push(temp)
//   // //触发图表更新
//   initEcharts(data)
//   console.log("实时数据2： ",  data)
// };
// //连接断开时触发
// ws.onclose = (event)=> {
//   console.log("websocket connection close.");
//   console.log(event.code);
// };
//
// //连接出错时触发
// ws.onerror = (event)=> {
//   console.log("websocket connection error.");
//   console.log(event);
// };
//
// onUnmounted(()=>{
//   ws.close()
//   console.log("websocket已关闭");
// })

</script>

<style scoped>
#collection-main {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #eeeeee;
  margin-top: 0;
}
</style>

<style scoped lang='scss'>
/*图表所需的盒子（必须设置dom大小）*/
.collection-echarts {
  /*width: 1500px;*/
  /*height: 600px;*/
}
</style>