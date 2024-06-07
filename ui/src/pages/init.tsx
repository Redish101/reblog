import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { LoginForm, ProFormText } from "@ant-design/pro-components";

const Init = () => {
  return (
    <LoginForm title="登录" subTitle="reblog dashboard">
      <ProFormText
        name="username"
        placeholder="用户名"
        fieldProps={{
          size: "large",
          prefix: <UserOutlined />,
        }}
        rules={[{ required: true, message: "请输入用户名" }]}
      />
      <ProFormText.Password
        name="password"
        placeholder="密码"
        fieldProps={{ size: "large", prefix: <LockOutlined /> }}
        rules={[{ required: true, message: "请输入密码" }]}
      />
    </LoginForm>
  );
};

export default Init;
