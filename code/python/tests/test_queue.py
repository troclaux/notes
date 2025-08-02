from ..queue import Node, Queue

def test_queue():
    q = Queue()

    print("\n--- Enqueuing 1, 2, 3 ---")
    q.enqueue(1)
    q.enqueue(2)
    q.enqueue(3)

    print("\n--- Current Queue ---")
    q.display()

    print("\n--- Dequeue Once ---")
    q.dequeue()

    print("\n--- Current Queue ---")
    q.display()

    print("\n--- Dequeue All ---")
    q.dequeue()
    q.dequeue()

    print("\n--- Dequeue from Empty ---")
    q.dequeue()

    print("\n--- Enqueue After Emptying ---")
    q.enqueue(99)
    q.display()

test_queue()
