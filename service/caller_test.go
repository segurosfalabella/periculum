package service_test

import (
	"testing"

	"github.com/mrsangrin/periculum/godogs/drivers"
	"github.com/mrsangrin/periculum/service"
)

func TestService_Request(t *testing.T) {
	defer drivers.CloseServer()
	drivers.RunApp()

	tests := []struct {
		name    string
		c       service.Caller
		want    service.Response
		wantErr bool
	}{
		{
			name:    "ShouldFailIfServiceUrlIsEmpty",
			c:       service.Caller{Endpoint: ""},
			want:    service.Response{},
			wantErr: true,
		},
		{
			name:    "ShouldFetchResponseFromRemoteService",
			c:       service.Caller{Endpoint: "http://localhost:3001/health"},
			want:    service.Response{StatusCode: 200, Message: "OK"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.c.Request()
			if (err != nil) == tt.wantErr {
				return
			}

			if res.StatusCode == tt.want.StatusCode {
				return
			}

		})
	}
}
