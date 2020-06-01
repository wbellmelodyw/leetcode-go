package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//归并排序
	n1,n2 := 0,0
	merge := make([]int,0)
	for n1<len(nums1) || n2 < len(nums2){
		if n1<len(nums1) && n2<len(nums2){
			if nums1[n1] <= nums2[n2]{
				merge = append(merge,nums1[n1])
				n1++
			}else{
				merge = append(merge,nums2[n2])
				n2++
			}
		}
		if n1<len(nums1) && n2>len(nums2)-1{
			merge = append(merge,nums1[n1])
			n1++
		}
		if n2<len(nums2)  && n1>len(nums1)-1{
			merge = append(merge,nums2[n2])
			n2++
		}
	}
	if len(merge) % 2 == 0{
		tag := len(merge)/2
		return float64(merge[tag]+merge[tag-1])/2
	}else{
		tag := len(merge)/2
		return float64(merge[tag])
	}
}

func main(){
	fmt.Println(findMedianSortedArrays([]int{1,3},[]int{2}))
}
