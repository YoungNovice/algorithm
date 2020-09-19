package expresstion

import (
	"fmt"
	"testing"
)

func TestPostfixExpression(t *testing.T) {
	str := "16 9 4 3 +*-"
	fmt.Println(PostfixExpression(str))
}
