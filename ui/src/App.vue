<template></template>

<script lang="ts" setup>
import api from "@/api/index";

api.article.articleDetail("test1").then((res) => {
  console.log(res.data);
});
</script>
