package apiserver

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				cfg: Config{
					BindAddress: "127.0.0.2",
					SecurePort:  8443,
					EtcdEndpoints: []string{
						"http://localhost:2379",
					},
				},
			},
			wantNil: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("New() should not got nil")
			}
		})
	}
}

func TestDefault(t *testing.T) {
	tests := []struct {
		name    string
		wantNil bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Default()
			if (err != nil) != tt.wantErr {
				t.Errorf("Default() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("New() should not got nil")
			}
		})
	}
}
