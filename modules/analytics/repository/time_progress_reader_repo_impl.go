package repository

import (
	"backend-golang/modules/analytics/domain/entity"
	database "backend-golang/pkgs/dbs/postgres"
)

type timeProgressReaderRepoImpl struct {
	db database.Database
}

var _ TimeProgressReaderRepo = (*timeProgressReaderRepoImpl)(nil)

func NewTimeProgressReaderRepository(db database.Database) TimeProgressReaderRepo {
	return &timeProgressReaderRepoImpl{
		db: db,
	}
}

func (repo timeProgressReaderRepoImpl) GetEachTaskProgress(userId int) (*[]entity.TaskProgressEntity, error) {
	var results *[]entity.TaskProgressEntity
	query := `
			SELECT task."taskId", "taskName", "status",
			SUM("sessionEnd" - "sessionStart") as "totalTimeSpent", "estimatedTime"
			FROM "Tasks" task JOIN "TimeProgressHistory" time ON task."taskId" = time."taskId"
			WHERE task."userId" = ?
			GROUP BY task."taskId", "taskName", "description", "estimatedTime"
			ORDER BY task."taskId"
		`

	if err := repo.db.Executor.Raw(query, userId).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (repo timeProgressReaderRepoImpl) GetTaskNumberByStatus(userId int, startTime string, endTime string) (*[]entity.TaskNumberByStatusEntity, error) {
	return nil, nil
}

func (repo timeProgressReaderRepoImpl) GetTimeSpentDaily(userId int, startTime string, endTime string) (*[]entity.DailyProgressEntity, error) {
	var results *[]entity.DailyProgressEntity
	query := `
        SELECT day,
            SUM(
                LEAST("sessionEnd", day::DATE + INTERVAL '1 day') - GREATEST("sessionStart", day::DATE)
            ) AS "TotalTime"
        FROM 
            generate_series(?::DATE, ?::DATE, '1 day') AS day LEFT JOIN "TimeProgressHistory"
        ON 
            "sessionStart" <= day::DATE + INTERVAL '1 day' AND 
            "sessionEnd" >= day::DATE
		WHERE "userId" = ?
		GROUP BY day
		ORDER BY day
    `

	if err := repo.db.Executor.Raw(query, startTime, endTime, userId).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (repo timeProgressReaderRepoImpl) GetTotalTimeSpent(userId int, startTime string, endTime string) (string, error) {
	var results string
	query := `
        SELECT
    `

	if err := repo.db.Executor.Raw(query, startTime, endTime).Scan(&results).Error; err != nil {
		return "", err
	}

	return results, nil
}
