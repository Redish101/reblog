import { PageContainer, ProCard, ProLayout } from "@ant-design/pro-components";
import { Outlet, useLocation } from "umi";

const Layout = () => {
  const location = useLocation();
  return (
    <ProLayout title="reblog" siderWidth={216} location={location}>
      <PageContainer>
        <ProCard style={{ minHeight: 700 }}>
          <Outlet />
        </ProCard>
      </PageContainer>
    </ProLayout>
  );
};

export default Layout;
