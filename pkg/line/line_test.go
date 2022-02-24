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
				str: "в034вв",
			},
			Line{
				Prefix:  "в",
				num:     34,
				postfix: "вв",
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
				str: "в015 ыв3m",
			},
			Line{
				Prefix:  "в",
				num:     15,
				postfix: " ыв3m",
			},
			false,
		},
		{
			"t01",
			args{
				str: "в18a015ц3",
			},
			Line{
				Prefix:  "в18a",
				num:     15,
				postfix: "ц3",
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
