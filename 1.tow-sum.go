/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
    mapValue := make(map[int]int)
    for k,v := range nums{
        if i,ok := mapValue[target - v];ok{
            return []int{i,k}
        }
        mapValue[v] = k
    }
    return []int{}
}
