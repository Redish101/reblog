import useApi from "@/utils/fetcher";
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { LoginForm, ProFormText } from "@ant-design/pro-components";
import { App, message } from "antd";
import { Link, history } from "umi";

interface LoginFormValues {
  username: string;
  password: string;
}

const Login = () => {
  const onFinish = async (values: LoginFormValues) => {
    const { username, password } = values;
    const formData = new FormData();
    formData.append("username", username);
    formData.append("password", password);

    const res = await useApi("/api/admin/login", {
      method: "POST",
      body: formData,
    });

    const data = await res.json();

    if (res.status != 200) {
      message.open({
        type: "error",
        content: data["msg"],
      });
    } else {
      localStorage.setItem("token", data["data"]["token"]);
      history.push("/");
    }
  };
  return (
    <LoginForm<LoginFormValues>
      title="登录"
      subTitle="reblog dashboard"
      onFinish={onFinish}
    >
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
      <div style={{ marginBottom: "16px", textAlign: "right" }}>
        <Link to="/init">尚未初始化?</Link>
      </div>
    </LoginForm>
  );
};

export default Login;
