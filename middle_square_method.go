package middle_square_method

import "time"

type MiddleSquareMethodRandomGenerator struct {
	length uint64
	seed   uint64
}

// NewMiddleSquareMethodRandomGenerator seed不指定的话默认为当前unix毫秒时间戳
func NewMiddleSquareMethodRandomGenerator(seed ...uint64) *MiddleSquareMethodRandomGenerator {
	if len(seed) == 0 {
		seed = append(seed, uint64(time.Now().UnixMilli()))
	}
	return &MiddleSquareMethodRandomGenerator{
		length: computeLength(seed[0]),
		seed:   seed[0],
	}
}

func (x *MiddleSquareMethodRandomGenerator) SetLength(length uint64) {
	x.length = length
}

func (x *MiddleSquareMethodRandomGenerator) SetSeed(seed uint64) {
	x.seed = seed
}

func computeLength(seed uint64) uint64 {
	lengthCount := 0
	for seed > 0 {
		lengthCount++
		seed /= 10
	}
	return uint64(lengthCount)
}

// Next 返回下一个随机数
func (x *MiddleSquareMethodRandomGenerator) Next() uint64 {
	x.seed = x.takeMiddle(x.seed * x.seed)
	return x.seed
}

func (x *MiddleSquareMethodRandomGenerator) takeMiddle(n uint64) uint64 {

	// 先丢弃右边的length/2个字符
	needDropCount := x.length / 2
	for needDropCount > 0 {
		n /= 10
		needDropCount--
	}

	// 然后取length个字符，作为获取到的中间的结果，如果长度不够的话则左边空着
	// 因为是数值类型，所以左边不能补0
	result := uint64(0)
	needTakeCount := x.length
	weight := 0
	for needTakeCount > 0 {
		result = (n%uint64(10))*x.pow(weight) + result
		n /= 10
		needTakeCount--
		weight++
	}
	return result
}

// 计算10的n次幂
func (x *MiddleSquareMethodRandomGenerator) pow(n int) uint64 {
	result := uint64(1)
	for n > 0 {
		result *= 10
		n--
	}
	return result
}

// IsZero 判断当前是否已经退化为0了，这是平方取中法会有的缺陷，当退化为0的时候生成器便失效无法生成新的随机数
func (x *MiddleSquareMethodRandomGenerator) IsZero() bool {
	return x.seed == 0
}
