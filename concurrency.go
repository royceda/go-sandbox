package main

import ("fmt"
  "time"
  "math/rand")


func handler(n int){
  for i :=0; i<10; i++{
    fmt.Println("Thread :", n, "\t i : ", i);
    amt := time.Duration(rand.Intn(250));
    time.Sleep(time.Millisecond * amt);
  }
}


//with channel direction : ping is only allowed to send c
func ping(c chan<- string){
  for i := 0; ; i++{
    c <- "ping"
  }
}

//without channel direction : pong i allowed to send/receive c
func pong(c chan string){
  for i := 0; ; i++{
    c <- "pong"
  }
}


//with channel direction : printer is only allowed to receive c
func printer(c <-chan string){
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
  for i := 0; i<100; i++{
    go handler(i);
  }


  // sync goroutines (thread)
  //channels : provide a way for a thread to communicate with another one
  /*var c chan string = make(chan string);
  go ping(c)
  go pong(c)
  go wait(c)
  go printer(c)
*/

  //channels : select, its like a switch for channels
  c1 := make(chan string);
  c2 := make(chan string);

  go func(){
    for{
      c1 <- "from 1";
      time.Sleep(time.Second*3)
    }
  }()


  go func(){
    for{
      c2 <- "from 2";
      time.Sleep(time.Second*2);
    }
  }()

  go func(){
    for{
      select{
      case msg1 := <- c1:
        fmt.Println(msg1);
      case msg2 := <- c2:
        fmt.Println(msg2);
      }
    }
  }()


  //Buffered channels
  c := make(chan string, 16);



  var input string;
  fmt.Scanln(&input);
}
