from python.heap import Heap


def test_heap():
    h = Heap()
    h.insert(3)
    assert h.head.value == 3


def test_heap_str():
    h = Heap()
    h.insert(3)
    h.insert(2)
    h.insert(1)

    expected_str = "3 -> 2 -> 1"
    assert str(h) == expected_str
