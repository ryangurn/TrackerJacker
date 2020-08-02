package windows

import (
	"testing"
	"time"
)

func TestSoftwareExist(t *testing.T) {
	type args struct {
		software string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Software Exist Valid Software", args: args{software: "Microsoft Edge"}, wantRetBool: true},
		{ name: "Software Exist Invalid Software", args: args{software: "Edge"}, wantRetBool: false},
		{ name: "Software Exist Empty Args", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := SoftwareExist(tt.args.software); gotRetBool != tt.wantRetBool {
				t.Errorf("SoftwareExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestSoftwareMeta(t *testing.T) {
	testingTime := time.Date(2020, time.August, 1, 0,0,0,0, time.UTC)

	type args struct {
		software string
		key      string
		value    interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Software Invalid Args",  args: args{software: "Microsoft Edge"}, wantRetBool: false},
		{ name: "Software No Args",  args: args{}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Display True",  args: args{software: "Microsoft Edge", key: "DisplayVersion", value: "84.0.522.52"}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Display False",  args: args{software: "Microsoft Edge", key: "DisplayVersion", value: "83.0.47.61"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Arch True",  args: args{software: "Microsoft Edge", key: "Arch", value: "X32"}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Arch False",  args: args{software: "Microsoft Edge", key: "Arch", value: "X31"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Publisher True",  args: args{software: "Microsoft Edge", key: "Publisher", value: "Microsoft Corporation"}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Publisher False",  args: args{software: "Microsoft Edge", key: "Publisher", value: "Inc"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - InstallDate True",  args: args{software: "Microsoft Edge", key: "InstallDate", value: testingTime}, wantRetBool: true},
		{ name: "Software Meta Valid Args - InstallDate False",  args: args{software: "Microsoft Edge", key: "InstallDate", value: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := SoftwareMeta(tt.args.software, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("SoftwareMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestSoftwareParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := SoftwareParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("SoftwareParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}