<script setup>
// import { getJsonSolveLogData } from '@/api/json.js'
import * as echarts from 'echarts';
import { onMounted, ref } from 'vue';
import { formatTime } from '@/utils/format.js'
import { Delete, Search } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';

const router = useRouter()
const loading = ref(false) // loading状态
const solveLogList = ref([]) // 解析记录列表

// 定义请求参数对象
const params = ref({
  userId: 1,
})

const articleList = ref([
  {
    id: 1,
    filename: '1- 你小子有点不服？',
    date: '2021-09-01 12:00:00',
    grade: 'Grade A ：82'
  },
  {
    id: 2,
    filename: '2- 来嘛，不服创我！',
    date: '2021-09-01 12:00:00',
    grade: 'Grade B+：72'
  },
  {
    id: 3,
    filename: '3- 试试就逝世！',
    date: '2021-09-01 12:00:00',
    grade: 'Grade A+：90'
  },
  {
    id: 4,
    filename: '4- Duang！！！！',
    date: '2021-09-01 12:00:00',
    grade: 'Grade A ：86'
  },
  {
    id: 5,
    filename: '5- 服了服了别创我了',
    date: '2021-09-01 12:00:00',
    grade: 'Grade A+：92'
  },
  {
    id: 6,
    filename: '6- 你小子有点不服？',
    date: '2021-09-01 12:00:00',
    grade: 'Grade A ：82'
  },
  {
    id: 7,
    filename: '7- 来嘛，不服创我！',
    date: '2021-09-01 12:00:00',
    grade: 'Grade B+：72'
  }
]) 

// 基于params参数，获取解析记录
const getSolveLogList = async () => {
  loading.value = true
  // const res = await getJsonSolveLogData(params.value)
  const res = await getJsonSolveLogData()
  solveLogList.value = res.data.data
  total.value = res.data.total

  loading.value = false
}
// getSolveLogList()


// 删除逻辑
const onDelete = async (row) => {
  // 提示用户是否要删除
  await ElMessageBox.confirm('此操作将永久删除该记录, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  await artDelService(row.id)
  ElMessage.success('删除成功')
  // 重新渲染列表
  getSolveLogList()
}

// 添加解析
const addJsonSolve = () => {
  router.push('/json/solve')
}
const onShow = (row) => {
  console.log(row.id)
  router.push('/json/show')
}



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
    <el-table :data="articleList" v-loading="loading">
      <el-table-column label="解析ID" prop="title">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{ row.filename }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="综合得分" prop="title">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{ row.grade }}</el-link>
        </template>
      </el-table-column>

      <el-table-column label="解析时间" prop="date">
        <template #default="{ row }">
          {{ formatTime(row.date) }}
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
