package res

type TaskAnalyzeResponse struct {
	FeedBack string         `json:"feedBack"`
	Data     []TaskResponse `json:"data"`
}
