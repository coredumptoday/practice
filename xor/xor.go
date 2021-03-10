package xor

func Swap(arr []int, i, j int) {
	if arr[i] == arr[j] {
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
