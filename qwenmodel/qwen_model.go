package qwenmodel

const (
	ChatUser = "user"
	ChatBot  = "assistant"
)

const (
	ChatQWenModel = "qwen-max"
	ChatBaseUrl   = "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
)

type Input struct {
	Messages []Messages `json:"messages"`
}

type QWenTurbo struct {
	Model      string     `json:"model"`
	Input      Input      `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	EnableSearch bool `json:"enable_search"`
	Seed int `json:"seed"`
    	TopP float64 `json:"top_p"`
	MaxTokens int `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	RepetitionPenalty float64 `json:"repetition_penalty"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Result

type Output struct {
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

type Usage struct {
	OutputTokens int `json:"output_tokens"`
	InputTokens  int `json:"input_tokens"`
}

type Response struct {
	Output    Output `json:"output"`
	Usage     Usage  `json:"usage"`
	RequestID string `json:"request_id"`
}
