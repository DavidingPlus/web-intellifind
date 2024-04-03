<script setup>
import {
  Iphone,
  Location,
  Tickets,
  Compass
} from '@element-plus/icons-vue';
import * as echarts from 'echarts';
import { ElDescriptions, ElDescriptionsItem, ElTable, ElTableColumn, ElTag } from 'element-plus';
import { onMounted, ref, computed } from 'vue';
import { getJsonSolveData } from '@/api/json.js';
import { useRoute } from 'vue-router';
import { formatTime } from '@/utils/format.js';


// 接收路由参数
const props = defineProps({
  file_name: {
    type: String,
    required: true
  }
})
const route = useRoute()
const file_name = ref('')

// 环境信息表数据
const envlist = ref([
  { name: '位置信息', value: '' },
  { name: '客户端', value: '' },
  { name: '解析时间', value: '' }
])
// 权重信息表数据
const ratioList = ref([
  { word: '页面加载报错', name: 'console_errors', value: '' },
  { word: '点击报错', name: 'error_count', value: '' },
  { word: '网络响应慢', name: 'feedback_interval', value: '' },
  { word: '页面出现白屏', name: 'is_blank', value: '' },
  { word: '页面无响应', name: 'no_reaction', value: '' },
  { word: '出现多类问题', name: 'occur_many', value: '' },
  { word: '页面加载慢', name: 'page_load', value: '' },
  { word: '重复点击', name: 'repeat_click', value: '' },
  { word: '跳出率较高', name: 'stay_time', value: '' }
])
// 综合得分
const gradeChartTotalScore = ref(0.00)
// 权重比例图数据列表
const ratioChartDataList = ref([])
// 各项得分详情表
const exactGradeChartDataList = ref([
  { name: '页面加载报错', word: 'console_errors', value: '' },
  { name: '点击报错', word: 'error_count', value: '' },
  { name: '网络响应慢', word: 'feedback_interval', value: '' },
  { name: '页面出现白屏', word: 'is_blank', value: '' },
  { name: '页面无响应', word: 'no_reaction', value: '' },
  { name: '出现多类问题', word: 'occur_many', value: '' },
  { name: '页面加载慢', word: 'page_load', value: '' },
  { name: '重复点击', word: 'repeat_click', value: '' },
  { name: '跳出率较高', word: 'stay_time', value: '' }
])
// 评价表
const tableData = ref([
  {  name: '页面加载报错', paramName: 'console_errors', value: '', comment: '' },
  {  name: '点击报错', paramName: 'error_count', value: '', comment: '' },
  {  name: '网络响应慢', paramName: 'feedback_interval', value: '', comment: '' },
  {  name: '页面出现白屏', paramName: 'is_blank', value: '', comment: '' },
  {  name: '页面无响应', paramName: 'no_reaction', value: '', comment: '' },
  {  name: '出现多类问题', paramName: 'occur_many', value: '', comment: '' },
  {  name: '页面加载慢', paramName: 'page_load', value: '', comment: '' },
  {  name: '重复点击', paramName: 'repeat_click', value: '', comment: '' },
  {  name: '跳出率较高', paramName: 'stay_time', value: '', comment: '' }
])
// 简要检测结果
const briefResultList = ref([])
// 综合分析及建议
const allAnalysis = ref('')


// 获取解析数据
const getJsonSolveChartData = async () => {
  console.log(typeof(file_name.value));
  const { data: { data } } = await getJsonSolveData(file_name.value)

  // 用户环境信息表数据
  envlist.value[0].value = data.basic_info.Province + ' ' + data.basic_info.City 
                              + '  ip: ' + data.basic_info.Ip
  envlist.value[1].value = data.basic_info.OperatingSystem + ' ' + data.basic_info.DeviceType 
                              + ' ' + data.basic_info.Broswer 
  envlist.value[2].value = formatTime(data.basic_info.CreateTime)
  // console.log(envlist.value);

  // 权重信息表数据
  ratioList.value[0].value = data.settings.console_errors
  ratioList.value[1].value = data.settings.error_count
  ratioList.value[2].value = data.settings.feedback_interval
  ratioList.value[3].value = data.settings.is_blank
  ratioList.value[4].value = data.settings.no_reaction
  ratioList.value[5].value = data.settings.occur_many
  ratioList.value[6].value = data.settings.page_load
  ratioList.value[7].value = data.settings.repeat_click
  ratioList.value[8].value = data.settings.stay_time
  // console.log(ratioList.value);

  // 综合得分表总得分     (.toFixed(2) 保留两位有效数字)
  gradeChartTotalScore.value = (data.result.total_score / 2.00 ).toFixed(2)
  // console.log(gradeChartTotalScore.value);

  // 权重比例图数据
  ratioChartDataList.value = ratioList.value.map(item => item.value)

  // 各项参数得分详细数据
  exactGradeChartDataList.value[0].value = data.result.console_errors
  exactGradeChartDataList.value[1].value = data.result.error_count
  exactGradeChartDataList.value[2].value = data.result.feedback_interval
  exactGradeChartDataList.value[3].value = data.result.is_blank
  exactGradeChartDataList.value[4].value = data.result.no_reaction
  exactGradeChartDataList.value[5].value = data.result.occur_many
  exactGradeChartDataList.value[6].value = data.result.page_load
  exactGradeChartDataList.value[7].value = data.result.repeat_click
  exactGradeChartDataList.value[8].value = data.result.stay_time

  // 评价表
  tableData.value[0].value = data.result.console_errors
  tableData.value[1].value = data.result.error_count
  tableData.value[2].value = data.result.feedback_interval
  tableData.value[3].value = data.result.is_blank
  tableData.value[4].value = data.result.no_reaction
  tableData.value[5].value = data.result.occur_many
  tableData.value[6].value = data.result.page_load
  tableData.value[7].value = data.result.repeat_click
  tableData.value[8].value = data.result.stay_time

  // 简要检测结果
  const str1 = data.result.brief_desc
  const regex = /(?<=，)(.*?)(?=，得分)/g;
  briefResultList.value = str1.match(regex);
  tableData.value[0].comment = briefResultList.value[6]
  tableData.value[1].comment = briefResultList.value[5]
  tableData.value[2].comment = briefResultList.value[3]
  tableData.value[3].comment = briefResultList.value[7]
  tableData.value[4].comment = briefResultList.value[4]
  tableData.value[5].comment = briefResultList.value[8]
  tableData.value[6].comment = briefResultList.value[2]
  tableData.value[7].comment = briefResultList.value[1]
  tableData.value[8].comment = briefResultList.value[0]

  // 综合分析及建议
  const str2 = data.result.detail_desc
  allAnalysis.value = str2.replace(/\n/g, '<br> &nbsp; &nbsp; &nbsp; &nbsp;')  // 将换行符替换为<br>标签 
}


// 图标大小
const size = ref('default')
const iconStyle = computed(() => {
  const marginMap = {
    large: '8px',
    default: '6px',
    small: '4px',
  }
  return {
    marginRight: marginMap[size.value] || marginMap.default,
  }
})

const chartContainRef = ref([]);
const gradeChartRef = ref(null);
const ratioChartRef = ref(null);
const exactGradeChartRef = ref(null);


onMounted(async () => {
  // 获取路由传参
  file_name.value = route.query.file_name
  console.log(file_name.value);

  // 获取解析数据
  await getJsonSolveChartData()

  // 图表数据
  if (chartContainRef.value) {
  // 设置容器宽高
  const chartContainerWidth = chartContainRef.value.clientWidth / 3; // 分成三份
  const chartContainerHeight = 390; 

  // 综合得分表展示
  if (gradeChartRef.value) {
    const gradeChart = echarts.init(gradeChartRef.value);
    gradeChart.resize({ width: chartContainerWidth, height: chartContainerHeight });
      const option1 = {
      title: {
        text: '     综合得分表',
      },
      series: [
        {
          type: 'gauge',
          startAngle: 180,
          endAngle: 0,
          center: ['50%', '75%'],
          radius: '90%',
          min: 0,
          max: 1,
          splitNumber: 8,
          axisLine: {
            lineStyle: {
              width: 6,
              color: [
                [0.25, '#FF6E76'],
                [0.5, '#FDDD60'],
                [0.75, '#58D9F9'],
                [1, '#7CFFB2']
              ]
            }
          },
          pointer: {
            icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z',
            length: '12%',
            width: 20,
            offsetCenter: [0, '-60%'],
            itemStyle: {
              color: 'auto'
            }
          },
          axisTick: {
            length: 12,
            lineStyle: {
              color: 'auto',
              width: 2
            }
          },
          splitLine: {
            length: 20,
            lineStyle: {
              color: 'auto',
              width: 5
            }
          },
          axisLabel: {
            color: '#464646',
            fontSize: 20,
            distance: -60,
            rotate: 'tangential',
            formatter: function (value) {
              if (value === 0.875) {
                return 'Grade A';
              } else if (value === 0.625) {
                return 'Grade B';
              } else if (value === 0.375) {
                return 'Grade C';
              } else if (value === 0.125) {
                return 'Grade D';
              }
              return '';
            }
          },
          title: {
            offsetCenter: [0, '-10%'],
            fontSize: 20
          },
          detail: {
            fontSize: 30,
            offsetCenter: [0, '-35%'],
            valueAnimation: true,
            formatter: function (value) {
              return Math.round(value * 100) + '';
            },
            color: 'inherit'
          },
          data: [
            {
              value: gradeChartTotalScore.value,
              name: 'Grade'
            }
          ]
        }
      ]
    }
gradeChart.setOption(option1)}

  // 比例图展示
  if (ratioChartRef.value) {
    const ratioChart = echarts.init(ratioChartRef.value);
    ratioChart.resize({ width: chartContainerWidth, height: chartContainerHeight });
    const option2 = {
    title: {
      text: '各项参数权重图    ',
      left: 'center'
    },
    radar: {
      // shape: 'circle',
      indicator: ratioList.value.map( item => ({
        name: item.word,
        max: 10
      }))
    },
    series: [
      {
        type: 'radar',
        data: [
          {
            value: ratioChartDataList.value
          }
        ]
      }
    ]
}
// console.log(option2.radar);
// console.log(option2.series.data);
ratioChart.setOption(option2)}

// 每项成绩具体得分表
if (exactGradeChartRef.value) {
  const exactGradeChart = echarts.init(exactGradeChartRef.value);
  exactGradeChart.resize({ width: chartContainerWidth, height: chartContainerHeight });
  const option3 = {
    title: { 
      text: '各项具体得分表         ',
      left: 'right' 
    },
    tooltip: {
    trigger: 'item'
  },
  series: [
    {
      name: '各项具体得分表',
      type: 'pie',
      radius: '50%',
      data: exactGradeChartDataList.value,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }
  ]
  }
  exactGradeChart.setOption(option3)}}})

// 表格数据优劣评判
const tableRowClassName = ({ row }) => {
  if (row.value >= 90) {
    // row.comment = '不戳不戳针不戳针不戳'
    return 'success-row';
  } else if (row.value <= 80) {
    // row.comment = 'emmmmm，海星海星'
    return 'warning-row';
  }
  // row.comment = '蛙趣凶滴，拟屎真滴捞哇'
  return '';
}




</script>

<template>
  <page-container title="JSON解析详情展示">
    <el-descriptions
    class="margin-top"
    title="用户环境配置信息"
    :column="3"
    border
  >
  <!-- <el-descriptions-item v-for="(item, index) in envlist" :key="index">
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
            <location />
          </el-icon>
          {{item.name}}
        </div>
      </template>
      {{item.value}}
    </el-descriptions-item> -->
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
            <location />
          </el-icon>
          {{ envlist[0].name }}
        </div>
      </template>
      {{ envlist[0].value }}
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
          <Compass />
          </el-icon>
          {{ envlist[1].name}}
        </div>
      </template>
      {{ envlist[1].value }}
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
            <iphone />
          </el-icon>
          {{ envlist[2].name}}
        </div>
      </template>
      {{ envlist[2].value }}
    </el-descriptions-item>
  </el-descriptions>




  <el-descriptions
    class="margin-top"
    title="用户权重配置信息"
    :column="3"
    border>
    <el-descriptions-item v-for="(item, index) in ratioList" :key="index">
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
            <tickets />
          </el-icon>
          {{item.word}}
        </div>
      </template>
      <el-tag>{{item.value}}</el-tag>
    </el-descriptions-item>
  </el-descriptions>
    <br>
    <br>
    <br>


    <div ref="chartContainRef"  style="display: flex; justify-content: space-between;">

    <!-- 综合得分表 -->
      <div ref="gradeChartRef"></div>

    <!-- 各项系数比例图 -->
      <div ref="ratioChartRef"></div>

    <!-- 各项具体得分表 -->
      <div ref="exactGradeChartRef"></div>

    </div>

    <el-table
      :data="tableData"
      style="width: 100%"
      :row-class-name="tableRowClassName"
    >
      <el-table-column prop="name" label="参数名" width="180" />
      <el-table-column prop="paramName" label="英文名" width="180" />
      <el-table-column prop="value" label="该项得分" width="180" />
      <el-table-column prop="comment" label="详细评判" />
    </el-table>

    <br>
    
    <!-- 此处添加 -->
 
          <div class="long-text-container">  
            <el-card>
            <h3>综合分析及建议</h3>
            <p v-html="allAnalysis" ></p>  
          </el-card>
          </div>  
 

  </page-container>
</template>

<style>
.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}
.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}

.el-descriptions {
  margin-top: 20px;
}
.cell-item {
  display: flex;
  align-items: center;
}
.margin-top {
  margin-top: 20px;
}

.long-text-container p {  
  font-size: 16px;  
  line-height: 1.8;  
  text-align: justify; /* 两端对齐 */  
  color: #333; /* 文本颜色 */  
  margin-bottom: 20px; /* 段间距 */  
  /* 其他你需要的样式 */  
}  
</style>
