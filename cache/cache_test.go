package cache

import (
	"testing"

	"github.com/muesli/cache2go"
	"github.com/stretchr/testify/assert"
)

func TestCache_Set(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set string to cache",
			args: args{
				key:   "string",
				value: "test",
			},
		},
		{
			name: "set bool to cache",
			args: args{
				key:   "bool",
				value: true,
			},
		},
		{
			name: "set float to cache",
			args: args{
				key:   "float",
				value: 12.12,
			},
		},
		{
			name: "set map to cache",
			args: args{
				key:   "float",
				value: map[string]bool{"first": true, "second": false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCache("testing")
			c.Set(tt.args.key, tt.args.value)
			cacheValue, err := c.Get(tt.args.key)
			assert.Nil(t, err)
			assert.Exactly(t, cacheValue, tt.args.value)
		})
	}
}

func TestCache_Flush(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	type fields struct {
		cache *cache2go.CacheTable
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set string to cache",
			args: args{
				key:   "string",
				value: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCache("testing")
			c.Set(tt.args.key, tt.args.value)
			c.Flush()
			cacheValue, err := c.Get(tt.args.key)
			assert.NotNil(t, err)
			assert.Nil(t, cacheValue)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "set string to cache",
			args: args{
				key:   "string",
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "set string to cache",
			args: args{
				key:   "string",
				value: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCache("testing")
			c.Set(tt.args.key, tt.args.value)
			if tt.wantErr == true {
				c.Flush()
				cacheValue, err := c.Get(tt.args.key)
				assert.NotNil(t, err)
				assert.Nil(t, cacheValue)
			} else {
				cacheValue, err := c.Get(tt.args.key)
				assert.Nil(t, err)
				assert.Exactly(t, cacheValue, tt.args.value)
			}
		})
	}
}
