package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestUsage(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Usage()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "xxx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_isDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "是目录",
			args: args{name: "/Users/yudong/www/goproject/src/webcrawler/helper/log/field/cmd"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDir(tt.args.name); got != tt.want {
				t.Errorf("isDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFieldTypePrefixes(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "field",
			args:    args{filePath: "/Users/yudong/www/goproject/src/webcrawler/helper/log/field/field.go"},
			want:    []string{"bool", "int64", "float64", "string", "object"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findFieldTypePrefixes(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("findFieldTypePrefixes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFieldTypePrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_reset(t *testing.T) {
	type fields struct {
		buf bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				buf: tt.fields.buf,
			}
			g.reset()
		})
	}
}

func TestGenerator_generate(t *testing.T) {
	type fields struct {
		buf bytes.Buffer
	}
	type args struct {
		pkgName  string
		prefixes []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				buf: tt.fields.buf,
			}
			got, err := g.generate(tt.args.pkgName, tt.args.prefixes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_genHeader(t *testing.T) {
	type fields struct {
		buf bytes.Buffer
	}
	type args struct {
		pkgName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				buf: tt.fields.buf,
			}
			g.genHeader(tt.args.pkgName)
		})
	}
}

func TestGenerator_genFieldDecls(t *testing.T) {
	type fields struct {
		buf bytes.Buffer
	}
	type args struct {
		prefixes []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				buf: tt.fields.buf,
			}
			if err := g.genFieldDecls(tt.args.prefixes...); (err != nil) != tt.wantErr {
				t.Errorf("Generator.genFieldDecls() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_format(t *testing.T) {
	type fields struct {
		buf bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				buf: tt.fields.buf,
			}
			got, err := g.format()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.format() = %v, want %v", got, tt.want)
			}
		})
	}
}
