package errors

import (
	"reflect"
	"testing"
)

func Test_myCrawlerError_Type(t *testing.T) {
	type fields struct {
		errType    ErrorType
		errMsg     string
		fullErrMsg string
	}
	tests := []struct {
		name   string
		fields fields
		want   ErrorType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := &myCrawlerError{
				errType:    tt.fields.errType,
				errMsg:     tt.fields.errMsg,
				fullErrMsg: tt.fields.fullErrMsg,
			}
			if got := ce.Type(); got != tt.want {
				t.Errorf("myCrawlerError.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myCrawlerError_Error(t *testing.T) {
	type fields struct {
		errType    ErrorType
		errMsg     string
		fullErrMsg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := &myCrawlerError{
				errType:    tt.fields.errType,
				errMsg:     tt.fields.errMsg,
				fullErrMsg: tt.fields.fullErrMsg,
			}
			if got := ce.Error(); got != tt.want {
				t.Errorf("myCrawlerError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myCrawlerError_genFullErrMsg(t *testing.T) {
	type fields struct {
		errType    ErrorType
		errMsg     string
		fullErrMsg string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := &myCrawlerError{
				errType:    tt.fields.errType,
				errMsg:     tt.fields.errMsg,
				fullErrMsg: tt.fields.fullErrMsg,
			}
			ce.genFullErrMsg()
		})
	}
}

func TestNewCrawlerError(t *testing.T) {
	type args struct {
		errType ErrorType
		errMsg  string
	}
	tests := []struct {
		name string
		args args
		want CrawlerError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCrawlerError(tt.args.errType, tt.args.errMsg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCrawlerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIllegalParameterError_Error(t *testing.T) {
	type fields struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipe := IllegalParameterError{
				msg: tt.fields.msg,
			}
			if got := ipe.Error(); got != tt.want {
				t.Errorf("IllegalParameterError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIllegalParameterError(t *testing.T) {
	type args struct {
		errMsg string
	}
	tests := []struct {
		name string
		args args
		want IllegalParameterError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIllegalParameterError(tt.args.errMsg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIllegalParameterError() = %v, want %v", got, tt.want)
			}
		})
	}
}
