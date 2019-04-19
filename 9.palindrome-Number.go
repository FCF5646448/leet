/*
	判断回文数
*/
func isPalindrome(x int) bool {
    if x < 0{
        return false
    }
    var y int
    temp := x
    for x > 0{
        y = y *10 +x % 10
        x = x / 10
    }
    if y== temp{
        return true
    }
    return false
}
