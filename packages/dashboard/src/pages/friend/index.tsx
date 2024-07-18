import { Link } from "umi";
import { Friend } from "@/types/friend";
import useApi from "@/utils/fetcher";
import { PageContainer, ProCard, ProForm } from "@ant-design/pro-components";
import { Avatar, Button, Col, Drawer, Row } from "antd";
import { useEffect, useState } from "react";
import { EditTwoTone } from "@ant-design/icons";

const FriendPage = () => {
  const [friends, setFriends] = useState<Friend[]>([]);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [drawerOpen, setDrawerOpen] = useState(false);

  const loadMore = async () => {
    if (!hasMore || loading) return;

    setLoading(true);
    try {
      const res = await useApi(
        `/api/friend/list/?pageIndex=${page}&pageSize=9`,
      );
      const data = await res.json();
      const newFriends = data["data"]["friends"];
      setFriends((prev) => [...prev, ...newFriends]);
      setPage((prevPage) => prevPage + 1);
      setHasMore(newFriends.length === 9);
    } catch (error) {
      console.error("Error loading more friends:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadMore();
  }, []);

  return (
    <PageContainer title="友情链接">
      <Drawer open={drawerOpen} onClose={() => setDrawerOpen(false)}>
        <ProForm></ProForm>
      </Drawer>
      <Row gutter={[16, 16]}>
        {friends.map((item) => (
          <Col xs={24} sm={12} md={8} key={item.name}>
            <ProCard
              title={item.name}
              extra={<EditTwoTone onClick={() => setDrawerOpen(true)} />}
            >
              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                }}
              >
                <Avatar src={item.avatar} size={64} />
                <div style={{ textAlign: "right", whiteSpace: "normal" }}>
                  <div>{item.desc}</div>
                  <Link to={item.url}>{item.url}</Link>
                  <div>{item.visible ? "可见" : "不可见"}</div>
                </div>
              </div>
            </ProCard>
          </Col>
        ))}
      </Row>
      <div style={{ textAlign: "center", marginTop: 16 }}>
        {hasMore && !loading && <Button onClick={loadMore}>加载更多</Button>}
        {loading && <p>正在加载...</p>}
      </div>
    </PageContainer>
  );
};

export default FriendPage;
