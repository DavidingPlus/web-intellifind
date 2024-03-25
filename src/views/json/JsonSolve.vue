<script setup>
import { UploadFilled } from '@element-plus/icons-vue';
import { ElUpload, ElIcon, ElButton } from 'element-plus';
import { Check } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';

const getShowId = () => {
  return 1
}

const router = useRouter()
// 跳转到新开页面并展示详情
const show = (id) => {
  console.log(id)
  router.push('/json/show')
}

</script>

<template>
  <page-container title="JSON解析">

    <!-- 上传文件组件 -->
    <!-- action ： 请求的URL -->
      <el-upload
        class="upload-demo"
        drag
        action= #
        multiple
      >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <br>
        <br>
        <div class="el-upload__text">
          拖动JSON文件到此 或<em> 点此上传 </em>
        </div>
      </el-upload>

      <el-button type="success" @click="show" round>点此查看解析详情</el-button>
  </page-container>
</template>

<style lang="scss" scoped>

</style>



