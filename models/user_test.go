package models

import (
	"testing"
)

func FuzzGenerateJWT(f *testing.F) {
	testcase := []string{"123456", "fJMTksCtDib9d2", "p8C5h5KzAPhYCQ"}
	for _, tc := range testcase {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, pass string) {
		user := &User{
			Password: pass,
		}

		if err := user.HashPassword(user.Password); err != nil {
			t.Errorf("HashPassword() error = %v", err)
		}

		if err := user.CheckPassword(pass); err != nil {
			t.Errorf("CheckPassword() error = %v", err)
		}

	})
}
