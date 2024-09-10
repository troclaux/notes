class Node:
    def __init__(self, value):
        self.value = value
        self.next = None


# max heap
class Heap:
    def __init__(self):
        self.head = None
        self.size = 0

    def insert(self, value):
        if self.size == 0:
            self.head = Node(value)
            self.size += 1
            return

        curr = self.head
        while value < curr.value:
            if curr.next is None:
                curr.next = Node(value)
                self.size += 1
                return
            prev = curr
            curr = curr.next
        prev.next = Node(value)
        prev.next.next = curr

    def __str__(self):
        curr = self.head
        while curr:
            print(curr.value)
            curr = curr.next
        return ""
