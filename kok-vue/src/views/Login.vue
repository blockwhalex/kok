<script lang="ts" setup>
import {ref} from 'vue';
import axios from 'axios';
import { ElMessage } from 'element-plus'
import router from '@/router';
import Cookies from 'js-cookie';

const name = ref('') //用户名
const loginemail = ref('') //登录邮箱
const registeremail = ref('')//注册邮箱
const loginpassword = ref('')//登录密码
const registerpassword = ref('')//注册密码
const telephone = ref('')//手机号
const leftValue = ref('0')//控制遮罩位置
const isLoading = ref(false)//加载动画
const isDisabled = ref(false)//禁止重复提交请求

//用户登录
const login = async () => {
    isDisabled.value = true;
    isLoading.value = true;
   if (loginemail.value.length != 0 && loginpassword.value.length > 6){
    if (validateEmail(loginemail.value)){
        try{
            const isLoginIn = await axiosLogin(loginemail.value,loginpassword.value);
        if (isLoginIn.data.code != 200){
            isLoading.value = false;
            isDisabled.value = false;
            return ElMessage.error(isLoginIn.data.msg)
        }else{
            isLoading.value = false;
            isDisabled.value = false;
            ElMessage({
                        message : '登录成功',
                        type: 'success',
                    });
                    
                    Cookies.set("token",isLoginIn.data.data.token,{expires:7,path : '/'})
                    localStorage.setItem('isLoggedIn', 'true');
                    return setTimeout(() => {
                        router.replace('/home')
                    },1700);
            }
        }catch{
            isLoading.value = false;
            isDisabled.value = false;
            return ElMessage.error("登录失败，请检查网络")
        }
    }else{
        isLoading.value = false;
        isDisabled.value = false;
        return ElMessage.error("邮箱格式错误")
    }
   }else{
    isLoading.value = false;
    isDisabled.value = false;
    return ElMessage.error("邮箱或密码不完整")
   }
    
};

//用户注册
const register = async() =>{
    isDisabled.value = true;
    isLoading.value = true;
    if (registeremail.value.length > 0 && telephone.value.length == 11 && registerpassword.value.length > 6){
        if (validateEmail(registeremail.value)){
            try{
                const isRegister = await axiosregister(registeremail.value,registerpassword.value,name.value,telephone.value);
                if (isRegister.data.code != 200){
                    isLoading.value = false;
                    isDisabled.value = false;
                    return ElMessage.error(isRegister.data.msg)
                }else{
                    isLoading.value = false;
                    isDisabled.value = false;
                    ElMessage({
                        message : '注册成功',
                        type: 'success',
                    });
                      leftValue.value = '0'
                      return
                    }
            }catch{
                    isLoading.value = false;
                    isDisabled.value = false;
                    return ElMessage.error("注册失败，请检查网络")
            }
        }else{
            isLoading.value = false;
            isDisabled.value = false;
            return ElMessage.error("邮箱格式错误")
        }
    }else{
        isLoading.value = false;
        isDisabled.value = false;
        return ElMessage.error("必填信息不能为空")
    }
}

//校验邮箱格式是否正确
const validateEmail = (email: string): boolean => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
};



//axios登录请求
const axiosLogin = (email :string ,password :string) => {
    const formData = new FormData();
    formData.append("email",email);
    formData.append("password",password);
    const response =  axios.post('http://127.0.0.1:8080/auth/login',formData);
    return response;//后端返回的登录信息
};

//axios注册请求
const axiosregister = (email :string ,password :string ,name :string ,telephone :string) => {
    const formData = new FormData();
    formData.append("name",name);
    formData.append("telephone",telephone);
    formData.append("email",email);
    formData.append("password",password);
    const response =  axios.post('http://127.0.0.1:8080/auth/register',formData);
    return response;//后端返回的登录信息
};

//控制遮罩位置
 const movecontainer= () => {
    if (leftValue.value == '0') {
        return leftValue.value = "35vw";
    }else{
        return leftValue.value = '0';
    }
    
};
</script>
<template>
    <!-- 背景板 -->
    <div class="content">
        <div class="login-wrapper">
            <div class="container-change" :style="{ left : leftValue}">
            <div class="container">
                <div class="photo-container">
                    <div class="photo-cont-item animation-1">
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>

                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                    </div>

                    <div class="photo-cont-item animation-2">
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>

                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                    </div>

                    <div class="photo-cont-item animation-3">
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>

                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                    </div>

                    <div class="photo-cont-item animation-4">
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>

                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                    </div>

                    <div class="photo-cont-item animation-5">
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>

                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                        <div class="photo-item photo-1"></div>
                        <div class="photo-item photo-2"></div>
                        <div class="photo-item photo-3"></div>
                        <div class="photo-item photo-4"></div>
                    </div>
                </div><!-- .photo-container -->
                <div class="blur">
                    <div class="tips">
                        <div class="title">
                            THE KOK XRAY-UI
                        </div>
                        <h1>Welcome to kok</h1>
                        <span>Scan the QR code below for more information.</span>
                        <div class="QRcode">
                            <img src="../assets/images/QRcode.png" alt="QRcode">
                            <span>Use your phone to scan codes</span>
                        </div>
                        
                    </div>
                </div>
            </div><!-- .container -->
            </div>
            <div class="left-register-form">
                <div class="left-form-wrapper">
                    <h1>Sign Up</h1>
                    <div class="input-items">
                        <span class="input-tips">
                            Account Name
                        </span>
                        <input type="text" class="inputs" placeholder="Enter your name" v-model="name">
                    </div>
                    <div class="input-items">
                        <span class="input-tips">
                            Email Address
                        </span>
                        <input type="text" class="inputs" placeholder="Enter your email" v-model="registeremail">
                    </div>
                    <div class="input-items">
                        <span class="input-tips">
                            Phone Number
                        </span>
                        <input type="text" class="inputs" placeholder="Enter your telephone" v-model="telephone">
                    </div>
                    <div class="input-items">
                        <span class="input-tips">
                            Password
                        </span>
                        <input type="password" class="inputs" placeholder="Enter password" v-model="registerpassword">
                    </div>
                    <button class="btn" @click="register" :disabled="isDisabled">
                        <div v-if="isLoading" class="loadingbox">
                        <div class="loading2"></div>
                    </div>
                    <div v-else>
                        <span>Sign Up</span>
                    </div>  
                    </button>
                    <div class="login-tips">
                        <span>Already has an account?&nbsp;&nbsp;&nbsp;&nbsp;</span>
                        <span  class="LoginButton" @click="movecontainer">Login</span>
                    </div>
                </div>
            </div>
            <div class="right-login-form">
                <div class="right-form-wrapper">
                    <h1>Log in</h1>              
                    <div class="input-items">
                        <span class="input-tips">
                            Email Address
                        </span>
                        <input type="text" class="inputs" placeholder="Enter your email" v-model="loginemail">
                    </div>
                    <div class="input-items">
                        <span class="input-tips">
                            Password
                        </span>
                        <input type="password" class="inputs" placeholder="Enter password" v-model="loginpassword">
                    </div>
                    <span class="forgot">Forgot Password</span>

                    <button class="btn" @click="login" :disabled="isDisabled">
                        <div v-if="isLoading" class="loadingbox">
                        <div class="loading2"></div>
                    </div>
                    <div v-else>
                        <span>Log in</span>
                    </div>  
                    </button>
                    <div class="siginup-tips">
                        <span>Don't Have An Account?&nbsp;&nbsp;&nbsp;&nbsp;</span>
                        <span @click="movecontainer" class="SingupButton">Signup</span>
                    </div>
                    <div class="other-login">
                        <div class="divider">
                            <span class="line"></span>
                            <span class="divider-text">or</span>
                            <span class="line"></span>
                        </div>
                        <div class="other-login-wrapper">
                            <div class="other-login-item">
                                <img src="../assets/images/app_popular_logo.png" alt="popular">
                            </div>
                            <div class="other-login-item">
                                <img src="../assets/images/logo_whatsapp.png" alt="whatsapp">
                            </div>
                            <div class="other-login-item">
                                <img src="../assets/images/social_facebook_media.png" alt="facebook">
                            </div>
                            <div class="other-login-item">
                                <img src="../assets/images/social_app_media.png" alt="twitter">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style>
/* 加载动画 */
@import '../assets/css/loading2.css';
/*页面整体css*/
@import '../assets/css/style.css';
/*icon流转控制css*/
@import '../assets/css/test.css';
</style>