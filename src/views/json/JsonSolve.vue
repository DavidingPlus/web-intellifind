<script setup>
import { UploadFilled } from '@element-plus/icons-vue';
import * as echarts from 'echarts';
import { onMounted, ref } from 'vue';

// const articleEditRef = ref()
// const jsonFile = ref(null)

// // 解析json文件
// const onParse = () => {
//   const file = jsonFile.value
//   if (!file) {
//     return
//   }
//   const reader = new FileReader()
//   reader.onload = (e) => {
//     const data = e.target.result
//     const jsonData = JSON.parse(data)
//     console.log(jsonData)
//   }
//   reader.readAsText(file)
// }

// 通过echarts 中的折线图进行历次数据展示
const chartRef = ref(null)
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
      type: 'value'
    },
    series: [
      {
        name: 'JSON1',
        type: 'line',
        stack: '总量',
        data: [12, 13, 10, 13, 9, 23, 21]
      },
      {
        name: 'JSON2',
        type: 'line',
        stack: '总量',
        data: [22, 18, 19, 23, 29, 33, 31]
      },
      {
        name: 'JSON3',
        type: 'line',
        stack: '总量',
        data: [15, 23, 20, 15, 19, 33, 41]
      },
      {
        name: 'JSON4',
        type: 'line',
        stack: '总量',
        data: [32, 33, 30, 33, 39, 33, 32]
      },
      {
        name: 'JSON5',
        type: 'line',
        stack: '总量',
        data: [82, 93, 90, 93, 129, 133, 132]
      }
    ]
  }
chart.setOption(option)}})



</script>

<template>
  <page-container title="JSON解析">

    <!-- <el-form>
        <el-form-item>
          <el-input type="file" v-model="jsonFile" size="large" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onParse">解析</el-button>
        </el-form-item>
      <br>
      <br>
      </el-form> -->


      <el-upload
        class="upload-demo"
        drag
        action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
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