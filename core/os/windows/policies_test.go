package windows

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"testing"
)

func TestPolicyParse(t *testing.T) {
	policy, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	i, _, err := policy.GetStringValue("SystemRoot")
	if err != nil {
		return
	}

	fmt.Println(i)

	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Policy Parse Invalid Args", args: args{args: []string{"-"}, result: true}, wantRetBool: false },
		{ name: "Policy Parse No Args", args: args{}, wantRetBool: false },
		{ name: "Policy Parse Valid Args - Value", args: args{args: []string{"value", "LOCAL_MACHINE", "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", "SystemRoot", i}, result: true}, wantRetBool: true },
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
	// string, int64, binary
	policy, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer policy.Close()

	// string
	s, _, err := policy.GetStringValue("SystemRoot")
	if err != nil {
		return
	}

	// int64
	i, _, err := policy.GetStringValue("UBR")
	if err != nil {
		return
	}

	// binary
	b, _, err := policy.GetStringValue("DigitalProductId")
	if err != nil {
		return
	}

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
		{name: "Policy Value No Args", args: args{}, wantRetBool: false},
		{name: "Policy Value Invalid Args", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "Systemoot", value: s}, wantRetBool: false},
		{name: "Policy Value Lack of Args", args: args{k: registry.LOCAL_MACHINE}, wantRetBool: false},
		{name: "Policy Value Valid Args - String True", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "SystemRoot", value: s}, wantRetBool: true},
		{name: "Policy Value Valid Args - String False", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "SystmRoot", value: s}, wantRetBool: false},
		{name: "Policy Value Valid Args - Int64 True", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "UBR", value: i}, wantRetBool: true},
		{name: "Policy Value Valid Args - Int64 False", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "BR", value: i}, wantRetBool: false},
		{name: "Policy Value Valid Args - Binary True", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "DigitalProductId", value: b}, wantRetBool: true},
		{name: "Policy Value Valid Args - Binary False", args: args{k: registry.LOCAL_MACHINE, path: "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", key: "DiitalProductId", value: b}, wantRetBool: false},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := PolicyValue(tt.args.k, tt.args.path, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("PolicyValue() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}
