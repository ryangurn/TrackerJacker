package windows

import "testing"

func TestSoftwareExist(t *testing.T) {
	type args struct {
		software string
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
			if gotRetBool := SoftwareExist(tt.args.software); gotRetBool != tt.wantRetBool {
				t.Errorf("SoftwareExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestSoftwareMeta(t *testing.T) {
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
		// TODO: Add test cases.
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