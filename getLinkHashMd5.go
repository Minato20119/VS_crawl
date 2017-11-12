package crawVirusshare

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const URL = "https://virusshare.com/hashes.4n6"

var RegexLinkMd5 = regexp.MustCompile("hashes/VirusShare_[0-9]{5}.md5")

// Get Content File
func getSourceUrl(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error getSourceUrl: ")
		return "", err
	}
	defer resp.Body.Close()

	sourcePage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error getSourceUrl: ")
		return "", err
	}
	return string(sourcePage), err
}

func getLinkMd5() []string {
	source, err := getSourceUrl(URL)
	if err != nil {
		fmt.Println(err)
	}
	linkMd5 := RegexLinkMd5.FindAllString(source, -1)

	for i := 0; i < len(linkMd5); i++ {
		linkMd5[i] = "https://virusshare.com/" + linkMd5[i]
	}

	return linkMd5
}
