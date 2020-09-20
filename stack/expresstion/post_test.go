package expresstion

import "testing"

func TestPostfixExpression1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"16 9 4 3 +*-"}, -47},
		{"case1", args{"5 7 3 * + 5 2 1 +*-"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostfixExpression(tt.args.str); got != tt.want {
				t.Errorf("PostfixExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}