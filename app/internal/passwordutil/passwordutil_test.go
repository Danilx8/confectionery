package passwordutil

import "testing"

func TestValidateClientPassword(t *testing.T) {
	type args struct {
		password string
		login    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid password (5-20 characters, contains uppercase and lowercase letters, does not contain login)",
			args: args{
				password: "Abc123Def",
				login:    "johndoe",
			},
			wantErr: false,
		},
		{
			name: "Password too short (less than 5 characters)",
			args: args{
				password: "Abc1",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Password too long (more than 20 characters)",
			args: args{
				password: "ThisPasswordIsTooLongAndShouldNotBeAccepted",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Password does not contain uppercase letters",
			args: args{
				password: "abc123def",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Password does not contain lowercase letters",
			args: args{
				password: "ABC123DEF",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Password contains login",
			args: args{
				password: "Johndoe123",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Password contains only digits",
			args: args{
				password: "12345",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Password contains only special characters",
			args: args{
				password: "!@#$%",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Empty password",
			args: args{
				password: "",
				login:    "johndoe",
			},
			wantErr: true,
		},
		{
			name: "Empty login and password",
			args: args{
				password: "",
				login:    "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateClientPassword(tt.args.password, tt.args.login); (err != nil) != tt.wantErr {
				t.Errorf("ValidateClientPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
