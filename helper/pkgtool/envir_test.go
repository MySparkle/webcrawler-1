package pkgtool

import (
	"reflect"
	"testing"
)

func TestGetGoroot(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGoroot(); got != tt.want {
				t.Errorf("GetGoroot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllGopath(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllGopath(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllGopath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSrcDirs(t *testing.T) {
	type args struct {
		fresh bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSrcDirs(tt.args.fresh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSrcDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}
