package cross

import "testing"

func TestHostExist(t *testing.T) {
	type args struct {
		desiredHost string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Host Exist Valid Host", args: args{desiredHost: "localhost"}, wantRetBool: true},
		{ name: "Host Exist Invalid Host", args: args{desiredHost: "google.com"}, wantRetBool: false},
		{ name: "Host Exist Invalid Args", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := HostExist(tt.args.desiredHost); gotRetBool != tt.wantRetBool {
				t.Errorf("HostExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestHostIpExist(t *testing.T) {
	type args struct {
		desiredIp string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Host Exist Valid IP", args: args{desiredIp: "127.0.0.1"}, wantRetBool: true},
		{ name: "Host Exist Invalid IP", args: args{desiredIp: "192.168.1.1"}, wantRetBool: false},
		{ name: "Host Exist Invalid Args", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := HostIpExist(tt.args.desiredIp); gotRetBool != tt.wantRetBool {
				t.Errorf("HostIpExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestHostParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Host Parse Valid Args", args: args{args: []string{"host", "exist", "localhost"}, result: true}, wantRetBool: true},
		{ name: "Host Parse Valid Args - Invalid Host", args: args{args: []string{"host", "exist", "local.host"}, result: false}, wantRetBool: true},
		{ name: "Host Parse Valid Args - No Host", args: args{args: []string{"host", "exist"}, result: false}, wantRetBool: false},
		{ name: "Host Parse Valid Args", args: args{args: []string{"ip", "exist", "127.0.0.1"}, result: true}, wantRetBool: true},
		{ name: "Host Parse Valid Args - Invalid Host", args: args{args: []string{"ip", "exist", "local."}, result: false}, wantRetBool: true},
		{ name: "Host Parse Valid Args - No Host", args: args{args: []string{"ip", "exist"}, result: false}, wantRetBool: false},
		{ name: "Host Parse Valid Args - No Args", args: args{args: []string{}, result: false}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := HostParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("HostParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}