package windows

import "testing"

func TestProcessExist(t *testing.T) {
	type args struct {
		processFullPath string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Process Exist Valid Process", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe"}, wantRetBool: true},
		{ name: "Process Exist Invalid Args", args: args{processFullPath: "-"}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ProcessExist(tt.args.processFullPath); gotRetBool != tt.wantRetBool {
				t.Errorf("ProcessExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestProcessMeta(t *testing.T) {
	type args struct {
		processFullPath string
		key             string
		value           interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Process Meta Invalid Args", args: args{processFullPath: "-"}, wantRetBool: false},
		{ name: "Process Meta No Args", args: args{}, wantRetBool: false},
		{ name: "Process Meta Valid Args - Pid True", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Pid", value: 7572}, wantRetBool: true},
		{ name: "Process Meta Valid Args - Pid False", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Pid", value: -1}, wantRetBool: false},
		{ name: "Process Meta Valid Args - Ppid True", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Ppid", value: 820}, wantRetBool: true},
		{ name: "Process Meta Valid Args - Ppid False", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Ppid", value: -1}, wantRetBool: false},
		{ name: "Process Meta Valid Args - Username True", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Username", value: "DESKTOP-6QGRMKO\\ryan gurnick"}, wantRetBool: true},
		{ name: "Process Meta Valid Args - Username False", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Username", value: "-"}, wantRetBool: false},
		{ name: "Process Meta Valid Args - Executable True", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Executable", value: "dllhost.exe"}, wantRetBool: true},
		{ name: "Process Meta Valid Args - Executable False", args: args{processFullPath: "C:\\Windows\\System32\\dllhost.exe", key: "Executable", value: "-"}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ProcessMeta(tt.args.processFullPath, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("ProcessMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestProcessParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Processes Parse Invalid Args", args: args{args: []string{"-a"}, result: "a"}, wantRetBool: false},
		{ name: "Processes Parse No Args", args: args{args: []string{}}, wantRetBool: false},
		{ name: "Processes Parse Valid Args - Exist", args: args{args: []string{"exist", "C:\\Windows\\System32\\dllhost.exe"}, result: true}, wantRetBool: true},
		{ name: "Processes Parse Valid Args - Does not exist", args: args{args: []string{"exist", "C:\\Windows\\System32\\dllhst.exe"}, result: true}, wantRetBool: false},
		{ name: "Processes Parse Valid Args - Invalid", args: args{args: []string{"exst", "C:\\Windows\\System32\\dllhost.exe"}, result: true}, wantRetBool: false},

		{ name: "Processes Parse Meta - Nominal Args", args: args{args: []string{"meta", "C:\\Windows\\System32\\dllhost.exe", "Username", "DESKTOP-6QGRMKO\\ryan gurnick"}, result: true}, wantRetBool: true},
		{ name: "Processes Parse Meta - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "Processes Parse Meta - Invalid Args", args: args{args: []string{"meta", "C:\\Windows\\Syste32\\dllhost.exe", "Username", "DESKTOP-6QGRMKO\\ryan gurnick"}, result: true}, wantRetBool: false},
		{ name: "Processes Parse Meta - Invalid Args", args: args{args: []string{"meta", "C:\\Windows\\System32\\dllhost.exe", "Usernae", "DESKTOP-6QGRMKO\\ryan gurnick"}, result: true}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ProcessParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("ProcessParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}