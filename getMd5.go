package crawVirusshare

import (
	"log"
	"regexp"
)

var RegexMd5 = regexp.MustCompile("[0-9|a-z]{32}")

//
//const url = "https://virusshare.com/hashes/VirusShare_00279.md5"
func getMd5(url string) []string {
	source, err := getSourceUrl(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	md5 := RegexMd5.FindAllString(source, -1)
	return md5
}
