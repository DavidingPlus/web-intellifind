<script setup>
import { ref } from 'vue'
import { Plus, Upload } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores'
import { userUpdateAvatarService } from '@/api/user'


const userStore = useUserStore();
const uploadRef = ref();
const imgUrl = ref(userStore.user.user_pic)

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

</script>

<template>
  <page-container title="更换头像">
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
