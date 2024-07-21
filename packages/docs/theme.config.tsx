export default {
  logo: <span>reblog</span>,
  project: {
    link: "https://github.com/redish101/reblog",
  },
  docsRepositoryBase:
    "https://github.com/redish101/reblog/tree/main/packages/docs/pages",
  useNextSeoProps() {
    return {
      titleTemplate: "%s – reblog",
    };
  },
  search: {
    placeholder: "搜索",
  },
  toc: {
    title: "目录",
  },
  feedback: {
    content: "发现问题？提交反馈 →",
  },
  editLink: {
    text: "编辑此页",
  },
  footer: {
    text: "© 2024-present reblog",
  },
  themeSwitch: {
    useOptions() {
      return {
        light: "亮色",
        dark: "暗色",
        system: "跟随系统",
      };
    },
  },
};
