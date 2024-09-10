import pytest


def test_heap():
    from python.heap import Heap

    h = Heap()
    h.insert(3)
    assert h.head.value == 3
