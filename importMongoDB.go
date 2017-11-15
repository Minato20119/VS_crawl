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

func importLinkMd5(linkMd5 []string, start int, end int) {
	defer wg.Done()
	for lineLinkMd5 := start; lineLinkMd5 < end; lineLinkMd5++ {
		log.Println(linkMd5[lineLinkMd5])

		md5 := getMd5(linkMd5[lineLinkMd5])
		lenMd5 := len(md5)
		temp := lenMd5 / config.Worker

		for worker := 0; worker < config.Worker-1; worker++ {
			wg.Add(1)
			go insertMd5ToDB(md5, worker*temp, (worker+1)*temp)
		}
		wg.Add(1)
		go insertMd5ToDB(md5, (config.Worker - 1)*temp, lenMd5)
	}
}

func insertMd5ToDB(md5 []string, start int, end int) {
	defer wg.Done()
	for lineMd5 := start; lineMd5 < end; lineMd5++ {
		err := data.Insert(Virusshare{"md5", md5[lineMd5], "virusshare.com"})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
