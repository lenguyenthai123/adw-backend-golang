package usecase

import (
	"backend-golang/core"
	"backend-golang/modules/analytics/constant"
	"backend-golang/modules/task/domain/entity"
	"backend-golang/utils"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/openai/openai-go"
)

type GetAIFeedbackUsecase interface {
	ExecuteGetAIFeedback(ctx context.Context) (*entity.TaskOpenaiResponse, error)
}

type getAIFeedbackUsecaseImpl struct {
	timeProgressReaderRepo TimeProgressReaderRepo
	openaiClient           *openai.Client
}

var _ GetAIFeedbackUsecase = (*getAIFeedbackUsecaseImpl)(nil)

func NewGetAIFeedbackUsecase(
	timeProgressReaderRepo TimeProgressReaderRepo,
	openaiClient *openai.Client,
) GetAIFeedbackUsecase {
	return &getAIFeedbackUsecaseImpl{
		timeProgressReaderRepo: timeProgressReaderRepo,
		openaiClient:           openaiClient,
	}
}

func (uc getAIFeedbackUsecaseImpl) ExecuteGetAIFeedback(ctx context.Context) (*entity.TaskOpenaiResponse, error) {
	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	// Convert userId to int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert userId to int: %v", err)
	}

	// Fetch data from repositories
	taskProgress, err := uc.timeProgressReaderRepo.GetEachTaskProgress(userIdInt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch task progress: %v", err)
	}

	if taskProgress == nil || len(*taskProgress) == 0 {
		return nil, constant.ErrorNotAnyProgressToAnalyze(err)
	}

	now := time.Now()

	// Calculate the start time (30 days ago from now)
	startTime := now.AddDate(0, 0, -30).Format(time.RFC3339)

	// Format the current time as the end time
	endTime := now.Format(time.RFC3339)
	dailyProgress, err := uc.timeProgressReaderRepo.GetTimeSpentDaily(userIdInt, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch daily progress: %v", err)
	}

	// Prepare AI prompt data
	taskDataJSON, _ := json.Marshal(taskProgress)
	dailyProgressJSON, _ := json.Marshal(dailyProgress)

	question := fmt.Sprintf(`
        Based on the user's study data, provide detailed feedback as follows (please do not break the format):
        1. Identify tasks or subjects where the user is excelling. 
        2. Highlight areas requiring more attention.
        3. Offer motivational feedback to encourage consistency and improvement.
		Note: When indicate a task, please use task name instead of id, and give feedback for in task in "Description" field of those tasks.
        Data:
        - Task Progress: %s
        - Daily Progress: %s
    `, taskDataJSON, dailyProgressJSON)

	// Generate schema for AI response
	schema := utils.GenerateSchema[entity.TaskOpenaiResponse]()
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("study_feedback"),
		Description: openai.F("AI feedback on user study progress and performance."),
		Schema:      openai.F(schema),
		Strict:      openai.Bool(true),
	}

	// Query OpenAI API
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
		Model: openai.F(openai.ChatModelGPT4o),
	})

	if err != nil {
		return nil, fmt.Errorf("AI feedback generation failed: %v", err)
	}

	// Parse AI response
	feedbackResponse := entity.TaskOpenaiResponse{}
	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &feedbackResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v", err)
	}

	return &feedbackResponse, nil
}
