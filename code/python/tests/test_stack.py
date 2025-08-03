
import pytest
from ..stack import Stack

def test_push_and_pop():
    s = Stack()
    s.push(1)
    s.push(2)
    s.push(3)
    s.pop()
    assert s.stack == [1, 2]

def test_peek():
    s = Stack()
    s.push(10)
    val = s.peek()
    assert val == 10

def test_pop_empty_list():
    s = Stack()
    s.push(5)
    s.pop()
    with pytest.raises(IndexError, match="can't pop when stack is empty"):
        s.pop()
