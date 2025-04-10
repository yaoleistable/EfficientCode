package ai

// TranslateText 英汉互译函数
func TranslateText(modelName, text string) (string, error) {
	prompt := "你是一个英汉互译专家。如果输入的是中文，请将其翻译成英文；如果输入的是英文，请将其翻译成中文。请只输出翻译结果，不要包含任何解释或额外的文字。"
	return AskAI(modelName, prompt, text)
}

// PolishText 文本润色函数
func PolishText(modelName, text string) (string, error) {
	prompt := "你是一个文本润色专家。请对输入的文本进行润色和优化，使其更加流畅、专业和优雅。保持原文的核心意思不变，但可以适当调整表达方式。请只输出润色后的文本，不要包含任何解释或额外的文字。"
	return AskAI(modelName, prompt, text)
}

// SummarizeText 文本总结函数
func SummarizeText(modelName, text string) (string, error) {
	prompt := "你是一个文本总结专家。请对输入的文本进行简明扼要的总结，突出重点内容。总结应当保持原文的主要观点，但要更加简洁。请只输出总结后的文本，不要包含任何解释或额外的文字。"
	return AskAI(modelName, prompt, text)
}

// AskAssistant AI助手问答函数
func AskAssistant(modelName, question string) (string, error) {
	prompt := `你是一个专业、友好的AI助手。你的任务是：
1. 理解用户的问题并提供准确、有帮助的回答
2. 使用清晰、易懂的语言
3. 如果不确定答案，诚实地表示不知道
4. 在合适的时候提供相关的建议或补充信息
请直接回答用户的问题，不需要重复问题内容。`
	return AskAI(modelName, prompt, question)
}
