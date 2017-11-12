package crawVirusshare

import (
	"fmt"
	"regexp"
)

var RegexMd5 = regexp.MustCompile("[0-9|a-z]{32}")

//
//const url = "https://virusshare.com/hashes/VirusShare_00279.md5"
func getMd5(url string) []string {
	source, err := getSourceUrl(url)
	if err != nil {
		fmt.Println("Error get Source Url from getMd5: ", err)
	}
	md5 := RegexMd5.FindAllString(source, -1)

	return md5
}
