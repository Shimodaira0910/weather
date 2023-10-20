package env

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {
}

func TestEnv_LoadEnv(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name string
		e    *Env
		args args
		want string
	}{
		{
			name: "Load existing environment variable",
			e:    &Env{},
			args: args{envName: "TEST"},
			want: "test_value",
		},
		
		{
			name: "Load non-existent environment variable",
			e:    &Env{},
			args: args{envName: "NON_EXISTENT_VARIABLE"},
			want: "読み込み失敗",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Env{}
			if got := e.LoadEnv(tt.args.envName); got != tt.want {
				t.Errorf("Env.LoadEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
