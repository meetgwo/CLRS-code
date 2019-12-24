package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	插入排序
	从第二个元素i开始，数组被分成左右两部分,左边[0,i]和右边[i+,len-1]
	左边始终是已排序好的，右边是未排序的。
	从右边取一个数，对左边已排序的数组从右向左比较，如果大于元素i，向右移动一位。

 */
func main() {
	//arr :=[...]int{3,3,5,2,1,4,8,5,1,9}
	const size = 10000
	var arr [size]int
	for i:=0 ; i< size;i++{
		arr[i] = rand.Intn(size)
	}
	start := time.Now().Unix()

	for i :=2; i < len(arr); i++ {
		key := arr[i] //暂存当前排序的key
		j := i - 1
		for j >= 0 && arr[j] > key {//j<0说明key是最小值，插入在arr[0]
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	consumingTime := time.Now().Unix()-start
	fmt.Printf("%v\n", arr)
	fmt.Print(consumingTime)

}