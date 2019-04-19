/* Sort a linked list in O(n log n) time using constant space complexity.

Input: 4->2->1->3
Output: 1->2->3->4

将比基元小的节点的值，依次和基元后的节点的值交换。 
如 5->3->6->4->7->2 
则 5 为基元 
3 < 5: swap(3, 3) ，(起始交换位置为基元的下一个节点，即第2个节点) 
6 > 5: continue; 
4 < 5: swap(6, 4) （交换位置后移，交换4和第3个节点的值） 
7 > 5: continue 
2 < 5: swap(4, 2) （交换位置后移，交换2和第4个节点的值）

循环结束 swap(5, 2) （交换基元值和第4个节点的值）

最后得到：2->3->4->5->6->7
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    quicksort(head,nil)
    return head
}

func quicksort(head,end *ListNode){
    if head == end || head.Next == end{
        return
    }
    pt := getPartion(head,end)
    quicksort(head,pt)
    quicksort(pt.Next,end)
}

func getPartion(head,end *ListNode)*ListNode{
    if head == end || head.Next == end{
        return head
    }
    temp := head.Val
    p := head
    q := head.Next
    for q != end{
        if q.Val < temp{
            p = p.Next // 起始交换位置为基元的下一个节点
            swap(&p.Val,&q.Val)
        }
        q = q.Next
    }
    swap(&p.Val,&head.Val)
    return p
}
func swap(a, b *int){
    temp := *a
    *a = *b
    *b = temp
}