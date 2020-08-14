package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"testing"
)

func TestUpdateHistoryExists(t *testing.T) {
	updates, err := wapi.UpdatesPending()
	if err != nil {
		t.Errorf("Unable to load updates: %s", err)
	}

	type args struct {
		update string
		status string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Update History Exists", args: args{update: updates.UpdateHistory[0].UpdateName, status: updates.UpdateHistory[0].Status}, wantRetBool: true},
		{ name: "Update History Does not Exists", args: args{update: updates.UpdateHistory[0].UpdateName, status: "-"}, wantRetBool: false},
		{ name: "Update History Empty", args: args{update: "", status: ""}, wantRetBool: false},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UpdateHistoryExists(tt.args.update, tt.args.status); gotRetBool != tt.wantRetBool {
				t.Errorf("UpdateHistoryExists() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestUpdateParse(t *testing.T) {
	updates, err := wapi.UpdatesPending()
	if err != nil {
		t.Errorf("Unable to load updates: %s", err)
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
		{ name: "Update Parse History Exist - Nominal Args", args: args{args: []string{"history_exist", updates.UpdateHistory[0].UpdateName, updates.UpdateHistory[0].Status}, result: true}, wantRetBool: true},
		{ name: "Update Parse History Exist - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "Update Parse History Exist - Incorrect Args", args: args{args: []string{"history_exist"}, result: true}, wantRetBool: false},
		{ name: "Update Parse History Exist - Incorrect Args", args: args{args: []string{"history_exist", "", ""}, result: true}, wantRetBool: false},

		{ name: "Update Parse Updated - Nominal Args", args: args{args: []string{"updated", strconv.FormatBool(!updates.UpdatesReq)}, result: true}, wantRetBool: true},
		{ name: "Update Parse Updated - No Args", args: args{args: []string{}, result: true}, wantRetBool: false},
		{ name: "Update Parse Updated - Incorrect Args", args: args{args: []string{"updated", "-"}, result: true}, wantRetBool: false},
		{ name: "Update Parse Updated - Incorrect Args", args: args{args: []string{"updated", "/"}, result: true}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := UpdateParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("UpdateParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestUpdated(t *testing.T) {
	updates, err := wapi.UpdatesPending()
	if err != nil {
		t.Errorf("Unable to get updates: %s", err)
	}

	type args struct {
		completed bool
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "Updates Updated Valid Args", args: args{completed: !updates.UpdatesReq}, wantRetBool: true},
		{ name: "Updates Updated Invalid Args", args: args{completed: updates.UpdatesReq}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := Updated(tt.args.completed); gotRetBool != tt.wantRetBool {
				t.Errorf("Updated() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}
