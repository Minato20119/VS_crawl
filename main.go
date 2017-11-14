package crawVirusshare

import (
	"log"
	"runtime"
	"time"
	"gopkg.in/mgo.v2"
)

func main() {
	start := time.Now()
	log.Println(start)
	var session *mgo.Session
	data, session = connectMongoDB()

	defer session.Close()

	linkMd5 := getLinkMd5()
	lenLinkMd5 := len(linkMd5)

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(config.Worker)

	temp := lenLinkMd5 / config.Worker

	// Vi link hash (_00000.md5 - temp) co 130k line, nen cho chay sau cung de no co the chay xong cac goroutine khac
	// Tu link hash 149-300 co 65k line
	for worker := 1; worker < config.Worker-1; worker++ {
		go importDB(linkMd5, worker*temp, (worker+1)*temp)
	}
	go importDB(linkMd5, (config.Worker-1)*temp, lenLinkMd5)
	// chay link 00000.md5 - temp
	importDB(linkMd5, 0, temp)

	log.Printf("Duration: %s\n", time.Since(start))
}
