package invoices_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"upsider-coding-test/cmd/api/view/request/invoices"
)

func TestNewGetRequest(t *testing.T) {
	tests := []struct {
		name    string
		query   map[string]string
		wantErr bool
	}{
		{
			name: "valid dates",
			query: map[string]string{
				"startAt": "20230101",
				"endAt":   "20230102",
			},
			wantErr: false,
		},
		{
			name: "invalid start date format",
			query: map[string]string{
				"startAt": "202301",
				"endAt":   "20230102",
			},
			wantErr: true,
		},
		{
			name: "invalid end date format",
			query: map[string]string{
				"startAt": "20230101",
				"endAt":   "202301",
			},
			wantErr: true,
		},
		{
			name: "start date after end date",
			query: map[string]string{
				"startAt": "20230102",
				"endAt":   "20230101",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			q := req.URL.Query()
			for key, value := range tt.query {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()
			ctx.Request = req

			_, err := invoices.NewGetRequest(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGetRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
