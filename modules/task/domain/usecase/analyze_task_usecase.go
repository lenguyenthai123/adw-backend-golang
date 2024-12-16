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
	// Initialize min and max with the first element
	minStartDate := taskList[0].StartDate
	maxDueDate := taskList[0].DueDate
	for _, task := range taskList {
		taskOpenaiList = append(taskOpenaiList, mapper.ConvertTaskToTaskOpenai(*task))
		if task.StartDate.Before(minStartDate) {
			minStartDate = task.StartDate
		}
		if task.DueDate.After(maxDueDate) {
			maxDueDate = task.DueDate
		}
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
	3. Reorganize the tasks by adjusting their startTime, dueDate, and priority to better align with the context.
	4. If necessary, split tasks into smaller subtasks based on the context to ensure better distribution and alignment.
	5. The startDate and dueDate must be generated based on the original startDate = %s and dueDate = %s provided for each task, ensuring they align appropriately with the context. AI should allocate hours and minutes (time of day) automatically to fit within the range 08:00 to 22:00 to suitable with context.
	`, tasksJSON, minStartDate, maxDueDate)

	fmt.Printf(question)

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("schedule_analysis"),
		Description: openai.F("Analysis and recommendations for a given schedule, including adjustments to tasks."),
		Schema:      openai.F(schema), // Schema đã cập nhật
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

	for _, taskResponse := range taskResponseOpenai.TaskList {
		taskResponse.DueDate, _ = mapper.ConvertToUTC7(taskResponse.DueDate)
		taskResponse.StartDate, _ = mapper.ConvertToUTC7(taskResponse.DueDate)
	}

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
