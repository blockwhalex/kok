<script lang="ts" setup>
import {ref} from 'vue';
import axios from 'axios';

const email = ref('') //邮箱
const password = ref('')//密码
const msg = ref('')//报错提示
const backgroundPosition = ref('0')
const login = async () => {
   if (email.value.length != 0 && password.value.length > 6){
    if (validateEmail(email.value)){
        const isLoginIn = await axiosLogin(email.value,password.value);
        if (isLoginIn.data.code != 200){
            return msg.value = isLoginIn.data.msg
        }else{
            return msg.value = "登录成功"
        }
    }else{
        return msg.value = "邮箱格式错误"
    }
   }else{
    return msg.value = "请输入邮箱和6位以上密码！"
   }
    
};

//校验邮箱格式是否正确
const validateEmail = (email: string): boolean => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
};

//axios异步请求
const axiosLogin = (email :string ,password :string) => {
    const formData = new FormData();
    formData.append("telephone",email);
    formData.append("password",password);
    const response =  axios.post('http://127.0.0.1:8080/auth/login',formData);
    return response;//后端返回的登录信息
};

 var moveBackground = () => {
     backgroundPosition.value = "35vw";
};
</script>
<template>
    <div class="content">
        
        <div class="login-wrapper">
            <div class="background" :style="{ left: backgroundPosition }">
            </div>
            <div class="left-register-form">
            </div>
            <div class="right-login-form">
                <button @click="moveBackground">click</button>
            </div>

        </div>
    </div>
</template>