package apiserver

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg *Configuration
	}
	var tests = []struct {
		name string
		args args
		want *Server
	}{
		{"", args{&Configuration{"1.1.1.1", uint16(1443)}}, &Server{&Configuration{"1.1.1.1", uint16(1443)}}},
		{"", args{&Configuration{"2.2.2.2", uint16(2443)}}, &Server{&Configuration{"2.2.2.2", uint16(2443)}}},
		{"", args{&Configuration{"3.3.3.3", uint16(3443)}}, &Server{&Configuration{"3.3.3.3", uint16(3443)}}},
		{"", args{&Configuration{"4.4.4.4", uint16(4443)}}, &Server{&Configuration{"4.4.4.4", uint16(4443)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultConfiguration(t *testing.T) {
	tests := []struct {
		name string
		want *Configuration
	}{
		{"", &Configuration{"0.0.0.0", uint16(443)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultConfiguration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
