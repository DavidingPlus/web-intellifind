<script setup>
import { getJsonSolveLogData, deleteJsonSolveLogData } from '@/api/json.js'
import * as echarts from 'echarts';
import { onMounted, ref } from 'vue';
import { formatTime } from '@/utils/format.js'
import { Delete, Search } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';

const emptyFlag = ref()
const router = useRouter()
const loading = ref(false) // loading状态
const solveLogList = ref([]) // 解析记录列表
const total = ref(0) // 总数
const totalPage = ref(0) // 总页数

// 折线图数据
const totalScoreList = ref([])
const pageErrorScoreList = ref([])
const pageLoadScoreList = ref([])
const pageExperienceScoreList = ref([])

// 通过echarts 中的折线图进行历次数据展示
const chartRef = ref(null)
const chartData = ref([])


// 定义请求参数对象
const params = ref({
  page: 1,
  size: 7
})

// 基于params参数，获取解析记录
const getSolveLogList = async () => {
  emptyFlag.value = false
  loading.value = true
  const res = await getJsonSolveLogData(params.value)
  // 七次解析数据
  solveLogList.value = res.data.data
  total.value = res.data.length
  totalPage.value = res.data.total_page
  if (!total.value) {
    emptyFlag.value = true
  } else {
    loading.value = false
    console.log(total.value);
  }

  // 折线图数据
  // console.log(res.data.data[0].total_score);
  totalScoreList.value = res.data.data.map(item => (item.total_score).toFixed(2))
  pageErrorScoreList.value = res.data.data.map(item => (item.page_error).toFixed(2))
  pageLoadScoreList.value = res.data.data.map(item => (item.page_load).toFixed(2))
  pageExperienceScoreList.value = res.data.data.map(item => (item.page_experience).toFixed(2))
  fetchData()

  
}
getSolveLogList()

// 处理分页逻辑
const onSizeChange = (size) => {
  params.value.page = 1
  params.value.size = size
  getSolveLogList()
}
const onCurrentChange = (page) => {
  params.value.page = page
  getSolveLogList()
}


// 删除逻辑
const onDelete = async (row) => {
  // 提示用户是否要删除
  await ElMessageBox.confirm('此操作将永久删除该记录, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  await deleteJsonSolveLogData(row.file_name)
  console.log('删除成功' + row.file_name);
  // 重新渲染列表
  setTimeout(() => {
    getSolveLogList()
    console.log('列表刷新成功' + row.file_name);
    ElMessage.success('删除成功')
    console.log(total.value);
  }, 300)
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




  // 获取数据
  const fetchData = () => {
    chartData.value = [{
        name: '综合得分',
        data: totalScoreList.value
      },
      {
        name: '页面报错得分',
        data: pageErrorScoreList.value
      },
      {
        name: '页面加载得分',
        data: pageLoadScoreList.value
      },
      {
        name: '页面体验得分',
        data: pageExperienceScoreList.value
      }]
    // console.log(chartData.value);

  // 图表
  if (chartRef.value) {
  const chart = echarts.init(chartRef.value)
  const option = {
    title: {
      text: '    七次解析反馈数据'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['综合得分', '页面报错得分', '页面加载得分', '页面体验得分']
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
      min: Math.floor(Math.min(...chartData.value.map(item => Math.min(...item.data))) / 5) * 5,  
      // 设置interval为5的倍数，这里假设我们想要每5个单位一个标签  
      interval: 5,  
      axisLabel: {  
        // 格式化标签，确保显示为整数  
        formatter: function (value) {  
          return value.toFixed(0); // 根据需要调整小数点位数  
        }  
      }
    },
    // 通过循环的方式，将历次数据进行展示
    series: chartData.value.reverse().map(item => ({
      name: item.name,
      data: item.data,
      type: 'line',
    }))
  }
chart.setOption(option)}  }

// 表格数据优劣评判
const tableRowClassName = ({ row }) => {
  if (row.total_score >= 89) {
    return 'success-row';
  } else if (row.total_score <= 86) {
    return 'warning-row';
  }
  return '';
}



</script>

<template>
  <page-container title="解析记录">
    <template #extra>
      <el-button type="primary" @click="addJsonSolve">添加解析</el-button>
    </template>

    <el-empty v-if="emptyFlag" description="暂无解析记录" />
    <!-- 表格区域 -->
    <el-table v-else 
      :data="solveLogList" v-loading="loading" 
      :row-class-name="tableRowClassName">
      <el-table-column label="解析文件名" align="center" prop="file_name">
        <template #default="{ row }">
          <el-link type="primary">{{ (row.file_name).substr(14) }}</el-link>
        </template>

      </el-table-column>
      <el-table-column label="页面报错得分" align="center" prop="page_error">
        <template #default="{ row }">
          <el-link type="primary">{{ (row.page_error).toFixed(2) }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="页面加载得分" align="center" prop="page_load">
        <template #default="{ row }">
          <el-link type="primary">{{ (row.page_load).toFixed(2) }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="页面体验得分" align="center" prop="page_expercience">
        <template #default="{ row }">
          <el-link type="primary">{{ (row.page_experience).toFixed(2) }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="综合得分" align="center" prop="total_score">
        <template #default="{ row }">
          <el-link type="primary">{{ (row.total_score).toFixed(2) }}</el-link>
        </template> 
      </el-table-column>

      <el-table-column label="解析时间" align="center" width="280">
        <template #default="{ row }">
          <el-link type="primary">{{ formatTime(row.create_time) }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="操作"  align="center" width="150">
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

    <!-- 分页区域 -->
    <!-- :page-sizes="[3, 5, 10]" -->
    <el-pagination
      :current-page="params.page"
      :page-size="params.size"
      :background="true"
      layout="jumper, total, prev, pager, next, ->"
      :total="total"
      @size-change="onSizeChange"
      @current-change="onCurrentChange"
      style="margin-top: 20px; justify-content: flex-end; margin-right: 80px;"
    />

    <br>
    <br>

    <!-- 图表区域 -->
    <div ref="chartRef" style="width: 100%; height: 400px;"></div>
    <br>
  </page-container>
</template>

<style>
.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}
.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}
</style>
