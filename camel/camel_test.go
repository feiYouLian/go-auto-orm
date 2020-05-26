package camel

import (
	"testing"
)

func TestUnMarshal(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{name: "ABC"}, "a_b_c"},
		{"test2", args{name: "AbCd"}, "ab_cd"},
		{"test2", args{name: "AbCD_efg"}, "ab_c_d_efg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnMarshal(tt.args.name); got != tt.want {
				t.Errorf("UnMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	type args struct {
		name       string
		firstUpper bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{name: "a_b_c"}, "aBC"},
		{"test3", args{name: "_a_b_c"}, "ABC"},
		{"test5", args{name: "a_bc_de"}, "aBcDe"},
		{"test7", args{name: "ab_cd_efGh"}, "abCdEfGh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Marshal(tt.args.name); got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
