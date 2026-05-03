type DynamicArray struct {
    array    []int
    size     int
    capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        array:    make([]int, capacity),
        size:     0,
        capacity: capacity,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
    // if i >= da.capacity {
    // 	da.resize()
    // }
    da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.capacity == da.size {
        da.resize()
    }
    da.array[da.size] = n
    da.size = da.size + 1
}

func (da *DynamicArray) Popback() int {
    da.size = da.size - 1
    return da.array[da.size]
}

func (da *DynamicArray) resize() {
    newDa := make([]int, da.capacity*2)
    for i := 0; i < da.size; i++ {
        newDa[i] = da.array[i]
    }
    da.array = newDa
    da.capacity = da.capacity * 2
}

func (da *DynamicArray) GetSize() int {
    return da.size
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}