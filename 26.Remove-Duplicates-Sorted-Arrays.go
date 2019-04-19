/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

*/
func removeDuplicates(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    curr := nums[0]   
    index := 1
    for i := 1; i < len(nums); i++{
        if nums[i] != curr{
            nums[index] = nums[i]	
            curr = nums[i]
            index ++
        } 
    }
    nums =nums[:index]
    return index
}