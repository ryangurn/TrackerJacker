package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"testing"
)

func TestServiceExist(t *testing.T) {
	type args struct {
		svc string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{  name: "Service Exist Valid Service", args: args{svc: "AxInstSV"}, wantRetBool: true},
		{  name: "Service Exist Invalid Service", args: args{svc: "AxnstSV"}, wantRetBool: false},
		{  name: "Service Exist Empty Args", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ServiceExist(tt.args.svc); gotRetBool != tt.wantRetBool {
				t.Errorf("ServiceExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestServiceMeta(t *testing.T) {
	services, err := wapi.GetServices()
	if err != nil {
		return
	}

	svc := "AxInstSV"
	DisplayName := ""
	StatusText := ""
	Status := "0"
	AcceptStop := ""
	IsRunning := ""
	RunningPid := ""
	ServiceType := ""

	for _, s := range services {
		if s.SCName == svc {
			DisplayName = s.DisplayName
			StatusText = s.StatusText
			Status = strconv.FormatUint(uint64(s.Status), 10)
			AcceptStop = strconv.FormatBool(s.AcceptStop)
			IsRunning = strconv.FormatBool(s.IsRunning)
			RunningPid = strconv.FormatUint(uint64(s.RunningPid), 10)
			ServiceType = strconv.FormatUint(uint64(s.ServiceType), 10)
		}
	}

	type args struct {
		svc   string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{  name: "Services Invalid Args", args: args{svc: svc}, wantRetBool: false},
		{  name: "Services No Args", args: args{}, wantRetBool: false},
		{  name: "Services Meta Valid Args - DisplayName True", args: args{svc: svc, key: "DisplayName", value: DisplayName}, wantRetBool: true},
		{  name: "Services Meta Valid Args - DisplayName False", args: args{svc: svc, key: "DisplayName", value: "ActveX Installer (AxInstSV)"}, wantRetBool: false},
		{  name: "Services Meta Valid Args - StatusText True", args: args{svc: svc, key: "StatusText", value: StatusText}, wantRetBool: true},
		{  name: "Services Meta Valid Args - StatusText False", args: args{svc: svc, key: "StatusText", value: "Provides User Account Control validation for the installation of ActiveX controls from the Internet and enables management of ActiveX control installation based on Group Policy settings. This service is started on demand and if disabled the installation of ActiveX controls will behave acding to default browser settings."}, wantRetBool: false},
		{  name: "Services Meta Valid Args - Status True", args: args{svc: svc, key: "Status", value: Status}, wantRetBool: true},
		{  name: "Services Meta Valid Args - Status False", args: args{svc: svc, key: "Status", value: "-5486"}, wantRetBool: false},
		{  name: "Services Meta Valid Args - AcceptStop True", args: args{svc: svc, key: "AcceptStop", value: AcceptStop}, wantRetBool: true},
		{  name: "Services Meta Valid Args - AcceptStop False", args: args{svc: svc, key: "AcceptStop", value: "-"}, wantRetBool: false},
		{  name: "Services Meta Valid Args - IsRunning True", args: args{svc: svc, key: "IsRunning", value: IsRunning}, wantRetBool: true},
		{  name: "Services Meta Valid Args - IsRunning False", args: args{svc: svc, key: "IsRunning", value: "-"}, wantRetBool: false},
		{  name: "Services Meta Valid Args - RunningPid True", args: args{svc: svc, key: "RunningPid", value: RunningPid}, wantRetBool: true},
		{  name: "Services Meta Valid Args - RunningPid False", args: args{svc: svc, key: "RunningPid", value: "-1834"}, wantRetBool: false},
		{  name: "Services Meta Valid Args - ServiceType True", args: args{svc: svc, key: "ServiceType", value: ServiceType}, wantRetBool: true},
		{  name: "Services Meta Valid Args - ServiceType False", args: args{svc: svc, key: "ServiceType", value: "-1834"}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ServiceMeta(tt.args.svc, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("ServiceMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestServiceParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Service Parse Invalid Args", args: args{args: []string{"-1"}, result: -1}, wantRetBool: false },
		{ name: "Service Parse No Args", args: args{args: []string{}}, wantRetBool: false},
		{ name: "Service Parse Valid Args - Exist", args: args{args: []string{"exist", "AxInstSV"}, result: true}, wantRetBool: true},
		{ name: "Service Parse Valid Args - Does not Exist", args: args{args: []string{"exist", "AxIstSV"}, result: true}, wantRetBool: false},
		{ name: "Service Parse Valid Args - Invalid", args: args{args: []string{"exist"}, result: true}, wantRetBool: false},

		{ name: "Service Parse Meta - Nominal Args", args: args{args: []string{"meta", "AxInstSV", "DisplayName", "ActiveX Installer (AxInstSV)"}, result: true}, wantRetBool: true},
		{ name: "Service Parse Meta - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "Service Parse Meta - Invalid Args", args: args{args: []string{"meta", "AxIntSV", "DisplayName", "ActiveX Installer (AxInstSV)"}, result: true}, wantRetBool: false},
		{ name: "Service Parse Meta - Invalid Args", args: args{args: []string{"meta", "AxInstSV", "DiplayName", "ActiveX Installer (AxInstSV)"}, result: true}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ServiceParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("ServiceParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}