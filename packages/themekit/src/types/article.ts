export type Article = {
  // 文章id
  id: number;

  // 文章创建时间
  created_at: string;

  // 文章更新时间
  updated_at: string;

  // 文章slug
  slug: string;

  // 文章标题
  title: string;

  // 文章描述
  desc: string;

  // 文章内容
  content: string;
};

export type ArticleList = {
  // 文章总数
  count: number;

  // 文章列表
  articles: Article[];
};
