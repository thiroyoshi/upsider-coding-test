package invoices_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"upsider-coding-test/cmd/api/view/request/invoices"
)

func TestNewPostRequest(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "valid request",
			body:    `{"partnerId":"123e4567-e89b-12d3-a456-426614174000","amount":100,"paymentDue":"20260101"}`,
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			body:    `{"partnerId":"123e4567-e89b-12d3-a456-426614174000","amount":100,"paymentDue":"20240101"`,
			wantErr: true,
		},
		{
			name:    "empty request body",
			body:    `{}`,
			wantErr: true,
		},
		{
			name:    "invalid UUID",
			body:    `{"partnerId":"invalid-uuid","amount":100,"paymentDue":"20240101"}`,
			wantErr: true,
		},
		{
			name:    "invalid amount",
			body:    `{"partnerId":"123e4567-e89b-12d3-a456-426614174000","amount":-100,"paymentDue":"20240101"}`,
			wantErr: true,
		},
		{
			name:    "invalid date format",
			body:    `{"partnerId":"123e4567-e89b-12d3-a456-426614174000","amount":100,"paymentDue":"202401"}`,
			wantErr: true,
		},
		{
			name:    "past date",
			body:    `{"partnerId":"123e4567-e89b-12d3-a456-426614174000","amount":100,"paymentDue":"20220101"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			_, err := invoices.NewPostRequest(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
