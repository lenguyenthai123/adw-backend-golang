package mapper

import (
	"backend-golang/modules/task/api/model/req"
	"backend-golang/modules/task/domain/entity"
	"fmt"
	"strconv"
	"time"
)

// req_to_entity maps the CreateTaskRequest or UpdateTaskRequest to the Task entity.
func ConvertCreateTaskRequestToTaskEntity(req req.CreateTaskRequest) entity.Task {
	// Map the request data to the Task entity
	task := entity.Task{
		TaskName:      req.TaskName,
		Description:   req.Description,
		Priority:      req.Priority,
		EstimatedTime: req.EstimatedTime,
		Status:        req.Status,
		DueDate:       parseDueDate(req.DueDate),
	}

	return task
}

func ConvertUpdateTaskRequestToTaskEntity(req req.UpdateTaskRequest, taskId string) entity.Task {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		panic(err)
	}

	// Map the request data to the Task entity
	task := entity.Task{
		TaskID:        id,
		TaskName:      req.TaskName,
		Description:   req.Description,
		Priority:      req.Priority,
		EstimatedTime: req.EstimatedTime,
		Status:        req.Status,
		DueDate:       parseDueDate(req.DueDate),
	}

	return task
}

// Helper function to parse the due date string into time.Time
func parseDueDate(dueDate string) time.Time {
	// Implement proper date parsing based on your expected format
	// For example, using time.Parse() to convert a string to time
	parsedTime, _ := time.Parse(time.RFC3339, dueDate)
	return parsedTime
}

// ConvertToUTC7 nhận vào một chuỗi thời gian và cố gắng chuyển sang UTC+7
func ConvertToUTC7(input string) (string, error) {
	// Danh sách các layout phổ biến
	layouts := []string{
		"2006-01-02T15:04:05Z07:00", // ISO 8601
		"2006-01-02T15:04:05Z",      // ISO 8601 với Z (UTC)
		"2006-01-02 15:04:05",       // YYYY-MM-DD HH:mm:ss
		"02/01/2006 15:04:05",       // DD/MM/YYYY HH:mm:ss
		"2006-01-02",                // YYYY-MM-DD (chỉ có ngày)
	}

	// Thử phân tích chuỗi input với từng layout
	var parsedTime time.Time
	var err error
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, input)
		if err == nil {
			break // Nếu phân tích thành công, thoát khỏi vòng lặp
		}
	}

	// Nếu không parse được, trả lỗi
	if err != nil {
		return "", fmt.Errorf("không thể nhận diện định dạng thời gian: %v", err)
	}

	// Tải múi giờ UTC+7
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return "", fmt.Errorf("lỗi khi tải múi giờ: %v", err)
	}

	// Chuyển đổi sang múi giờ UTC+7
	timeInUTC7 := parsedTime.In(location)

	// Trả về chuỗi ngày giờ trong định dạng chuẩn
	return timeInUTC7.Format("2006-01-02 15:04:05"), nil
}
