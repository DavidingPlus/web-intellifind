<script setup>
import { UploadFilled } from '@element-plus/icons-vue';
import { ElButton, ElIcon, ElLoading, ElUpload } from 'element-plus';
import { ref } from 'vue';
import { useRouter } from 'vue-router';


const router = useRouter()
// 跳转到新开页面并展示详情
const show = (id) => {
  console.log(id)
  router.push('/json/show')
}

const flag = ref(true)
const percentage = ref(0.0)
const perFlag = ref(true)
const loading = ref(false)
// 上传文件
const upload = () => {
  console.log('上传文件')
  setTimeout(() => {
    flag.value = false
    loading.value = true
    progress()
  }, 3000)
  
}
// 进度函数
const progress = () => {
  percentage.value = 0.0
  setInterval(() => {
    if (perFlag.value) {
      percentage.value += 2.0
      if (percentage.value === 100) {
        perFlag.value = false
        setTimeout(() => {
          ElMessage.success('解析成功')
          loading.value = false
        }, 300)
    }
  }
}, 100)
}

// 用户配置表单
const formRef = ref()
const form = ref({
  city: '',
  system: '',
  ratio1: '',
  ratio1Score: '',
  ratio2: '',
  ratio2Score: ''
})
// 表单校验规则
const rules = {
  ratio1Score: [
    // 不是必填项，如果有，则数据应在0~25之间
    { required: false, validator: (value) => value >= 0 && 
      value <= 25 ? Promise.resolve() : Promise.reject('请输入0~25之间的数值'), trigger: 'blur' }
  ],
  ratio2Score: [
  { required: false, validator: (value) => value >= 0 && 
      value <= 25 ? Promise.resolve() : Promise.reject('请输入0~25之间的数值'), trigger: 'blur' }
  ]
}
// 上传用户配置信息
const submitForm = async () => {
  // 等待校验结果
  await formRef.value.validate()
  // 提交修改
  await userUpdateInfoService(form.value)
  // 通知 user 模块，进行数据的更新
  getUser()
  // 提示用户
  ElMessage.success('修改成功')
}

</script>

<template>
  <page-container title="JSON解析">

    <!-- 上传文件组件 -->
    <!-- action ： 请求的URL -->
    <!-- action= "8.137.100.0" -->
    <div style="position: relative; width: 100%; max-height:565px;">
        <el-upload
        class="upload-demo"
        drag
        
        multiple
        v-if="flag"
        @click="upload"
      >
        <el-icon class="el-icon--upload" @click="upload"><upload-filled @click="upload" /></el-icon>
        <br>
        <br>
        <div class="el-upload__text" @click="upload">
          拖动JSON文件到此 或<em @click="upload"> 点此上传 </em>
        </div>
      </el-upload>

      <div v-else="flag" class="demo-progress" 
      style="display: flex; flex-direction: column; 
      justify-content: center; align-items: center;">
        
        <!-- <el-loading content="文件上传中..."  />   -->
        <br>
        <el-progress
          :text-inside="true"
          :stroke-width="24"
          :percentage="percentage"
          status="success"
          style="width: 80%;"
        />
        <br>
        <el-table v-if="loading" v-loading="loading" style="width: 100%">
        </el-table>
        <el-button v-else="loading" type="success" @click="show" round>点此查看解析详情</el-button>
      </div>
    </div>
      <br>
      <br>

      <el-form ref="formRef" :model="form" :rules="rules">
      <el-row>
        <el-col :span="24">
          <h2 style="text-align: left;">自定义配置信息</h2>
        </el-col>
      </el-row>
      <el-form-item label="所在城市">
        <el-input v-model="form.city" placeholder="请输入您当前所在城市"></el-input>
      </el-form-item>
      <el-form-item label="所使用系统">
        <el-input v-model="form.system" placeholder="请输入您当前使用的主机系统（Windows、MacOS、Linux）"></el-input>
      </el-form-item>
      <el-form-item label="权重点1" style="display: flex; justify-content: space-between;">
          <el-select style="flex: 1" v-model="form.ratio1" placeholder="请选择你对本次解析看重的点"></el-select>
          <el-input style="flex: 1" v-model="form.ratio1Score" placeholder="请输入您为其设定的权重值（0~25）"></el-input>
      </el-form-item>
      <el-form-item label="权重点2" style="display: flex; align-items: center;">
          <el-select style="flex: 1" v-model="form.ratio2" placeholder="请选择你对本次解析看重的点"></el-select>
          <el-input style="flex: 1" v-model="form.ratio2Score" placeholder="请输入您为其设定的权重值（0~25）"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm">保存配置</el-button>
      </el-form-item>
    </el-form>


  </page-container>
</template>

<style lang="scss" scoped>
.demo-progress .el-progress--line {
  margin-bottom: 15px;
  max-width: 600px;
}
</style>



