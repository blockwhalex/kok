<script lang="ts" setup>
import Homepage from '../components/Homepage.vue'
import Servepage from '../components/Servepage.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {ref , shallowRef , onMounted} from 'vue';
import router from '@/router';
import axios from 'axios';
import Cookies from 'js-cookie';

const currentComponent = shallowRef(Homepage)
const modetext = ref('Light mode')
const isDarkMode = ref(false);
const isClose = ref(true);
const username = ref('');
const email = ref('');
const tabid = ref(1);

onMounted(() => {
      axiosInfo(); // 在页面加载时调用axiosInfo函数
    });

const axiosInfo = () => {
  const token = Cookies.get('token');
  axios.get('http://127.0.0.1:8080/auth/info',{
    headers: {
      Authorization: `Bearer ${token}`, // 添加Bearer Token到请求头
    },
  }).then(response => {
    username.value = response.data.data.user.name;
    email.value = response.data.data.user.email;
    console.log(response.data.data.user);
  }).catch(error => {
    console.log(error);
  });
};

//侧边栏隐藏
const closebar = () =>{
    isClose.value = !isClose.value
};

//亮色模式&黑夜模式
const toggleDarkMode = () =>{
    if (isDarkMode.value){
        document.body.classList.remove('dark');
        isDarkMode.value = !isDarkMode.value
        modetext.value = 'Light mode'
    }else{
        document.body.classList.add('dark');
        isDarkMode.value = !isDarkMode.value
        modetext.value = 'Dark mode'
    }

};

//侧边栏跳转
const changeComponent = (id : number) => {
    debugger
    tabid.value = id
    if (id == 1) {
        currentComponent.value = Homepage
    }else if (id == 2) {
        currentComponent.value = Servepage
    }
};

//用户登出
const logoutpage = () =>{
    ElMessageBox.confirm(
    '这将会退出当前账号，确定登出吗？',
    '登出确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
        localStorage.removeItem('isLoggedIn');
        router.replace('/login')
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消登出',
      })
    })
};


</script>
<template>
    <div class="home-wapper">
    <nav :class="{sidebar : true , close : isClose}">
        <header>
            <div class="image-text">
                <span class="image">
                    <img src="../assets/images/logo.png" alt="">
                </span>

                <div class="text logo-text">
                    <span class="name">姓名:{{ username }}</span>
                    <span class="profession">邮箱:{{ email }}</span>
                </div>
            </div>

            <i class='bx bx-chevron-right toggle' @click="closebar"></i>
        </header>

        <div class="menu-bar">
            <div class="menu">

                <!-- 搜索框暂时没用上 -->
                <!-- <li class="search-box">
                    <i class='bx bx-search icon'></i>
                    <input type="text" placeholder="Search..." disabled>
                </li> -->

                <ul class="menu-links">
                    <li class="nav-link">
                        <a @click="changeComponent(1)" :class="tabid===1? 'selected':''">
                            <i class='bx bx-home-alt icon' ></i>
                            <span class="text nav-text" >Home</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a @click="changeComponent(2)" :class="tabid===2? 'selected':''">  
                            <i class='bx bx-bar-chart-alt-2 icon' ></i>
                            <span class="text nav-text" >Serve</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a @click="changeComponent(3)" :class="tabid===3? 'selected':''">
                            <i class='bx bx-bell icon'></i>
                            <span class="text nav-text">Notifications</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a >
                            <i class='bx bx-pie-chart-alt icon' ></i>
                            <span class="text nav-text">Analytics</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a>
                            <i class='bx bx-heart icon' ></i>
                            <span class="text nav-text">Likes</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a>
                            <i class='bx bx-wallet icon' ></i>
                            <span class="text nav-text">Wallets</span>
                        </a>
                    </li>

                </ul>
            </div>

            <div class="bottom-content">
                <li class="">
                    <a @click="logoutpage">
                        <i class='bx bx-log-out icon' ></i>
                        <span class="text nav-text">Logout</span>
                    </a>
                </li>

                <li class="mode" >
                    <div class="sun-moon">
                        <i class='bx bx-sun icon sun'></i>
                        <i class='bx bx-moon icon moon'></i>   
                    </div>
                    <span class="mode-text text">{{ modetext }}</span>

                    <div class="toggle-switch">
                        <span class="switch" @click="toggleDarkMode"></span>
                    </div>
                </li>
                
            </div>
        </div>

    </nav>
    <section class="home">
        <div class="text">
            <component :is="currentComponent"></component>
            <!-- <currentComponent /> -->
        </div>
    </section>
    </div>
</template>

<style>
@import url('../assets/css/saidebar_style.css');
@import url('https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css');
.selected{
    background-color: var(--primary-color) !important;
}

.selected>span{
    color: #fff !important;
}

.selected>i{
    color: #fff !important;
}
</style>