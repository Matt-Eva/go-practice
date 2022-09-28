package main

import "fmt"

func main(){

}

// REDUCING GARBAGE COLLECTION

// "garbage" refers to data that has no more pointers pointing to it.
// Once there are no more pointers pointing to some data, the
// memory that that data takes up can be reused. If this memory
// isn't recovered, the program's memory usage would continue
// to grow until the computer ran out of RAM.

// The garbage collector's job is to automatically
// detect unused memory and recover it so that it can be
// reused.

// Go's built-in garbage collector makes this process much easier
// since i'ts difficult more people to properly manage memory 
// manually. 

// However, we shouldn't take this garbage collector for granted -
// when possible, we should minimize the amount of garbage we
// create in the first place.


// STACKS AND HEAPS

// STACK

//A stack is a consecutive block of memory, and every function 
// call in thread of execution shares the same stack. Allocating 
// memory on the stack is fast and simple. A stack pointer tracks 
// the last location where memory was allocated; allocating 
// additional memory is done by moving the stack pointer. When a 
// function is invoked, a new stack frame is created for the
//  function’s data. Local variables are stored on the stack, 
//  along with parameters passed into a function. Each new variable 
//  moves the stack pointer by the size of the value. When a 
//  function exits, its return values are copied back to the 
//  calling function via the stack and the stack pointer is moved
//  back to the beginning of the stack frame for the exited function,
//  deallocating all of the stack memory that was used by that 
//  function’s local variables and parameters.

// NOTE
// Go is unusual in that it can actually increase the size of a 
// stack while the program is running. This is possible because 
// each goroutine has its own stack and goroutines are managed by 
// the Go runtime, not by the underlying operating system 
// (we discuss goroutines when we talk about concurrency in 
// Chapter 10). This has advantages (Go stacks start small 
// and use less memory) and disadvantages (when the stack needs to 
// grow, all of the data on the stack needs to be copied, which 
// is slow). It’s also possible to write worst-case scenario code 
// that causes the stack to grow and shrink over and over.

//To store something on the stack, you have to know exactly how big it is at compile time. When you look at the value types in Go (primitive values, arrays, and structs), they all have one thing in common: we know exactly how much memory they take at compile time. This is why the size is considered part of the type for an array. Because their sizes are known, they can be allocated on the stack instead of the heap. The size of a pointer type is also known, and it is also stored on the stack.

// The rules are more complicated when it comes to the data that 
// the pointer points to. In order for Go to allocate the data the
// pointer points to on the stack, several conditions must be true.
// It must be a local variable whose data size is known at compile 
// time. The pointer cannot be returned from the function. If the 
// pointer is passed into a function, the compiler must be able to 
// ensure that these conditions still hold. If the size isn’t known, 
// you can’t make space for it by simply moving the stack pointer. 
// If the pointer variable is returned, the memory that the pointer 
// points to will no longer be valid when the function exits. When 
// the compiler determines that the data can’t be stored on the 
// stack, we say that the data the pointer points to escapes the 
// stack and the compiler stores the data on the heap.


// HEAP

// The heap is the memory managed by the garbage collector (or
// manually, in the case of C and C++). Any data that's stored 
// on the heap is valid as long as it can be tracked back to a
// pointer type variabel on the stack. Once there are no more
// pointers pointing to that data (or to data that points to 
// that data), the data becomes garbage and is cleared out by the
// garbage collector.

// NOTE
// A common source of bugs in C programs is returning a pointer to a
// local variable. In C, this results in a pointer pointing to invalid
// memory. The Go compiler is smarter. When it sees that a pointer to
// a local variable is returned, the local variable’s value is stored
// on the heap.



// The escape analysis done by the Go compiler isn’t perfect. 
// There are some cases where data that could be stored on the stack 
// escapes to the heap. However, the compiler has to be conservative;
//  it can’t take the chance of leaving a value on the stack when it 
//  might need to be on the heap because leaving a reference to invalid
//   data causes memory corruption. Newer Go releases improve escape 
//   analysis.

// You might be wondering: what’s so bad about storing things on the 
// heap? There are two problems related to performance. First is that 
// the garbage collector takes time to do its work. It isn’t trivial to 
// keep track of all of the available chunks of free memory on the heap 
// or tracking which used blocks of memory still have valid pointers. 
// This is time that’s taken away from doing the processing that your 
// program is written to do. Many garbage collection algorithms have 
// been written, and they can be placed into two rough categories: 
// those that are designed for higher throughput (find the most garbage 
// possible in a single scan) or lower latency (finish the garbage 
// scan as quickly as possible). Jeff Dean, the genius behind 
// many of Google’s engineering successes, co-wrote a paper in 
// 2013 called The Tail at Scale. It argues that systems should be 
// optimized for latency, to keep response times low. The garbage 
// collector used by the Go runtime favors low latency. Each garbage 
// collection cycle is designed to take less than 500 microseconds. 
// However, if your Go program creates lots of garbage, then the 
// garbage collector won’t be able to find all of the garbage during a 
// cycle, slowing down the collector and increasing memory usage.

// NOTE
// If you are interested in the implementation details, you may want 
// to listen to the talk Rick Hudson gave at the International 
// Symposium on Memory Management in 2018, describing the history 
// and implementation of the Go garbage collector.

// The second problem deals with the nature of computer hardware. 
// RAM might mean “random access memory,” but the fastest way to 
// read from memory is to read it sequentially. A slice of structs 
// in Go has all of the data laid out sequentially in memory. This 
// makes it fast to load and fast to process. A slice of pointers to 
// structs (or structs whose fields are pointers) has its data 
// scattered across RAM, making it far slower to read and process. 
// Forrest Smith wrote an in-depth blog post that explores how much 
// this can affect performance. His numbers indicate that it’s roughly 
// two orders of magnitude slower to access data via pointers that are 
// stored randomly in RAM.

// This approach of writing software that’s aware of the hardware 
// it’s running on is called mechanical sympathy. The term comes 
// from the world of car racing, where the idea is that a driver who 
// understands what the car is doing can best squeeze the last bits 
// of performance out of it. In 2011, Martin Thompson began applying 
// the term to software development. Following best practices in Go 
// gives it to you automatically.

// Compare Go’s approach to Java’s. In Java, local variables and 
// parameters are stored in the stack, just like Go. However, as 
// we discussed earlier, objects in Java are implemented as pointers. 
// That means for every object variable instance, only the pointer to 
// it is allocated on the stack; the data within the object is 
// allocated on the heap. Only primitive values (numbers, booleans, 
// and chars) are stored entirely on the stack. This means that 
// the garbage collector in Java has to do a great deal of work. 
// It also means that things like Lists in Java are actually a 
// pointer to an array of pointers. Even though it looks like a 
// linear data structure, reading it actually involves bouncing 
// through memory, which is highly inefficient. There are similar 
// behaviors in Python, Ruby, and JavaScript. To work around all 
// of this inefficiency, the Java Virtual Machine includes some 
// very clever garbage collectors that do lots of work, some 
// optimized for throughput, some for latency, and all with 
// configuration settings to tune them for the best performance. 
// The virtual machines for Python, Ruby, and JavaScript are less 
// optimized and their performance suffers accordingly.

// Now you can see why Go encourages you to use pointers sparingly. 
// We reduce the workload of the garbage collector by making sure 
// that as much as possible is stored on the stack. Slices of 
// structs or primitive types have their data lined up 
// sequentially in memory for rapid access. And when the 
// garbage collector does do work, it is optimized to return 
// quickly rather than gather the most garbage. The key to 
// making this approach work is to simply create less garbage in 
// the first place. While focusing on optimizing memory allocations 
// can feel like premature optimization, the idiomatic approach in 
// Go is also the most efficient.