package postgres

import (
	"backend-golang/pkgs/dbs/postgres/migration"
	"backend-golang/pkgs/log"
	"gorm.io/gorm/logger"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Executor *gorm.DB
}

// InitDatabase initializes a new database instance and performs a database migration if allowed.
//
// Parameters:
// - conn: a Connection object representing the database connection
//
// Returns:
// - a pointer to the initialized Database object
func InitDatabase(conn Connection) *Database {

	// Create a new database instance
	db, err := NewDatabase(conn)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	allowMigrate, _ := strconv.ParseBool(os.Getenv("CONFIG_ALLOW_MIGRATION"))

	// Perform database migration if allowed
	if allowMigrate {
		err := migration.Migration(db.Executor)
		if err != nil {
			log.JsonLogger.Error(err.Error())
			return nil
		}
	}

	return db
}

// NewDatabase opens a database connection using the gorm library and returns a *Database object.
//
// It takes a Connection object as a parameter and returns a *Database object and an error.
func NewDatabase(conn Connection) (*Database, error) {
	// Open a database connection using the gorm library
	db, err := gorm.Open(postgres.Open(conn.ToPostgresConnectionString()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	// Get the underlying *sql.DB instance from the gorm.DB object
	settingDb, err := db.DB()
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	// Ping the database to check the connection
	if pingError := settingDb.Ping(); pingError != nil {
		log.JsonLogger.Error(pingError.Error())
		panic(pingError)
	}

	// Configure the connection pool settings
	settingDb.SetMaxOpenConns(conn.MaxOpenConnections)
	settingDb.SetMaxIdleConns(conn.MaxIdleConnections)
	settingDb.SetConnMaxIdleTime(conn.ConnectionMaxIdleTime)
	settingDb.SetConnMaxLifetime(conn.ConnectionMaxLifeTime)

	// Return the Database struct with the gorm.DB object
	return &Database{Executor: db}, nil
}
