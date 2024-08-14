import { Link } from "umi";
import { Friend } from "@/types/friend";
import useApi from "@/utils/fetcher";
import {
  PageContainer,
  ProCard,
  ProForm,
  ProFormSwitch,
  ProFormText,
} from "@ant-design/pro-components";
import { Avatar, Button, Col, Drawer, Popconfirm, Row } from "antd";
import { useEffect, useState } from "react";
import { DeleteTwoTone, EditTwoTone } from "@ant-design/icons";

const FriendPage = () => {
  const [friends, setFriends] = useState<Friend[]>([]);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [currentFriend, setCurrentFriend] = useState<Friend>();

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

  const handleDrawerSubmit = async (values: Friend) => {
    let updatedFriends: Friend[] = [];

    try {
      if (currentFriend !== undefined) {
        await useApi(`/api/friend/${currentFriend?.id}`, {
          method: "PUT",
          data: values,
        });
        updatedFriends = friends.map((friend) =>
          friend.id === currentFriend?.id ? { ...friend, ...values } : friend,
        );
      } else {
        await useApi(`/api/friend`, {
          method: "POST",
          data: values,
        });
        updatedFriends = [...friends, values];
      }
    } catch (error) {
      console.error("Error updating or adding friend:", error);
    } finally {
      setFriends(updatedFriends);
      setDrawerOpen(false);
      setCurrentFriend(undefined);
    }
  };

  return (
    <PageContainer title="友情链接">
      <Drawer
        open={drawerOpen}
        onClose={() => setDrawerOpen(false)}
        title={
          currentFriend === undefined
            ? "添加友情链接"
            : `修改 ${currentFriend?.name}`
        }
      >
        <ProForm onFinish={handleDrawerSubmit}>
          <ProFormText
            name={"name"}
            label={"名称"}
            initialValue={currentFriend?.name}
            rules={[{ required: true, message: "请输入名称" }]}
          />
          <ProFormText
            name={"desc"}
            label={"描述"}
            initialValue={currentFriend?.desc}
            rules={[{ required: true, message: "请输入描述" }]}
          />
          <ProFormText
            name={"url"}
            label={"链接"}
            initialValue={currentFriend?.url}
            rules={[{ required: true, message: "请输入链接", type: "url" }]}
          />
          <ProFormText
            name={"avatar"}
            label={"头像"}
            initialValue={currentFriend?.avatar}
            rules={[{ required: true, message: "请输入头像URL", type: "url" }]}
          />
          <ProFormSwitch
            name={"visible"}
            label={"可见性"}
            initialValue={
              currentFriend?.visible !== undefined
                ? currentFriend.visible
                : true
            }
            rules={[{ required: true, message: "请选择可见性" }]}
          />
        </ProForm>
      </Drawer>
      <Row gutter={[16, 16]}>
        {friends.map((item) => (
          <Col xs={24} sm={12} md={8} key={item.name}>
            <ProCard
              title={item.name}
              extra={
                <div>
                  <EditTwoTone
                    onClick={() => {
                      setCurrentFriend(item);
                      setDrawerOpen(true);
                    }}
                  />
                  <Popconfirm
                    title="确认删除？"
                    onConfirm={() => {
                      useApi(`/api/friend/${item.id}`, {
                        method: "DELETE",
                      }).then(() => {
                        setFriends((prev) =>
                          prev.filter((f) => f.id !== item.id),
                        );
                      });
                    }}
                  >
                    <DeleteTwoTone style={{ marginLeft: 16 }} />
                  </Popconfirm>
                </div>
              }
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
      <div
        style={{
          marginTop: 16,
          display: "flex",
          justifyContent: "center",
          gap: 16,
        }}
      >
        {hasMore && !loading && <Button onClick={loadMore}>加载更多</Button>}
        {loading && <p>正在加载...</p>}
        <Button onClick={() => setDrawerOpen(true)}>添加</Button>
      </div>
    </PageContainer>
  );
};

export default FriendPage;
