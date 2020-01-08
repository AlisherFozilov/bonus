package bonus

import "testing"

func Test_bonus(t *testing.T) {

	tests := []struct {
		name  string
		sales []int
		want  int
	}{
		{"Has bonus", []int{12_000, 8_000, 15_000, 8_000}, 350},
		{"Hasn't bonus", []int{9_000, 500, 5_000}, 0},
	}

	for _, test := range tests {
		got := bonus(test.sales)
		if got != test.want {
			t.Error("For bonus test:", test.name, "got:", got, "want:", test.want)
		}
	}
}
