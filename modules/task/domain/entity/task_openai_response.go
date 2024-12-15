package entity

// AIResponse là cấu trúc trả về của API
type TaskOpenaiResponse struct {
	TaskList []TaskOpenai `json:"taskList" jsonschema_description:"The list of tasks"`
	Feedback string       `json:"feedback" jsonschema_description:"AI-generated feedback or recommendations"`
}
