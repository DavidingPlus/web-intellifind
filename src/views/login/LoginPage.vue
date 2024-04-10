<template>
     <!-- 最外层的大盒子 -->
     <div class="body">
      <div class="box">

            <!-- 滑动盒子 -->
            <div class="pre-box" ref="preRef">
                <h1>Web IntelliFind</h1>
                <h1 class="title">智寻系统</h1>
                <div class="img-box">
                    <img :src="flag == true ? imgList[1] : imgList[0]" alt="" />
                </div>
            </div>

                             <!-- 注册盒子 -->
                             <div class="login-form">
                  <!-- 标题盒子 -->
                  <div class="title-box">
                    <h1>注册</h1>
                  </div>
                  <!-- 输入框盒子 -->
                    <el-form
                    class="login-container"
                    :model="formData2"
                    :rules="rules"
                    ref="formDataRef2"
                    @submit.prevent
                    >
                    <el-form-item  prop="username">
                    <el-input 
                    size="large"
                    clearable 
                    placeholder="请输入用户名" 
                    v-model.trim="formData2.username">
                    <template #prefix>
                        <span class="iconfont icon-yonghuming"></span>
                    </template>
                    </el-input>
                    </el-form-item>
                    <el-form-item  prop="email">
                    <div class="send-email-panel">
                        <el-input 
                        size="large"
                        clearable 
                        placeholder="请输入邮箱" 
                        v-model.trim="formData2.email">
                        <template #prefix>
                            <span class="iconfont icon-email"></span>
                        </template>
                        </el-input>
                        <el-form-item>
                            <el-button @click="sendEmail" class="send-mail-btn" size = "large" :disabled="countdown > 0">{{ countdown > 0 ? `${countdown}秒后重试` : '发送' }}</el-button>
                        </el-form-item>
                    </div>
                    </el-form-item>
                    <el-form-item  prop="verify_code">
                    <el-input 
                    size="large"    
                    clearable 
                    placeholder="请输入邮箱验证码" 
                    v-model.trim="formData2.verify_code">
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
                    placeholder="请输入密码" 
                    v-model.trim="formData2.password">
                    <template #prefix>
                        <span class="iconfont icon-password"></span>
                    </template>
                    </el-input>
                    </el-form-item>
                    
                    <!-- <el-form-item>
                        <el-button type="primary" class="op-btn" @click="doRegis" size="large">注册</el-button>
                    </el-form-item> -->
                    </el-form>
                

          <!-- 按钮盒子 -->
          <div class="btn-box">
            <el-button type="primary" class="op-btn" @click="doRegis" size="large">注册</el-button>
            <!-- 绑定点击事件 -->
            <p @click="mySwitch">已有账号?去登录</p>

          </div>
        </div>

        <!-- 登录盒子 -->
        <div class="register-form">
            <!-- 标题盒子 -->
            <div class="title-box">
                <h1>登录</h1>
            </div>
                    <el-form
                    class="login-container"
                    :model="formData"
                    :rules="rules"
                    label-width="5px"
                    ref="formDataRef"
                    @submit.prevent
                    >
                    <el-form-item  prop="email">
                    <el-input 
                    size="large"
                    clearable 
                    placeholder="请输入邮箱" 
                    v-model.trim="formData.email">
                    <template #prefix>
                        <span class="iconfont icon-email"></span>
                    </template>
                    </el-input>
                    </el-form-item>

                    <el-form-item  prop="password">
                    <el-input 
                    size="large"    
                    clearable 
                    show-password
                    placeholder="请输入密码" 
                    v-model.trim="formData.password">
                    <template #prefix>
                        <span class="iconfont icon-password"></span>
                    </template>
                    </el-input>
                    </el-form-item>

                    <el-form-item  prop="answer">
                        <div class="check-code-panel">
                            <el-input 
                            size="large"    
                            clearable 
                            placeholder="请输入图形验证码" 
                            v-model.trim="formData.answer">
                            <template #prefix>
                                <span class="iconfont icon-password"></span>
                            </template>
                            </el-input>
                            <img :src="checkCodeUrl" class="check-code" @click="changeCheckCode"/>
                        </div>          
                    </el-form-item>

                    <el-form-item>
                    <div class="remenberme-panel">
                        <el-checkbox v-model="formData.rememberMe">记住我</el-checkbox>
                    </div>
                    </el-form-item>
                    </el-form>
                

           <!-- 按钮盒子 -->
          <div class="btn-box">
            <el-button type="primary" class="op-btn" @click="doSubmit" size="large">登录</el-button>
            <!-- 绑定点击事件 -->
            <p @click="mySwitch">没有账号?去注册</p>
          </div>
        </div>

        </div>
    </div>
</template>

<script setup>
import {ref,getCurrentInstance,nextTick,onMounted} from "vue";
import { useUserStore } from "../../stores/modules/user";
const {proxy} = getCurrentInstance();
import { useRouter } from "vue-router";

import registerImg from "@/assets/logo2.png";
import loginImg from "@/assets/logo.png";

const router = useRouter(); 
const userStore = useUserStore();

const preRef = ref('')
const imgList = ref([ registerImg, loginImg ])
  let flag = ref(false)
  /* 登录框与注册框滑动切换效果 */
  const mySwitch = () => {
    if (flag.value) {
      preRef.value.style.background = 'rgb(191, 227, 241)'
      preRef.value.style.transform = 'translateX(100%)'
    } else {
      preRef.value.style.background = 'rgb(247, 243, 209)'
      preRef.value.style.transform = 'translateX(0%)'
    }
    flag.value = !flag.value
  }

const api = 
{
    checkCode:"/capatcha-code",
    login:"/login",
    signup:"/signup",
    sendcode:"/send-code",
    reset:"/reset",
    isexist:"/is-exist"
}

const formData = ref({});
const formDataRef = ref();
const rules=
{
    username:[{required: true, message:"请输入用户名"}],
    email: [{required: true, message:"请输入邮箱"}],
    password: [{required: true, message:"请输入密码"}],
    answer:[{required: true, message:"请输入图形验证码"}],
    verify_code:[{required: true, message:"请输入邮箱验证码"}],
}

//checkcode
const checkCodeUrl = ref(null);
//?time=${new Date().getTime()}
let  captcha_id ="";
const changeCheckCode = async ()=>
{
    let params =
    {
        time:new Date().getTime()
    };
    let result = await proxy.Request(
            {
                url:api.checkCode,
                params,
            });
            if(result)
            {
                checkCodeUrl.value = result.captcha_image;
                captcha_id = result.captcha_id;
            }
}

//login
const doSubmit = () =>
{
    formDataRef.value.validate(async(valid) =>
    {
        if(!valid)
        {
            return;
        }
        let params = {};
        Object.assign(params,formData.value,{captcha_id:captcha_id});
        let cookieLoginInfo = proxy.VueCookies.get("loginInfo");
        let cookiePassword = cookieLoginInfo == null? null:cookieLoginInfo.password;
        // if(params.password !== cookiePassword)
        // {
        //     params.password = md5(params.password);
        // }
        let result = await proxy.Request(
            {
                url:api.login,
                params:JSON.stringify(params),
                dataType:'json',
                // params,
                errorCallback:()=>
                {
                    changeCheckCode();
                },
            });
        if(!result)
        {
            return;
        }
        if(params.rememberMe)
        {
            const loginInfo = 
            {
                email:params.email,
                password:params.password,
                rememberMe:params.rememberMe,
            };
            proxy.VueCookies.set("loginInfo",loginInfo,"7d");
        }else
        {
            proxy.VueCookies.remove("loginInfo");
        }
        proxy.Message.success("登陆成功");
        sessionStorage.setItem("userInfo",JSON.stringify(result.data));
        // 将Token存入userStore
        console.log(result.token);
        userStore.setToken(result.token);
        console.log(userStore.token);
        router.push("/user/profile");
    });
};

onMounted(()=>
{
    init();
})

const init = ()=>
{
    nextTick(()=>   
    {
        changeCheckCode();
        formDataRef.value.resetFields();
        formData.value = {};
        const cookieLoginInfo =  proxy.VueCookies.get("loginInfo");
        if(cookieLoginInfo)
        {
            formData.value = cookieLoginInfo;
        }
        
    })
}

const formData2 = ref({});
const formDataRef2 = ref();
//regist
const countdown = ref(0);
const sendEmail = ()=>  
{
    formDataRef2.value.validateField("email",async(valid) =>
    {
        if(!valid)
        {
            return;
        }
        if (countdown.value === 0) 
        {
            countdown.value = 60;
            const intervalId = setInterval(() => 
            {
            if (countdown.value > 0) 
            {
                countdown.value--;
            } else 
            {
                clearInterval(intervalId);
            }
            }, 1000);
        }
        let params = {};
        Object.assign(params,{email:formData2.value.email});
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
        if(result1.code == 2)
        {
            let result2 = await proxy.Request(
            {
                url:api.sendcode,
                params:JSON.stringify(params),
                dataType:'json',
            });
            if(!result2)
            {
                return;
            }
            proxy.Message.success("验证码发送成功，请回邮箱查看");
        }else
        {
            proxy.Message.warning("邮箱已注册，请重新输入邮箱");
        }
        
    });
};

const doRegis = ()=>
{
    formDataRef2.value.validate(async(valid)=>
    {
        if(!valid)
        {
            return;
        }
        let params = {};
        Object.assign(params,formData2.value);
        let result = await proxy.Request(
            {
                url:api.signup,
                params:JSON.stringify(params),
                dataType:'json',
            });
        if(!result)
        {
            return;
        }
        proxy.Message.success("注册成功");
        router.push("/login");
    });
};

// const formData3 = ref({});
// const formDataRef3 = ref();
// //change
// const countdown2 = ref(0);
// const sendEmail2 = ()=>  
// {
//     formDataRef3.value.validateField("email",async(valid) =>
//     {
//         if(!valid)
//         {
//             return;
//         }
//         if (countdown2.value === 0) 
//         {
//             countdown2.value = 60;
//             const intervalId = setInterval(() => 
//             {
//             if (countdown2.value > 0) 
//             {
//                 countdown2.value--;
//             } else 
//             {
//                 clearInterval(intervalId);
//             }
//             }, 1000);
//         }
//         let params = {};
//         Object.assign(params,{email:formData3.value.email});
//         let result1 = await proxy.Request(
//             {
//                 url:api.isexist,
//                 params:JSON.stringify(params),
//                 dataType:'json',
//             });
//         if(!result1)
//         {
//             return;
//         }
//         if(result1.code == 1)
//         {
//             let result = await proxy.Request(
//             {
//                 url:api.sendcode,
//                 params:JSON.stringify(params),
//                 dataType:'json',
//             })
//             if(!result)
//             {
//                 return;
//             }
//         }else
//         {
//             proxy.Message.warning("邮箱未注册，请输入正确邮箱");
//         }
//     });
// };
// const doChange = ()=>
// {
//     formDataRef3.value.validate(async(valid)=>
//     {
//         if(!valid)
//         {
//             return;
//         }
//         let params = {};
//         Object.assign(params,formData3.value);
//         let result = await proxy.Request(
//             {
//                 url:api.reset,
//                 params:JSON.stringify(params),
//                 dataType:'json',
//             });
//         if(!result)
//         {
//             return;
//         }
//         proxy.Message.success("修改成功");
//     })
// };
</script>

<style lang="less" scoped>
/* 去除input的轮廓 */
input {
  outline: none;
}

.body {
  /* 溢出隐藏 */
  height: 100vh;
  overflow-x: hidden;
  display: flex;
  /* 渐变方向从左到右 */
  background: linear-gradient(to right, rgb(247, 243, 209), rgb(191, 227, 241));
}

/* 最外层的大盒子 */
.box {
  width: 1050px;
  height: 600px;
  display: flex;
  /* 相对定位 */
  position: relative;
  z-index: 2;
  margin: auto;
  /* 设置圆角 */
  border-radius: 8px;
  /* 设置边框 */
  border: 1px solid rgba(255, 255, 255, 0.6);
  /* 设置盒子阴影 */
  box-shadow: 2px 1px 19px rgba(0, 0, 0, 0.1);
}

/* 滑动的盒子 */
.pre-box {
  /* 宽度为大盒子的一半 */
  width: 50%;
  /* width: var(--width); */
  height: 100%;
  /* 绝对定位 */
  position: absolute;
  /* 距离大盒子左侧为0 */
  left: 0;
  /* 距离大盒子顶部为0 */
  top: 0;
  z-index: 99;
  border-radius: 4px;
  background-color: rgb(247, 243, 209);
  box-shadow: 2px 1px 19px rgba(0, 0, 0, 0.1);
  /* 动画过渡，先加速再减速 */
  transition: 0.5s ease-in-out;
}

/* 滑动盒子的标题 */
.pre-box h1 {
  margin-top: 150px;
  text-align: center;
  /* 文字间距 */
  letter-spacing: 5px;
  color: white;
  /* 禁止选中 */
  user-select: none;
  /* 文字阴影 */
  text-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
}

/* 滑动盒子的文字 */
.pre-box .title{
  height: 30px;
  line-height: 30px;
  text-align: center;
  margin: 20px 0;
  /* 禁止选中 */
  user-select: none;
  font-weight: bold;
  color: white;
  text-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
}

/* 图片盒子 */
.img-box {
  width: 200px;
  height: 200px;
  margin: 20px auto;
  /* 设置为圆形 */
  border-radius: 50%;
  /* 设置用户禁止选中 */
  user-select: none;
  overflow: hidden;
  box-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
}

/* 图片 */
.img-box img {
  width: 100%;
  transition: 0.5s;
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
  .check-code-panel
        {
            width: 100%;
            display: flex;
            .check-code
            {   
                width:150px;
                margin-left: 5px;
                cursor: pointer;
            }
        }

/* 登录和注册盒子 */
.login-form,
.register-form {
  flex: 1;
  height: 100%;
}

/* 登录盒子 */
.login-container {
  width: 500px;
  background-color: rgba(255, 255, 255, 0);
  border-radius: 15px;
  padding: 0px 35px 15px 35px;
  margin: 0px auto;
  :deep(.el-input__wrapper) {
    background-color: rgba(255, 255, 0, 0);
    border: 1px solid #000;
  }
}

/* 标题盒子 */
.title-box {
  height: 200px;
  line-height: 200px;
}

/* 标题 */
.title-box h1 {
  text-align: center;
  color: white;
  /* 禁止选中 */
  user-select: none;
  letter-spacing: 5px;
  text-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
}

/* 输入框盒子 */
.el-form {
  display: flex;
  /* 纵向布局 */
  flex-direction: column;
  /* 水平居中 */
  align-items: center;
}

/* 输入框 */
.el-form-item {
  width: 65%;
  height: 40px;
}

input {
  /* width: 60%; */
  /* height: 40px; */
  margin-bottom: 20px;
  text-indent: 10px;
  border: 1px solid #fff;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 120px;
  /* 增加磨砂质感 */
  backdrop-filter: blur(10px);
}

input:focus {
  /* 光标颜色 */
  color: #b0cfe9;
}

/* 聚焦时隐藏文字 */
input:focus::placeholder {
  opacity: 0;
}

/* 按钮盒子 */
.btn-box {
  display: flex;
  justify-content: center;
}

/* 按钮 */
button {
  width: 100px;
  height: 30px;
  margin: 0 7px;
  line-height: 30px;
  border: none;
  border-radius: 4px;
  background-color: #69b3f0;
  color: white;
}

/* 按钮悬停时 */
button:hover {
  /* 鼠标小手 */
  cursor: pointer;
  /* 透明度 */
  opacity: 0.8;
}

/* 按钮文字 */
.btn-box p {
  height: 30px;
  line-height: 30px;
  /* 禁止选中 */
  user-select: none;
  font-size: 14px;
  color: white;
}

.btn-box p:hover {
  cursor: pointer;
  border-bottom: 1px solid white;
}
</style>
