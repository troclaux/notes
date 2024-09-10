class BSTNode:
    def __init__(self, val=None):
        self.left = None
        self.right = None
        self.val = val

    def __eq__(self, other):
        return self.val == other.val

    def __lt__(self, other):
        return self.val < other.val

    def __gt__(self, other):
        return self.val > other.val

    def insert(self, val):
        if not self.val:
            self.val = val
            return

        if self.val == val:
            return

        if val < self.val:
            if self.left:
                self.left.insert(val)
                return
            self.left = BSTNode(val)
            return

        if self.right:
            self.right.insert(val)
            return
        self.right = BSTNode(val)

    def delete(self, val):

        if self.val is None:
            return None

        if val < self.val:

            if self.left:
                self.left = self.left.delete(val)

            return self

        if val > self.val:

            if self.right:
                self.right = self.right.delete(val)

            return self

        if self.right is None:
            return self.left

        if self.left is None:
            return self.right

        min_larger_node = self.right

        while min_larger_node.left:
            min_larger_node = min_larger_node.left

        self.val = min_larger_node.val
        self.right = self.right.delete(min_larger_node.val)

        return self
