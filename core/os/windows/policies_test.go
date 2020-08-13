package windows

import (
	"golang.org/x/sys/windows/registry"
	"testing"
)

func TestPolicyParse(t *testing.T) {
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
			if gotRetBool := PolicyParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("PolicyParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestPolicyValue(t *testing.T) {
	type args struct {
		k     registry.Key
		path  string
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
			if gotRetBool := PolicyValue(tt.args.k, tt.args.path, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("PolicyValue() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func Test_checkBinary(t *testing.T) {
	type args struct {
		k     registry.Key
		path  string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkBinary(tt.args.k, tt.args.path, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("checkBinary() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func Test_checkInteger64(t *testing.T) {
	type args struct {
		k     registry.Key
		path  string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkInteger64(tt.args.k, tt.args.path, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkInteger64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("checkInteger64() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func Test_checkMUI(t *testing.T) {
	type args struct {
		k     registry.Key
		path  string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkMUI(tt.args.k, tt.args.path, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkMUI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("checkMUI() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func Test_checkString(t *testing.T) {
	type args struct {
		k     registry.Key
		path  string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkString(tt.args.k, tt.args.path, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("checkString() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func Test_checkStrings(t *testing.T) {
	type args struct {
		k     registry.Key
		path  string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkStrings(tt.args.k, tt.args.path, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("checkStrings() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}