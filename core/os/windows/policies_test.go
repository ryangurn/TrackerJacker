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
	tests := []struct {
		name        string
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkBinary()
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

func Test_checkInteger32(t *testing.T) {
	tests := []struct {
		name        string
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkInteger32()
			if (err != nil) != tt.wantErr {
				t.Errorf("checkInteger32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("checkInteger32() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func Test_checkMUI(t *testing.T) {
	tests := []struct {
		name        string
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkMUI()
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
	tests := []struct {
		name        string
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkString()
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
	tests := []struct {
		name        string
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := checkStrings()
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

func Test_getInteger64(t *testing.T) {
	tests := []struct {
		name        string
		wantRetBool bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetBool, err := getInteger64()
			if (err != nil) != tt.wantErr {
				t.Errorf("getInteger64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetBool != tt.wantRetBool {
				t.Errorf("getInteger64() gotRetBool = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}