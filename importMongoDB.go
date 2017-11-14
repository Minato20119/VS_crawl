package crawVirusshare

import (
	"log"
	"sync"
	"gopkg.in/mgo.v2"
)

var data *mgo.Collection

type Virusshare struct {
	Type   string
	Value  string
	Source string
}

var wg sync.WaitGroup

func importDB(linkMd5 []string, start int, end int) {

	for lineLinkMd5 := start; lineLinkMd5 < end; lineLinkMd5++ {
		md5 := getMd5(linkMd5[lineLinkMd5])
		lenMd5 := len(md5)

		log.Println(linkMd5[lineLinkMd5])

		for lineMd5 := 0; lineMd5 < lenMd5; lineMd5++ {
			err := data.Insert(Virusshare{"md5", md5[lineMd5], "virusshare.com"})
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
