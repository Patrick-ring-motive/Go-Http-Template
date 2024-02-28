package submodules

import (
  "sync"
  . "unsafe"
  "fmt"
)

var initPromiseSync = Pass(sync.Mutex{})

// short example
 // i := GenericFunction(1)
 // prom := Go(AsyncGenericFunction,1);
 // x := Await(prom).(int);

func Collect(a ...Any)[]Any{
  return a
}

func Itemize(a Any)[]Any{
  return a.([]Any)
} 

var AsyncMap = ExtendedMap{Map: MakeMap()}

func AsyncRegister(fun Any, afun func(Any) Any) Any {
  AsyncMap.set(fun, afun)
  return afun
}


/******************Promise Structures*************************/
func GenericFunction(i int) int {
  fmt.Println("GenericFunction ")
  fmt.Println(i)
  return i+1
}

func AsyncGenericFunction(i Any) Any{
  return GenericFunction(i.(int))
}



type Promise struct {
  PromiseChannel chan (Any)
  Error          error
  Result         any
  Resolved       bool
  Rejected       bool
}

func Asunc(a Any) Any { return a }

func AsyncifyUnsafe(a Any) (func(Any)Any) {
  fn:=*(* func(Any) Any)(Pointer(&a))
  return fn
}

func Asyncify(afun Any) func(Any) Any {
  switch af := afun.(type) {
  case func(Any) Any:
    return af
  default:
    return AsyncifyUnsafe(afun)
  }
}



func Go(fn func(Any)Any, payload Any) Promise {
  promiseChannel := make(chan Any)
  promise := Promise{PromiseChannel: promiseChannel, Error: nil, Result: nil, Resolved: false, Rejected: false}
  go Async(fn, payload, promise)
  return promise
}

func Async(function func(Any) Any, payload Any, promise Promise) Promise {
  promise.Result = function(payload)
  promise.Resolved = true
  promise.PromiseChannel <- promise.Result
  return promise
}

func Await(promise Promise) (Any) {
  if promise.Resolved || promise.Rejected {
    return promise.Result
  } else {
    promise.Result = <-promise.PromiseChannel
    promise.Resolved = true
    close(promise.PromiseChannel)
    return promise.Result
  }
}

