<template>
  <el-menu
    default-active=""
    router
    class="side_menu"
    background-color="#545c64"
    text-color="#fff"
  >
    <el-menu-item index="">
      <el-icon><HomeFilled /></el-icon>
      <span> 首页 </span>
    </el-menu-item>
  </el-menu>
</template>

<script lang="ts" setup>
import { HomeFilled } from "@element-plus/icons-vue";
</script>

<style>
.side_menu {
  height: 100%;
}
</style>
