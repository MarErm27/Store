package main

import (
	"github.com/go-errors/errors"
	"testing"
)

func TestDelete(t *testing.T) {
	type args struct {
		key string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{"read-key", "read-value"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"read-key", "read-value"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"0", "1"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"", ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var contains bool
			defer delete(store, tt.args.key)

			store[tt.args.key] = tt.args.value

			_, contains = store[tt.args.key]
			if !contains {
				t.Error("key/value doesn't exist")
			}

			if err := Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, contains = store[tt.args.key]
			if contains {
				t.Error("Delete failed")
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		key string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{"read-key", "read-value"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"read-key", "read-value"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"0", "1"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"", ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var err error

			defer delete(store, tt.args.key)

			// Read a non-thing
			_, err = Get(tt.args.key)
			if err == nil {
				t.Error("expected an error")
			}
			if !errors.Is(err, ErrorNoSuchKey) {
				t.Error("unexpected error:", err)
			}
			store[tt.args.key] = tt.args.value

			got, err := Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.args.value {
				t.Errorf("Get() got = %v, want %v", got, tt.args.value)
			}

		})
	}
}

func TestPut(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{"create-key", "create-value"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"create-key", "create-value"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"0", "1"},
			wantErr: false,
		},
		{
			name: "",
			args: args{"", ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var contains bool
			var val interface{}

			defer delete(store, tt.args.key)

			// Sanity check
			_, contains = store[tt.args.key]
			if contains {
				t.Error("key/value already exists")
			}

			if err := Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}

			val, contains = store[tt.args.key]
			if !contains {
				t.Error("create failed")
			}

			if val != tt.args.value {
				t.Error("val/value mismatch")
			}
		})
	}
}
