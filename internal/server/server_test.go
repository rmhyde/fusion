package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rmhyde/fusion/internal/helpers"
)

func TestServer(t *testing.T) {

	tcs := []struct {
		name         string
		folder       string
		serverPath   string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "TestServer_RootMessage",
			folder:       "",
			serverPath:   "",
			expectedCode: http.StatusOK,
			expectedBody: "Boards can be found in /api/boards",
		},
		{
			name:         "TestServer_404",
			folder:       "",
			serverPath:   "/api",
			expectedCode: http.StatusNotFound,
			expectedBody: "404 page not found\n",
		},
		{
			name:         "TestServer_GetBoards",
			folder:       "testdata",
			serverPath:   "/api/boards",
			expectedCode: http.StatusOK,
			expectedBody: "{\"Boards\":[{\"Name\":\"A1-100X\",\"Vendor\":\"Boards R Us\",\"Core\":\"Cortex-M7\",\"has_wifi\":true},{\"Name\":\"B7-400X\",\"Vendor\":\"Boards R Us\",\"Core\":\"Cortex-M7\",\"has_wifi\":true},{\"Name\":\"C1-100X\",\"Vendor\":\"Boards R Us\",\"Core\":\"Cortex-M7\",\"has_wifi\":false},{\"Name\":\"Low_Power\",\"Vendor\":\"Tech Corp.\",\"Core\":\"Cortex-M0+\",\"has_wifi\":false}],\"_metadata\":{\"Totals\":{\"Vendors\":2,\"Boards\":4,\"wifi_enabled\":2},\"Errors\":{\"has_errors\":false}}}",
		},
		{
			name:         "TestServer_GetBoardsWithInvalidFolder",
			folder:       "invalid/board/folder",
			serverPath:   "/api/boards",
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Something went wrong\n",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			options := Options{
				Folder: tc.folder,
				Ctx:    helpers.NewTestWriterContext(t),
			}
			s := httptest.NewServer(options.newRouter())
			defer s.Close()

			resp, err := s.Client().Get(s.URL + tc.serverPath)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Errorf("Expected status OK, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if string(body) != tc.expectedBody {
				t.Errorf("Expected body %q, got %q", tc.expectedBody, string(body))
			}
		})
	}
}
