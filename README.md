# 平方取中法随机数生成器

# 一、平方取中法介绍

平方取中方法是1946年由John Von Neumann，S. Ulm和N. Metropolis 在Los Alamos实验室研究中子碰撞时提出的，这是一种用来生成随机数的算法，但是因为存在致命的缺陷，因此不适合用在生产环境中。

# 二、添加依赖

```bash
go get -u github.com/cryptography-research-lab/middle-square-method
```

注：此算法过于久远，仅作为技术研究，已经不适合用于生产环境，请谨慎使用。

# 三、API Example 

```go
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
```

# 四、平方取中法生成过程

1. 选择一个m位的数字作为随机数种子seed
2. 把seed开方，得到结果r = seed * seed 
3. 如果r的位数不足2*m位，则左边补0到2 * m 位
4. 然后取中间的m位作为生成的结果，同时作为下一次生成的seed
5. 如此往复，便能生成随机数序列

# 五、 缺点

当多次生成的时候可能会退化为0

TODO 证明退化为0的过程 

# 六、参考资料

- https://baike.sogou.com/v71147797.htm
- https://www.cnblogs.com/hanyu1995/p/14658844.html











