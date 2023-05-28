package authjwt

import "testing"

func TestGenerateJWT(t *testing.T) {
	type args struct {
		email string
		pass  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Case 1",
			args{"teste@teste.com", "teste"},
			false,
		},
		{
			"Case 2",
			args{"rodinei@gmail.com", "123456"},
			false,
		},
		{
			"Case 3",
			args{"pessoa@gmail.com", "123teste"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenString, err := GenerateJWT(tt.args.email, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = ValidateToken(gotTokenString)
			if err != nil {
				t.Errorf("ValidateToken() error %v", err)
			}
		})
	}
}

func FuzzGenerateJWT(f *testing.F) {
	testcase := []string{"teste@teste.com", "rodinei@gmail.com", "pessoa@gmail.com"}
	for _, tc := range testcase {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, email string) {
		gotTokenString, err := GenerateJWT(email, "pass")
		if err != nil {
			t.Errorf("GenerateJWT() error = %v", err)
			return
		}
		err = ValidateToken(gotTokenString)
		if err != nil {
			t.Errorf("ValidateToken() error %v", err)
		}
	})
}
