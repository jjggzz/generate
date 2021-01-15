// author: JGZ
// time:   2021-01-15 17:47
package build

import (
	"fmt"
	"testing"
)

func Test_StrToFirstLetterCapital(t *testing.T) {
	capital := StrFirstLetterToLowercase("Abc")
	fmt.Println(capital)
}

func Test_UnderlineToLowCamel(t *testing.T) {
	capital := StrFirstLetterToLowercase(UnderlineToUpCamel(""))
	fmt.Println(capital)
}
