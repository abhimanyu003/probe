package parser

import (
	"reflect"
	"testing"
)

func TestParseYaml(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "read bytes",
			args: args{
				data: []byte(`name: SetHeaders`),
			},
			want: []byte(`{"name":"SetHeaders"}`),
		},
		{
			name: "should return error on wrong yaml input",
			args: args{
				data: []byte(`name -- SetHeaders`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYaml(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYaml() got = %v, want %v", got, tt.want)
			}
		})
	}
}
