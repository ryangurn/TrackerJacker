package windows

import "testing"

func TestBitlockerDriveExist(t *testing.T) {
	type args struct {
		drive string
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
			if gotRetBool := BitlockerDriveExist(tt.args.drive); gotRetBool != tt.wantRetBool {
				t.Errorf("BitlockerDriveExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestBitlockerDriveMeta(t *testing.T) {
	type args struct {
		drive string
		key   string
		value interface{}
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
			if gotRetBool := BitlockerDriveMeta(tt.args.drive, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("BitlockerDriveMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestBitlockerParse(t *testing.T) {
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
			if gotRetBool := BitlockerParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("BitlockerParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}
