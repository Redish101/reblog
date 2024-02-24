<template>
  <h2>登录</h2>
  <el-form :model="form" label-width="120px" class="login_form">
    <el-form-item label="用户名">
      <el-input v-model="form.username" />
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="form.password" type="password" />
    </el-form-item>
    <el-form-item>
      <el-button @click="submit(form)" type="primary"> 登录 </el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import api from "@/api";
import router from "@/router";
import { ElMessage } from "element-plus";
import "element-plus/theme-chalk/el-message.css"; // 为什么要加一行如此突兀的css?因为element的自动导入不会导入message组件的css, WTF?
import { reactive } from "vue";

api.site.siteList().then(res => {
  console.log(res)
})

interface RealForm {
  username: string;
  password: string;
}

const form = reactive<RealForm>({
  username: "",
  password: "",
});

const submit = async (form: RealForm) => {
  if (form.username == "" || form.password == "") {
    ElMessage.error("用户名或密码不能为空");

    return false;
  }

  api.admin
    .loginCreate({ username: form.username, password: form.password })
    .then((res) => res.json())
    .then((data) => {
      const token = data.data.token;
      api.setSecurityData(token);
      localStorage.setItem("token", token);
      ElMessage.success("登录成功");
      router.push("/");
    })
    .catch((err) => {
      if (err.status == 401) {
        ElMessage.error("用户名或密码错误");
      }
      return false;
    });
};
</script>

<style>
.login_form {
  max-width: 500px;
}
</style>
