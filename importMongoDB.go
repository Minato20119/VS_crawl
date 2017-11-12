package crawVirusshare

import (
	"fmt"
	"runtime"
	"sync"
	"gopkg.in/mgo.v2"
)

type Virusshare struct {
	Type   string
	Value  string
	Source string
}

var wg sync.WaitGroup
var countLink = 0

func importDB(data *mgo.Collection, md5 []string, start int, end int) {

	for lineMd5 := start; lineMd5 < end; lineMd5++ {
		err := data.Insert(&Virusshare{"md5", md5[lineMd5], "virusshare.com"})
		if err != nil {
			fmt.Println("Error Insert Into Database: ", err)
		}
	}
}

func allLinkMd5(data *mgo.Collection, linkMd5 []string, start int, end int) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for lineLinkMd5 := start; lineLinkMd5 < end; lineLinkMd5++ {
		md5 := getMd5(linkMd5[lineLinkMd5])
		lenMd5 := len(md5)
		fmt.Println(countLink)
		countLink++
		fmt.Println(linkMd5[lineLinkMd5])

		workers := 20
		temp := lenMd5 / workers
		wg.Add(20)

		for worker := 0; worker < 19; worker++ {
			go importDB(data, md5, worker*temp, (worker+1)*temp)
		}
		importDB(data, md5, 19*temp, lenMd5)
	}
}
