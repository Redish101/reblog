import useApi from "@/utils/fetcher";
import {
  ProCard,
  ProFormInstance,
  ProFormText,
  StepsForm,
} from "@ant-design/pro-components";
import { message } from "antd";
import { useEffect, useRef } from "react";
import { history } from "umi";

interface Init {
  name: string;
  url: string;
  desc: string;
  icon: string;
  email: string;
  username: string;
  nickname: string;
  password: string;
}

const InitPage = () => {
  const refresh = async () => {
    const siteRes = await useApi("/api/site");
    const siteData = await siteRes.json();

    if (siteData["data"]) {
      message.open({
        type: "warning",
        content: "站点信息已存在，请勿重复初始化！",
      });

      history.push("/login");
    }
  };

  useEffect(() => {
    refresh();
  }, []);

  const fromRef = useRef<ProFormInstance<Init>>();

  const handleSubmit = async (values: Init) => {
    fromRef.current?.validateFields();

    const formData = new FormData();

    formData.append("name", values.name);
    formData.append("url", values.url);
    formData.append("desc", values.desc);
    formData.append("email", values.email);
    formData.append("username", values.username);
    formData.append("nickname", values.nickname);
    formData.append("password", values.password);

    const res = await useApi("/api/init", {
      method: "POST",
      body: formData,
    });

    const data = await res.json();

    if (res.status == 200) {
      message.open({
        type: "success",
        content: "站点初始化成功",
      });
      history.push("/login");
    } else {
      message.open({
        type: "error",
        content: `站点初始化失败: ${data["msg"]}`,
      });
    }
  };

  return (
    <div
      style={{ marginTop: "100px", marginInline: "auto", maxWidth: "600px" }}
    >
      <ProCard title="初始化站点" bordered>
        <StepsForm<Init> onFinish={handleSubmit} formRef={fromRef}>
          <StepsForm.StepForm title="管理员信息">
            <ProFormText
              label="用户名"
              name="username"
              rules={[{ required: true, message: "请输入用户名" }]}
            />
            <ProFormText
              label="昵称"
              name="nickname"
              rules={[{ required: true, message: "请输入昵称" }]}
            />
            <ProFormText
              label="邮箱"
              name="email"
              rules={[{ required: true, message: "请输入邮箱" }]}
            />
            <ProFormText.Password
              label="密码"
              name="password"
              rules={[{ required: true, message: "请输入密码" }]}
            />
          </StepsForm.StepForm>
          <StepsForm.StepForm title="站点信息">
            <ProFormText
              label="站点名称"
              name="name"
              rules={[{ required: true, message: "请输入站点名称" }]}
            />
            <ProFormText
              label="站点URL"
              name="url"
              rules={[{ required: true, message: "请输入站点URL" }]}
              tooltip="请输入前端URL而非本页面URL"
            />
            <ProFormText
              label="站点描述"
              name="desc"
              rules={[{ required: true, message: "请输入站点描述" }]}
            />
            {/* TODO: 站点图标上传 */}
          </StepsForm.StepForm>
        </StepsForm>
      </ProCard>
    </div>
  );
};

export default InitPage;
