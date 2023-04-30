package main

import (
	"fmt"
	middle_square_method "github.com/cryptography-research-lab/middle-square-method"
)

func main() {

	// 指定seed
	//generator := middle_square_method.NewMiddleSquareMethodRandomGenerator(123456788)
	// 不指定的话默认会把当前unix毫毛时间戳作为随机数序列的seed
	generator := middle_square_method.NewMiddleSquareMethodRandomGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(generator.Next())
	}
	// Output:
	// 8368732329101
	// 5929387542414
	// 5544795373933
	// 112994856646
	// 2690729443080
	// 9928319934722
	// 3156013888483
	// 1967977745017
	// 5593118728487
	// 833319471388

}
