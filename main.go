package crawVirusshare

import (
	"gopkg.in/mgo.v2"
	"log"
	"runtime"
	"time"
	"sync"
)
var wg sync.WaitGroup

func main() {
	start := time.Now()
	log.Println("Begin program: ", start)
	session := connectMongoDB()
	data = session.DB(config.Database).C(config.Collection)
	defer session.Close()

	linkMd5 := getLinkMd5()
	lenLinkMd5 := len(linkMd5)

	runtime.GOMAXPROCS(runtime.NumCPU())

	temp := lenLinkMd5 / config.Worker

	for worker := 0; worker < config.Worker-1; worker++ {
		wg.Add(1)
		go importLinkMd5(linkMd5, worker*temp, (worker+1)*temp)
	}
	wg.Add(1)
	go importLinkMd5(linkMd5, (config.Worker-1)*temp, lenLinkMd5)
	wg.Wait()

	log.Printf("Duration: %s\n", time.Since(start))
}
