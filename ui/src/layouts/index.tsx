import { HomeFilled, SmileTwoTone } from "@ant-design/icons";
import { PageContainer, ProCard, ProLayout } from "@ant-design/pro-components";
import { notification } from "antd";
import { useEffect, useState } from "react";
import { Outlet, useLocation, Icon } from "umi";
import useApi from "@/utils/fetcher";

const Layout = () => {
  const location = useLocation();

  const [userInfo, setUserInfo] = useState();

  const refresh = async () => {
    try {
      const res = await useApi("/api/admin/userInfo");

      setUserInfo(await res.json());

      if (res.status != 200) throw new Error("获取用户信息失败");
    } catch (err) {
      notification.error({
        message: `无法获取用户信息: ${err}`,
      });
    }
  };

  useEffect(() => {
    refresh();
  }, []);

  return (
    <ProLayout
      title="reblog"
      logo={<Icon icon="local:logo" />}
      siderWidth={216}
      location={location}
      route={{
        path: "/",
        routes: [
          {
            path: "/",
            name: "首页",
            icon: <HomeFilled />,
          },
        ],
      }}
      avatarProps={{
        src: <SmileTwoTone />,
        title: userInfo ? userInfo["data"]["nickname"] : "Loading",
        size: "small",
      }}
    >
      <PageContainer>
        <ProCard style={{ minHeight: 700 }}>
          <Outlet />
        </ProCard>
      </PageContainer>
    </ProLayout>
  );
};

export default Layout;
