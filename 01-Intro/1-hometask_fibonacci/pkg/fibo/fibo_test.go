package fibo_test

import (
	"go-core-2/01-Intro/1-hometask_fibonacci/pkg/fibo"
	"testing"
)

func TestNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "negative",
			args: args{
				n: -1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "zero",
			args: args{
				n: 0,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "one",
			args: args{
				n: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "eight",
			args: args{
				n: 8,
			},
			want:    21,
			wantErr: false,
		},
		{
			name: "twenty",
			args: args{
				n: 20,
			},
			want:    6765,
			wantErr: false,
		},
		{
			name: "twenty one",
			args: args{
				n: 21,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fibo.Num(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Num() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}
