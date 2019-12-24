package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//arr :=[]int{3,6,4,2,8,5}
	//
	const size = 10000
	arr:=make([]int,size)
	for i:=0 ; i< size;i++{
		arr[i] = rand.Intn(size)
	}
	start := time.Now().Unix()
	//fmt.Printf("%v",arr[0:])
	sort(arr,0, len(arr)-1)
	fmt.Printf("%v\n",arr)
	fmt.Println(time.Now().Unix()-start)

}
func sort(arr []int,start,end int){
	//fmt.Printf("s=%v ,end= %v\n",start,end)
	if start<end { //如果start>=end 说明左右两边仅有一个元素
		mid := (end+start)/2
		sort(arr,start,mid)
		sort(arr,mid+1,end)
		merge(arr,start,mid,end)
	}
}
//merge过程就是排序的过程
func merge(arr []int,start,mid,end int ){
	//fmt.Printf("s=%v ,mid=%v,end= %v\n",start,mid,end)
	//p和q分别是左右两个子数组的起始索引
	leftIndex:=start
	rightIndex:=mid+1
	temp:=make([]int,end-start+1)
	k:=0
	for i:=start;i<=end;i++{
		if leftIndex>mid{//左边已没有数据，直接把右边元素copy到temp。这种方式可以避免使用哨兵元素（最大值 ）
			temp[k]=arr[rightIndex]
			k++
			rightIndex++
		}else if rightIndex>end{
			temp[k]=arr[leftIndex]
			k++
			leftIndex++
		}else if arr[leftIndex]<arr[rightIndex]{
			temp[k]=arr[leftIndex]
			k++
			leftIndex++
		}else{
			temp[k]=arr[rightIndex]
			k++
			rightIndex++
		}
	}

	for i:=0;i<k;i++{ //把merge的临时数组copy回原数组
		arr[start]=temp[i]
		start++
	}
}