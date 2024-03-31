<script setup>
import { userGetInfoService, userUpdateInfoService } from '@/api/user';
import PageContainer from '@/components/PageContainer.vue';
import { ref } from 'vue';
import { useUserStore } from '@/stores';

const userStore = useUserStore()
const formRef = ref()
// // 获取当前时间并填入输入框
// const currentTime = new Date().toLocaleString();


const form = ref({
  user_name: '蓬鱼咽',
  email: '9517538462@pp.com',
  city: '山河四省',
  // date: '2025年41月31日'
  avatar: ''
})


const rules = ref({
  username: [
    { required: true, message: '请输入用户姓名', trigger: 'blur' },
    {
      pattern: /^\S{2,10}/,
      message: '用户姓名长度在2-5个非空字符',
      trigger: 'blur'
    }
  ],
  email: [
    { required: true, message: '请输入用户邮箱', trigger: 'blur' },
    {
      type: 'email',
      message: '请输入正确的邮箱格式',
      trigger: ['blur', 'change']
    }
  ]
})

// 更新当前表单用户信息
const getUser = async () => {
  const {data}  = await userGetInfoService()
  form.value = data.data
  userStore.user.user_pic = data.data.avatar
}

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

getUser()

</script>
<template>
  <page-container title="基本资料">
    <!-- 表单部分 -->
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item label="用户姓名" prop="username">
        <el-input v-model="form.user_name" placeholder="蓬鱼咽"></el-input>
      </el-form-item>
      <el-form-item label="用户邮箱" prop="email">
        <el-input v-model="form.email" placeholder="9517538462@pp.com"></el-input>
      </el-form-item>
      <el-form-item label="所在城市">
        <el-input v-model="form.city" placeholder="山河四省"></el-input>
      </el-form-item>
      <el-form-item label="注册时间">
        <el-input></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm">提交修改</el-button>
      </el-form-item>
      
    </el-form>
  </page-container>
</template>
