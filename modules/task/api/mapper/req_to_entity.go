package mapper

import (
	"backend-golang/modules/task/api/model/req"
	"backend-golang/modules/task/constant"

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

func ConvertUpdateTaskListToTaskEntityList(updateTaskList []req.UpdateTaskRequest) []*entity.Task {
	var taskEntityList []*entity.Task
	for _, updateTask := range updateTaskList {
		taskEntity := ConvertUpdateTaskRequestToTaskEntity(updateTask, updateTask.TaskID)
		taskEntityList = append(taskEntityList, &taskEntity)
	}
	return taskEntityList
}

func ConvertListIDRequestToListIDEntity(listID []string) []int {
	var listIDEntity []int
	for _, id := range listID {
		intID, err := strconv.Atoi(id)
		if err != nil {
			panic(constant.ErrrorTaskIDNotInteger(err))
		}
		listIDEntity = append(listIDEntity, intID)
	}
	return listIDEntity
}

func ConvertApplyAnalyzedTaskRequestToTaskApplyAnalyzedEntity(req req.ApplyAnalyzedTaskRequest) entity.TaskApplyAnalyzedTaskEntity {
	var taskEntityList []*entity.Task
	for _, task := range req.TaskList {
		taskEntity := ConvertCreateTaskRequestToTaskEntity(task)
		taskEntityList = append(taskEntityList, &taskEntity)
	}

	taskApplyAnalyzed := entity.TaskApplyAnalyzedTaskEntity{
		TaskList:  taskEntityList,
		StartTime: parseDueDate(req.StartTime),
		EndTime:   parseDueDate(req.EndTime),
	}
	return taskApplyAnalyzed

}

func ConvertUpdateTaskProgressRequestToTaskEntity(req req.UpdateTaskProgressRequest, taskId string) entity.TaskProgress {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		panic(err)
	}

	// Map the request data to the Task entity
	task := entity.TaskProgress{
		TaskID:       id,
		SessionStart: parseDueDate(req.SessionStart),
		SessionEnd:   parseDueDate(req.SessionEnd),
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
