/*
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    result := []int{}
    for _,list := range lists{
        for list != nil {
            result = append(result,list.Val)
            list = list.Next
        }
    }
    sort.Ints(result) // 对数组排序
    
    head := new(ListNode)
    ptail := head
    for _,v := range result{
        fmt.Println(v)
        l := new(ListNode)
        l.Val = v
        l.Next = nil
        ptail.Next = l
        ptail = l
    }   
    head = head.Next
    ptail.Next = nil
    return head
}