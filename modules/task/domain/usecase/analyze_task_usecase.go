package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/domain/entity"
	"backend-golang/pkgs/log"
	utils "backend-golang/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/openai/openai-go"
)

type AnalyzeTaskUsecase interface {
	ExecAnalyzeTask(ctx context.Context, startTime, endTime string) (*entity.TaskOpenaiResponse, error)
}

type analyzeTaskUsecaseImpl struct {
	readerRepo   TaskReaderRepository
	openaiClient *openai.Client
}

var _ AnalyzeTaskUsecase = (*analyzeTaskUsecaseImpl)(nil)

func NewAnalyzeTaskUsecase(
	readerRepo TaskReaderRepository,
	openaiClient *openai.Client) AnalyzeTaskUsecase {
	return &analyzeTaskUsecaseImpl{
		readerRepo:   readerRepo,
		openaiClient: openaiClient,
	}
}

func (uc analyzeTaskUsecaseImpl) ExecAnalyzeTask(ctx context.Context, startTime, endTime string) (*entity.TaskOpenaiResponse, error) {

	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	taskList, err := uc.readerRepo.FindTaskListByRangeTime(ctx, userId, startTime, endTime)

	taskOpenaiList := make([]entity.TaskOpenai, 0)
	for _, task := range taskList {
		taskOpenaiList = append(taskOpenaiList, mapper.ConvertTaskToTaskOpenai(*task))
	}

	// Tạo JSON Schema từ AIResponse
	schema := utils.GenerateSchema[entity.TaskOpenaiResponse]()

	// Tạo prompt (chuỗi câu hỏi gửi lên OpenAI)
	tasksJSON, err := json.Marshal(taskOpenaiList)

	question := fmt.Sprintf(`
	Analyze the following schedule and provide feedback.
	Tasks: %s.
	Feedback should include:
	1. Warnings about tight schedules.
	2. Recommendations for prioritization to ensure balance and focus.
	`, tasksJSON)

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("biography"),
		Description: openai.F("Notable information about a person"),
		Schema:      openai.F(schema),
		Strict:      openai.Bool(true),
	}

	// Query the Chat Completions API
	chat, err := uc.openaiClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		// Only certain models can perform structured outputs
		Model: openai.F(openai.ChatModelGPT4o2024_08_06),
	})

	// The model responds with a JSON string, so parse it into a struct
	taskResponseOpenai := entity.TaskOpenaiResponse{}
	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &taskResponseOpenai)

	fmt.Printf("Feedback: %v\n", taskResponseOpenai.Feedback)
	for _, taskResponse := range taskResponseOpenai.TaskList {
		fmt.Printf("Name: %v\n", taskResponse.TaskName)
		fmt.Printf("Year: %v\n", taskResponse.DueDate)
		fmt.Printf("Org: %v\n", taskResponse.Description)
		fmt.Printf("Legacy: %v\n", taskResponse.Priority)
	}
	log.JsonLogger.Info("ExecAnalyzeTask.find_task_list_by_time_range", chat, err)
	if err != nil {
		return nil, err
	}
	return &taskResponseOpenai, nil
}
