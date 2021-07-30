package os

import (
	"context"
	"log"
	"testing"
)

func TestRunBashCommand(t *testing.T) {
	err := RunBashCommand("ls")
	if err != nil {
		t.Error("RunBashCommand error must be nil")
	}

	err = RunBashCommand("asasa")
	if err == nil {
		t.Error("RunBashCommand error must not be nil")
	}
}

func TestCmdForLog(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ls -l",
			args: args{
				command: "ls",
				args: []string{
					"-l",
				},
			},
			want: "ls -l",
		},
		{
			name: "curl -v https://baidu.com",
			args: args{
				command: "curl",
				args: []string{
					"-v",
					"https://baidu.com",
				},
			},
			want: "curl -v https://baidu.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CmdForLog(tt.args.command, tt.args.args...); got != tt.want {
				t.Errorf("CmdForLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunCommand(t *testing.T) {
	res, err := RunCommand(context.Background(), "", "ls", "-l")
	if err != nil {
		t.Error("RunCommand error must be nil")
	}
	log.Print(res)

	res, err = RunCommand(context.Background(), "", "asasa")
	if err == nil {
		t.Error("RunCommand error must not be nil")
	}
	log.Print(res)
}
