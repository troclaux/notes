
class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

class Queue:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def enqueue(self, item):
        if not isinstance(item, Node):
            item = Node(item)

        if self.tail:
            self.tail.next = item
        else:
            self.head = item

        self.tail = item
        self.length += 1
        print("Inserted node", item.value, "into queue's end")
        return self.head

    def dequeue(self):
        if self.length == 0:
            print("Queue is empty, can't dequeue")
            return None

        dequeued = self.head
        print("Removed node", dequeued.value, "from the front of the queue")

        self.head = self.head.next
        if self.head is None:
            self.tail = None

        self.length -= 1
        return dequeued

    def display(self):
        current = self.head
        while current:
            print(current.value)
            current = current.next
