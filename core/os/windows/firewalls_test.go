package windows

import "testing"

func TestFirewallEnabled(t *testing.T) {
	type args struct {
		firewall string
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
			if gotRetBool := FirewallEnabled(tt.args.firewall); gotRetBool != tt.wantRetBool {
				t.Errorf("FirewallEnabled() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestFirewallMeta(t *testing.T) {
	type args struct {
		firewall string
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
			if gotRetBool := FirewallMeta(tt.args.firewall, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("FirewallMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestFirewallParse(t *testing.T) {
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
			if gotRetBool := FirewallParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("FirewallParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}