package log

import "testing"

func TestL(t *testing.T) {
	E("Helloworld")
	W("Helloworld")
	H("Helloworld")
	I("Helloworld")
	V("Helloworld")

	I("Hello%v", "world")
}
