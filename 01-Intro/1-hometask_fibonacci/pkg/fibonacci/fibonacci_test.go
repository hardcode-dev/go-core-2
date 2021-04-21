package fibonacci

import "testing"

func TestFindFibonacci(t *testing.T) {
	type args struct {
		number int
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
				number: -1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "zero",
			args: args{
				number: 0,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "one",
			args: args{
				number: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "eight",
			args: args{
				number: 8,
			},
			want:    21,
			wantErr: false,
		},
		{
			name: "twenty",
			args: args{
				number: 20,
			},
			want:    6765,
			wantErr: false,
		},
		{
			name: "twenty one",
			args: args{
				number: 21,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFibonacci(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindFibonacci() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
