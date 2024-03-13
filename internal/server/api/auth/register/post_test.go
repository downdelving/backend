package register_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/downdelving/backend/internal/server/api/auth/register"
	"github.com/downdelving/backend/internal/storage"
)

// setup sets up a test HTTP server along with a register.Handler.
// Returns the server and the response recorder.
func setup(t *testing.T) (*httptest.ResponseRecorder, *http.ServeMux) {
	t.Helper()
	rec := httptest.NewRecorder()
	accountStorage := storage.NewInMemoryAccountStorage()
	registerHandler := register.NewHandler(accountStorage)
	server := http.NewServeMux()
	server.HandleFunc("/api/auth/register", registerHandler.Post)
	return rec, server
}

func TestRegisterEndpoint(t *testing.T) {
	tests := []struct {
		name           string
		payload        []byte
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Empty Body",
			payload:        []byte(""),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing request body\n",
		},
		{
			name:           "Invalid Payload",
			payload:        []byte(`{"invalid": "json"`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request payload\n",
		},
		{
			name:           "Missing Fields (all)",
			payload:        []byte(`{}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing required fields\n",
		},
		{
			name:           "Missing Fields (password, email)",
			payload:        []byte(`{"username":"test"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing required fields\n",
		},
		{
			name:           "Missing Fields (username, email)",
			payload:        []byte(`{"password":"password"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing required fields\n",
		},
		{
			name:           "Missing Fields (email)",
			payload:        []byte(`{"username":"test","password":"password"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing required fields\n",
		},
		{
			name:           "Missing Fields (username, password)",
			payload:        []byte(`{"email":"user@domain.com"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing required fields\n",
		},
		{
			name:           "Valid Request",
			payload:        []byte(`{"username":"test","password":"password","email":"user@domain.com"}`), // Adjust this payload as necessary.
			expectedStatus: http.StatusCreated,
			expectedBody:   "{\"message\": \"Successfully registered!\"}", // Adjust or check for the expected body of a valid request.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec, server := setup(t)
			reader := bytes.NewReader(tt.payload)
			req, err := http.NewRequest("POST", "/api/auth/register", reader)
			if err != nil {
				t.Fatal(err)
			}

			server.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q; got %q", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
