# queue implementation in golang using slices
Queue - follows FIFO principle - first in first out - first element to enter the queue is evicted the first
Operations - Enqueue(Add an element to queue), DeQueue(Delete an element from queue), Peek(Returns front of the queue)
Special conditions - return error in case when no element is present in queue and dequeue operation is requested. 
This implementation is go-routine safe as well and that is because we are making use of mutexes.

Note : Once a queue is 
