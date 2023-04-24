package pwd

import (
	"fmt"
	"testing"
)

func TestBcryptPw(t *testing.T) {
	fmt.Println(BcryptPw("01312934a"))
}

func TestVerifyPwd(t *testing.T) {
	fmt.Println(VerifyPwd("01312934a", "$2a$04$uVOra4WkkTszPpiGpwKnkOmxeBFdA1ErwqdDQXIiSdOvWX2uU2fJO"))
}
