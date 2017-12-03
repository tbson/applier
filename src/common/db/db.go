package db

import (
	"log"
	"database/sql"
	_"github.com/lib/pq"
	"gopkg.in/yaml.v2"
    "io/ioutil"
)

var Db *sql.DB;

type DbInfo struct {
	Adapter string `yaml:adapter`
    Host string `yaml:host`
    Name string `yaml:name`
    User string `yaml:user`
    Pass string `yaml:pass`
    Port int `yaml:port`
    Charset string `yaml:charset`
}

type Conf struct {
	Environments struct {
		Production DbInfo `yaml:Production`
		Development DbInfo `yaml:development`
		Testing DbInfo `yaml:testing`
	}
}

func (conf *Conf) getConf() *Conf {

    yamlFile, err := ioutil.ReadFile("../migration/phinx.yml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, conf)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return conf
}

func init() {
	var err error = nil
	var conf Conf
	conf.getConf()
	user := conf.Environments.Development.User
	pass := conf.Environments.Development.Pass
	host := conf.Environments.Development.Host
	name := conf.Environments.Development.Name
	connStr := "postgres://" + user + ":" + pass + "@" + host + "/" + name + "?sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
