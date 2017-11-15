package crawVirusshare

import (
	"gopkg.in/mgo.v2"
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
	Source     string `toml:"Source"`
	Worker     int    `toml:"Worker"`
}

func LoadFileConfig(path string) Configuration {
	var config Configuration
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}
	return config
}

func connectMongoDB() *mgo.Session {
	infoMongoDb := &mgo.DialInfo{
		Addrs:    []string{config.HostName + ":" + config.PortName},
		Timeout:  60 * time.Second,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
		Source:   config.Source,
	}

	session, err := mgo.DialWithInfo(infoMongoDb)
	if err != nil {
		log.Println("Error connect database...")
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}
