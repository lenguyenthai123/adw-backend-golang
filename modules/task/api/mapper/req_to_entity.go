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
		UserID:        req.UserID,
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
