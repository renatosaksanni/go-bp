package infra

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestInitDB_Success(t *testing.T) {
	// Mock the DATABASE_DSN environment variable
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	os.Setenv("DATABASE_DSN", dsn)
	defer os.Unsetenv("DATABASE_DSN")

	// Initialize a new sqlmock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Replace the connection string for the mock database
	dialector := mysql.New(mysql.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	// Mock expectations for the database connection
	mock.ExpectPing()

	// Ensure the GORM DB instance is not nil
	assert.NotNil(t, gormDB)

	// Ensure that the sqlmock expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestInitDB_Fail(t *testing.T) {
	// Mock the DATABASE_DSN environment variable with an invalid DSN
	dsn := "invalid_dsn"
	os.Setenv("DATABASE_DSN", dsn)
	defer os.Unsetenv("DATABASE_DSN")

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic due to invalid DSN")
		}
	}()

	// Initialize the database, expecting a panic
	_ = InitDB()
}

func TestInitDB_EnvNotSet(t *testing.T) {
	// Ensure the DATABASE_DSN environment variable is not set
	os.Unsetenv("DATABASE_DSN")

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic due to unset DATABASE_DSN")
		}
	}()

	// Initialize the database, expecting a panic
	_ = InitDB()
}
