package runner

import (
	"os"
	"probe/cache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseVariables(t *testing.T) {
	type args struct {
		input string
		cache *cache.Cache
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parse user variable",
			args: args{
				input: "${var}",
				cache: cache.NewCache("testing"),
			},
			want: "var",
		},
		{
			name: "parse env variable",
			args: args{
				input: "${env:var}",
				cache: cache.NewCache("testing"),
			},
			want: "var",
		},

		{
			name: "parse user variable with spaces",
			args: args{
				input: "${ var  }",
				cache: cache.NewCache("testing"),
			},
			want: "var",
		},
		{
			name: "parse env variable with spaces",
			args: args{
				input: "${ env:var  }",
				cache: cache.NewCache("testing"),
			},
			want: "var",
		},
		{
			name: "should not parse variable",
			args: args{
				input: "env:var",
				cache: cache.NewCache("testing"),
			},
			want: "env:var",
		},
		{
			name: "should not parse variable",
			args: args{
				input: "${env:var",
				cache: cache.NewCache("testing"),
			},
			want: "${env:var",
		},
		{
			name: "user cache value not present",
			args: args{
				input: "${something}",
				cache: cache.NewCache("testing"),
			},
			want: "${something}",
		},
		{
			name: "env value not present",
			args: args{
				input: "${env:something}",
				cache: cache.NewCache("testing"),
			},
			want: "${env:something}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("var", "var")
			tt.args.cache.Set("var", "var")

			assert.Equalf(t, tt.want, parseVariables(tt.args.input, tt.args.cache), "parseVariables(%v, %v)", tt.args.input, tt.args.cache)
		})
	}
}
