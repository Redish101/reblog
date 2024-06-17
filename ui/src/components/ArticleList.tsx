import { Article } from "@/types";
import useApi from "@/utils/fetcher";
import { ProColumns, ProTable } from "@ant-design/pro-components";
import { Button, Popconfirm } from "antd";
import { Link, history } from "umi";

const ArticleList = () => {
  const columns: ProColumns<Article>[] = [
    {
      dataIndex: "id",
      title: "序号",
      valueType: "indexBorder",
    },
    {
      title: "slug",
      dataIndex: "slug",
      width: 200,
    },
    {
      title: "标题",
      dataIndex: "title",
      copyable: true,
    },
    {
      title: "创建时间",
      dataIndex: "created_at",
      valueType: "dateTime",
    },
    {
      title: "更新时间",
      dataIndex: "updated_at",
      valueType: "dateTime",
    },
    {
      title: "操作",
      valueType: "option",
      key: "option",
      render: (text, record, _, action) => [
        <Link
          key={`edit-${record.slug}-link`}
          to={`/article/${record.slug}/edit`}
        >
          编辑
        </Link>,
        <Popconfirm
          title="确认删除?"
          description={`确认删除文章 ${record.title}?此操作不可逆。`}
          onConfirm={async () => {
            useApi(`/api/article/${record.slug}`, { method: "DELETE" });
            action?.reload();
          }}
          key={`delete-${record.slug}-popconfirm`}
        >
          <Link key={`delete-${record.slug}-link`} to="#">
            删除
          </Link>
        </Popconfirm>,
      ],
    },
  ];
  return (
    <ProTable
      columns={columns}
      search={false}
      pagination={{ pageSize: 10 }}
      rowKey={"slug"}
      request={async (params: { pageSize: number; current: number }) => {
        const { pageSize, current } = params;

        const res = await useApi(
          `/api/article/list/?pageIndex=${current}&pageSize=${pageSize}`,
        );
        const data = await res.json();

        return {
          success: data["success"],
          data: data["data"]["articles"],
          total: data["data"]["count"],
        };
      }}
      toolBarRender={() => [
        <Button type="primary" onClick={() => history.push("/article/create")}>
          新建
        </Button>,
      ]}
    ></ProTable>
  );
};

export default ArticleList;
