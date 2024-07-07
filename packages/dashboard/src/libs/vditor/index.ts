import Vditor from "vditor";
import "vditor/dist/index.css";
import "./index.less";

const useVditor = (value?: string) => {
  const vditor = new Vditor("vditor", {
    mode: "wysiwyg",
    minHeight: 650,
    after: () => {
      vditor.setValue(value || "# 你好, 世界! \n Markdown Supported!");
    },
    toolbar: [
      "emoji",
      "headings",
      "bold",
      "italic",
      "strike",
      "link",
      "|",
      "list",
      "ordered-list",
      "check",
      "outdent",
      "indent",
      "|",
      "quote",
      "line",
      "code",
      "inline-code",
      "insert-before",
      "insert-after",
      "|",
      "table",
      "|",
      "undo",
      "redo",
      "|",
      "fullscreen",
      "edit-mode",
      {
        name: "more",
        toolbar: [
          "both",
          "code-theme",
          "content-theme",
          "export",
          "outline",
          "preview",
          "devtools",
          "info",
          "help",
        ],
      },
    ],
  });

  return vditor;
};

export default useVditor;
