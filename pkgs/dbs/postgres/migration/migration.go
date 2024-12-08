package migration

import (
	"gorm.io/gorm"

	user "backend-golang/modules/user/domain/entity"
)

// Migration is a function that performs migrations on the given database.
//
// It takes a pointer to a gorm.DB object as a parameter.
// It returns an error indicating the success or failure of the migration operation.
func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(
		user.UserEntity{},
	)

	return err
}
