package cliintrepeter_test

import (
	"github.com/bilfash/azwraith/cliintrepeter"
	"testing"
)

func Test_cliInterpreter_Execute(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		ci      cliintrepeter.CliInterpreter
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "TestPositiveEcho",
			ci:      cliintrepeter.NewCliInterpreter(),
			args:    args{
				command: "echo",
				args:    []string{
					"This",
					"one",
					"is",
					"valid",
					"string",
				},
			},
			want:    "This one is valid string\n",
			wantErr: false,
		},
		{
			name:    "TestNegativeEcho",
			ci:      cliintrepeter.NewCliInterpreter(),
			args:    args{
				command: "invalid-command-xxx-you-should-never-have-a-command-or-alias-like-this",
			},
			want:    "exec: \"invalid-command-xxx-you-should-never-have-a-command-or-alias-like-this\": executable file not found in $PATH",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ci.Execute(tt.args.command, tt.args.args...)
			if got != tt.want {
				t.Errorf("cliInterpreter.Execute() = \"%v\", want \"%v\"", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("cliInterpreter.Execute() error = \"%v\", wantErr \"%v\"", err, tt.wantErr)
				return
			}
		})
	}
}
