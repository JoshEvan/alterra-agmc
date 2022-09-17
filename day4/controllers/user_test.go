package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JoshEvan/alterra-agmc-day4/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func fixturesSetUp() {
	config.InitDB()
}

func TestGetUsers(t *testing.T) {
	fixturesSetUp()
	type args struct {
		c   echo.Context
		rec *httptest.ResponseRecorder
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantStatusCode int
	}{
		{
			name: "success",
			args: func() args {
				rec := httptest.NewRecorder()
				return args{
					c:   echo.New().NewContext(httptest.NewRequest(http.MethodGet, "/", nil), rec),
					rec: rec,
				}
			}(),
			wantErr:        false,
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUsers(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantStatusCode, tt.args.rec.Code)
		})
	}
}
