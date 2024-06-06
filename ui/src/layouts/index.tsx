import { HomeFilled } from "@ant-design/icons";
import {
  DefaultFooter,
  PageContainer,
  ProCard,
  ProLayout,
} from "@ant-design/pro-components";
import { Outlet, useLocation, Icon } from "umi";

const Layout = () => {
  const location = useLocation();
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
