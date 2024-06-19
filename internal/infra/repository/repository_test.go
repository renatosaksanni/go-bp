package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewRepository(t *testing.T) {
	// Initialize a new sqlmock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Use the sqlmock database connection to create a GORM DB instance
	dialector := mysql.New(mysql.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	// Create a new repository using the GORM DB instance
	repo := NewRepository(gormDB)

	// Ensure that the repository's DB instance matches the GORM DB instance
	assert.Equal(t, gormDB, repo.db)

	// Ensure that the sqlmock expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
