package xerrs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultHTTPStatus(t *testing.T) {
	tests := []struct {
		name      string
		errorType ErrorType
		expected  int
	}{
		{"Validation Error", ErrorTypeValidation, http.StatusBadRequest},
		{"Authentication Error", ErrorTypeAuthentication, http.StatusUnauthorized},
		{"Authorization Error", ErrorTypeAuthorization, http.StatusForbidden},
		{"Not Found Error", ErrorTypeNotFound, http.StatusNotFound},
		{"Conflict Error", ErrorTypeConflict, http.StatusConflict},
		{"Rate Limit Error", ErrorTypeRateLimit, http.StatusTooManyRequests},
		{"Internal Error", ErrorTypeInternal, http.StatusInternalServerError},
		{"External Error", ErrorTypeExternal, http.StatusBadGateway},
		{"Unavailable Error", ErrorTypeUnavailable, http.StatusServiceUnavailable},
		{"Unknown Error Type", ErrorType("UNKNOWN"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.errorType.DefaultHTTPStatus()
			assert.Equal(t, tt.expected, result)
		})
	}
}
