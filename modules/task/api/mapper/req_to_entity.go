package mapper

import (
	"backend-golang/modules/task/api/model/req"
	"backend-golang/modules/task/domain/entity"
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
		StartDate:     parseDueDate(req.StartDate),
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
		StartDate:     parseDueDate(req.StartDate),
		DueDate:       parseDueDate(req.DueDate),
	}

	return task
}

func parseDueDate(dueDate string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, dueDate)
	// Chuyển đổi sang múi giờ UTC+7 (Asia/Ho_Chi_Minh)
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	localTime := parsedTime.In(location)

	return localTime
}

// ConvertToUTC7 nhận vào một chuỗi thời gian và cố gắng chuyển sang UTC+7
func ConvertToUTC7(input string) (string, error) {
	// Danh sách các layout phổ biến
	// Chuyển đổi time.Time sang string với định dạng chuẩn RFC3339
	convertedTime := parseDueDate(input)
	formattedTime := convertedTime.Format(time.RFC3339)
	return formattedTime, nil
}
