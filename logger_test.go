package levellogger

import (
	"testing"
)

func TestLogger_New(t *testing.T) {
	type fields struct {
		LogLevel int
	}
	tests := []struct {
		name   string
		fields fields
		want   Log
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				LogLevel: tt.fields.LogLevel,
			}
			if got := l.New(); got == nil {
				t.Errorf("Logger.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		LogLevel int
	}
	type args struct {
		val []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				LogLevel: InfoLevel,
			},
		},
		{
			name: "test 2",
			fields: fields{
				LogLevel: AllLevel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				LogLevel: tt.fields.LogLevel,
			}
			log := l.New()
			//log.Info(tt.args.val)
			log.Info("this is a test", "this is too")
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		LogLevel int
	}
	type args struct {
		val []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				LogLevel: DebugLevel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				LogLevel: tt.fields.LogLevel,
			}
			log := l.New()
			log.Debug("this is debug logging", "this too")
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		LogLevel int
	}
	type args struct {
		val []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				LogLevel: DebugLevel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				LogLevel: tt.fields.LogLevel,
			}
			log := l.New()
			log.Error("this is a error", "and another")
		})
	}
}

func TestLogger_SetLogLevel(t *testing.T) {
	type fields struct {
		LogLevel int
	}
	type args struct {
		level int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args:args{
				level: OffLevel,
			},
		},
		{
			name: "test 2",
			args:args{
				level: DebugLevel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				LogLevel: tt.fields.LogLevel,
			}
			log:= l.New()
			log.SetLogLevel(tt.args.level)
		})
	}
}
