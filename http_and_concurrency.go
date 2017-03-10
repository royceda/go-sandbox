package main


/* An sample which check who is the biggest website among a list.
*  It's a simple use case in the using of channel and go routine
*/



import(
  "fmt"
  "io/ioutil"
  "net/http"
)


type HomePageSize struct {
  URL string
  Size int
}



func handler(url string, results chan HomePageSize){
  res, err := http.Get(url);
  if err != nil {
    panic(err);
  }
  defer res.Body.Close();

  bs, err := ioutil.ReadAll(res.Body);
  if err != nil{
    panic(err);
  }
  results <- HomePageSize{ URL: url, Size: len(bs)}
}

func main(){
 urls := []string{
    "http://www.google.fr",
    "http://www.amazon.com",
    "http://www.apple.com",
    "http://www.google.fr",
  }

  results := make(chan HomePageSize);

  for _,url := range urls{
    go handler(url, results);
  }

  var biggest HomePageSize;
  for range urls {
    result := <- results;
    if result.Size > biggest.Size{
      biggest = result;
    }
  }

  fmt.Println("The biggest is ", biggest.URL);
}
