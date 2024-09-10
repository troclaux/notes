class Stack:
    def __init__(self):
        self.arrows = []

    def push(self, arrow):
        self.arrows.append(arrow)

    def pop(self):
        if len(self.arrows) < 1:
            return None
        return self.arrows.pop()

    def peek(self):
        if len(self.arrows) < 1:
            return None
        return self.arrows[-1]

    def size(self):
        return len(self.arrows)
