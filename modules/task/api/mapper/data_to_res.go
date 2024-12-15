package mapper

import (
	"backend-golang/modules/task/api/model/res"
	"backend-golang/modules/task/domain/entity"
)

// DataToRes maps the Task entity to the TaskResponse DTO
func ConvertTaskEntityToTaskRes(task entity.Task) res.TaskResponse {
	return res.TaskResponse{
		TaskID:        task.TaskID,
		UserID:        task.UserID,
		TaskName:      task.TaskName,
		Description:   task.Description,
		Priority:      task.Priority,
		EstimatedTime: task.EstimatedTime,
		Status:        task.Status,
		CreatedAt:     task.CreatedAt,
		DueDate:       task.DueDate,
		LastUpdated:   task.LastUpdated,
	}
}

func ConvertTaskAna(taskList []entity.Task) []res.TaskResponse {
	var taskResList []res.TaskResponse
	for _, task := range taskList {
		taskResList = append(taskResList, ConvertTaskEntityToTaskRes(task))
	}
	return taskResList
}

func ConvertTaskToTaskOpenai(task entity.Task) entity.TaskOpenai {
	return entity.TaskOpenai{
		TaskID:        task.TaskID,
		UserID:        task.UserID,
		TaskName:      task.TaskName,
		Description:   task.Description,
		Priority:      task.Priority,
		EstimatedTime: task.EstimatedTime,
		Status:        task.Status,
		DueDate:       task.DueDate.Format("2006-01-02 15:04:05"),
		LastUpdated:   task.LastUpdated.Format("2006-01-02 15:04:05"),
	}
}
