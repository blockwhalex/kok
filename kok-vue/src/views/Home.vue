<script lang="ts" setup>
import Homepage from '../components/Homepage.vue'
import Servepage from '../components/Servepage.vue'
import {ref , shallowRef} from 'vue';

const currentComponent = shallowRef(Homepage)
const modetext = ref('Light mode')
const isDarkMode = ref(false);
const isClose = ref(true);
const closebar = () =>{
    isClose.value = !isClose.value
};

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

const changeComponent = (id : number) => {
    if (id == 1) {
        currentComponent.value = Homepage
    }else if (id == 2) {
        currentComponent.value = Servepage
    }
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
                    <span class="name">Codinglab</span>
                    <span class="profession">Web developer</span>
                </div>
            </div>

            <i class='bx bx-chevron-right toggle' @click="closebar"></i>
        </header>

        <div class="menu-bar">
            <div class="menu">

                <li class="search-box">
                    <i class='bx bx-search icon'></i>
                    <input type="text" placeholder="Search...">
                </li>

                <ul class="menu-links">
                    <li class="nav-link">
                        <a href="#" @click="changeComponent(1)">
                            <i class='bx bx-home-alt icon' ></i>
                            <span class="text nav-text" >Home</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a href="#" @click="changeComponent(2)">
                            <i class='bx bx-bar-chart-alt-2 icon' ></i>
                            <span class="text nav-text" >Serve</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a href="#">
                            <i class='bx bx-bell icon'></i>
                            <span class="text nav-text">Notifications</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a href="#">
                            <i class='bx bx-pie-chart-alt icon' ></i>
                            <span class="text nav-text">Analytics</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a href="#">
                            <i class='bx bx-heart icon' ></i>
                            <span class="text nav-text">Likes</span>
                        </a>
                    </li>

                    <li class="nav-link">
                        <a href="#">
                            <i class='bx bx-wallet icon' ></i>
                            <span class="text nav-text">Wallets</span>
                        </a>
                    </li>

                </ul>
            </div>

            <div class="bottom-content">
                <li class="">
                    <a href="#">
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
            <!-- <components :is="currentComponent"></components> -->
            <currentComponent />
        </div>
    </section>
    </div>
</template>

<style>
@import url('../assets/css/saidebar_style.css');
@import url('https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css');
</style>