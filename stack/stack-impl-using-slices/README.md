# stack implementation in golang using slices
Stack - follows LIFO principle - last in first out - last element to enter the stack is evicted the first
Operations - Push(Add an element to stack), Pop(Delete an element from stack), Peek(Returns top most element of stack)
Special conditions - Underflow(When stack is empty and we try to delete an element) and Overflow(When stack is full and we try to insert an element : Not implemented in this example)
This implementation is go-routine safe as well and that is because we are making use of mutex