package windows

import "testing"

func TestGroupExist(t *testing.T) {
	type args struct {
		grp string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Group Exist Invalid Args", args: args{"-"}, wantRetBool: false},
		{ name: "Group Exist No Args", args: args{}, wantRetBool: false},
		{ name: "Group Exist Valid Args - Exist", args: args{"Users"}, wantRetBool: true},
		{ name: "Group Exist Valid Args - Does not Exist", args: args{"Usr"}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := GroupExist(tt.args.grp); gotRetBool != tt.wantRetBool {
				t.Errorf("GroupExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestGroupMeta(t *testing.T) {
	type args struct {
		grp   string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Group Meta Invalid Args", args: args{grp: "=", key: "Commet", value: false}, wantRetBool: false },
		{ name: "Group Meta No Args", args: args{}, wantRetBool: false },
		{ name: "Group Meta Valid Args - Comment True", args: args{grp: "Users", key: "Comment", value: "Users are prevented from making accidental or intentional system-wide changes and can run most applications"}, wantRetBool: true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := GroupMeta(tt.args.grp, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("GroupMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestGroupParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{  name: "Group Parse Invalid Args", args: args{args: []string{"-1"}, result: true}, wantRetBool: false },
		{  name: "Group Parse No Args", args: args{}, wantRetBool: false },
		{  name: "Group Parse Valid Args - Exist", args: args{args: []string{"exist", "Users"}, result: true}, wantRetBool: true },
		{  name: "Group Parse Valid Args - Does not Exist", args: args{args: []string{"exist", "Usrs"}, result: true}, wantRetBool: false },
		{  name: "Group Parse Valid Args - Invalid", args: args{args: []string{"exist"}, result: true}, wantRetBool: false },

		{  name: "Group Parse Meta - Nominal Args", args: args{args: []string{"meta", "Users", "Comment", "Users are prevented from making accidental or intentional system-wide changes and can run most applications"}, result: true}, wantRetBool: true },
		{  name: "Group Parse Meta - No Args", args: args{args: []string{"meta"}, result: false}, wantRetBool: false },
		{  name: "Group Parse Meta - Invalid Args", args: args{args: []string{"meta", "-1"}, result: false}, wantRetBool: false },
		{  name: "Group Parse Meta - Invalid Args", args: args{args: []string{"meta", "Users", "Codas", "false"}, result: true}, wantRetBool: false },

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := GroupParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("GroupParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}