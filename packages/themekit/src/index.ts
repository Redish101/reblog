import { Config } from "./config";
import { Article, ArticleList, Friend, Site } from "./types";
import { Api, objToFormData } from "./utils";

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

  public async getSite() {
    const data = await this.api.get<Site>("/api/site");

    return data.data;
  }

  public async getFriendList(opts?: { pageIndex: number; pageSize: number }) {
    const params = opts
      ? `/?pageIndex=${opts?.pageIndex}&pageSize=${opts?.pageSize}`
      : "";
    const data = await this.api.get<ArticleList>(`/api/friend/list/${params}`);

    return data.data;
  }

  public async addFriend(friend: Friend) {
    const data = objToFormData(friend);
    const res = await this.api.post<Friend>("/api/friend", data);

    return res;
  }
}

export * from "./types";
export * from "./config";
export * from "./utils";
