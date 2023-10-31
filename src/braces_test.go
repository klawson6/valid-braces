package main

import "testing"

func TestBraceYourself(t *testing.T) {
	type args struct {
		braces string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "When brace string is a series of valid braces, then returns true",
			args: args{braces: "(){}[]"},
			want: true,
		},
		{
			name: "When brace string is a nest of valid braces, then returns true",
			args: args{braces: "([{}])"},
			want: true,
		},
		{
			name: "When brace string has both series and nest of valid braces, then returns true",
			args: args{braces: "([{}])[]{}"},
			want: true,
		},
		{
			name: "When brace string has both series and nest of valid braces, then returns true",
			args: args{braces: "([{}])[]{}"},
			want: true,
		},
		{
			name: "When brace string has only unmatched braces, then returns false",
			args: args{braces: "(}"},
			want: false,
		},
		{
			name: "When brace string has matched braces which are not correctly nested, then returns false",
			args: args{braces: "[(])"},
			want: false,
		},
		{
			name: "When brace string has a valid nest as well as unmatched braces, then returns false",
			args: args{braces: "[({})](]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BraceYourself(tt.args.braces); got != tt.want {
				t.Errorf("BraceYourself() = %v, want %v", got, tt.want)
			}
		})
	}
}
