class Node {
    constructor(value) {
        this.value = value;
        this.next = null; // Pointer to the next node
    }
}

class LinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
        this.size = 0;
    }

    /**
     * @param {number} index
     * @return {number}
     */
    get(index) {
        if (index < 0 || index >= this.size) {
            return -1;
        }

        let current = this.head;
        for (let i = 0; i < index; i++) {
            current = current.next
        }
        return current.value
    }

    /**
     * @param {number} val
     * @return {void}
     */
    insertHead(val) {
        let newNode = new Node(val)
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode
        }else{
            newNode.next = this.head;
            this.head = newNode;
        }
        this.size++
    }

    /**
     * @param {number} val
     * @return {void}
     */
    insertTail(val) {
        let newNode = new Node(val)
        if (!this.head) {
            this.head = newNode
            this.tail = newNode
        }else{
            this.tail.next = newNode
            this.tail = newNode
        }

        this.size++
    }

    /**
     * @param {number} index
     * @return {boolean}
     */
        remove(index) {
        if (index < 0 || index >= this.size) return false;

        if (index === 0) {
            this.head = this.head.next;
            if (!this.head) this.tail = null;
            this.size--;
            return true;
        }

        let prev = this.head;
        for (let i = 0; i < index - 1; i++) {
            prev = prev.next;
        }
        prev.next = prev.next.next;

        if (index === this.size - 1) {
            this.tail = prev;
        }

        this.size--;
        return true;
    }


    /**
     * @return {number[]}
     */
    getValues() {
        let valuesArray = [];
        if (this.size < 0) {
            return valuesArray;
        }

        let current = this.head;
        for (let i = 0; i < this.size; i++) {
            valuesArray.push(current.value)
            current = current.next
        }

        return valuesArray;
    }
}
