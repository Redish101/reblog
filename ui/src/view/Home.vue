<template>
  <h2>首页</h2>
</template>

<script lang="ts" setup>
import api from "@/api";
import router from "@/router";
import { token } from "@/utils/auth";

api.admin.tokenStateList(token()).then((res) => {
  if (!res.data.data) {
    router.push({ name: "login" });
  }
});
</script>
