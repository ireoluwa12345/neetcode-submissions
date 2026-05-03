type Node struct {
    value int
    next  *Node
}

type LinkedList struct {
    head *Node
    tail *Node
    size int
}

func NewLinkedList() *LinkedList {
    // Initialize with nil so the list starts truly empty
    return &LinkedList{
        head: nil,
        tail: nil,
        size: 0,
    }
}

func (ll *LinkedList) Get(index int) int {
    // Corrected condition: index must be within [0, size-1]
    if index < 0 || index >= ll.size {
        return -1
    }

    current := ll.head
    for i := 0; i < index; i++ {
        current = current.next
    }

    return current.value
}

func (ll *LinkedList) InsertHead(val int) {
    newNode := &Node{value: val}

    if ll.head == nil {
        ll.head = newNode
        ll.tail = newNode
    } else {
        newNode.next = ll.head
        ll.head = newNode
    }
    ll.size++
}

func (ll *LinkedList) InsertTail(val int) {
    newNode := &Node{value: val}

    if ll.head == nil {
        ll.head = newNode
        ll.tail = newNode
    } else {
        ll.tail.next = newNode
        ll.tail = newNode
    }
    ll.size++
}

func (ll *LinkedList) Remove(index int) bool {
    if index < 0 || index >= ll.size {
        return false
    }

    if index == 0 {
        ll.head = ll.head.next
        if ll.head == nil {
            ll.tail = nil
        }
        ll.size--
        return true
    }

    current := ll.head
    for i := 0; i < index-1; i++ {
        current = current.next
    }

    // If removing the tail, update the tail pointer
    if index == ll.size-1 {
        ll.tail = current
        current.next = nil
    } else {
        current.next = current.next.next
    }

    ll.size--
    return true
}

func (ll *LinkedList) GetValues() []int {
    values := make([]int, 0, ll.size)
    current := ll.head

    for current != nil {
        values = append(values, current.value)
        current = current.next
    }

    return values
}