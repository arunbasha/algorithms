//The checkpoint synchronization is a problem of synchronizing multiple tasks. Consider a workshop where several workers
// assembling details of some mechanism. When each of them completes his work, they put the details together. There is
// no store, so a worker who finished its part first must wait for others before starting another one. Putting details
// together is the checkpoint at which tasks synchronize themselves before going their paths apart
package Algorithm

import (
	"fmt"
	"sync"
	"time"
)

func worker(w string, wg *sync.WaitGroup) {
	fmt.Println(w, "Started work")
	time.Sleep(time.Second/100)
	fmt.Println(w, "Finished work")
	wg.Done()
}
func checkPoint(workers []string, numAssemblies int) bool{
	var wg sync.WaitGroup
	counter := 0
	for i:=1; i<= numAssemblies; i++ {
		wg.Add(len(workers))
		for _, w := range workers {
			go worker(w, &wg)
		}
		wg.Wait()
		fmt.Printf("%d. Assembled after work\n", i)
		counter++
	}
	if counter == numAssemblies {
		return true
	}
	return false
}