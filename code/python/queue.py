class Queue:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def enqueue(self, item):
        if self.length > 0:
            self.tail.next = item
        if self.length == 0:
            self.head = item
        self.tail = item
        self.length += 1
        print("inserted node", item.value, "into queue's end")
        return self.head

    def dequeue(self):
        if self.length == 0:
            print("queue is empty, can't dequeue")
            return None
        if self.length == 1:
            print("removed node", self.head, "from the front of the queue")
            self.tail = None
            self.head = None
        if self.length > 0:
            print("removed node", self.head, "from the front of the queue")
            dequeued = self.head
            self.head = dequeued.next
        self.length -= 1
        return dequeued

    def qprint(self):
        current = self.head
        for i in range(self.length):
            print(current.value)
            if current.next is not None:
                current = current.next


class Node:
    def __init__(self, value):
        self.value = value
        self.next = None
