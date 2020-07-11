package cross

import "testing"

const validUsername = "Ryan Gurnick"
const invalidUsername = "Ryan"

func TestUserExist(t *testing.T) {
	type args struct {
		usr string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "User Exist Valid User", args: args{usr: validUsername}, wantRetBool: true},
		{ name: "User Exist Invalid User", args: args{usr: invalidUsername}, wantRetBool: false},
		{ name: "User Exist Invalid User", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserExist(tt.args.usr); gotRetBool != tt.wantRetBool {
				t.Errorf("UserExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestUserParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "No Arguments", args: args{}, wantRetBool: false },
		{ name: "User Parse Valid User - Valid Args", args: args{args: []string{"exist", validUsername}, result: true}, wantRetBool: true},
		{ name: "User Parse Valid User - Invalid Args", args: args{args: []string{"exist", invalidUsername}, result: true}, wantRetBool: false},
		{ name: "User Parse Valid User - Invalid Args", args: args{args: []string{"exist", invalidUsername}, result: false}, wantRetBool: true},
		{ name: "User Parse Valid User - Invalid Args", args: args{args: []string{"exist"}, result: false}, wantRetBool: false},
		{ name: "User Parse Valid User - No Args", args: args{args: []string{}, result: false}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("UserParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}