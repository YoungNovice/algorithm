package brackets

import (
	"fmt"
	"testing"
)

func TestPostfixExpression(t *testing.T) {
	str := "16 9 4 3 +*-"
	fmt.Println(PostfixExpression(str))
}

func TestMatch(t *testing.T) {
	b := Match2("(asdf(s)(())()){(){}[nihao]}")
	if !b {
		t.Error("not match")
	} else {
		t.Log("match")
	}
}

func TestMatch1(t *testing.T) {
	var1 := new(int32)
	fmt.Printf("%T->%v\n", var1, var1)
	var2 := (*int32)(var1)
	fmt.Printf("%T->%v\n", var2, var2)
}