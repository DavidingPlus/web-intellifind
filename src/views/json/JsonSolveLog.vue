<script setup>
import { getJsonSolveLogData } from '@/api/json.js'
import * as echarts from 'echarts';
import { onMounted, ref } from 'vue';
import { formatTime } from '@/utils/format.js'
import { Delete, Search } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';

const router = useRouter()
const loading = ref(false) // loading状态
const solveLogList = ref([]) // 解析记录列表
const total = ref(0) // 总数

// 定义请求参数对象
const userId = 1


// 基于params参数，获取解析记录
const getSolveLogList = async () => {
  loading.value = true
  const res = await getJsonSolveLogData(userId)
  solveLogList.value = res.data.data
  total.value = res.data.length
  loading.value = false
}
getSolveLogList()


// 删除逻辑
const onDelete = async (row) => {
  // 提示用户是否要删除
  await ElMessageBox.confirm('此操作将永久删除该记录, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  await artDelService(row.file_name)
  ElMessage.success('删除成功')
  // 重新渲染列表
  getSolveLogList()
}

// 添加解析
const addJsonSolve = () => {
  router.push('/json/solve')
}
// 路由挑战并携带file_name参数
const onShow = (row) => {
  console.log(row.file_name);
  console.log(typeof(row.file_name))
  router.push({ path: '/json/show', query: {file_name: row.file_name} })
}

// 通过echarts 中的折线图进行历次数据展示
const chartRef = ref(null)
const chartData = ref([])
onMounted(() => {

  // 图表
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
  <page-container title="解析记录">
    <template #extra>
      <el-button type="primary" @click="addJsonSolve">添加解析</el-button>
    </template>

    <!-- 表格区域 -->
    <el-table :data="solveLogList" v-loading="loading">
      <el-table-column label="解析文件名" prop="title">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{ row.file_name }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="综合得分" prop="title">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{ row.total_score }}</el-link>
        </template>
      </el-table-column>

      <el-table-column label="解析时间" prop="date">
        <template #default="{ row }">
          {{ formatTime(row.create_time) }}
        </template>
      </el-table-column>
      <!-- 利用作用域插槽 row 可以获取当前行的数据 => v-for 遍历 item -->
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button
            circle
            plain
            type="primary"
            :icon="Search"
            @click="onShow(row)"
          ></el-button>
          <el-button
            circle
            plain
            type="danger"
            :icon="Delete"
            @click="onDelete(row)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>

    <br>
    <br>

    <!-- 图表区域 -->
    <div ref="chartRef" style="width: 100%; height: 400px;"></div>
    <br>
  </page-container>
</template>

<style lang="scss" scoped></style>
