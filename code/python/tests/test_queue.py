import pytest
from ..queue import Queue

def test_enqueue_adds_items():
    q = Queue()
    q.enqueue(1)
    q.enqueue(2)
    assert q.items == [1, 2]

def test_dequeue_removes_and_returns_front_item():
    q = Queue()
    q.enqueue(1)
    q.enqueue(2)
    val = q.dequeue()
    assert val == 1
    assert q.items == [2]

def test_peek_returns_front_without_removing():
    q = Queue()
    q.enqueue(5)
    q.enqueue(10)
    front = q.peek()
    assert front == 5
    assert q.items == [5, 10]

def test_peek_on_empty_queue_returns_none():
    q = Queue()
    assert q.peek() is None

def test_dequeue_from_empty_queue_raises_error():
    q = Queue()
    with pytest.raises(IndexError, match="Dequeue from empty queue"):
        q.dequeue()

