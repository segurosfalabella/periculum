package application_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/mrsangrin/periculum/application"
)

var tempFilePATH = "../tmp/apps.yml"
var yamlFile = `version: 3.2	
services:
  - name: dummy
    description: dummy
    endpoint: dummy`

func TestRemoteServices_GetApps(t *testing.T) {
	ioutil.WriteFile(tempFilePATH, []byte(yamlFile), 0644)

	tests := []struct {
		name    string
		c       application.RemoteServices
		want    []application.Service
		wantErr bool
	}{
		{
			name:    "MustFetchServicesFromYmlFile",
			c:       application.RemoteServices{},
			want:    []application.Service{{Name: "dummy", Description: "dummy", Endpoint: "dummy"}},
			wantErr: false,
		},
		{
			name:    "ShouldBeFailIfYAMLHasNotService",
			c:       application.RemoteServices{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetApps(tempFilePATH).Services; !reflect.DeepEqual(got, tt.want) != tt.wantErr {
				t.Errorf("RemoteServices.GetApps() = %v, want %v", got, tt.want)
			}

			if got := tt.c.GetApps("../missing-file.yml").Services; (len(got) == 0) == tt.wantErr {
				return
			}
		})
	}
}
