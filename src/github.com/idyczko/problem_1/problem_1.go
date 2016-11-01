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

func vertex(v int, n []int, channels []chan Message) (int, []int){
  fmt.Println("Starting node: ", v);
  parent:=-1;
  children := []int{};
  var unexplored Slice = n;
  if(v==root){
    parent = v;
    if(len(unexplored)>0){
      val:=unexplored[0];
      unexplored=append(unexplored[:0],unexplored[1:]...);
      channels[val]<-Message{v, "M"};
      fmt.Println(v," sent M to ", val);
    }else{
      if(parent!=v){
        channels[parent]<-Message{v, "parent"};
        fmt.Println(v," sent parent to ", parent);
      }
      fmt.Println("Exiting from ", v,"Parent: ", parent,"Children: ", children);
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
          if(len(unexplored)>0){
            val:=unexplored[0];
            unexplored=append(unexplored[:0],unexplored[1:]...);
            channels[val]<-Message{v, "M"};
            fmt.Println(v," sent M to ", val);
          }else{
            if(parent!=v){
              channels[parent]<-Message{v, "parent"};
              fmt.Println(v," sent parent to ", parent);
            }
            fmt.Println("Exiting from ", v,"Parent: ", parent,"Children: ", children);
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
        if(len(unexplored)>0){
          val:=unexplored[0];
          unexplored=append(unexplored[:0],unexplored[1:]...);
          channels[val]<-Message{v, "M"};
          fmt.Println(v," sent M to ", val);
        }else{
          if(parent!=v){
            channels[parent]<-Message{v, "parent"};
            fmt.Println(v," sent parent to ", parent);
          }
          fmt.Println("Exiting from ", v,"Parent: ", parent,"Children: ", children);
          return parent, children;
        }
      case "already":
        if(len(unexplored)>0){
          val:=unexplored[0];
          unexplored=append(unexplored[:0],unexplored[1:]...);
          channels[val]<-Message{v, "M"};
          fmt.Println(v," sent M to ", val);
        }else{
          if(parent!=v){
            channels[parent]<-Message{v, "parent"};
            fmt.Println(v," sent parent to ", parent);
          }
          fmt.Println("Exiting from ", v,"Parent: ", parent,"Children: ", children);
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
