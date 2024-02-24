<template>
  <el-container class="container">
    <!-- <el-header>
            <Header />
        </el-header> -->
    <el-aside>
      <NavBar />
    </el-aside>
    <el-main>
      <!-- <Header /> -->
      <div class="content">
        <slot />
      </div>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
// import Header from './Header.vue'
import NavBar from "./NavBar.vue";
</script>

<style>
.container {
  height: 100vh;
}

.content {
  padding-inline: 1rem;
}
</style>
