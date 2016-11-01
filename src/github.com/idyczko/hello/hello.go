package main

import "fmt"

func change(tab *[]int){
  (*tab)=append((*tab), 120);
}

func change2(tab []int){
  tab=append(tab,120);
}

func main() {
    tab := []int{12,10,23,42};
    change(&tab);
    fmt.Println(tab)
}
