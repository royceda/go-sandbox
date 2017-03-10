package main

import ("fmt"
  "time"
  "math/rand")


func handler(n int){
  for i :=0; i<10; i++{
    fmt.Println("Thread :", n, " i : ", i);
    amt := time.Duration(rand.Intn(250));
    time.Sleep(time.Millisecond * amt);
  }
}

func ping(c chan string){
  for i := 0; ; i++{
    c <- "ping"
  }
}

func pong(c chan string){
  for i := 0; ; i++{
    c <- "pong"
  }
}



func printer(c chan string){
  for{
    msg := <- c;
    fmt.Println(msg);
    time.Sleep(time.Second * 1);
  }
}


func wait(c chan string){
  for i := 0; ; i++{
    c <- "wait"
  }
}


func main(){
  // async goroutines (thread)
  for i := 0; i<10; i++{
    go handler(i);
  }


  // sync goroutines (thread)
  //channels
  var c chan string = make(chan string);
  go ping(c)
  go pong(c)
  go wait(c)
  go printer(c)

  var input string;
  fmt.Scanln(&input);
}
