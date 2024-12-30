package req

type GetAnalyticsRequest struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
