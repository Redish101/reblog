import { Button, Result } from "antd";

const NotFound = () => (
  <Result
    status="404"
    title="404 - Not Found"
    subTitle="这里什么都木有"
    extra={
      <Button type="primary" href="/">
        返回首页
      </Button>
    }
  />
);

export default NotFound;
