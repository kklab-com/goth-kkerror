package kkerror

import (
	"reflect"
	"testing"
)

func TestDefaultKKError_Level(t *testing.T) {
	type fields struct {
		ErrorLevel    Level
		ErrorCategory Category
		ErrorCode     string
		ErrorMessage  string
		error         KKError
	}

	tests := []struct {
		name   string
		fields fields
		want   Level
	}{
		{
			fields: fields{
				ErrorLevel:    Critical,
				ErrorCategory: Server,
				ErrorCode:     DefaultErrorCode,
				ErrorMessage:  "",
				error:         nil,
			},
			want: Critical,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &DefaultKKError{
				ErrorLevel:    tt.fields.ErrorLevel,
				ErrorCategory: tt.fields.ErrorCategory,
				ErrorCode:     tt.fields.ErrorCode,
				ErrorMessage:  tt.fields.ErrorMessage,
				error:         tt.fields.error,
			}

			if got := k.Level(); got != tt.want {
				t.Errorf("Level() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestWrappedError(t *testing.T) {
	type args struct {
		error KKError
	}
	tests := []struct {
		name string
		args args
		want *DefaultKKError
	}{
		{
			args: args{error: &DefaultKKError{
				ErrorLevel:    Normal,
				ErrorCategory: Undefined,
				ErrorCode:     DefaultErrorCode,
				ErrorMessage:  "error",
				error:         nil,
			}},
			want: WrappedError(Error("error")),
		},
		{
			args: args{
				error: &DefaultKKError{
					ErrorLevel:    Normal,
					ErrorCategory: Undefined,
					ErrorCode:     DefaultErrorCode,
					ErrorMessage:  "",
					error: &DefaultKKError{
						ErrorLevel:    Normal,
						ErrorCategory: Undefined,
						ErrorCode:     DefaultErrorCode,
						ErrorMessage:  "error",
						error:         nil,
					},
				},
			},
			want: WrappedError(WrappedError(Error("error"))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WrappedError(tt.args.error); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrappedError() = %v, want %v", got, tt.want)
			}
		})
	}
}
