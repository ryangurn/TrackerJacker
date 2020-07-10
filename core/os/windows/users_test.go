package windows

import "testing"

const validUsername = "Ryan Gurnick"
const invalidUsername = "Ryan"
const emptyUsername = ""

func TestUserExist(t *testing.T) {
	type args struct {
		usr string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{name: "User Exists", args: args{usr: validUsername}, wantRetBool: true },
		{name: "User Does Not Exist", args: args{usr: invalidUsername}, wantRetBool: false },
		{name: "User Empty", args: args{usr: emptyUsername}, wantRetBool: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserExist(tt.args.usr); gotRetBool != tt.wantRetBool {
				t.Errorf("UserExist() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestUserLoggedIn(t *testing.T) {
	type args struct {
		usr string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "User Logged In", args: args{usr: validUsername}, wantRetBool: true},
		{ name: "User Not Logged In", args: args{usr: invalidUsername}, wantRetBool: false},
		{ name: "User Logged In Empty", args: args{usr: emptyUsername}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserLoggedIn(tt.args.usr); gotRetBool != tt.wantRetBool {
				t.Errorf("UserLoggedIn() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestUserMeta(t *testing.T) {
	type args struct {
		usr   string
		key   string
		value interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		// general
		{ name: "User Meta No Args", args: args{}, wantRetBool: false},

		// badPasswordCount
		{ name: "User Meta BadPasswordCount Valid User - Invalid Count", args: args{usr: validUsername, key: "BadPasswordCount", value: uint32(1)}, wantRetBool: false},
		{ name: "User Meta BadPasswordCount Valid User - Valid Count", args: args{usr: validUsername, key: "BadPasswordCount", value: uint32(0)}, wantRetBool: true},
		{ name: "User Meta BadPasswordCount Invalid User", args: args{usr: invalidUsername, key: "BadPasswordCount", value: uint32(99)}, wantRetBool: false},
		{ name: "User Meta BadPasswordCount Invalid User - No Count", args: args{usr: invalidUsername, key: "BadPasswordCount"}, wantRetBool: false},

		// fullName
		{ name: "User Meta FullName Valid User - Valid FullName", args: args{usr: validUsername, key: "FullName", value: "Test"}, wantRetBool: true},
		{ name: "User Meta FullName Valid User - Invalid FullName", args: args{usr: validUsername, key: "FullName", value: ""}, wantRetBool: false},
		{ name: "User Meta FullName Invalid User", args: args{usr: invalidUsername, key: "FullName", value: ""}, wantRetBool: false},

		// isAdmin
		{ name: "User Meta IsAdmin Valid User - Valid Bool", args: args{usr: validUsername, key: "IsAdmin", value: "true"}, wantRetBool: true},
		{ name: "User Meta IsAdmin Valid User - Invalid Bool", args: args{usr: validUsername, key: "IsAdmin", value: "false"}, wantRetBool: false},
		{ name: "User Meta IsAdmin Invalid User - Valid Bool", args: args{usr: validUsername, key: "IsAdmin", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsAdmin Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsAdmin", value: "t"}, wantRetBool: false},
		{ name: "User Meta IsAdmin Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsAdmin", value: ""}, wantRetBool: false},

		// isEnabled
		{ name: "User Meta IsEnabled Valid User - Valid Bool", args: args{usr: validUsername, key: "IsEnabled", value: "true"}, wantRetBool: true},
		{ name: "User Meta IsEnabled Valid User - Invalid Bool", args: args{usr: validUsername, key: "IsEnabled", value: "false"}, wantRetBool: false},
		{ name: "User Meta IsEnabled Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "IsEnabled", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsEnabled Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsEnabled", value: "f"}, wantRetBool: false},
		{ name: "User Meta IsEnabled Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsEnabled", value: " "}, wantRetBool: false},

		// isLocked
		{ name: "User Meta IsLocked Valid User - Valid Bool", args: args{usr: validUsername, key: "IsLocked", value: "false"}, wantRetBool: true},
		{ name: "User Meta IsLocked Valid User - Invalid Bool", args: args{usr: validUsername, key: "IsLocked", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsLocked Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "IsLocked", value: "false"}, wantRetBool: false},
		{ name: "User Meta IsLocked Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsLocked", value: "t"}, wantRetBool: false},
		{ name: "User Meta IsLocked Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsLocked", value: " -"}, wantRetBool: false},

		// TODO: Add LastLogin, NoChangePassword, NumberOfLogons, PasswordAge, PasswordNeverExpires
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserMeta(tt.args.usr, tt.args.key, tt.args.value); gotRetBool != tt.wantRetBool {
				t.Errorf("UserMeta() = %v, want %v", gotRetBool, tt.wantRetBool)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("UserParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}