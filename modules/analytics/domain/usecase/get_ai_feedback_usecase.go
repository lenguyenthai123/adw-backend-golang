package usecase

type GetAIFeedbackUsecase interface {
	ExecuteGetAIFeedback() (string, error)
}

type getAIFeedbackUsecaseImpl struct {
	timeProgressReaderRepo TimeProgressReaderRepo
}

var _ GetAIFeedbackUsecase = (*getAIFeedbackUsecaseImpl)(nil)

func NewGetAIFeedbackUsecase(
	timeProgressReaderRepo TimeProgressReaderRepo,
) GetAIFeedbackUsecase {
	return &getAIFeedbackUsecaseImpl{
		timeProgressReaderRepo: timeProgressReaderRepo,
	}
}

func (uc getAIFeedbackUsecaseImpl) ExecuteGetAIFeedback() (string, error) {
	return "", nil
}
