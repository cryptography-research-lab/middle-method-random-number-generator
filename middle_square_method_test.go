package middle_square_method

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	//generator := NewMiddleSquareMethodRandomGenerator( 2041)
	generator := NewMiddleSquareMethodRandomGenerator(675248)
	for i := 0; i < 20; i++ {
		fmt.Println(generator.Next())
	}
	// Seed:
	// 2041
	//Output:
	// 1656
	// 7423
	// 1009
	// 180
	// 324
	// 1049
	// 1004
	// 80
	// 64
	// 40
	// 16
	// 2
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0

}
