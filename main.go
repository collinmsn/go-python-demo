package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

func main() {
	err := python.Initialize()
	if err != nil {
		panic(err)
	}
	defer python.Finalize()

	demo := python.PyImport_ImportModule("demo")
	mdict := python.PyModule_GetDict(demo)
	clsFoo := python.PyDict_GetItemString(mdict, "Foo")
	fmt.Println(clsFoo.Type().String())

	args := python.PyList_New(0)
	err = python.PyList_Append(args, python.PyInt_FromLong(1))
	if err != nil {
		panic(err)
	}
	python.PyList_Append(args, python.PyString_FromString("2"))

	instFoo := python.PyInstance_New(clsFoo, python.PyList_AsTuple(args), nil)
	fmt.Println(instFoo.Type().String())

	resp := instFoo.CallMethodObjArgs(python.PyString_FromString("invoke").String(), python.PyInt_FromLong(100))
	fmt.Printf("resp: %d\n", python.PyInt_AsLong(resp))
}
