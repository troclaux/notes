class Stack:
    def __init__(self):
        self.stack = []
        self.size = 0

    def push(self, value):
        self.stack.append(value)
        self.size += 1

    def pop(self):
        if len(self.stack) == 0:
            raise IndexError("can't pop when stack is empty")
        removed_value = self.stack.pop()
        self.size -= 1
        return removed_value

    def peek(self):
        if len(self.stack) == 0:
            raise IndexError("can't peek when stack is empty")
        return self.stack[-1]

    def __str__(self):
        return str(self.stack)
