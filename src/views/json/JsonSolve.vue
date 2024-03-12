<script set up>
import {
  ElButton,
  ElForm,
  ElFormItem,
  ElInput,
  ElCard,
} from 'element-plus'
import { ref } from 'vue'
import ArticleEdit from './components/JsonEdit.vue'
const articleEditRef = ref()
const onAddArticle = () => {
  articleEditRef.value.open({})
}
const onSuccess = () => {
  getArticleList()
}
const content = ref('')

</script>

<template>
  <page-container title="JSON解析">

    <!-- 直接输入JSON文本进行解析 -->
    <!-- <template #extra>
      <el-button type="primary" @click="onAddArticle">添加解析</el-button>
    </template> -->
    <article-edit ref="articleEditRef" @success="onSuccess"></article-edit>

    <el-form>
    <el-input type="textarea" v-model="content" rows="5" placeholder="请输入JSON文本">
    </el-input>
    <el-form-item>
      <el-button type="primary" @click="onParse">解析</el-button>
    <br>
    <br>
    </el-form-item>
    
    <br>


    <!-- 通过传递JSON文件进行解析 -->
      <el-card>
        <el-form-item label="JSON文件解析">
        </el-form-item>
        <el-form-item>
          <el-input type="file" v-model="jsonFile" size="large" />
        </el-form-item>
        <br>
        <el-form-item>
          <el-button type="primary" @click="onParse">解析</el-button>
        </el-form-item>
      </el-card>
      </el-form>
    

  </page-container>
</template>

<style lang="scss" scoped>

</style>