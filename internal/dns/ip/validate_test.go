package ip

import "testing"

func Test_validate(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "regular IP should pass validation",
			args:    args{ip: "8.8.8.8"},
			wantErr: false,
		},
		{
			name:    "invalid IP range should fail validation",
			args:    args{ip: "155.264.3.100"},
			wantErr: true,
		},
		{
			name:    "wrong input should fail validation",
			args:    args{ip: "foo"},
			wantErr: true,
		},
		{
			name:    "valid ip with spaces should pass validation",
			args:    args{ip: "  192.123.0.204  "},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
