<script setup>
// import { artDelService, artGetListService } from '@/api/json.js'
import * as echarts from 'echarts';
import { onMounted, ref } from 'vue';
import { formatTime } from '@/utils/format.js'
import { Delete, Edit } from '@element-plus/icons-vue'
import ArticleEdit from './components/JsonEdit.vue'
import ChannelSelect from './components/ChannelSelect.vue'
const total = ref(0) // 总条数
const loading = ref(false) // loading状态

// 定义请求参数对象
const params = ref({
  pagenum: 1, // 当前页
  pagesize: 5, // 当前生效的每页条数
  cate_id: '',
  grade: '',
})

const articleList = ref([
  {
    id: 1,
    title: '用户1- 你小子有点不服？',
    pub_date: '2021-09-01 12:00:00',
    grade: 'Grade A ：82'
  },
  {
    id: 2,
    title: '用户2- 来嘛，不服创我！',
    pub_date: '2021-09-01 12:00:00',
    grade: 'Grade B+：72'
  },
  {
    id: 3,
    title: '用户3- 试试就逝世！',
    pub_date: '2021-09-01 12:00:00',
    grade: 'Grade A+：90'
  },
  {
    id: 4,
    title: '用户4- Duang！！！！',
    pub_date: '2021-09-01 12:00:00',
    grade: 'Grade A ：86'
  },
  {
    id: 5,
    title: '用户5- 服了服了别创我了',
    pub_date: '2021-09-01 12:00:00',
    grade: 'Grade A+：92'
  }
]) 


// 基于params参数，获取文章列表
// const getArticleList = async () => {
//   loading.value = true

//   const res = await artGetListService(params.value)
//   articleList.value = res.data.data
//   total.value = res.data.total

//   loading.value = false
// }
// getArticleList()

// 处理分页逻辑
const onSizeChange = (size) => {
  // console.log('当前每页条数', size)
  // 只要是每页条数变化了，那么原本正在访问的当前页意义不大了，数据大概率已经不在原来那一页了
  // 重新从第一页渲染即可
  params.value.pagenum = 1
  params.value.pagesize = size
  // 基于最新的当前页 和 每页条数，渲染数据
  getArticleList()
}
const onCurrentChange = (page) => {
  // 更新当前页
  params.value.pagenum = page
  // 基于最新的当前页，渲染数据
  getArticleList()
}

// 搜索逻辑 => 按照最新的条件，重新检索，从第一页开始展示
const onSearch = () => {
  params.value.pagenum = 1 // 重置页面
  getArticleList()
}

// 重置逻辑 => 将筛选条件清空，重新检索，从第一页开始展示
const onReset = () => {
  params.value.pagenum = 1 // 重置页面
  params.value.cate_id = ''
  params.value.state = ''
  getArticleList()
}

// const articleEditRef = ref()
// // 添加逻辑
// const onAddArticle = () => {
//   articleEditRef.value.open({})
// }
// // 编辑逻辑
// const onEditArticle = (row) => {
//   articleEditRef.value.open(row)
// }

// 删除逻辑
const onDeleteArticle = async (row) => {
  // 提示用户是否要删除
  await ElMessageBox.confirm('此操作将永久删除该记录, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  await artDelService(row.id)
  ElMessage.success('删除成功')
  // 重新渲染列表
  getArticleList()
}

// // 添加或者编辑 成功的回调
// const onSuccess = (type) => {
//   if (type === 'add') {
//     // 如果是添加，最好渲染最后一页
//     const lastPage = Math.ceil((total.value + 1) / params.value.pagesize)
//     // 更新成最大页码数，再渲染
//     params.value.pagenum = lastPage
//   }

//   getArticleList()
// }




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
  <page-container title="解析记录">
    <template #extra>
      <el-button type="primary" @click="onAddArticle">添加解析</el-button>
    </template>

    <!-- 表单区域 -->
    <el-form inline style="display: flex; justify-content: space-between;">
      <el-form-item label="用户 :" style="flex: 1">
        <channel-select v-model="params.cate_id" width="100px"></channel-select>
      </el-form-item>
      <el-form-item label="时间 :" style="flex: 1">
        <!-- 这里后台标记发布状态，就是通过中文标记的，已发布 / 草稿 -->
        <channel-select v-model="params.state" width="100px"></channel-select>
      </el-form-item>
      <el-form-item style="flex: 1">
        <el-button @click="onSearch" type="primary">搜索</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
      <el-form-item style="flex: 1"></el-form-item>
    </el-form>

    <!-- 表格区域 -->
    <el-table :data="articleList" v-loading="loading">
      <el-table-column label="用户" prop="title">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{ row.title }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="得分" prop="title">
        <template #default="{ row }">
          <el-link type="primary" :underline="false">{{ row.grade }}</el-link>
        </template>
      </el-table-column>
      <!-- <el-table-column label="分类" prop="cate_name"></el-table-column> -->
      <el-table-column label="解析时间" prop="pub_date">
        <template #default="{ row }">
          {{ formatTime(row.pub_date) }}
        </template>
      </el-table-column>
      <!-- <el-table-column label="状态" prop="state"></el-table-column> -->
      <!-- 利用作用域插槽 row 可以获取当前行的数据 => v-for 遍历 item -->
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button
            circle
            plain
            type="primary"
            :icon="Edit"
            @click="onEditArticle(row)"
          ></el-button>
          <el-button
            circle
            plain
            type="danger"
            :icon="Delete"
            @click="onDeleteArticle(row)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页区域 -->
    <el-pagination
      v-model:current-page="params.pagenum"
      v-model:page-size="params.pagesize"
      :page-sizes="[2, 3, 5, 10]"
      :background="true"
      layout="jumper, total, sizes, prev, pager, next"
      :total="total"
      @size-change="onSizeChange"
      @current-change="onCurrentChange"
      style="margin-top: 20px; justify-content: flex-end"
    />

    <!-- 添加编辑的抽屉 -->
    <article-edit ref="articleEditRef" @success="onSuccess"></article-edit>

    <div
          ref="chartRef"
          style="width: 100%; height: 300px;"
        > </div>

  </page-container>
</template>

<style lang="scss" scoped></style>
