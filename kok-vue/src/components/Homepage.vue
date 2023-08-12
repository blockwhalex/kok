<script lang="ts" setup>

import { ref, onMounted } from 'vue';
import VanillaTilt from 'vanilla-tilt';
const value = ref(new Date());

const cardRef = ref<HTMLElement[]>([]);

onMounted(() => {
  VanillaTilt.init(cardRef.value, {
    max: 15,
    speed: 400,
    glare: true,
    'max-glare': 1
  });
});

</script>
<template>

      <!-- 第一部分 -->

            <div class="card" ref="cardRef">
            <div class="content" >
                <h2>03</h2>
                <h3>Card One</h3>
                <p>Realistic glass card hover effect, realistic glass card hover effect, realistic glass card hover effect.</p>
            </div>
            </div>

        

        <!-- 第二部分 -->
        <div class="minicard">
            <div class="clockin">
            <h1>签到</h1>
            </div>
            <div class="integral">
            <h1>积分</h1>
            </div>
        </div>

        <!-- 第三部分 -->
        <!-- <div class="calendar">
        <div class="calendar-header">
            <span class="month-picker" id="month-picker">February</span>
            <div class="year-picker">
                <span class="year-change" id="prev-year">
                    <pre>&lt;</pre>
                </span>
                <span id="year">2021</span>
                <span class="year-change" id="next-year">
                    <pre>></pre>
                </span>
            </div>
        </div>
        <div class="calendar-body">
            <div class="calendar-week-day">
                <div>周日</div>
                <div>周一</div>
                <div>周二</div>
                <div>周三</div>
                <div>周四</div>
                <div>周五</div>
                <div>周六</div>
            </div>
            <div class="calendar-days"></div>
        </div>
        <div class="calendar-footer">
            <div class="toggle">
                <span>夜间模式</span>
                <div class="dark-mode-switch">
                    <div class="dark-mode-switch-ident"></div>
                </div>
            </div>
        </div>
        <div class="month-list"></div>
    </div> -->
    

</template>
<style>
@import url(../assets/css/homepage.css);
@import url(../assets/css/calendar.css);

</style>