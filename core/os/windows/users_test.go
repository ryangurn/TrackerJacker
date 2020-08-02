package windows

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	wapi "github.com/iamacarpet/go-win64api"
)

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
	// truthy values
	users,  err := wapi.ListLocalUsers()
	if err != nil {
		fmt.Println("Error with listing local users.")
	}

	var badPasswordCount, fullName, isAdmin, isEnabled, isLocked, noChangePassword, passwordNeverExpires = "", "", "", "", "", "", ""
	var lastLogon = time.Time{}
	//var passwordAge = time.Duration(0)
	var numberOfLogons = uint32(0)

	for _, u := range users {
		if u.Username == validUsername {
			badPasswordCount = strconv.FormatUint(uint64(u.BadPasswordCount), 10)
			fullName = u.FullName
			isAdmin = strconv.FormatBool(u.IsAdmin)
			isEnabled = strconv.FormatBool(u.IsEnabled)
			isLocked = strconv.FormatBool(u.IsLocked)
			passwordNeverExpires = strconv.FormatBool(u.PasswordNeverExpires)
			lastLogon = u.LastLogon
			noChangePassword = strconv.FormatBool(u.NoChangePassword)
			numberOfLogons = u.NumberOfLogons
			//passwordAge = u.PasswordAge
		}
	}

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
		{ name: "User Meta BadPasswordCount Valid User - Invalid Count", args: args{usr: validUsername, key: "BadPasswordCount", value: "1"}, wantRetBool: false},
		{ name: "User Meta BadPasswordCount Valid User - Valid Count", args: args{usr: validUsername, key: "BadPasswordCount", value: badPasswordCount}, wantRetBool: true},
		{ name: "User Meta BadPasswordCount Invalid User", args: args{usr: invalidUsername, key: "BadPasswordCount", value: "99"}, wantRetBool: false},
		{ name: "User Meta BadPasswordCount Invalid User - No Count", args: args{usr: invalidUsername, key: "BadPasswordCount"}, wantRetBool: false},

		// fullName
		{ name: "User Meta FullName Valid User - Valid FullName", args: args{usr: validUsername, key: "FullName", value: fullName}, wantRetBool: true},
		{ name: "User Meta FullName Valid User - Invalid FullName", args: args{usr: validUsername, key: "FullName", value: ""}, wantRetBool: false},
		{ name: "User Meta FullName Invalid User", args: args{usr: invalidUsername, key: "FullName", value: ""}, wantRetBool: false},

		// isAdmin
		{ name: "User Meta IsAdmin Valid User - Valid Bool", args: args{usr: validUsername, key: "IsAdmin", value: isAdmin}, wantRetBool: true},
		{ name: "User Meta IsAdmin Valid User - Invalid Bool", args: args{usr: validUsername, key: "IsAdmin", value: "false"}, wantRetBool: false},
		{ name: "User Meta IsAdmin Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "IsAdmin", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsAdmin Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsAdmin", value: "t"}, wantRetBool: false},
		{ name: "User Meta IsAdmin Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsAdmin", value: ""}, wantRetBool: false},

		// isEnabled
		{ name: "User Meta IsEnabled Valid User - Valid Bool", args: args{usr: validUsername, key: "IsEnabled", value: isEnabled}, wantRetBool: true},
		{ name: "User Meta IsEnabled Valid User - Invalid Bool", args: args{usr: validUsername, key: "IsEnabled", value: "false"}, wantRetBool: false},
		{ name: "User Meta IsEnabled Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "IsEnabled", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsEnabled Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsEnabled", value: "f"}, wantRetBool: false},
		{ name: "User Meta IsEnabled Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsEnabled", value: " "}, wantRetBool: false},

		// isLocked
		{ name: "User Meta IsLocked Valid User - Valid Bool", args: args{usr: validUsername, key: "IsLocked", value: isLocked}, wantRetBool: true},
		{ name: "User Meta IsLocked Valid User - Invalid Bool", args: args{usr: validUsername, key: "IsLocked", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsLocked Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "IsLocked", value: "true"}, wantRetBool: false},
		{ name: "User Meta IsLocked Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsLocked", value: "t"}, wantRetBool: false},
		{ name: "User Meta IsLocked Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "IsLocked", value: " -"}, wantRetBool: false},

		// lastLogon (there is not a great way to tell if it is the true case right now)
		{ name: "User Meta LastLogin Valid User - Valid Time", args: args{usr: validUsername, key: "LastLogon", value: time.Now().Add(time.Hour * time.Duration(1))}, wantRetBool: false},
		{ name: "User Meta LastLogin Valid User - Valid Time", args: args{usr: validUsername, key: "LastLogon", value: lastLogon}, wantRetBool: true},
		{ name: "User Meta LastLogin Invalid User - Valid Time", args: args{usr: invalidUsername, key: "LastLogon", value: time.Now().Add(time.Hour * time.Duration(1))}, wantRetBool: false},
		{ name: "User Meta LastLogin Invalid User - Empty Time", args: args{usr: invalidUsername, key: "LastLogon", value: time.Time{}}, wantRetBool: false},

		// noChangePassword
		{ name: "User Meta NoChangePassword Valid User - Valid Bool", args: args{usr: validUsername, key: "NoChangePassword", value: noChangePassword}, wantRetBool: true},
		{ name: "User Meta NoChangePassword Valid User - Valid Bool", args: args{usr: validUsername, key: "NoChangePassword", value: "true"}, wantRetBool: false},
		{ name: "User Meta NoChangePassword Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "NoChangePassword", value: "true"}, wantRetBool: false},
		{ name: "User Meta NoChangePassword Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "NoChangePassword", value: "t"}, wantRetBool: false},
		{ name: "User Meta NoChangePassword Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "NoChangePassword", value: ".."}, wantRetBool: false},

		// numberOfLogons
		{ name: "User Meta NumberOfLogons Valid User - Valid Int", args: args{usr: validUsername, key: "NumberOfLogons", value: uint32(1)}, wantRetBool: false},
		{ name: "User Meta NumberOfLogons Valid User - Valid Int", args: args{usr: validUsername, key: "NumberOfLogons", value: numberOfLogons}, wantRetBool: true},
		{ name: "User Meta NumberOfLogons Invalid User - Valid Int", args: args{usr: invalidUsername, key: "NumberOfLogons", value: uint32(5)}, wantRetBool: false},
		{ name: "User Meta NumberOfLogons Invalid User - No Count", args: args{usr: invalidUsername, key: "NumberOfLogons"}, wantRetBool: false},

		// passwordAge (since this is a duration it will be hard to test the true case)
		{ name: "User Meta PasswordAge Valid User - Valid Duration", args: args{usr: validUsername, key: "PasswordAge", value: time.Duration(time.Hour * 24)}, wantRetBool: false},
		//{ name: "User Meta PasswordAge Valid User - Valid Duration", args: args{usr: validUsername, key: "PasswordAge", value: passwordAge}, wantRetBool: true},
		{ name: "User Meta PasswordAge Invalid User - Valid Duration", args: args{usr: invalidUsername, key: "PasswordAge", value: time.Duration(time.Hour * 24)}, wantRetBool: false},
		{ name: "User Meta PasswordAge Invalid User - Empty Duration", args: args{usr: invalidUsername, key: "PasswordAge"}, wantRetBool: false},
		
		{ name: "User Meta PasswordNeverExpires Valid User - Valid Bool", args: args{usr: validUsername, key: "PasswordNeverExpires", value: "true"}, wantRetBool: false},
		{ name: "User Meta PasswordNeverExpires Valid User - Valid Bool", args: args{usr: validUsername, key: "PasswordNeverExpires", value: passwordNeverExpires}, wantRetBool: true},
		{ name: "User Meta PasswordNeverExpires Invalid User - Valid Bool", args: args{usr: invalidUsername, key: "PasswordNeverExpires", value: "false"}, wantRetBool: false},
		{ name: "User Meta PasswordNeverExpires Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "PasswordNeverExpires", value: "f"}, wantRetBool: false},
		{ name: "User Meta PasswordNeverExpires Invalid User - Invalid Bool", args: args{usr: invalidUsername, key: "PasswordNeverExpires", value: "/.."}, wantRetBool: false},
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
		{ name: "User Parse Exist - Nominal Args", args: args{args: []string{"exist", validUsername}, result: true}, wantRetBool: true},
		{ name: "User Parse Exist - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "User Parse Exist - Incorrect Args", args: args{args: []string{"exists"}, result: true}, wantRetBool: false},
		{ name: "User Parse Exist - Incorrect Args", args: args{args: []string{"exists", "", ""}, result: true}, wantRetBool: false},

		{ name: "User Parse Meta - Nominal Args", args: args{args: []string{"meta", validUsername, "BadPasswordCount", "2"}, result: true}, wantRetBool: false},
		{ name: "User Parse Meta - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "User Parse Meta - Incorrect Args", args: args{args: []string{"meta", validUsername}, result: true}, wantRetBool: false},
		{ name: "User Parse Meta - Incorrect Args", args: args{args: []string{"meta", validUsername, "IsLocked"}, result: true}, wantRetBool: false},

		{ name: "User Parse Logged - Nominal Args", args: args{args: []string{"logged", "exist", validUsername}, result: true}, wantRetBool: true},
		{ name: "User Parse Logged - Invalid User", args: args{args: []string{"logged", "exist", invalidUsername}, result: false}, wantRetBool: true},
		{ name: "User Parse Logged - Invalid Args", args: args{args: []string{"logged", validUsername}, result: false}, wantRetBool: false},
		{ name: "User Parse Logged - Invalid Args", args: args{args: []string{"logged", validUsername}, result: false}, wantRetBool: false},
		{ name: "User Parse Logged - No Args", args: args{args: []string{}, result: false}, wantRetBool: false},

		{ name: "Unrecognized Command", args: args{args: []string{}, result: false}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UserParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("UserParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}