import useApi from "@/utils/fetcher";
import {
  PageContainer,
  ProForm,
  ProFormInstance,
  ProFormText,
} from "@ant-design/pro-components";
import { message } from "antd";
import { useEffect, useRef, useState } from "react";

interface Site {
  name: string;
  url: string;
  desc: string;
}

const SitePage = () => {
  const formRef = useRef<ProFormInstance<Site>>();

  const refresh = () => {
    useApi("/api/site")
      .then((res) => res.json())
      .then((data) => {
        const { name, url, desc } = data["data"];
        formRef.current?.setFieldsValue({
          name,
          url,
          desc,
        });
      })
      .catch((err) => {
        message.open({
          content: `获取站点信息失败: ${err}`,
          type: "error",
        });
      });
  };

  useEffect(() => {
    refresh();
  }, []);

  const handleSubmit = async (value: Site) => {
    const formData = new FormData();

    formData.append("name", value.name);
    formData.append("url", value.url);
    formData.append("desc", value.desc);

    useApi("/api/admin/site", {
      method: "PUT",
      body: formData,
    })
      .then((res) => res.json())
      .then(() => {
        message.success("站点信息更新成功");
        refresh();
      })
      .catch((err) => {
        message.error(`站点信息更新失败: ${err}`);
      });
  };

  return (
    <PageContainer title="站点信息">
      <ProForm formRef={formRef} onFinish={handleSubmit}>
        <ProFormText name="name" label="站点名称" width="md" />
        <ProFormText name="url" label="站点地址" width="md" />
        <ProFormText name="desc" label="站点描述" width="md" />
      </ProForm>
    </PageContainer>
  );
};
export default SitePage;
