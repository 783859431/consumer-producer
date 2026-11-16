package main
import (
	"fmt"
    "time"
    "sync")
func producer(prod chan<-int ,id int ){
  for i:=0 ; i<10;i++{
	fmt.Printf("proid:%d,val:%d\n", id,i)
    prod <- i
   // 
  }
  close(prod)
}
func consumer(con <-chan int ,id int, wg *sync.WaitGroup){

   for i:=range con{ // if con is closed , for loop will  exit, if con is empty , for loop will be blocked ; 

     fmt.Printf("conid:%d,val:%d\n", id,i)
     time.Sleep(1*time.Second)
   }
  wg.Done();
}
func main(){
  ch := make(chan int,10)
  var wg sync.WaitGroup;
  wg.Add(1);
  go consumer(ch,1,&wg)
  go producer(ch,1)
  wg.Wait()
}

