package main

import "testing"

func TestCreateHandlerName(t *testing.T) {
	type args struct {
		sName []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CreateHandlerName",
			args: args{
				sName: []string{"test", "lambda"},
			},
			want: "testLambda",
		},
		{
			name: "CreateHandlerName",
			args: args{
				sName: []string{"delete", "test", "lambda"},
			},
			want: "deleteTestLambda",
		},
		{
			name: "CreateHandlerName",
			args: args{
				sName: []string{"lambda"},
			},
			want: "lambda",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateHandlerName(tt.args.sName); got != tt.want {
				t.Errorf("CreateHandlerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
