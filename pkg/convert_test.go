package pkg

import "testing"

func Test_secondsToHumna(t *testing.T) {
	type args struct {
		seconds string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "60 sec",
		args: args{seconds: "60"},
		want: "* 01:00",
	}, {
		name: "50 sec",
		args: args{seconds: "50"},
		want: "* 00:50",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondsToHuman(tt.args.seconds); got != tt.want {
				t.Errorf("secondsToHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}
