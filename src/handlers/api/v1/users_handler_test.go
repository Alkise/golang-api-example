package v1_test

import (
	"database"
	"handlers/api/v1"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const timeFormat = "Jan 01, 2009 at 00:00am (UTC)"

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error ‘%s’ was not expected when opening a stub database connection", err)
	}
	database.Pool = db
	defer database.Pool.Close()

	userTime, _ := time.Parse("", "")

	countRow := sqlmock.NewRows([]string{"count"}).AddRow(1)
	userRow := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "created_at", "updated_at"}).AddRow(1, "email@email.email", "Name", "Surname", userTime, userTime)
	mock.ExpectQuery(`SELECT COUNT\(\*\) AS count FROM users`).WillReturnRows(countRow)
	mock.ExpectQuery(`SELECT (.+) FROM users`).WillReturnRows(userRow)

	req, err := http.NewRequest("GET", "/api/v1/users", nil)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(v1.UsersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `[{"id":1,"created_at":"0000-01-01T00:00:00Z","updated_at":"0000-01-01T00:00:00Z","email":"email@email.email","first_name":"Name","last_name":"Surname"}]`
	if rr.Body.String() != expected {
		t.Errorf("expected %s, got %s", expected, rr.Body.String())
	}
}
