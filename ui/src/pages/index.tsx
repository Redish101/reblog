import useApi from "@/utils/fetcher";
import { Button } from "antd";
import { useEffect, useState } from "react";

const HomePage = () => {
  const [version, setVersion] = useState();

  async function refresh() {
    try {
      const res = await fetch("/api/version");
      setVersion(await res.json());
    } catch (err) {
      console.error(err);
    }
  }

  useEffect(() => {
    refresh();
  }, []);

  const handleClick = async () => {
    const formData = new FormData();
    formData.append("title", "测试文章");
    formData.append("desc", "这是一个测试文章");
    formData.append(
      "content",
      "wow，这是一个测试文章，他里面什么都没有，但是是一个来自前端的测试文章！哇！！！",
    );
    const res = await useApi("/api/article/test", {
      method: "POST",
      body: formData,
    });
  };

  return (
    <div>
      <h1>Placeeeeeeeeee Holder!</h1>
      <p>伟大的占位符！</p>
      <p>占位符静静地看着你。</p>
      <p>Hurr durr, i'ma placeholder!</p>
      <Button type="primary">我是一个按钮</Button>
      <p>他是一个按钮，尽管你按下他不会发生什么，但他的确是一个按钮。</p>
      <p>而且是一个尊贵的primary按钮。</p>
      <Button type="default" onClick={handleClick}>
        我也是一个按钮
      </Button>
      <p>他也是一个按钮，但是你按下这个按钮后会尝试新建一篇文章</p>
      <h2>版本信息</h2>
      {!version && <p>正在获取版本信息...</p>}
      {version && (
        <div>
          <p>{version["data"]["app_name"]}</p>
          <p>{version["data"]["runtime"]}</p>
        </div>
      )}
    </div>
  );
};

export default HomePage;
