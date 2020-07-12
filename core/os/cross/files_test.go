package cross

import (
	"os"
	"testing"
)

const validPath = "C:"+string(os.PathSeparator)+"Users"+string(os.PathSeparator)+"Ryan Gurnick"+string(os.PathSeparator)+"Desktop"+string(os.PathSeparator)+"hi.txt"
const invalidPath = "C:"+string(os.PathSeparator)+"Users"+string(os.PathSeparator)+"Ryan Gurnick"+string(os.PathSeparator)+"Desktop"+string(os.PathSeparator)+"nope.txt"

func TestFileExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{ name: "File Exist Valid Path", args: args{path: validPath}, wantRetBool: true},
		{ name: "File Exist Invalid Path", args: args{path: invalidPath}, wantRetBool: false},
		{ name: "File Exist No Args", args: args{}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := FileExists(tt.args.path); gotRetBool != tt.wantRetBool {
				t.Errorf("FileExists() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}

func TestFileHash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{ name: "Fle Hash Valid Path", args: args{path: validPath}, want: "f74f41add5ea120114de0145b3fb6d87", wantErr: false},
		{ name: "Fle Hash Invalid Path", args: args{path: invalidPath}, want: "", wantErr: true},
		{ name: "Fle Hash Empty Args", args: args{}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileHash(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileParse(t *testing.T) {
	type args struct {
		args   []string
		result interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantRetBool bool
	}{
		{name: "File Parse Valid Args Exist", args: args{args: []string{"exist", validPath}, result: true}, wantRetBool: true},
		{name: "File Parse Valid Args Exist - Invalid Path", args: args{args: []string{"exist", invalidPath}, result: false}, wantRetBool: true},
		{name: "File Parse Empty Args Exist", args: args{args: []string{}, result: false}, wantRetBool: false},

		{name: "File Parse Valid Args Hash", args: args{args: []string{"hash", validPath}, result: "f74f41add5ea120114de0145b3fb6d87"}, wantRetBool: true},
		{name: "File Parse Valid Args Hash - Invalid Path", args: args{args: []string{"hash", invalidPath}, result: "anything here"}, wantRetBool: false},
		{name: "File Parse Empty Args Hash", args: args{args: []string{}, result: "anything here"}, wantRetBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetBool := FileParse(tt.args.args, tt.args.result); gotRetBool != tt.wantRetBool {
				t.Errorf("FileParse() = %v, want %v", gotRetBool, tt.wantRetBool)
			}
		})
	}
}