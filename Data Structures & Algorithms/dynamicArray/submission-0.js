class DynamicArray {
    /**
     * @constructor
     * @param {number} capacity
     */
    constructor(capacity) {
        if (capacity > 0){
            this.capacity = capacity;
            this.size = 0;
            this.array = [];
            return null;
        }

        return null;
    }

    /**
     * @param {number} i
     * @returns {number}
     */
    get(i) {
        if (this.size < 0){
            return null
        }
        return this.array[i] 
    }

    /**
     * @param {number} i
     * @param {number} n
     * @returns {void}
     */
    set(i, n) {
        for (let j = 0; j < this.capacity; j++) {
            if (j == i){
                this.array[i] = n
            }
        }

        return null
    }

    /**
     * @param {number} n
     * @returns {void}
     */
    pushback(n) {
        if (this.size === this.capacity) {
            this.resize();
        }
        this.array[this.size] = n;
        this.size++;

        return null
    }

    /**
     * @returns {number}
     */
    popback() {
        let value = this.array[this.size - 1];
        this.size--;
        return value;
    }

    /**
     * @returns {void}
     */
    resize() {
        this.capacity = this.capacity * 2
    }

    /**
     * @returns {number}
     */
    getSize() {
        return this.size
    }

    /**
     * @returns {number}
     */
    getCapacity() {
        return this.capacity
    }
}
