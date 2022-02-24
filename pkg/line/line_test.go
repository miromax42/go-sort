package line

import (
	"reflect"
	"testing"
)

func Test_new(t *testing.T) {
	type args struct {
		str   string
		group string
		date  string
	}
	tests := []struct {
		name     string
		args     args
		wantLine Line
		wantErr  bool
	}{
		{
			"t01",
			args{
				str: "d034",
			},
			Line{
				Prefix: "d",
				num:    34,
			},
			false,
		},
		{
			"t01without0",
			args{
				str: "sd34",
			},
			Line{},
			true,
		},
		{
			"t01withcomment",
			args{
				str: "sd034qwe",
			},
			Line{
				Prefix:  "sd",
				num:     34,
				postfix: "qwe",
			},
			false,
		},
		{
			"withotherNum",
			args{
				str: "s2005dsf05d034q34we",
			},
			Line{
				Prefix:  "s2005dsf05d",
				num:     34,
				postfix: "q34we",
			},
			false,
		},
		{
			"test02",
			args{
				str: "d015    3m",
			},
			Line{
				Prefix:  "в",
				num:     15,
				postfix: " 3ц",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLine, err := newLine(tt.args.str, tt.args.group, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("new() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLine, tt.wantLine) {
				t.Errorf("new() = %v, want %v", gotLine, tt.wantLine)
			}
		})
	}
}
