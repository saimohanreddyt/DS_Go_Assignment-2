//Queue
//-x-x-x-x-
package main
import "fmt"

func enqueue(queue[] int, element int) []int {
  queue = append(queue, element); 
  fmt.Println("Enqueued:", element);
  return queue
}

func dequeue(queue[] int) ([]int) {
  element := queue[0]; 
  fmt.Println("Dequeued:", element)
  return queue[1:]; 
}

func main() {
  var queue[] int; 

  queue = enqueue(queue, 30);
  queue = enqueue(queue, 40);
  queue = enqueue(queue, 50);

  fmt.Println("Queue:", queue);

  queue = dequeue(queue);
  fmt.Println("Queue:", queue);

  queue = enqueue(queue, 40);
  fmt.Println("Queue:", queue);
}

/*

Introduction
A queue is a basic data structure that organized items in first-in-first-out
(FIFO) manner. First element inserted into a queue will be the first to be
removed. It is also known as "first-come-first-served".

*/

/*

The Queue Abstract Data Type
-x-x-x-x-x-x-x-x-x-x-x-x-x-x-x-
Queue abstract data type is defined as a class whose object follows FIFO or
first-in-first-out for the elements, added to it.
Queue should support the following operation:
1. add(): Which add a single element at the back of a queue
2. remove(): Which remove a single element from the front of a queue.
3. isEmpty(): Returns 1 if the queue is empty
4. size(): Returns the number of elements in a queue.

*/
