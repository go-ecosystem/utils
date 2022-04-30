// Package crypto useful crypto functions
// Reference: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
package crypto

import "testing"

func TestEncryptPwd(t *testing.T) {
	hashedPwd := "$2a$04$/91HZrhIr2.X0p2DZx6UZetnaoZyO6E.0JK8j.KBZw.7WPRlBqGGa"

	if !ComparePwd(hashedPwd, "test1") {
		t.Errorf("ComparePwd() error")
	}

	pwd := "asdkhasuifywerfnjsdkayhfriwejl;fmksajhcfier"
	hashedPwd = EncryptPwd(pwd)

	if !ComparePwd(hashedPwd, pwd) {
		t.Errorf("ComparePwd() error")
	}
}
