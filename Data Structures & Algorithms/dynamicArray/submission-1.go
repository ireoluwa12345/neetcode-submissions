type DynamicArray struct {
	array []int
	size int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	if capacity <= 0 {
        capacity = 1
    }
	da := &DynamicArray{
		array: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}

	return da
}

func (da *DynamicArray) Get(i int) int {
	return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size == da.capacity {
		da.resize()
	}

	da.array[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	value := da.array[da.size - 1]
	da.size--
	return value
}

func (da *DynamicArray) resize() {
	newCap := da.capacity * 2
	newArr := make([]int, newCap)

	copy(newArr, da.array[:da.size])
	da.array = newArr
	da.capacity = newCap
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
