<template>
  <div class="common-layout">
    <van-swipe style="height: 100%;" vertical>
      <van-swipe-item>1</van-swipe-item>
      <van-swipe-item>2</van-swipe-item>
      <van-swipe-item>3</van-swipe-item>
      <van-swipe-item>4</van-swipe-item>
    </van-swipe>
  </div>
</template>
<script setup lang="ts">
const onChange = (index) => showToast('当前 Swipe 索引：' + index);
</script>
<style scoped>
 .common-layout{
   width: 100%;
   height: 100%;
 }
</style>
