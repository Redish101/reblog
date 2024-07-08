import { Config } from "./config";
import { Api } from "./utils/api";
import { objToFormData } from "./utils/form";

type Article = {
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

type ArticleList = {
  // 文章总数
  count: number;

  // 文章列表
  articles: Article[];
};

export default class ThemeKit {
  private api: Api;

  constructor(private config: Config) {
    this.api = new Api(this.config.server.url);
  }

  public async getArticleList(opts?: { pageIndex: number; pageSize: number }) {
    const params = opts
      ? `/?pageIndex=${opts?.pageIndex}&pageSize=${opts?.pageSize}`
      : "";
    const data = await this.api.get<ArticleList>(`/api/article/list/${params}`);

    return data.data;
  }

  public async getArticle(slug: string) {
    const data = await this.api.get<Article>(`/api/article/${slug}`);

    return data.data;
  }
}
