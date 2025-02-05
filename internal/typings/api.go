package types

type APIRequest struct {
	Messages []api_message `json:"messages"`
	Stream   bool          `json:"stream"`
}

type api_message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
