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

  return (
    <div>
      <h1>Placeeeeeeeeee Holder!</h1>
      <p>伟大的占位符！</p>
      <p>占位符静静地看着你。</p>
      <p>Hurr durr, i'ma placeholder!</p>
      <Button type="primary">我是一个按钮</Button>
      <p>他是一个按钮，尽管你按下他不会发生什么，但他的确是一个按钮。</p>
      <p>而且是一个尊贵的primary按钮。</p>
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
