package main

import "fmt"

type Slice []int;

const root = 3;

type Message struct{
  sender int
  message string
}

func (slice Slice) indexOf(x int) (int){
  for i,v:=range slice{
    if(v==x){
      return i;
    }
  }
  return -1;
}

func explore(v, parent int, children Slice, unexplored *Slice, channels []chan Message) bool{
  if(len((*unexplored))>0){
    val:=(*unexplored)[0];
    (*unexplored)=append((*unexplored)[:0],(*unexplored)[1:]...);
    channels[val]<-Message{v, "M"};
    fmt.Println(v," sent M to ", val);
    return false;
  }else{
    if(parent!=v){
      channels[parent]<-Message{v, "parent"};
      fmt.Println(v," sent parent to ", parent);
    }
    fmt.Println("Exiting from ", v,"Parent: ", parent,"Children: ", children);
    return true;
  }
}

func vertex(v int, n []int, channels []chan Message) (int, []int){
  fmt.Println("Starting node: ", v);
  parent:=-1;
  var children Slice;
  var unexplored Slice = n;

  if(v==root){
    parent = v;
    if(explore(v, parent, children, &unexplored, channels)){
      return parent, children;
    }
  }
  for{
    select{
    case x:=<-channels[v]:
      fmt.Println(v," got ",x.message, " from ", x.sender);
      switch x.message{
      case "M":
        if(parent==-1){
          parent=x.sender;
          val:=unexplored.indexOf(x.sender);
          unexplored = append(unexplored[:val],unexplored[val+1:]...)
          if(explore(v, parent, children, &unexplored, channels)){
            return parent, children;
          }
        }else{
          channels[x.sender]<-Message{v, "already"};
          fmt.Println(v," sent already to ", x.sender);
          val:=unexplored.indexOf(x.sender);
          unexplored = append(unexplored[:val],unexplored[val+1:]...)
        }
      case "parent":
        children = append(children, x.sender);
        if(explore(v, parent, children, &unexplored, channels)){
          return parent, children;
        }
      case "already":
        if(explore(v, parent, children, &unexplored, channels)){
          return parent, children;
        }
      }
    }
  }
}

func main() {
    n := [][]int{
      {1,2},
      {0},
      {0,3,4},
      {2},
      {2},
    }
    var channels []chan Message;
    for i:=0;i<len(n);i++{
      channels = append(channels, make(chan Message));
    }

    for i, v := range n{
      go vertex(i, v, channels);
    }

    for{
    }
}
