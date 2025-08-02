
class Queue:
    def __init__(self):
        self.items = []

    def enqueue(self, value):
        self.items.append(value)

    def dequeue(self):
        if len(self.items) == 0:
            raise IndexError("Dequeue from empty queue")
        return self.items.pop(0)

    def peek(self):
        if len(self.items) == 0:
            return None
        return self.items[0]

