<script setup>
import { Delete, Edit } from '@element-plus/icons-vue';
import { ref } from 'vue';
// import { artDelChannelService, artGetChannelsService } from '../../api/json';
import ChannelEdit from './components/ChannelEdit.vue';
const channelList = ref([])
const loading = ref(false)
const dialog = ref()

// const getChannelList = async () => {
//   loading.value = true
//   const res = await artGetChannelsService()
//   channelList.value = res.data.data
//   loading.value = false
// }
// getChannelList()

const onDelChannel = async (row) => {
  await ElMessageBox.confirm('你确认要删除该分类么', '温馨提示', {
    type: 'warning',
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  })
  await artDelChannelService(row.id)
  ElMessage.success('删除成功')
  getChannelList()
}
const onEditChannel = (row) => {
  dialog.value.open(row)
}
const onAddChannel = () => {
  // dialog.value.open({})
  router.push('/json/solve')
}
const onSuccess = () => {
  getChannelList()
}
</script>

<template>
  <page-container title="图表展示">
    <el-table v-loading="loading" :data="channelList" style="width: 100%">
      <!-- <el-table-column type="index" label="序号" width="100"></el-table-column>
      <el-table-column prop="cate_name" label="分类名称"></el-table-column>
      <el-table-column prop="cate_alias" label="分类别名"></el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="{ row, $index }">
          <el-button
            :icon="Edit"
            circle
            plain
            type="primary"
            @click="onEditChannel(row, $index)"
          ></el-button>
          <el-button
            :icon="Delete"
            circle
            plain
            type="danger"
            @click="onDelChannel(row, $index)"
          ></el-button>
        </template>
      </el-table-column> -->
      
      <template #empty>
        <el-empty description="没有数据" style="height: 100%;"></el-empty>
      </template>
    </el-table>
  </page-container>
</template>

<style lang="scss" scoped></style>
