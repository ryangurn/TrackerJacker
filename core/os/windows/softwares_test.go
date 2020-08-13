package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"testing"
	"time"
)

func TestSoftwareExist(t *testing.T) {
	softwares, err := wapi.InstalledSoftwareList()
	if err != nil{
		t.Errorf("Unable to load software: %s", err)
	}

	type args struct {
		software string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Software Exist Valid Software", args: args{software: softwares[0].DisplayName}, wantRetBool: true},
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
	softwares, err := wapi.InstalledSoftwareList()
	if err != nil{
		t.Errorf("Unable to load software: %s", err)
	}

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
		{ name: "Software Invalid Args",  args: args{software: softwares[0].DisplayName}, wantRetBool: false},
		{ name: "Software No Args",  args: args{}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Display True",  args: args{software: softwares[0].DisplayName, key: "DisplayVersion", value: softwares[0].DisplayVersion}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Display False",  args: args{software: softwares[0].DisplayName, key: "DisplayVersion", value: "83.0.47.61"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Arch True",  args: args{software: softwares[0].DisplayName, key: "Arch", value: softwares[0].Arch}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Arch False",  args: args{software: softwares[0].DisplayName, key: "Arch", value: "X31"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Publisher True",  args: args{software: softwares[0].DisplayName, key: "Publisher", value: softwares[0].Publisher}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Publisher False",  args: args{software: softwares[0].DisplayName, key: "Publisher", value: "Inc"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - InstallDate True",  args: args{software: softwares[0].DisplayName, key: "InstallDate", value: softwares[0].InstallDate}, wantRetBool: true},
		{ name: "Software Meta Valid Args - InstallDate False",  args: args{software: softwares[0].DisplayName, key: "InstallDate", value: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}, wantRetBool: false},
		{ name: "Software Meta Valid Args - EstimatedSize True", args: args{software: softwares[0].DisplayName, key: "EstimatedSize", value: strconv.FormatUint(softwares[0].EstimatedSize, 10)}, wantRetBool: true},
		{ name: "Software Meta Valid Args - EstimatedSize False", args: args{software: softwares[0].DisplayName, key: "EstimatedSize", value: "-1"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - Contact True", args: args{software: softwares[0].DisplayName, key: "Contact", value: softwares[0].Contact}, wantRetBool: true},
		{ name: "Software Meta Valid Args - Contact False", args: args{software: softwares[0].DisplayName, key: "Contact", value: "-1"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - HelpLink True", args: args{software: softwares[0].DisplayName, key: "HelpLink", value: softwares[0].HelpLink}, wantRetBool: true},
		{ name: "Software Meta Valid Args - HelpLink False", args: args{software: softwares[0].DisplayName, key: "HelpLink", value: "-1"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - InstallSource True", args: args{software: softwares[0].DisplayName, key: "InstallSource", value: softwares[0].InstallSource}, wantRetBool: true},
		{ name: "Software Meta Valid Args - InstallSource False", args: args{software: softwares[0].DisplayName, key: "InstallSource", value: "-1"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - InstallLocation True", args: args{software: softwares[0].DisplayName, key: "InstallLocation", value: softwares[0].InstallLocation}, wantRetBool: true},
		{ name: "Software Meta Valid Args - InstallLocation False", args: args{software: softwares[0].DisplayName, key: "InstallLocation", value: "-1"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - VersionMajor True", args: args{software: softwares[0].DisplayName, key: "VersionMajor", value: strconv.FormatUint(softwares[0].VersionMajor, 10)}, wantRetBool: true},
		{ name: "Software Meta Valid Args - VersionMajor False", args: args{software: softwares[0].DisplayName, key: "VersionMajor", value: "-1"}, wantRetBool: false},
		{ name: "Software Meta Valid Args - VersionMinor True", args: args{software: softwares[0].DisplayName, key: "VersionMinor", value: strconv.FormatUint(softwares[0].VersionMinor, 10)}, wantRetBool: true},
		{ name: "Software Meta Valid Args - VersionMinor False", args: args{software: softwares[0].DisplayName, key: "VersionMinor", value: "-1"}, wantRetBool: false},
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
	softwares, err := wapi.InstalledSoftwareList()
	if err != nil{
		t.Errorf("Unable to load software: %s", err)
	}

	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{name: "Software Parse Invalid Args", args: args{args: []string{"-1"}, result: -1}, wantRetBool: false},
		{name: "Software Parse No Args", args: args{}, wantRetBool: false},
		{name: "Software Parse Valid Args - Exist", args: args{args: []string{"exist", softwares[0].DisplayName}, result: true}, wantRetBool: true},
		{name: "Software Parse Valid Args - Does Not Exist", args: args{args: []string{"exist", "Microsoft"}, result: false}, wantRetBool: true},
		{name: "Software Parse Valid Args - Invalid", args: args{args: []string{"exist"}, result: false}, wantRetBool: false},

		{name: "Software Parse Meta - Nominal Args", args: args{args: []string{"meta", softwares[0].DisplayName, "DisplayVersion", softwares[0].DisplayVersion}, result: true}, wantRetBool: true},
		{name: "Software Parse Meta - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{name: "Software Parse Meta - Invalid Args", args: args{args: []string{"meta", "Microsoft", "DisplayVersion", "84.0.522.52"}, result: false}, wantRetBool: true},
		{name: "Software Parse Meta - Invalid Args", args: args{args: []string{"meta", "Microsoft Edge", "1235", "84.0.522.52"}, result: false}, wantRetBool: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := SoftwareParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("SoftwareParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}