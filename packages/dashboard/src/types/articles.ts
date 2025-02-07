interface Article {
  created_at: string;
  updated_at: string;
  id: number;
  slug: string;
  title: string;
  desc: string;
  cover?: string;
  content: string;
  draft: boolean;
}

interface Articles {
  count: number;
  articles: Article[];
}

export { Article, Articles };
