<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import { reactive } from "vue";
const config = reactive
(
{max: 1,}
);
</script>

<template>
  <div>
    <!-- App.vue只需要留一个路由出口 router-view即可 -->
    <el-config-provider :locale="zhCn" :message="config">
      <router-view></router-view>
    </el-config-provider>
  </div>
</template>

<style scoped></style>
