import useApi from "@/utils/fetcher";
import { ArrowDownOutlined } from "@ant-design/icons";
import { Card, Col, Row, Statistic, Typography, message } from "antd";
import { useEffect, useState } from "react";
import { history } from "umi";

interface Article {
  created_at: string;
  updated_at: string;
  id: number;
  title: string;
  desc: string;
  content: string;
}

interface Articles {
  count: number;
  articles: Article[];
}

const HomePage = () => {
  const token = localStorage.getItem("token");

  if (!token) {
    history.push("/login");
  }

  useApi("/api/admin/tokenState")
    .then((res) => res.json())
    .then((data) => {
      if (!data["data"]) {
        localStorage.removeItem("token");
        history.push("/login");
      }
    });

  const calculateDaysAgo = (updatedAt: string) => {
    const lastUpdatedDate = new Date(updatedAt);
    const now = new Date();
    const diff = now.getTime() - lastUpdatedDate.getTime();
    const daysDiff = diff / (1000 * 3600 * 24);
    return Math.floor(daysDiff);
  };

  const [articles, setArticles] = useState<Articles>();
  const [lastUpdatedDay, setLastUpdatedDay] = useState<string>("暂无数据");
  const [version, setVersion] = useState<string>("暂无数据");

  const refresh = async () => {
    const articleListRes = await useApi("/api/article/list/?page=1&size=1");
    const articleListData = await articleListRes.json();

    if (!articleListData["success"]) {
      message.open({
        type: "error",
        content: `获取文章列表失败: ${articleListData["msg"]}`,
      });
    }

    setArticles(articleListData["data"]);

    const daysAgo = calculateDaysAgo(
      articleListData["data"].articles[0].updated_at,
    );
    setLastUpdatedDay(daysAgo === 0 ? "今天" : `${daysAgo} 天`);

    const versionRes = await useApi("/api/version");
    const versionData = await versionRes.json();

    if (!versionData["success"]) {
      message.open({
        type: "error",
        content: `获取版本信息失败: ${versionData["msg"]}`,
      });
    }

    setVersion(versionData["data"]["version"]);
  };

  useEffect(() => {
    refresh();
  }, []);

  return (
    <div>
      <div>
        <Row gutter={18}>
          <Col span={6}>
            <Card bordered={false}>
              <Statistic title="文章数量" value={articles?.count || 0} />
            </Card>
          </Col>
          <Col span={6}>
            <Card bordered={false}>
              <Statistic
                title="距上次更新"
                value={lastUpdatedDay || "暂无数据"}
              />
            </Card>
          </Col>
          <Col span={6}>
            <Card bordered={false}>
              <Statistic title="reblog 版本" value={version || "暂无数据"} />
            </Card>
          </Col>
        </Row>
      </div>
      <div style={{ marginTop: 20 }}>
        <Typography.Text type="secondary">今天也要多写文章</Typography.Text>
      </div>
    </div>
  );
};

export default HomePage;
