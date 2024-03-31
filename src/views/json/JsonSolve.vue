<script setup>
import { UploadFilled, Tickets } from '@element-plus/icons-vue';
import { ElButton, ElIcon, ElLoading, ElMessage, ElUpload } from 'element-plus';
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { uploadJsonFile, updateUserConfig } from '@/api/json';
import axios from 'axios';

const params = {
  file_name: ''
}
const uploadRef = ref()
// 获取上传的文件
const onSelectFile = (uploadFile) => {
  if (!uploadFile.raw) {
    console.error('未选择文件')
    return
  }
  uploadRef.value = uploadFile.raw
  onupload()
}

// 上传json文件函数
const upload = async (res) => {
  if (!uploadRef.value) {
    ElMessage.error('未选择文件')
    return
  }
  const formData = new FormData()
  formData.append('json_file', uploadRef.value)
  try {
    res = await axios.post('http://8.137.100.0:8080/core/upload-file', formData, {
          headers: {
            'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTMiLCJ1c2VyX25hbWUiOiJtYXgiLCJleHBpcmVfdGltZSI6MjQzMDE2MzQ1NSwiZXhwIjoyNDMwMTYzNDU1LCJpYXQiOjE3MTAxNjM0NTUsImlzcyI6InRlc3RfYmFja2VuZCIsIm5iZiI6MTcxMDE2MzQ1NX0.9cVo0XpzZiwlwLJ1oepX03LsPdh7XEIeqIFoBv2ZBgI'
          }
        })
    params.file_name = res.data.data.file_name
    console.log('上传文件成功:', params.file_name);
  } catch (error) {
    console.error('上传文件失败:', error)
    ElMessage.error('上传文件失败，请稍后重试')
  }
}

const flag = ref(true)
const percentage = ref(0.0)
const perFlag = ref(true)
const loading = ref(false)

// 进度函数
const progress = () => {
  percentage.value = 0.0
  setInterval(() => {
    if (perFlag.value) {
      percentage.value += 2.0
      if (percentage.value === 100) {
        perFlag.value = false
        setTimeout(() => {
          ElMessage.success('解析成功')
          loading.value = false
        }, 300)
    }
  }
}, 250)
}
// 上传文件
const onupload = () => {
  upload()
  setTimeout(() => {
    flag.value = false
    loading.value = true
    progress()
  }, 3000)
}

const router = useRouter()
// 跳转到新开页面并展示详情
const show = () => {
  console.log(params.file_name);
  router.push({ path: '/json/show', query: {file_name: params.file_name} })
}






// 用户配置表单
const formRef = ref()

const ratioList = [{
  name: 'stay_time',
  value: 5
  }, {
  name: 'repeat_click',
  value: 5
  }, {
  name: 'page_load',
  value: 5
  }, {
  name: 'feedback_interval',
  value: 5
  }, {
  name: 'no_reaction',
  value: 5
  }, {
  name: 'error_count',
  value: 5
  }, {
  name: 'console_errors',
  value: 5
  }, {
  name: 'is_blank',
  value: 5
  }, {
  name: 'occur_many',
  value: 5
  }
]
  
const formModel = ref({  
  items: [{
    name: '',
    value: ''
  }]  
});  

const count = ref(1);
const addItem = () => {  
  if (count.value >= 9) {  
    ElMessage.error('最多只能添加9个权重配置项');  
    return;  
  }
  formModel.value.items.push({ name: '', value: '' });  
  selectedItems.value.push({ name: '' });
  count.value ++;
};  
  
const removeItem = (index) => {  
  formModel.value.items.splice(index, 1);  
};  

// 上传用户配置信息
const submitForm = async () => {
  if (!formRef.value) {
    ElMessage.error('表单未初始化')
    return
  }
  // 等待校验结果
  await formRef.value.validate()
  // 提交修改
  await updateUserConfig(formModel.value)
  // 提示用户
  ElMessage.success('修改成功')
}

// 选中的选项
const selectedItems = ref(formModel.value.items 
    ? formModel.value.items.map(item => ({ name: item.name ?? '' })) 
    : [])


const handleSelectChange = (index) => {  
  // 获取当前选中的选项
  const selectedItem = selectedItems.value[index];
  const duplicateNames = selectedItems.value.filter((item, idx) => idx !== index && item.name === selectedItem.name);  
    
  if (duplicateNames.length > 0) {  
    ElMessage.error(`配置中不能存在重复的权重属性名：${selectedItem.name}`);  
    selectedItems.value[index].name = '' 
  }  else {
    selectedItems.value[index] = selectedItem // 更新选中的选项
  }

  // 更新表单项的校验规则
  const selectedName = selectedItems[index].name;
  const item = formModel.items[index];
  
  if (selectedName) {
    item.name = selectedName;
    item.rules = [{ required: false, message: '请输入权重值', trigger: 'change' }];
  } else {
    item.name = '';
    item.rules = [{ required: true, message: '请输入权重值', trigger: 'change' }];
  }
};

// 获取可选的权重属性(去除已经被选中的)
const getAvailableOptions = computed(() => {  
  return (index) => {  
    // 获取当前所有被选中的name，并过滤掉空字符串  
    const selectedNames = selectedItems.value.map(item => item.name)
                            .filter(name => name);
    // 如果没有被选中的数据，则返回ratioList中的所有数据  
    if (selectedNames.length === 0) {  
      return ratioList.map(option => ({ name: option.name}))
    }  
  
    // 否则，过滤 ratioList排除已经被选中的name  
    const excludedNames = [...selectedNames.slice(0, index), ...selectedNames.slice(index + 1)];
    return ratioList.map(option => ({ name: option.name})).filter(item => !excludedNames.includes(item.name));
  };  
});  

// 控制输入值的范围
const checkNumber = (item) => {
  const value = parseFloat(item.value);
  if (isNaN(value) || value < 0 || value > 10) {
    item.value = ''; // 清空输入框
    ElMessage.error('请输入0~10之间的数值');
  }
};

const iconStyle = {  
  fontSize: '15px',  
  marginRight: '10px'  
};

</script>

<template>
  <page-container title="JSON解析">

    <!-- 上传文件组件 -->
    <!-- action ： 请求的URL -->
    <!-- action= "8.137.100.0" -->
    <div style="position: relative; width: 100%; max-height:565px;">
        <el-upload
        type="file"
        class="upload-demo"
        drag
        multiple
        action="#"
        v-if="flag"
        :auto-upload="false"
        :on-change="onSelectFile"
      >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <br>
        <br>
        <div class="el-upload__text">
          拖动JSON文件到此 或<em> 点此上传 </em>
        </div>
      </el-upload>

      <div v-else="flag" class="demo-progress" 
      style="display: flex; flex-direction: column; 
      justify-content: center; align-items: center;">
        
        <!-- <el-loading content="文件上传中..."  />   -->
        <br>
        <el-progress
          :text-inside="true"
          :stroke-width="24"
          :percentage="percentage"
          status="success"
          style="width: 80%;"
        />
        <br>
        <el-table v-if="loading" v-loading="loading" style="width: 100%">
        </el-table>
        <el-button v-else="loading" type="success" @click="show" round>点此查看解析详情</el-button>
      </div>
    </div>
      <br>
      <br>




<el-form ref="formRef" :model="formModel" style="display: flex; flex-direction: column;">  

    <el-row>  
      <el-col :span="24">  
        <h2 style="text-align: left;">自定义配置信息</h2>  
      </el-col>  
    </el-row>  
    
    <el-descriptions
        class="margin-top"
        title="用户权重配置信息"
        :column="3"
        border>
        <el-descriptions-item v-for="(item, index) in ratioList" :key="index">
          <template #label>
            <div class="cell-item">
              <el-icon :style="iconStyle">
                <tickets />
              </el-icon>
              {{item.name}}
            </div>
          </template>
          <el-tag>{{item.value}}</el-tag>
        </el-descriptions-item>
      </el-descriptions>
  
      <br>
      <br>
    
    <el-form-item
    v-for="(item, index) in formModel.items"
    :key="index"
    :prop="'items.' + index + '.name'" 
    
  >
    <div style="display: flex; align-items: center;">
      <el-select v-model="selectedItems[index].name" 
      @change="handleSelectChange(index)" 
      placeholder="请选择权重属性">
        <el-option 
          v-for="(option, optionIndex) in getAvailableOptions(index)" 
          :key="optionIndex" 
          :label="option.name" 
          :value="option.name">
        </el-option>
      </el-select>
      <el-input 
        v-model="item.value" 
        type="number" 
        @input="checkNumber(item)" 
        placeholder="请输入0~10之间的数值"></el-input>
      <el-button type="danger" @click="removeItem(index)" class="cancel">删除</el-button>
    </div>
  </el-form-item>
  
    <el-form-item>  
      <el-button type="primary" @click="submitForm">保存配置</el-button>  
      <el-button @click="addItem">新增权重配置</el-button>  
    </el-form-item>  
  </el-form> 


  </page-container>
</template>

<style lang="scss" scoped>
.demo-progress .el-progress--line {
  margin-bottom: 15px;
  max-width: 600px;
}

/* 使用深度选择器来覆盖 Element UI 组件的内部样式 */  
::v-deep(.el-select .el-input__inner)  {  
  width: 200px; 
}  
::v-deep(.el-input__inner)  {
  width: 200px;
  margin-left: 10px;
}

.cancel {
  margin-left: 10px; /* 添加一些间距，使得按钮与 select/input 分开 */  
}

</style>



