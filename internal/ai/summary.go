package ai

import (
	"context"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/log"
	"github.com/yankeguo/zhipu"
)

const prompt = `
【角色定义】
你是一位资深技术文档工程师，专注于为复杂的技术博客生成精准摘要。请严格基于原文内容生成专业摘要，保持技术严谨性。

【核心要求】

内容忠实度：

完全基于原文事实，禁止添加任何外部知识或个人见解

保留关键专业术语（如Transformer、CUDA核心、RAG架构等）

准确呈现技术方案的核心创新点

结构规范：

首段：总览句（20字内）点明文章核心价值

主体内容包含：
• 待解决的技术问题
• 采用的解决方案架构
• 关键算法/工具链（如使用PyTorch 2.1+TorchDynamo）
• 实验结果（量化指标需精确，如"QPS提升37.2%"）
• 实际应用场景

语言风格：

专业而不晦涩，避免过度学术化表达

使用技术领域标准表达（如"微调"而非"调整"，"消融实验"而非"对比测试"）

保持客观中立，区分作者主张和客观事实

格式规范：

纯文本输出，禁用任何Markdown格式、禁止换行

中文字数严格控制在280-320字

段落间自然过渡，保持技术逻辑连贯性

不要出现任何文章内没有出现的内容，包括类似“实验结果：无具体量化指标“

【典型示例】
[合格摘要]
"本文探讨PyTorch 2.1的编译优化技术。针对动态图执行效率问题，提出基于TorchDynamo的图捕获方案，配合TensorRT的后端优化，实现训练速度提升。核心创新包括：1）通过字节码分析实现Python动态特性兼容 2）选择性子图JIT编译策略 3）与CUDA 11.7的深度适配。在ResNet-152训练任务中，相较原生PyTorch获得37.2%的加速比，内存占用降低19%。该方案适用于需要兼顾开发灵活性和执行效率的DL训练场景。"

[不合格摘要]
"这篇文章讲PyTorch的优化方法，作者搞了个新的编译方案，速度提升明显，这对AI训练很有帮助。"（缺乏技术细节和量化指标）

**GN(Generate Ninja) 快速入门指南**\n\n本文介绍了 Google 开发的元构建系统 GN 的使用方法。GN 用于生成 build.ninja 文件，类似于 cmake。文章通过构建一个基于 glfw + glad 的 C++ 程序，详细讲解了 GN 的安装、配置编译工具链、声明构建配置以及构建过程。\n\n**待解决的技术问题**：GN 的使用方法和配置流程。\n\n**解决方案架构**：通过构建一个简单的 C++ 程序，展示 GN 的使用步骤。\n\n**关键算法/工具链**：GN、ninja、glfw、glad。\n\n**实验结果**：无具体量化指标。\n\n**实际应用场景**：GN 适用于规模较大的项目中，更为清晰的描述构建过程。（使用了markdown语法以及换行）

你需要牢记： 禁止使用任何markdown语法、禁止换行！
你需要牢记： 禁止使用任何markdown语法、禁止换行！
你需要牢记： 禁止使用任何markdown语法、禁止换行！

请严格遵循以上规范生成专业技术摘要，现在请处理以下技术博客内容：
`
func Summary(app *core.App, text string) string {
    log.Debug("[AI] 正生成摘要")
    client, err := zhipu.NewClient(zhipu.WithAPIKey(app.Config().Ai.ApiKey))

	service := client.ChatCompletion("glm-4-flash").
		AddMessage(zhipu.ChatCompletionMessage{
			Role:    "user",
			Content: prompt + text,
            
		})

	res, err := service.Do(context.Background())

	if err != nil {
		zhipu.GetAPIErrorCode(err)
        log.Warnf("[AI] 摘要生成失败： %s", err)
	} else {
        log.Debugf("[AI] 成功生成摘要： %s", res.Choices[0].Message.Content)
	    return res.Choices[0].Message.Content
	}

    return "生成摘要失败"
}
