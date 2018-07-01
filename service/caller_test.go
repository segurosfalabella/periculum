package service_test

import (
	"testing"

	"github.com/mrsangrin/periculum/service"
)

func TestService_Request(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.c.Request()
			if (err != nil) == tt.wantErr {
				return
			}
		})
	}
}
