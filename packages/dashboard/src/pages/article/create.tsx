import { history } from "umi";
import useVditor from "@/libs/vditor";
import useApi from "@/utils/fetcher";
import { SendOutlined } from "@ant-design/icons";
import {
  PageContainer,
  ProForm,
  ProFormSwitch,
  ProFormText,
} from "@ant-design/pro-components";
import { Drawer, FloatButton, message } from "antd";
import { useEffect, useState } from "react";
import Vditor from "vditor";

interface ArticleFormValues {
  title: string;
  desc: string;
  slug: string;
  draft: boolean;
}

const CreateArticlePage = () => {
  const [vd, setVd] = useState<Vditor>();
  const [drawerOpen, setDrawerOpen] = useState<boolean>(true);
  const [articleMeta, setArticleMeta] = useState<ArticleFormValues>();

  useEffect(() => {
    const vditor = useVditor();
    setVd(vditor);

    return () => {
      vd?.destroy();
    };
  }, []);

  const handleDrawerFormFinish = (values: ArticleFormValues) => {
    setArticleMeta(values);
    setDrawerOpen(false);
  };

  const handleSubmit = async () => {
    if (!articleMeta) {
      setDrawerOpen(true);

      message.open({
        type: "info",
        content: "请补全文章基本信息",
      });

      return;
    }

    const { title, desc, slug, draft } = articleMeta;
    const content = vd?.getValue() || "";

    const res = await useApi(`/api/article/${slug}`, {
      method: "POST",
      data: {
        title,
        desc,
        slug,
        content,
        draft,
      },
    });

    const data = await res.json().catch((err) => {
      message.open({
        type: "error",
        content: `发布失败: ${err.message}`,
      });
    });

    if (data["success"]) {
      message.open({
        type: "success",
        content: "发布成功",
      });
      history.push("/article");
    } else {
      message.open({
        type: "error",
        content: `发布失败: ${data["msg"]}`,
      });
    }
  };

  return (
    <PageContainer title="新文章">
      <FloatButton.Group>
        <FloatButton
          onClick={() => setDrawerOpen(true)}
          tooltip="修改基本信息"
        />
        <FloatButton
          onClick={handleSubmit}
          tooltip="发布"
          type="primary"
          icon={<SendOutlined />}
        />
      </FloatButton.Group>
      <Drawer
        open={drawerOpen}
        title="文章信息"
        onClose={() => setDrawerOpen(false)}
      >
        <ProForm onFinish={handleDrawerFormFinish}>
          <ProFormText
            label="标题"
            name="title"
            rules={[{ required: true, message: "请填写标题" }]}
          />
          <ProFormText
            label="简介"
            name="desc"
            rules={[{ required: true, message: "请填写简介" }]}
          />
          <ProFormText
            label="slug"
            name="slug"
            rules={[{ required: true, message: "请填写slug" }]}
          />
          <ProFormSwitch label="草稿" name="draft" />
        </ProForm>
      </Drawer>
      <div id="vditor" className="vditor" />
    </PageContainer>
  );
};

export default CreateArticlePage;
