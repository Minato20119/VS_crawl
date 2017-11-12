package crawVirusshare

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"time"
	"toml"
)

var config = LoadFileConfig("./crawlVirusshareConfig.toml")

type Configuration struct {
	HostName   string `toml:"HostName"`
	PortName   string `toml:"PortName"`
	Username   string `toml:"Username"`
	Password   string `toml:"Password"`
	Database   string `toml:"Database"`
	Collection string `toml:"Collection"`
}

func LoadFileConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	_, err = toml.Decode(string(file), &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}

func connectMongoDB() (*mgo.Collection, *mgo.Session){
	infoMongoDb := &mgo.DialInfo{
		Addrs:    []string{config.HostName + ":" + config.PortName},
		Timeout:  60 * time.Second,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
		Source:   "admin",
	}

	session, err := mgo.DialWithInfo(infoMongoDb)
	if err != nil {
		fmt.Println("Error connect database...")
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	data := session.DB(config.Database).C(config.Collection)
	return data, session
}
