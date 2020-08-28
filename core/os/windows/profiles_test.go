package windows

import "testing"

func TestProfileDirectory(t *testing.T) {
	type args struct {
		dir string
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
			if gotRetBool := ProfileDirectory(tt.args.dir); gotRetBool != tt.wantRetBool {
				t.Errorf("ProfileDirectory() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestProfileParse(t *testing.T) {
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
			if gotRetBool := ProfileParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("ProfileParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestProfileUserDirectory(t *testing.T) {
	type args struct {
		dir string
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
			if gotRetBool := ProfileUserDirectory(tt.args.dir); gotRetBool != tt.wantRetBool {
				t.Errorf("ProfileUserDirectory() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}
