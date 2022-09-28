package main

import "fmt"

func main(){

}

// MAPS AND POINTERS

//As seen previously, any modifications made to a map when it's
// passed to a function are reflected in the original variable that
// was passed in. This occurs because the Go runtime, a map is
// implemented as a pointer to a struct. Passing a map to a function
// means that you're copying a pointer.

// For this reason, it's best to avoid using maps for input parameters
// or return values, in particular when dealing with public APIs.

// Rationale:

// When thinking in terms of API design, maps are a bad choice 
// because they say nothing about the values they contain - there's
// nothing that explicitly defines what keys are in the map, so the
// only way to know what they are is to trace through the code.

// When considered immutability, maps are bad because the only way to
// know what values ended up in a map is to trace through all
// of the functions that interact with it, since each can modify it.

// Rather than using a map, use a struct.

// SLICES AND POINTERS

// Passing a slice to a function also has some complex behavior:
// Any modifiction to the contents of a slice is reflected in the
// original variable, but using append to change the length is NOT
// relfected in the original variable, even if the slice has a capacity
// greater than its length.

// This is due tot he fact that a slice is implemented as a struct
// with 3 fields: an int field for length, an int field for capacity
// and a pointer to a block of memory.

// When a slice is copied to a different variable or a function
// parameter, the copy is made of the length, capacity, and the pointer.

// Changing the values in the slice changes the memory
// that the pointer is pointing to, so the changes are reflected
// in both places.

// However, changes in length and capacity are not reflected in the
// original, because those changes are only in the copy.

// Changing the capacity just means the pointer is pointing to
// a bigger block of memory.

// Changing the length in the copy doesn't adjust the length of the
// original.

// Because of this phenomenon, a slice that's passed to a function
// can have its contents modified, but the slice itself can't be
// resized. Slices are frequently passed around in Go programs -
// you should default to assuming that a slice is not modified by 
// a function. If a function does modify a slice, your documentation
// should specify that it does.

// NOTE
// The reason you can pass a slice of any size to a function is that 
// the data that’s passed to the function is the same for any size 
// slice: two int values and a pointer. The reason that you can’t 
// write a function that takes an array of any size is because the 
// entire array is passed to the function, not just a pointer to the 
// data.

// SLICES AS BUFFERS


// When reading data from an external resource (like a file or a
// network connection), many languages use code like this:

// r = open_resource()
// while r.has_data() {
// 	data_chunk = r.next_chunk()
// 	process(data_chunk)
//   }
//   close(r)

//   The problem with this pattern is that every time we 
//   iterate through that while loop, we allocate another data_chunk 
//   even though each one is only used once. This creates lots of 
//   unnecessary memory allocations. Garbage-collected languages 
//   handle those allocations for you automatically, but the work 
//   still needs to be done to clean them up when you are done
//    processing.
  
//   Even though Go is a garbage-collected language, writing 
//   idiomatic Go means avoiding unneeded allocations. Rather 
//   than returning a new allocation each time we read from a 
//   data source, we create a slice of bytes once and use it as a 
//   buffer to read data from the data source:
  
//   file, err := os.Open(fileName)
//   if err != nil {
// 	  return err
//   }
//   defer file.Close()
//   data := make([]byte, 100)
//   for {
// 	  count, err := file.Read(data)
// 	  if err != nil {
// 		  return err
// 	  }
// 	  if count == 0 {
// 		  return nil
// 	  }
// 	  process(data[:count])
//   }

//   Remember that we can’t change the length or capacity 
//   of a slice when we pass it to a function, but we can 
//   change the contents up to the current length. In this code, 
//   we create a buffer of 100 bytes and each time through the 
//   loop, we copy the next block of bytes (up to 100) into the 
//   slice. We then pass the populated portion of the buffer to 
//   process. We’ll look at more details about I/O in “io and Friends”.