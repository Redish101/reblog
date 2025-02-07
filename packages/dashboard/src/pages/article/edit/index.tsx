import { history, useParams } from "umi";
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
import { Article } from "@/types";

interface ArticleFormValues {
  title: string;
  desc: string;
  slug: string;
  cover?: string;
  draft: boolean;
}

const EditArticlePage = () => {
  const articleSlug = useParams()["slug"];

  const [vd, setVd] = useState<Vditor>();
  const [drawerOpen, setDrawerOpen] = useState<boolean>(false);
  const [articleMeta, setArticleMeta] = useState<ArticleFormValues>();
  const [articleData, setArticleData] = useState<Article>();

  const fetchArticle = async () => {
    if (!articleSlug) {
      message.open({
        type: "warning",
        content: "文章不存在",
      });

      return;
    }

    const res = await useApi(`/api/article/${articleSlug}`);
    const data = await res.json();

    if (data["success"]) {
      setArticleData(data["data"]);
      setArticleMeta({
        title: data["data"]["title"],
        desc: data["data"]["desc"],
        slug: data["data"]["slug"],
        cover: data["data"]["cover"],
        draft: data["data"]["draft"],
      });
    } else {
      message.open({
        type: "warning",
        content: data["msg"],
      });
    }
  };

  useEffect(() => {
    fetchArticle();
  }, []);

  useEffect(() => {
    if (articleData) {
      const vditor = useVditor(articleData.content);
      setVd(vditor);
    }

    return () => {
      if (vd) {
        vd.destroy();
      }
    };
  }, [articleData]);

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

    const { title, desc, slug, cover, draft } = articleMeta;
    const content = vd?.getValue() || "";

    const res = await useApi(`/api/article/${slug}`, {
      method: "PUT",
      data: {
        title,
        desc,
        slug,
        cover,
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
        content: "更新成功",
      });
      history.push("/article");
    } else {
      message.open({
        type: "error",
        content: `更新失败: ${data["msg"]}`,
      });
    }
  };

  return (
    <PageContainer title="编辑">
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
            initialValue={articleMeta?.title}
          />
          <ProFormText
            label="简介"
            name="desc"
            rules={[{ required: true, message: "请填写简介" }]}
            initialValue={articleMeta?.desc}
          />
          <ProFormText
            label="slug"
            name="slug"
            rules={[{ required: true, message: "请填写slug" }]}
            initialValue={articleMeta?.slug}
            disabled
          />
          <ProFormSwitch
            label="草稿"
            name="draft"
            initialValue={articleMeta?.draft}
            rules={[{ required: true }]}
          />
        </ProForm>
      </Drawer>
      <div id="vditor" className="vditor" />
    </PageContainer>
  );
};
export default EditArticlePage;
