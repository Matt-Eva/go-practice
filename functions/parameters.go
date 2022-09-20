package main

import "fmt"

func main(){
	myFunc(MyFuncParams {ParamOne: "Corp"}) // 

	myVariadicFunc("string", 1, []int{1,2,3})

	anySlice := []any{"adsf", 12, []int{6, 3}} // slice with the any datatype
	// fmt.Println(anySlice)
	myVariadicFunc(anySlice...)
	myVariadicFunc([]any{3,5,1}...)
}


// NAMED AND OPTIONAL PARAMETERS
// Go does not allow for named or optional parameters.
// You must supply all the parameters for a function.
// to emulate named or optional parameters, define a struct 
// whose fields match the desired parameters, as pass the struc
// as the argument to your function.

// Ex:

type MyFuncParams struct{
	ParamOne string
	ParamTwo string
	ParamThree int
}

func myFunc(parameter MyFuncParams) { // notice that the "type" here is MyFuncParams
	fmt.Println(parameter)
}

// Note that this doesn't assign default values - whatever
// fields are left out in the struct declaration simply have
// that types zero value. Generally, if you feel that you 
// need optional or named inputs, your function is too complicated.


// VARIADIC PARAMETERS

// Variadic parameters can be declared to allow
// a function to receive any number of inputs as arguments.
// the fmt.Println package uses variadic parameters to accomplish
// this exact purpose.

func myVariadicFunc(values ...any){ // note that the any keyword allows you to use any datatype in go.
	fmt.Println(values)
}

// We still have to give variadic parameters a name. The 
// name in the example above is "values". It can have any
// datatype.

// The variable created for the param within the funciton
// is a slice of the specified type.

// Because the variadic param is input as a slice, you can 
// pass your arguments as a slice. However, when passing
// the slice argument - either by slice literal or variable name - 
// we must follow it with three dots ...