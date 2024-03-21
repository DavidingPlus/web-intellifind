<script setup>
import { UploadFilled } from '@element-plus/icons-vue';
import * as echarts from 'echarts';
import { onMounted, ref } from 'vue';

// 通过echarts 中的折线图进行历次数据展示
const chartRef = ref(null)
const chartData = ref([])
onMounted(() => {
  if (chartRef.value) {
  const chart = echarts.init(chartRef.value)
  const option = {
    title: {
      text: '最近七次解析反馈数据'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['JSON1', 'JSON2', 'JSON3', 'JSON4', 'JSON5']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '0%',
      containLabel: true
    },
    toolbox: {
      feature: {
        saveAsImage: {}
      }
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: ['第一次解析', '第二次解析', '第三次解析', '第四次解析', '第五次解析', '第六次解析', '第七次解析']
    },
    yAxis: {
      type: 'value',
      // min: Math.min(...chartData.value.map(item => Math.min(...item.data))),  
      // max: Math.max(...chartData.value.map(item => Math.max(...item.data))),  
    },
    // 通过循环的方式，将历次数据进行展示
    series: chartData.value.map(item => ({
      name: item.name,
      data: item.data,
      type: 'line',
    }))
  }
chart.setOption(option)}})

const fetchData = () => {
  chartData.value = [{
      name: 'JSON1',
      data: [10, 42, 11, 14, 90, 23, 51]
    },
    {
      name: 'JSON2',
      data: [42, 8, 91, 60, 29, 33, 10]
    },
    {
      name: 'JSON3',
      data: [15, 32, 20, 54, 87, 60, 61]
    },
    {
      name: 'JSON4',
      data: [32, 13, 1, 34, 39, 42, 20]
    },
    {
      name: 'JSON5',
      data: [82, 23, 10, 24, 29, 36, 40]
    }]
}
fetchData()

</script>

<template>
  <page-container title="JSON解析">

    <!-- 上传文件组件 -->
    <!-- action ： 请求的URL -->
      <el-upload
        class="upload-demo"
        drag
        action= #
        multiple
      >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <br>
        <br>
        <div class="el-upload__text">
          拖动JSON文件到此 或<em> 点此上传 </em>
        </div>
      </el-upload>

        <!--  通过echarts 中的折线图进行历次数据展示  -->
        <div
          ref="chartRef"
          style="width: 100%; height: 300px;"
        > </div>
      
  </page-container>
</template>

<style lang="scss" scoped>

</style>



