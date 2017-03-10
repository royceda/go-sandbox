package main


import("fmt"; "container/list")


type Graph struct{
  weight [][]uint;
  length uint;
  adj []list.List;
}

/**
* Constructor
**/
func (g Graph) new(n uint) Graph{
  g.length = n;
  g.weight = make([][]uint, n, n);
  g.adj = make([]list.List, n);
  return g;
}




/**
* Dijskstra for a weighted 0-1 graph
**/
func (g Graph) dist01(source uint, target uint) ([]uint, []uint) {
  n  :=  g.length;

  grey := new(list.List);
  grey.PushFront(source)

  black := make([]bool, n);
  dist  := make([]uint, n);
  pred  := make([]uint, n);
  dist[source] = 0;

  for grey.Len() > 0{
    a := grey.Front();
    node := a.Value.(uint);
    black[node] = true;
    if(node == target){
      break;
    }

    for e := g.adj[node].Front(); e != nil; e = e.Next() {
      neighoor := e.Value.(uint);
      ell      := dist[node] + g.weight[node][neighoor];

      if black[neighoor] || dist[neighoor] <= ell {
        continue
      }
      dist[neighoor] = ell;
      pred[neighoor] = node;

      if g.weight[node][neighoor] == 0 {
        grey.PushFront(neighoor);
        } else {
          grey.PushBack(neighoor);
        }
      }
    }
    return dist, pred;
  }



  func (g Graph) Dijskstra(source uint, target uint) ([]uint, []uint) {
    n     :=  g.length;
    grey  := new(list.List);
    black := make([]bool, n);
    dist  := make([]uint, n);
    pred  := make([]uint, n);
    grey.PushFront(source)
    dist[source] = 0;


    return dist, pred;
  }


  func main(){
    fmt.Println("ok")
  }
