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
  Birthday: '2056-10-10',
  Tel: '123456789',
  Gender: '男',
  avatar: '',
  create_time: '2021-10-10',
})

// 格式化时间
const formatDate = (dateTimeString) => {  
    const date = new Date(dateTimeString);  
    const year = date.getFullYear();  
    const month = (1 + date.getMonth()).toString().padStart(2, '0'); // 月份从0开始，所以加1  
    const day = date.getDate().toString().padStart(2, '0');  
    return `${year}-${month}-${day}`;  
}

const rules = ref({
  user_name: [
    { required: true, message: '请输入用户姓名', trigger: 'blur' },
    {
      pattern: /^\S{2,10}/,
      message: '用户姓名长度在2-10个非空字符',
      trigger: 'blur'
    }
  ],
  Tel: [
    { required: true, message: '请输入电话号码', trigger: 'blur' },
    {
      pattern: /^1[3-9]\d{9}$/,
      message: '请输入正确的电话号码',
      trigger: 'blur'
    }
  ]
})

// 更新当前表单用户信息
const getUser = async () => {
  const {data}  = await userGetInfoService()
  form.value = data.data
  userStore.user.user_pic = data.data.avatar
  form.value.create_time = formatDate(data.data.create_time)
  console.log(form.value);
}

const submitForm = async () => {
  // 等待校验结果
  await formRef.value.validate()
  form.value.Birthday = formatDate(form.value.Birthday)
  // 提交修改
  console.log(form.value);
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
        
        <el-form-item label="用户姓名">
          <el-input v-model="form.user_name" placeholder="请输入姓名"></el-input>
        </el-form-item>
        <el-form-item label="所在城市">
          <el-input v-model="form.city" placeholder="请输入所在城市"></el-input>
        </el-form-item>
        <el-form-item label="生日">
          <el-date-picker
            v-model="form.Birthday"
            type="date"
            placeholder="请选择日期"
          />
        </el-form-item>
        <el-form-item label="性别">
          <div>
            <el-radio-group v-model="form.Gender">
              <el-radio label="男" size="large"> 男 </el-radio>
              <el-radio label="女" size="large"> 女 </el-radio>
            </el-radio-group>
          </div>
        </el-form-item>
        <el-form-item label="电话">
          <el-input v-model="form.Tel" placeholder="请输入电话号码"></el-input>
        </el-form-item>
        <el-form-item label="用户邮箱">
          <el-input v-model="form.email" disabled></el-input>
        </el-form-item>
        <el-form-item label="注册时间">
          <el-input v-model="form.create_time" disabled></el-input>
        </el-form-item>
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