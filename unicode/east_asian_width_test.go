package unicode

import "testing"

func TestGetEastAsianWidth(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				c: 'l',
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				c: '❤',
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				c: '龘',
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				c: '宽',
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEastAsianWidth(tt.args.c); got != tt.want {
				t.Errorf("GetEastAsianWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}
