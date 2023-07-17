package threads

import (
	"reflect"
	"testing"
)

func TestInitConfig(t *testing.T) {
	type args struct {
		configs []ConfigFn
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name:    "update credential",
			args:    args{[]ConfigFn{WithCridential("username", "password")}},
			want:    &Config{},
			wantErr: false,
		},
		{
			name: "update device id",
			args: args{[]ConfigFn{WithDeviceID("deviceID1")}},
			want: &Config{
				DeviceID: "deviceID1",
			},
			wantErr: false,
		},
		{
			name: "update user id",
			args: args{[]ConfigFn{WithUserID(1)}},
			want: &Config{
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitConfig(tt.args.configs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
