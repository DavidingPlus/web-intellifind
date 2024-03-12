<script setup>
// import { userUpdateInfoService } from '@/api/user';
import PageContainer from '@/components/PageContainer.vue';
import { useUserStore } from '@/stores';
import { ref } from 'vue';

const formRef = ref()

// 是在使用仓库中数据的初始值 (无需响应式) 解构无问题
const {
  user: { email, id, username, nickname },
  getUser
} = useUserStore()

const form = ref({
  id,
  nickname,
  username,
  email,
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
  <page-container title="基本资料">
    <!-- 表单部分 -->
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item label="登录昵称">
        <el-input v-model="form.nickname" disabled></el-input>
      </el-form-item>
      <el-form-item label="用户姓名" prop="username">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="用户邮箱" prop="email">
        <el-input v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="所在城市">
        <el-input ></el-input>
      </el-form-item>
      <el-form-item label="注册时间">
        <el-input ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm">提交修改</el-button>
      </el-form-item>
    </el-form>
  </page-container>
</template>
