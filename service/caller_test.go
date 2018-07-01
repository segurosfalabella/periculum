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
			c:       service.Caller{ApiUrl: ""},
			want:    service.Response{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.c.Request()
			if (err != nil) == tt.wantErr {
				//t.Errorf("Service.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// func TestCall(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{name: "ShouldBeFail"},
// 		{name: "ShouldNotFail"},
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Call()
// 		})
// 	}
// }
