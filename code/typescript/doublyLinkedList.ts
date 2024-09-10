import { NodeTS } from "./NodeTS";

export class DoublyLinkedList<T> {
    public length: number;
    public head?: NodeTS<T>;
    public tail?: NodeTS<T>;

    constructor() {
        this.length = 0;
        this.head = undefined
        this.tail = undefined
    }

    prepend(item: T): void {

        const node = { value: item } as NodeTS<T>;

        if (!this.head) {
            this.head = node;
            this.tail = node;
            return
        }

        this.head.prev = node;
        node.next = this.head;
        this.head = node;
        this.length++;
        return
    }

    insertAt(item: T, idx: number): void {
        if (idx > this.length) {
            throw new Error("insertion at index is not possible, index is higher than length");
        }

        if (idx === length) {
            this.append(item);
            return;
        } else if (idx === 0) {
            this.prepend(item);
            return;
        }

        this.length++;

        let curr = this.head;
        for (let i = 0; curr && i < idx; i++) {
            curr = curr.next;
        }
        curr = curr as NodeTS<T>;
        const node = { value: item } as NodeTS<T>;

        node.next = curr;
        node.prev = curr.prev;
        curr.prev = node;
        if (curr.prev) {
            curr.prev.next = curr;
        }
    }

    preprend(item: T): void {
        const node = { value: item } as NodeTS<T>;
        this.length++;

        if (!this.head) {
            this.head = node;
            this.tail = node;
            return;
        }

        node.next = this.head;
        this.head.prev = node;
        this.head = node;
        return;

    }

    append(item: T): void {
        const node = { value: item } as NodeTS<T>;
        this.length++;

        if (!this.tail) {
            this.head = node;
            this.tail = node;
            return
        }

        node.prev = this.tail;
        this.tail.next = node;
        this.tail = node;

    }

    remove(item: T): T | undefined {

        if (!this.head) {
            throw new Error("could not delete item");
        }

        let curr = this.head;
        let removedItem;

        for (let i = 0; curr && i < this.length; i++) {

            if (curr.value === item) {

                this.length--;
                removedItem = curr;

                if (curr.prev) {
                    curr.prev.next = curr.next;
                }

                if (curr.next) {
                    curr.next.prev = curr.prev;
                }

                curr.prev = curr.next = undefined;

                if (curr === this.head) {
                    this.head = curr.next
                }

                if (curr === this.tail) {
                    this.tail = curr.prev
                }
            }

            if (curr.value !== item && curr.next) {
                curr = curr.next;
            }
        }
        return removedItem?.value;
    }

    indexOf(item: T) {

        if (!this.head) {
            return -1;
        }

        let curr = this.head;

        for (let i = 0; i < this.length; i++) {

            if (curr.value == item) {
                return i;
            }

            if (!curr.next) {
                return -1;
            }
            curr = curr.next;
        }

        return -1;
    }


    print(): string {

        if (!this.head) {
            console.log("Empty list");
            return "";
        }

        let curr = this.head;
        let str: string = '';

        for (let i = 0; curr && i < this.length; i++) {

            str = str + String(curr.value) + ' -> ';
            if (curr.next) {
                curr = curr.next;
            }
        }
        let output: string = str.substring(0, str.length - 4)
        console.log(output);
        return output;
    }

}


