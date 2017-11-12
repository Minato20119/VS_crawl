package crawVirusshare

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {
	start := time.Now()
	fmt.Println("Begin Program: ", time.Since(start))

	data, session := connectMongoDB()
	defer session.Close()

	linkMd5 := getLinkMd5()
	lenLinkMd5 := len(linkMd5)

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(25)

	var linkMd5Second []string
	var linkMd5First []string

	if lenLinkMd5 > 148 {
		// Hash 65k line
		linkMd5Second = linkMd5[149:]
		// Hash 130k line
		linkMd5First = linkMd5[0:149]

		workers := 10
		lenLinkMd5Second := len(linkMd5Second)
		temp2 := lenLinkMd5Second / workers

		for worker := 0; worker < workers - 1; worker++ {
			go allLinkMd5(data, linkMd5Second, worker*temp2, (worker+1)*temp2)
		}
		go allLinkMd5(data, linkMd5Second, (workers - 1) *temp2, lenLinkMd5Second)

		workers = 15
		lenLinkMd5First := len(linkMd5First)
		temp1 := lenLinkMd5First / workers

		for worker := 0; worker < workers - 1; worker++ {
			go allLinkMd5(data, linkMd5First, worker*temp1, (worker+1)*temp1)
		}
		allLinkMd5(data, linkMd5First, (workers - 1)*temp1, lenLinkMd5First)
	}

	fmt.Printf("Duration: %s\n", time.Since(start))
}