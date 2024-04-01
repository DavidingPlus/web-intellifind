<script setup>
import { userGetInfoService, userUpdateInfoService, userUpdateAvatarService } from '@/api/user';
import PageContainer from '@/components/PageContainer.vue';
import { ref } from 'vue';
import { useUserStore } from '@/stores';
import { Plus, Upload } from '@element-plus/icons-vue'

const userStore = useUserStore()
const formRef = ref()
const uploadRef = ref();
const imgUrl = ref(userStore.user.user_pic)


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

// 更换头像
const onSelectFile = (uploadFile) => {
  if (!uploadFile.raw) {
    console.error('未选择图片');
    return;
  }
  uploadRef.value = uploadFile.raw;
  // 基于 FileReader 读取图片做预览
  const reader = new FileReader()
  reader.readAsDataURL(uploadFile.raw)
  reader.onload = () => {
    imgUrl.value = reader.result
  }
}

const onUpdateAvatar = async () => {
  if (!uploadRef.value) {
    ElMessage.error('请选择图片');
    return;
  }
  const formData = new FormData();
  formData.append('avatar', uploadRef.value);
  try {
    // 发送请求更新头像
    await userUpdateAvatarService(formData);
    // userStore 重新渲染
    userStore.getUser();
    // 提示用户
    ElMessage.success('头像更新成功');
  } catch (error) {
    console.error('更新头像失败:', error);
    ElMessage.error('更新头像失败，请稍后重试');
  }
};

getUser()

</script>
<template>
  <page-container title="基本资料">
    <!-- 表单部分 -->
    <div style="display: flex; align-items: center; justify-content: space-between;">
      <div style="flex: 1">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item label="用户姓名" prop="username">
        <el-input v-model="form.user_name" placeholder="蓬鱼咽"></el-input>
      </el-form-item>
      <el-form-item label="用户邮箱" prop="email">
        <el-input v-model="form.email" placeholder="9517538462@pp.com" disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="所在城市">
        <el-input v-model="form.city" placeholder="山河四省"></el-input>
      </el-form-item>
      <!-- <el-form-item label="注册时间">
        <el-input></el-input>
      </el-form-item> -->
      <el-form-item>
        <el-button type="primary" @click="submitForm">提交修改</el-button>
      </el-form-item>
      </el-form>
    </div>

    <div style="flex: 1; text-align: center;">
      <el-upload
      type="file"
      ref="uploadRef"
      :auto-upload="false"
      class="avatar-uploader"
      :show-file-list="false"
      :on-change="onSelectFile"
    >

      <img v-if="imgUrl" :src="imgUrl" class="avatar" />
      <img v-else src="../../assets/default.png" class="avatar" />
      <!-- <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon> -->
    </el-upload>

    <br />

    <el-button
      @click="uploadRef.$el.querySelector('input').click()"
      type="primary"
      :icon="Plus"
      size="large"
      >选择图片</el-button
    >
    <el-button
      @click="onUpdateAvatar"
      type="success"
      :icon="Upload"
      size="large"
      >上传头像</el-button
    >
  </div>

  </div>
  </page-container>
</template>


<style lang="scss" scoped>
.avatar-uploader {
  :deep() {
    .avatar {
      width: 278px;
      height: 278px;
      display: block;
    }
    .el-upload {
      border: 1px dashed var(--el-border-color);
      border-radius: 6px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
      transition: var(--el-transition-duration-fast);
    }
    .el-upload:hover {
      border-color: var(--el-color-primary);
    }
    .el-icon.avatar-uploader-icon {
      font-size: 28px;
      color: #8c939d;
      width: 278px;
      height: 278px;
      text-align: center;
    }
  }
}
</style>