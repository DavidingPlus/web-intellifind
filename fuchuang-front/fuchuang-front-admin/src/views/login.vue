<template>
    <div class="login-body">
        <div class="bg"></div>
        <div class="login-panel">
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
                    placeholder="请输入密码" 
                    v-model.trim="formData.password">
                    <template #prefix>
                        <span class="iconfont icon-password"></span>
                    </template>
                    </el-input>
                </el-form-item>

                <el-form-item  prop="checkcode_admin">
                    <div class="check-code-panel">
                        <el-input 
                        size="large"    
                        clearable 
                        placeholder="请输入验证码" 
                        v-model.trim="formData.checkcode_admin">
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

                <el-form-item>
                    <el-button type="primary" class="op-btn" @click="doSubmit" size="large">登录</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script setup>
import {ref,reactive,getCurrentInstance,nextTick,onMounted} from "vue";
const {proxy} = getCurrentInstance();
import md5 from "js-md5";
import { useRouter } from "vue-router";
import { routeLocationKey } from "vue-router";
const router = useRouter();

const api = 
{
    checkCode:"/capatcha-code",
    login:"/login"
}

const formData = ref({});
const formDataRef = ref();
const rules=
{
    email: [{required: true, message:"请输入邮箱"}],
    password: [{required: true, message:"请输入密码"}],
    checkcode_admin:[{required: true, message:"请输入验证码"}],
}

//checkcode
const checkCodeUrl = ref(null);
//?time=${new Date().getTime()}
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
            }
}

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
        // Object.assign(params,formData.value);
        let cookieLoginInfo = proxy.VueCookies.get("loginInfo");
        let cookiePassword = cookieLoginInfo == null? null:cookieLoginInfo.password;
        // if(params.password !== cookiePassword)
        // {
        //     params.password = md5(params.password);
        // }
        // params.password = md5(params.password);
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
        let firstMenu = result.data.menulist[0];
        if(firstMenu.children.length > 0)
        {
            firstMenu = firstMenu.children[0];
        }
        router.push(firstMenu.menuUrl);
    });
};

onMounted(()=>
{
    init();
})
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