<script setup>
import { reactive } from "vue";
import zhCn from 'element-plus/dist/locale/zh-cn.mjs';
const locale = zhCn;
const config = reactive
(
  {max: 1,}
);
</script>

<template>
<el-config-provider :locale="locale" :message="config">
  <RouterView />
</el-config-provider>
</template>

<style scoped>
</style>
