<template>
  <div class="common-layout">
    <el-container>
      <el-header style="background: #333333">
        <el-row  justify="space-between"  align="middle">
          <el-col span="6">
            <div class="mt5 tl">
              <el-row justify="space-between" align="middle">
                <el-col span="12">
                  <el-image :src="Logo" style="width: 50px; height: 50px"/>
                </el-col>
                <el-col span="12">
                  <div class="ml10 white bold">刷视频</div>
                </el-col>
              </el-row>
            </div>
          </el-col>
          <el-col span="5">
            <div class="ml10">
              <el-input
                  v-model="state"
                  size="large"
                  style="width: 500px;"
                  placeholder="请输入感兴趣的关键词"
                  suffix-icon="Search"
              />
            </div>
          </el-col>
          <el-col span="5">
            <el-button type="primary">
              <el-icon class="el-icon--left"><Position /></el-icon>
              登录
            </el-button>
          </el-col>
        </el-row>
      </el-header>
      <el-container>
        <el-aside width="200px"  style="height: 100%;background:#545c64">
          <el-menu
              default-active="2"
              class="el-menu-vertical-demo"
              @open="handleOpen"
              @close="handleClose"
              background-color="#545c64"
              style="height: 93.5vh"
              text-color="#fff"
              active-text-color="#ffd04b">
            <el-menu-item index="1">
              <el-icon><House /></el-icon>
              <span slot="title">首页</span>
            </el-menu-item>
            <el-menu-item index="2">
              <el-icon><Star /></el-icon>
              <span slot="title">推荐</span>
            </el-menu-item>
            <el-menu-item index="3" >
              <el-icon><Ship /></el-icon>
              <span slot="title">朋友</span>
            </el-menu-item>
            <el-menu-item index="4">
              <el-icon><User /></el-icon>
              <span slot="title">我的</span>
            </el-menu-item>
            <el-menu-item index="5">
              <el-icon><Discount /></el-icon>
              <span slot="title">业务合作</span>
            </el-menu-item>
            <el-menu-item index="6">
              <el-icon><MessageBox /></el-icon>
              <span slot="title">直播</span>
            </el-menu-item>
            <el-menu-item index="7">
              <el-icon><Fries /></el-icon>
              <span slot="title">热点</span>
            </el-menu-item>
            <el-menu-item index="8">
              <el-icon><ShoppingTrolley /></el-icon>
              <span slot="title">购物</span>
            </el-menu-item>
            <el-menu-item index="9">
              <el-icon><Baseball /></el-icon>
              <span slot="title">游戏</span>
            </el-menu-item>
            <el-menu-item index="10">
              <el-icon><Pointer /></el-icon>
              <span slot="title">娱乐</span>
            </el-menu-item>
            <el-menu-item index="11">
              <el-icon><Coordinate /></el-icon>
              <span slot="title">二次元</span>
            </el-menu-item>
            <el-menu-item index="12">
              <el-icon><Headset /></el-icon>
              <span slot="title">音乐</span>
            </el-menu-item>
            <el-menu-item index="13">
              <el-icon><Dish /></el-icon>
              <span slot="title">美食</span>
            </el-menu-item>
            <el-menu-item index="14">
              <el-icon><Football /></el-icon>
              <span slot="title">体育</span>
            </el-menu-item>
            <el-menu-item index="15">
              <el-icon><Mic /></el-icon>
              <span slot="title">时尚</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-main>
          <RouterView />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script setup lang="ts">
import Logo from '../src/assets/image/logo.png'
import {ref} from "vue";
import {RouterView} from "vue-router";
let state=ref('');
</script>
<style scoped>
.common-layout{
  width: 98vw;
  height: 90vh;
  background: #2c3e50;
}
</style>
