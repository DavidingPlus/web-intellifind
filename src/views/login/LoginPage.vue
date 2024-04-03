<template>
    <div class="login-body">
        <div class="bg"></div>
        <div class="login-panel">
            <el-tabs type = "border-card">
                <el-tab-pane label="登录">
                    <el-form
                    class="login-form"
                    :model="formData"
                    :rules="rules"
                    ref="formDataRef"
                    @submit.prevent
                    >
                    <div class="login-title">服创</div>
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
                    </el-form-item>

                    <el-form-item>
                    <div class="remenberme-panel">
                        <el-checkbox v-model="formData.rememberMe">记住我</el-checkbox>
                    </div>
                    </el-form-item>

                    <el-form-item>
                    <el-button type="primary" class="op-btn" @click="doSubmit" size="large">登录</el-button>
                    </el-form-item>
                    </el-form>
                </el-tab-pane>
                
                <el-tab-pane label="注册">
                    <el-form
                    class="login-form"
                    :model="formData2"
                    :rules="rules"
                    ref="formDataRef2"
                    @submit.prevent
                    >
                    <div class="login-title">服创</div>
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
                    
                    <el-form-item>
                        <el-button type="primary" class="op-btn" @click="doRegis" size="large">注册</el-button>
                    </el-form-item>
                    </el-form>
                </el-tab-pane>

                <el-tab-pane label="修改密码">
                    <el-form
                    class="login-form"
                    :model="formData3"
                    :rules="rules"
                    ref="formDataRef3"
                    @submit.prevent
                    >
                    <div class="login-title">服创</div>
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
                </el-tab-pane>
            </el-tabs>
        </div>
    </div>
</template>

<script setup>
import { use } from "echarts";
import {ref,reactive,getCurrentInstance,nextTick,onMounted} from "vue";
import { useUserStore } from "../../stores/modules/user";
const {proxy} = getCurrentInstance();
import { useRouter } from "vue-router";
const router = useRouter(); 
const userStore = useUserStore();

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
        Object.assign(params,{emial:formData3.value.email});
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
        if(result1.data == 1)
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

<style lang="scss" scoped>
.login-body
{
    height: calc(100vh);
    background: url("../assets/backg.webp");
    background-size: cover;
    display: flex;
    .bg
    {
        flex: 1;
        background-size:cover;
        background-position: center;
        background-size:800px;
        background-repeat: no-repeat;
        background-image: url("../assets/tou.jpg");
    }
    .login-panel
    {
        width:430px;
        margin-right:15%;
        margin-top: calc((100vh - 500px) / 2);

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

        .check-code-panel
        {
            width: 100%;
            display: flex;
            .check-code
            {
                margin-left: 5px;
                cursor: pointer;
            }
        }
    }


}

</style>