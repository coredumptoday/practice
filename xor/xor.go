package xor

func Swap(arr []int, i, j int) {
	if arr[i] == arr[j] { //相同的数字异或结果为0，该条件要独立判断
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

//获取唯一出现奇数次的数字
func SearchOnlyOneShowOddNumFromSlice(arr []int) int {
	num := 0
	for _, v := range arr {
		num = num ^ v
	}
	return num
}

//获取两个出现奇数次的数字
func SearchTwoShowOddNumFromSlice(arr []int) (int, int) {
	xor := 0
	for _, v := range arr {
		xor = xor ^ v
	}

	rightOne := xor & (-xor)	//取出最右的1，也就是最小的不同位
	x := 0
	for _, v := range arr {
		if v & rightOne != 0 {	//根据最小不同位将数组分位两组
			x = x ^ v
		}
	}

	return x, xor ^ x
}