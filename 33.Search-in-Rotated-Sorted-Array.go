/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

旋转数组的二分查找 时间复杂度 O(log n)
*/
func search(nums []int, target int) int {
    l := 0
    h := len(nums) - 1
    for l <= h{
        if nums[l] == target {
            return l
        }
        if nums[h] == target {
            return h
        }
        mid := l + (h-l)/2
        if nums[mid] == target{
            return mid
		} 
		// 5 6 1 2 3 4
        if nums[l] > nums[mid]{ 
            if target > nums[mid] && target < nums[l]{
                l = mid +1 //右边查找
            } else{
                h = mid -1 
            }
        }else{
            if target > nums[l] && target < nums[mid]{
                h = mid -1 // 左边查找
            }else{
                l = mid +1
            }
        }
    }
    return -1
}