package crawVirusshare

import (
	"log"
	"runtime"
	"time"
	"sync"
)
var wg sync.WaitGroup

func main() {
	start := time.Now()
	log.Println("Begin program: ", start)

	session, err := connectMongoDB()
	if err != nil {
		panic(err)
	}

	data = session.DB(config.Database).C(config.Collection)
	defer session.Close()

	linkMd5 := getLinkMd5()
	lenLinkMd5 := len(linkMd5)

	runtime.GOMAXPROCS(runtime.NumCPU())

	subWorker := lenLinkMd5 / config.Worker

	if subWorker > 1 {
		for worker := 0; worker < config.Worker-1; worker++ {
			wg.Add(1)
			subLinkMd5 := linkMd5[worker*subWorker: (worker+1)*subWorker]
			go importLinkMd5(subLinkMd5)
		}

		if lenLinkMd5 > (config.Worker-1)*subWorker {
			wg.Add(1)
			subLinkMd5 := linkMd5[(config.Worker-1)*subWorker:]
			go importLinkMd5(subLinkMd5)
		}
		wg.Wait()
	}

	log.Printf("Duration: %s\n", time.Since(start))
}
