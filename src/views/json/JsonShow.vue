<script setup>
import {
  Iphone,
  Location,
  Tickets,
  Compass
} from '@element-plus/icons-vue';
import * as echarts from 'echarts';
import { ElDescriptions, ElDescriptionsItem, ElTable, ElTableColumn, ElTag } from 'element-plus';
import { onMounted, ref } from 'vue';


const chartContainRef = ref(null);
const gradeChartRef = ref(null);
const ratioChartRef = ref(null);
const exactGradeChartRef = ref(null);

onMounted(() => {

  // 设置容器宽高
  const chartContainerWidth = chartContainRef.value.clientWidth / 3; // 分成两份
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
              value: 0.85,
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
      indicator: [
        { name: '参数1', max: 25 },
        { name: '参数2', max: 25 },
        { name: '参数3', max: 25 },
        { name: '参数4', max: 25 },
        { name: '参数5', max: 25 },
        { name: '参数6', max: 25 },
        { name: '参数7', max: 25 },
        { name: '参数8', max: 25 },
        { name: '参数9', max: 25 }
      ]
    },
    series: [
      {
        name: 'ratio',
        type: 'radar',
        data: [
          {
            value: [5, 15, 25, 10, 5, 20, 5, 15, 10]
          }
        ]
      }
    ]
}
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
      data: [
        { value: 86, name: '参数1' },
        { value: 95, name: '参数2' },
        { value: 90, name: '参数3' },
        { value: 84, name: '参数4' },
        { value: 80, name: '参数5' },
        { value: 84, name: '参数6' },
        { value: 90, name: '参数7' },
        { value: 78, name: '参数8' },
        { value: 79, name: '参数9' }
      ],
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
  exactGradeChart.setOption(option3)}})

// 表格数据优劣评判
const tableRowClassName = ({ row }) => {
  if (row.score >= 85) {
    return 'success-row';
  } else if (row.score <= 70) {
    return 'warning-row';
  }
  return '';
}

const tableData = [
  { paramName: '参数1', score: 66, comment: '评价较差' },
  { paramName: '参数2', score: 95, comment: '评价很好' },
  { paramName: '参数3', score: 90, comment: '评价较好' },
  { paramName: '参数4', score: 74, comment: '评价一般' },
  { paramName: '参数5', score: 80, comment: '评价一般' },
  { paramName: '参数6', score: 84, comment: '评价较好' },
  { paramName: '参数7', score: 90, comment: '评价较好' },
  { paramName: '参数8', score: 68, comment: '评价较差' },
  { paramName: '参数9', score: 79, comment: '评价一般' }
]

const pointlist = [
  { pointname: '参数1', point: 5 },
  { pointname: '参数2', point: 15 },
  { pointname: '参数3', point: 25 },
  { pointname: '参数4', point: 10 },
  { pointname: '参数5', point: 5 },
  { pointname: '参数6', point: 20 },
  { pointname: '参数7', point: 5 },
  { pointname: '参数8', point: 15 },
  { pointname: '参数9', point: 10 }
]

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
          位置信息
        </div>
      </template>
      country + province + city
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
          <Compass />
          </el-icon>
          浏览器版本
        </div>
      </template>
      Chrome  88.0.4324.190
    </el-descriptions-item>
    <el-descriptions-item>
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
            <iphone />
          </el-icon>
          设备操作系统
        </div>
      </template>
      PC端  Windows 11
    </el-descriptions-item>
  </el-descriptions>




  <el-descriptions
    class="margin-top"
    title="用户权重配置信息"
    :column="3"
    border>
    <el-descriptions-item v-for="(item, index) in pointlist" :key="index">
      <template #label>
        <div class="cell-item">
          <el-icon :style="iconStyle">
            <tickets />
          </el-icon>
          {{item.pointname}}
        </div>
      </template>
      <el-tag>{{item.point}}</el-tag>
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
      <el-table-column prop="paramName" label="参数名" width="180" />
      <el-table-column prop="score" label="该项得分" width="180" />
      <el-table-column prop="comment" label="详细评判" />
    </el-table>

    <!-- <br>
    <el-button type="primary" @click="saveAsImage" round>保存为图片</el-button> -->
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
</style>
