<template>
  <div id="collection-main">
<!--    <div class="collection-echarts" id="collectionEcharts" :style="collectionEchartsSize"> </div>-->
    <div id="collectionEcharts" :style="collectionEchartsSize"></div>
    <div>
      <el-row>
        ss
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      collectionEchartsSize: {
        width: "",
        height: "",
      },
      collection_data: {
        current_list: [631,633,435,539,631,733,634,433,336,631,633,435,539,631,733,634,433,336,631,633,435,539,631,733,634,433,336,631],
        voltage_list: [1,13,25,19,11,13,14,23,16,21,13,25,19,11,13,14,23,16,21,13,25,19,11,13,14,23,16,21],
        temperature_list: [21,33,45,39,51,33,44,53,66,31,33,45,39,51,33,44,53,66,31,33,45,39,51,33,44,53,66,31],
        count_list: [0.91,1.01,1.03,0.93,0.98,1.06,1.02,0.96,0.93,1.05,1.01,1.03,0.93,0.98,1.06,1.02,0.96,0.93,1.05,1.01,1.03,0.93,0.98,1.06,1.02,0.96,0.93,1.05],
        count_rate_list: [91,93,95,99,91,93,94,93,96,91,93,95,99,91,93,94,93,96,91,93,95,99,91,93,94,93,96,91],
        switch_list: [1,1,1,0,0,0,1,1,0,0,1,1,0,0,0,1,1,0,0,1,1,0,0,0,1,1,0,0],
      },
      colors: ['#c93756','#177cb0','#16a951','#8d4bbb','#ffb61e','#1baeae'],
    };
  },
  mounted() {
    this.myEcharts(this.collection_data);
    this.getWindowSize();
  },
  methods: {
    myEcharts(es_data) {
      const myChart = this.$echarts.init(document.getElementById("collectionEcharts"));

      const option = {
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: 'line',
            snap: false,
            lineStyle: {
              color: '#16a951'
            },
          },
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        legend: {
          show: true,
          textStyle: {
            color: '#eeeeee',
          },
          data: ["供电电流", "供电电压", "采集盒温度", "计数", "计数率", "电源开关"],
        },
        xAxis: {
          name: "时间(ms)",
          nameLocation: 'middle',
          type: "category",
          axisTick: {
            show: false
          },
          nameTextStyle: {
            fontSize: 12,
            color: '#eeeeee',
          },
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
                color: this.colors[0],
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
                color: this.colors[1],
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
                color: this.colors[2],
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
                color: this.colors[3],
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
                color: this.colors[4],
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
                color: this.colors[5],
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
            showSymbol: false,
            itemStyle: {
              normal: {
                color: this.colors[0],
                lineStyle: {
                  width: 2, // 设置线宽
                  color: this.colors[0]
                }
              }
            },
            yAxisIndex: 0,//使用第一个y轴，序号从0开始
            data: es_data.current_list,
          },
          {
            name: "供电电压",
            type: "line",
            showSymbol: false,
            itemStyle: {
              normal: {
                color: this.colors[1],
                lineStyle: {
                  width: 2, // 设置线宽
                  color: this.colors[1]
                }
              }
            },
            yAxisIndex: 1,//使用第一个y轴，序号从0开始
            data: es_data.voltage_list,
          },
          {
            name: "采集盒温度",
            type: "line",
            showSymbol: false,
            itemStyle: {
              normal: {
                color: this.colors[2],
                lineStyle: {
                  width: 2, // 设置线宽
                  color: this.colors[2]
                }
              }
            },
            yAxisIndex: 2,//使用第一个y轴，序号从0开始
            data: es_data.temperature_list,
          },
          {
            name: "计数",
            type: "line",
            showSymbol: false,
            itemStyle: {
              normal: {
                color: this.colors[3],
                lineStyle: {
                  width: 2, // 设置线宽
                  color: this.colors[3]
                }
              }
            },
            yAxisIndex: 3,//使用第一个y轴，序号从0开始
            data: es_data.count_list,
          },
          {
            name: "计数率",
            type: "line",
            showSymbol: false,
            itemStyle: {
              normal: {
                color: this.colors[4],
                lineStyle: {
                  width: 2, // 设置线宽
                  color: this.colors[4]
                }
              }
            },
            yAxisIndex: 4,//使用第一个y轴，序号从0开始
            data: es_data.count_rate_list,
          },
          {
            name: "电源开关",
            type: "line",
            showSymbol: false,
            itemStyle: {
              normal: {
                color: this.colors[5],
                lineStyle: {
                  width: 2, // 设置线宽
                  color: this.colors[5]
                }
              }
            },
            yAxisIndex: 5,//使用第一个y轴，序号从0开始
            data: es_data.switch_list,
          },
        ],
      };

      // 使用刚指定的配置项和数据显示图表。
      myChart.setOption(option);
    },

    getWindowSize() {
      this.collectionEchartsSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2 - 2 + "px"
      this.collectionEchartsSize.height = window.innerHeight - 70 - 2 -60 - 250 - 2 - 40 -2 + "px"
      console.log("echarts-size:", this.collectionEchartsSize)
    },
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getWindowSize);
    this.getWindowSize();
  },
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