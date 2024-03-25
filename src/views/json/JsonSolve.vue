<script setup>
import { UploadFilled } from '@element-plus/icons-vue';
import { ElUpload, ElIcon, ElButton } from 'element-plus';
import { useRouter } from 'vue-router';
import { ref } from 'vue';


const router = useRouter()
// 跳转到新开页面并展示详情
const show = (id) => {
  console.log(id)
  router.push('/json/show')
}

const flag = ref(true)
const percentage = ref(0.0)
const perFlag = ref(true)
// 上传文件
const upload = () => {
  console.log('上传文件')
  setTimeout(() => {
    flag.value = false
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
      <br>
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
      

      <div v-else="flag" class="demo-progress">
        
        <el-progress
          :text-inside="true"
          :stroke-width="24"
          :percentage=percentage
          status="success"
        />
        <br>
        <el-button type="success" @click="show" round>点此查看解析详情</el-button>
      </div>
      <br>
      <br>

      <el-form ref="formRef" :model="form" :rules="rules" >
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



