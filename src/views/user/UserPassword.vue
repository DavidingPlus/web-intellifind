<script setup>
import {ref,getCurrentInstance} from "vue";
const {proxy} = getCurrentInstance();

const api = 
{
    reset:"/reset",
    isexist:"/is-exist"
}

const formData3 = ref({});
const formDataRef3 = ref();
//change
const countdown2 = ref(0);
const sendEmail2 = ()=>  
{
    formDataRef3.value.validateField("email",async(valid) =>
    {
        if(!valid)
        {
            return;
        }
        if (countdown2.value === 0) 
        {
            countdown2.value = 60;
            const intervalId = setInterval(() => 
            {
            if (countdown2.value > 0) 
            {
                countdown2.value--;
            } else 
            {
                clearInterval(intervalId);
            }
            }, 1000);
        }
        let params = {};
        Object.assign(params,{email:formData3.value.email});
        let result1 = await proxy.Request(
            {
                url:api.isexist,
                params:JSON.stringify(params),
                dataType:'json',
            });
        if(!result1)
        {
            return;
        }
        if(result1.code == 1)
        {
            let result = await proxy.Request(
            {
                url:api.sendcode,
                params:JSON.stringify(params),
                dataType:'json',
            })
            if(!result)
            {
                return;
            }
        }else
        {
            proxy.Message.warning("邮箱未注册，请输入正确邮箱");
        }
    });
};
const doChange = ()=>
{
    formDataRef3.value.validate(async(valid)=>
    {
        if(!valid)
        {
            return;
        }
        let params = {};
        Object.assign(params,formData3.value);
        let result = await proxy.Request(
            {
                url:api.reset,
                params:JSON.stringify(params),
                dataType:'json',
            });
        if(!result)
        {
            return;
        }
        proxy.Message.success("修改成功");
    })
};
</script>

<template>
  <page-container title="修改密码">
    
    <el-row>
      <el-col :span="12">
        <el-card>
    <el-form
      class="login-form"
      :model="formData3"
      :rules="rules"
      ref="formDataRef3"
      @submit.prevent
      >
      <div class="login-title">Web智寻</div>
      <el-form-item  prop="email">
      <div class="send-email-panel">
          <el-input 
          size="large"
          clearable 
          placeholder="请输入邮箱" 
          v-model.trim="formData3.email">
          <template #prefix>
              <span class="iconfont icon-email"></span>
          </template>
          </el-input>
          <el-form-item>
              <el-button @click="sendEmail2" class="send-mail-btn" size = "large" :disabled="countdown2 > 0">{{ countdown2 > 0 ? `${countdown2}秒后重试` : '发送' }}</el-button>
          </el-form-item>
      </div>
      </el-form-item>
      <el-form-item  prop="verify_code">
      <el-input 
      size="large"    
      clearable 
      placeholder="请输入邮箱验证码" 
      v-model.trim="formData3.verify_code">
      <template #prefix>
          <span class="iconfont icon-youxiangyanzhengma"></span>
      </template>
      </el-input>
      </el-form-item>
      <el-form-item  prop="password">
      <el-input 
      size="large"    
      clearable
      show-password
      placeholder="请输入重置后密码" 
      v-model.trim="formData3.password">
      <template #prefix>
          <span class="iconfont icon-password"></span>
      </template>
      </el-input>
      </el-form-item>
      <el-form-item>
          <el-button type="primary" class="op-btn" @click="doChange" size="large">修改密码</el-button>
      </el-form-item>
      </el-form>
    </el-card>
    </el-col>
    </el-row>

  </page-container>
</template>

<style lang="scss" scoped>
.login-form
        {
            padding: 25px;
            background: #fff;
            border-radius: 5px;
            .login-title
            {
                text-align:center;
                font-size:18px;
                font-weight:bold;
                margin-bottom: 20px;
            }
            .send-email-panel
            {
                display: flex;
                width: 100%;
                justify-content: space-between;
                .send-mail-btn
                {
                    margin-left: 5px;
                }
            }
            .rememberme-panel
            {
                width: 100%;
            }
            .no-account
            {
                width: 100%;
                display: flex;
                justify-content: space-between;
            }
            .op-btn
            {
                width: 100%;
            }
        }
</style>