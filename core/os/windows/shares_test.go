package windows

import "testing"

func TestShareExist(t *testing.T) {
	type args struct {
		share string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Share Exist Valid Args", args: args{"C$"}, wantRetBool: true},
		{ name: "Share Exist Invalid Args", args: args{"123$"}, wantRetBool: false},
		{ name: "Share Exist No Args", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ShareExist(tt.args.share); gotRetBool != tt.wantRetBool {
				t.Errorf("ShareExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestShareMeta(t *testing.T) {
	type args struct {
		share string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Share Invalid Args", args: args{share: ""}, wantRetBool: false},
		{ name: "Share No Args", args: args{}, wantRetBool: false},
		{ name: "Share Meta Valid Args - Status True", args: args{share: "C$", key: "Status", value: "OK"}, wantRetBool: true},
		{ name: "Share Meta Valid Args - Status False", args: args{share: "C$", key: "Status", value: "-"}, wantRetBool: false},
		{ name: "Share Meta Valid Args - Caption True", args: args{share: "C$", key: "Caption", value: "Default share"}, wantRetBool: true},
		{ name: "Share Meta Valid Args - Caption False", args: args{share: "C$", key: "Caption", value: "-"}, wantRetBool: false},
		{ name: "Share Meta Valid Args - Description True", args: args{share: "C$", key: "Description", value: "Default share"}, wantRetBool: true},
		{ name: "Share Meta Valid Args - Description False", args: args{share: "C$", key: "Description", value: "-"}, wantRetBool: false},
		{ name: "Share Meta Valid Args - Path True", args: args{share: "C$", key: "Path", value: "C:\\"}, wantRetBool: true},
		{ name: "Share Meta Valid Args - Path False", args: args{share: "C$", key: "Path", value: "D:/"}, wantRetBool: false},
		{ name: "Share Meta Valid Args - AllowMaximum True", args: args{share: "C$", key: "AllowMaximum", value: "true"}, wantRetBool: true},
		{ name: "Share Meta Valid Args - AllowMaximum False", args: args{share: "C$", key: "AllowMaximum", value: "-"}, wantRetBool: false},
		{ name: "Share Meta Valid Args - Type True", args: args{share: "C$", key: "Type", value: "2147483648"}, wantRetBool: true},
		{ name: "Share Meta Valid Args - Type False", args: args{share: "C$", key: "Type", value: "-"}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ShareMeta(tt.args.share, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("ShareMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestShareParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Share Parse Invalid Args", args: args{args: []string{"-1"}, result: -1}, wantRetBool: false},
		{ name: "Share Parse No Args", args: args{args: []string{}}, wantRetBool: false},
		{ name: "Share Parse Valid Args - Exist", args: args{args: []string{"exist", "C$"}, result: true}, wantRetBool: true},
		{ name: "Share Parse Valid Args - Does not Exist", args: args{args: []string{"exist", "Z$"}, result: true}, wantRetBool: false},
		{ name: "Share Parse Valid Args - Invalid", args: args{args: []string{"exist", "$"}, result: true}, wantRetBool: false},

		{ name: "Share Parse Meta - Nominal Args", args: args{args: []string{"meta", "C$", "Status", "OK"}, result: true}, wantRetBool: true},
		{ name: "Share Parse Meta - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "Share Parse Meta - Invalid Args", args: args{args: []string{"meta", "C$", "Status", "Onn"}, result: true}, wantRetBool: false},
		{ name: "Share Parse Meta - Invalid Args", args: args{args: []string{"meta", "C$", "Sttus", "OK"}, result: true}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := ShareParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("ShareParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}