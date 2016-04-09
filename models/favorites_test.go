package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/eirka/eirka-libs/db"
	e "github.com/eirka/eirka-libs/errors"
)

func TestFavoritesValidateInput(t *testing.T) {

	var err error

	favorite := FavoritesModel{
		UID:   1,
		Image: 1,
	}

	err = favorite.ValidateInput()
	if assert.Error(t, err, "An error was expected") {
		assert.Equal(t, err, e.ErrInvalidParam, "Error should match")
	}

}

func TestFavoritesIsValid(t *testing.T) {

	favorite := FavoritesModel{
		UID:   1,
		Image: 1,
	}

	assert.False(t, favorite.IsValid(), "Should be false")

}

func TestFavoritesStatus(t *testing.T) {

	var err error

	mock, err := db.NewTestDb()
	assert.NoError(t, err, "An error was not expected")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery(`SELECT count\(1\) FROM favorites`).WillReturnRows(rows)

	favorite := FavoritesModel{
		UID:   2,
		Image: 1,
	}

	err = favorite.Status()
	assert.NoError(t, err, "An error was not expected")

	assert.NoError(t, mock.ExpectationsWereMet(), "An error was not expected")

}

func TestFavoritesStatusRemove(t *testing.T) {

	var err error

	mock, err := db.NewTestDb()
	assert.NoError(t, err, "An error was not expected")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery(`SELECT count\(1\) FROM favorites`).WillReturnRows(rows)

	mock.ExpectExec("DELETE FROM favorites").
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	favorite := FavoritesModel{
		UID:   2,
		Image: 1,
	}

	err = favorite.Status()
	if assert.Error(t, err, "An error was expected") {
		assert.Equal(t, err, e.ErrFavoriteRemoved, "Error should match")
	}

	assert.NoError(t, mock.ExpectationsWereMet(), "An error was not expected")

}

func TestFavoritesPost(t *testing.T) {

	var err error

	mock, err := db.NewTestDb()
	assert.NoError(t, err, "An error was not expected")

	mock.ExpectExec("INSERT into favorites").
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	favorite := FavoritesModel{
		UID:   2,
		Image: 1,
	}

	err = favorite.Post()
	assert.NoError(t, err, "An error was not expected")

	assert.NoError(t, mock.ExpectationsWereMet(), "An error was not expected")

}

func TestFavoritesPostInvalid(t *testing.T) {

	var err error

	favorite := FavoritesModel{
		UID:   1,
		Image: 1,
	}

	err = favorite.Post()
	if assert.Error(t, err, "An error was expected") {
		assert.Equal(t, err, errors.New("FavoritesModel is not valid"), "Error should match")
	}

}
