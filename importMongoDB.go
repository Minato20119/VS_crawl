package crawVirusshare

import (
	"gopkg.in/mgo.v2"
	"log"
)

var data *mgo.Collection

type Virusshare struct {
	Type   string
	Value  string
	Source string
}

func importLinkMd5(linkMd5 []string) {
	defer wg.Done()

	for lineLinkMd5 := range linkMd5{
		log.Println(linkMd5[lineLinkMd5])

		md5 := getMd5(linkMd5[lineLinkMd5])
		lenMd5 := len(md5)

		subWorker := lenMd5 / config.Worker

		if subWorker > 1 {
			for worker := 0; worker < config.Worker-1; worker++ {
				wg.Add(1)
				subLineMd5 := linkMd5[worker*subWorker: (worker+1)*subWorker]
				go insertMd5ToDB(subLineMd5)
			}

			if lenMd5 > (config.Worker-1)*subWorker {
				wg.Add(1)
				subLineMd5 := linkMd5[(config.Worker-1)*subWorker:]
				go insertMd5ToDB(subLineMd5)
			}
			wg.Wait()
		}
	}
}

func insertMd5ToDB(md5 []string) {
	defer wg.Done()
	for lineMd5 := range md5 {
		err := data.Insert(Virusshare{"md5", md5[lineMd5], "virusshare.com"})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
