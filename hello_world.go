package main
import "fmt"




var global string = "hello"

func f() (int, string){
  fmt.Println(global)
  return 3, "toto"
}


func fact(x uint) uint{
  if x == 0{
    return 1
  }else{
    return x * fact(x-1)
  }
}



func def(){
  fmt.Println("defer")
}



//structure and interface
type Circle struct{
  x, y, r float64;
}


func init() Circle {
  c := new(Circle);
  a := Circle{0, 1, 9};
  b := Circle(x: 0, y: 1, r:9)
  c := &Circle{0,0,9} //ptr of circle
}


// comment
func main(){
  fmt.Println("Hello world");
  fmt.Println("2+3 = ", 2+3.0);
  fmt.Println("len(toto) = ", len("toto"))

  //scanf
  var in float64;
  fmt.Println("Write a float...")
  fmt.Scanf("%f", &in)



  //Variables and types
  const x string = "toto";
  var a int = 12;
  variableName := "name"
  a++
  fmt.Println(x, " ", a)
  caca := "ouii"
  caca = "OK"
  fmt.Println(caca, " ", variableName)


  //control statement
  //For
  i := 1
  for i<=10{
    fmt.Println(i)
    i++;
  }

  //if, else and switch
  if i%2 == 0{
    fmt.Println("even", i)
  }else if i%3 == 0{
    fmt.Println("3-multiple")
  }else{
    fmt.Println("odd", i)
  }


  switch i {
    case 0: fmt.Println(i)
    case 1: fmt.Println(i)
    default: fmt.Println("Unknown Number")
  }

  //Arrays, slice and map
  var y [5]int;
  y[3] = 777;
  fmt.Println(y);

  for i := 0; i<len(y); i++{
    y[i] = i

  }
  fmt.Println(y);

  //slices == ArrayList (Java) or vector (c++)
  //var z []int;
  slice1 := []int{0,1,2,3}
  slice2 := append(slice1, 4,5,6)
  //slice := append(slice1, slice2)

  //z := make([]int, 7)
  fmt.Println(slice2)

  //map
  m := make(map[string]int)
  m["toto"] = 777
  m["H"] = 123
  fmt.Println(m)

  num, str := f();
  fmt.Println(num, str);


  //Private variadic parameter
  add := func(args ...int) int {
    sum := 0;
    for _, v := range args {
      sum += v;
    }
    return sum
  }

  sum := add(1, 2, 3, 65, 76);
  fmt.Println("sum : ", sum)
  fmt.Println("fact : ", fact(30))


  //pointer
  zero := func(ptr *int) {
    *ptr = 0;
  }

  ze := 4;
  zero(&ze)
  fmt.Println("ze ",ze, "address : ", &ze)

  ptr := new(string);
  *ptr = "hello world"
  fmt.Println("ptr : ", *ptr)


    //defer, panic, and recover
    defer def()

    defer func(){
      str := recover()
      fmt.Println(str)
    }()
    panic("PANIC")

}
