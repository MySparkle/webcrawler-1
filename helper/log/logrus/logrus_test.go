package logrus

import (
	"bytes"
	"reflect"
	"testing"
	"webcrawler/helper/log/base"
	"webcrawler/helper/log/field"

	"github.com/Sirupsen/logrus"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
		want base.MyLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoggerBy(t *testing.T) {
	type args struct {
		level   base.LogLevel
		format  base.LogFormat
		options []base.Option
	}
	tests := []struct {
		name       string
		args       args
		want       base.MyLogger
		wantWriter string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if got := NewLoggerBy(tt.args.level, tt.args.format, writer, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoggerBy() = %v, want %v", got, tt.want)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("NewLoggerBy() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func Test_initInnerLogger(t *testing.T) {
	type args struct {
		level  logrus.Level
		format base.LogFormat
	}
	tests := []struct {
		name       string
		args       args
		want       *logrus.Entry
		wantWriter string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if got := initInnerLogger(tt.args.level, tt.args.format, writer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initInnerLogger() = %v, want %v", got, tt.want)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("initInnerLogger() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func Test_loggerLogrus_Name(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			if got := logger.Name(); got != tt.want {
				t.Errorf("loggerLogrus.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loggerLogrus_Level(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	tests := []struct {
		name   string
		fields fields
		want   base.LogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			if got := logger.Level(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loggerLogrus.Level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loggerLogrus_Format(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	tests := []struct {
		name   string
		fields fields
		want   base.LogFormat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			if got := logger.Format(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loggerLogrus.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loggerLogrus_Options(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	tests := []struct {
		name   string
		fields fields
		want   []base.Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			if got := logger.Options(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loggerLogrus.Options() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loggerLogrus_Debug(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Debug(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Debugf(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		format string
		v      []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Debugln(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Debugln(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Error(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Error(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Errorf(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		format string
		v      []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Errorf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Errorln(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Errorln(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Fatal(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Fatal(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Fatalf(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		format string
		v      []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Fatalf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Fatalln(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Fatalln(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Info(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Info(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Infof(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		format string
		v      []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Infof(tt.args.format, tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Infoln(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Infoln(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Panic(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Panic(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Panicf(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		format string
		v      []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Panicf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Panicln(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Panicln(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Warn(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Warn(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Warnf(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		format string
		v      []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Warnf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_loggerLogrus_Warnln(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		v []interface{}
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
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			logger.Warnln(tt.args.v...)
		})
	}
}

func Test_loggerLogrus_WithFields(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	type args struct {
		fields []field.Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   base.MyLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			if got := logger.WithFields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loggerLogrus.WithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loggerLogrus_getInner(t *testing.T) {
	type fields struct {
		level           base.LogLevel
		format          base.LogFormat
		optWithLocation base.OptWithLocation
		inner           *logrus.Entry
	}
	tests := []struct {
		name   string
		fields fields
		want   *logrus.Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &loggerLogrus{
				level:           tt.fields.level,
				format:          tt.fields.format,
				optWithLocation: tt.fields.optWithLocation,
				inner:           tt.fields.inner,
			}
			if got := logger.getInner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loggerLogrus.getInner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLocation(t *testing.T) {
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name string
		args args
		want *logrus.Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLocation(tt.args.entry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
