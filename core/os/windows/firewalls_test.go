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
		{ name: "Firewall Enabled Invalid Args", args: args{"-"}, wantRetBool: false},
		{ name: "Firewall Enabled No Args", args: args{}, wantRetBool: false},
		{ name: "Firewall Enabled Valid Args - Public Exist", args: args{"Public"}, wantRetBool: true},
		{ name: "Firewall Enabled Valid Args - Domain Exist", args: args{"Domain"}, wantRetBool: true},
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
		{  name: "Firewall Meta Invalid Args", args: args{firewall: "Pu", key: "Enabled", value: "0"}, wantRetBool: false},
		{  name: "Firewall Meta No Args", args: args{}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - Profile True", args: args{firewall: "Public", key: "Profile", value: "Public"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - Profile False", args: args{firewall: "Public", key: "Profile", value: "g"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - Enabled True", args: args{firewall: "Public", key: "Enabled", value: "1"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - Enabled False", args: args{firewall: "Public", key: "Enabled", value: "0"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - DefaultInboundAction True", args: args{firewall: "Public", key: "DefaultInboundAction", value: "0"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - DefaultInboundAction False", args: args{firewall: "Public", key: "DefaultInboundAction", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - DefaultOutboundAction True", args: args{firewall: "Public", key: "DefaultOutboundAction", value: "0"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - DefaultOutboundAction False", args: args{firewall: "Public", key: "DefaultOutboundAction", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - AllowInboundRules True", args: args{firewall: "Public", key: "AllowInboundRules", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - AllowInboundRules False", args: args{firewall: "Public", key: "AllowInboundRules", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - AllowLocalFirewallRules True", args: args{firewall: "Public", key: "AllowLocalFirewallRules", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - AllowLocalFirewallRules False", args: args{firewall: "Public", key: "AllowLocalFirewallRules", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - AllowLocalIPsecRules True", args: args{firewall: "Public", key: "AllowLocalIPsecRules", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - AllowLocalIPsecRules False", args: args{firewall: "Public", key: "AllowLocalIPsecRules", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - AllowUserApps True", args: args{firewall: "Public", key: "AllowUserApps", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - AllowUserApps False", args: args{firewall: "Public", key: "AllowUserApps", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - AllowUserPorts True", args: args{firewall: "Public", key: "AllowUserPorts", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - AllowUserPorts False", args: args{firewall: "Public", key: "AllowUserPorts", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - AllowUnicastResponseToMulticast True", args: args{firewall: "Public", key: "AllowUnicastResponseToMulticast", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - AllowUnicastResponseToMulticast False", args: args{firewall: "Public", key: "AllowUnicastResponseToMulticast", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - NotifyOnListen True", args: args{firewall: "Public", key: "NotifyOnListen", value: "1"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - NotifyOnListen False", args: args{firewall: "Public", key: "NotifyOnListen", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - EnableStealthModeForIPsec True", args: args{firewall: "Public", key: "EnableStealthModeForIPsec", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - EnableStealthModeForIPsec False", args: args{firewall: "Public", key: "EnableStealthModeForIPsec", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - LogMaxSizeKilobytes True", args: args{firewall: "Public", key: "LogMaxSizeKilobytes", value: "4096"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - LogMaxSizeKilobytes False", args: args{firewall: "Public", key: "LogMaxSizeKilobytes", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - LogAllowed True", args: args{firewall: "Public", key: "LogAllowed", value: "0"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - LogAllowed False", args: args{firewall: "Public", key: "LogAllowed", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - LogBlocked True", args: args{firewall: "Public", key: "LogBlocked", value: "0"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - LogBlocked False", args: args{firewall: "Public", key: "LogBlocked", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - LogIgnored True", args: args{firewall: "Public", key: "LogIgnored", value: "2"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - LogIgnored False", args: args{firewall: "Public", key: "LogIgnored", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - Caption True", args: args{firewall: "Public", key: "Caption", value: nil}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - Caption False", args: args{firewall: "Public", key: "Caption", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - Description True", args: args{firewall: "Public", key: "Description", value: nil}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - Description False", args: args{firewall: "Public", key: "Description", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - ElementName True", args: args{firewall: "Public", key: "ElementName", value: ""}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - ElementName False", args: args{firewall: "Public", key: "ElementName", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - InstanceID True", args: args{firewall: "Public", key: "InstanceID", value: "MSFT|FW|FirewallProfile|Public"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - InstanceID False", args: args{firewall: "Public", key: "InstanceID", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - LogFileName True", args: args{firewall: "Public", key: "LogFileName", value: "%systemroot%\\system32\\LogFiles\\Firewall\\pfirewall.log"}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - LogFileName False", args: args{firewall: "Public", key: "LogFileName", value: "-1"}, wantRetBool: false},
		{  name: "Firewall Meta Valid Args - PSComputerName True", args: args{firewall: "Public", key: "PSComputerName", value: nil}, wantRetBool: true},
		{  name: "Firewall Meta Valid Args - PSComputerName False", args: args{firewall: "Public", key: "PSComputerName", value: "-1"}, wantRetBool: false},
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
		{ name: "Firewall Parse Invalid Args", args: args{args: []string{"-a"}, result: true}, wantRetBool: false},
		{ name: "Firewall Parse No Args", args: args{}, wantRetBool: false},
		{ name: "Firewall Parse Valid Args - Enabled", args: args{args: []string{"enabled", "Public"}, result: true}, wantRetBool: true},
		{ name: "Firewall Parse Valid Args - Not Enabled", args: args{args: []string{"enabled", "Pubic"}, result: true}, wantRetBool: false},


		{ name: "Firewall Parse Meta - Nominal Args", args: args{args: []string{"meta", "Public", "Enabled", "1"}, result: true}, wantRetBool: true},
		{ name: "Firewall Parse Meta - No Args", args: args{}, wantRetBool: false},
		{ name: "Firewall Parse Meta - Invalid Args", args: args{args: []string{"meta", "-1"}, result: true}, wantRetBool: false},
		{ name: "Firewall Parse Meta - Invalid Args", args: args{args: []string{"meta", "Puic", "Enabled", "1"}, result: true}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := FirewallParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("FirewallParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}