package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func InitConfig() Conf {
	var c Conf
	if c.isGetConfigFromEnv() {
		c.getConfigFromEnv()
	} else {
		c.getConf()
	}
	return c
}

type db struct {
	Postgres postgres `yaml:"postgres"`
}

type postgres struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type server struct {
	Port   string `yaml:"port"`
	Domain string `yaml:"domain"`
}

type Conf struct {
	Db     db     `yaml:"db"`
	Server server `yaml:"server"`
	Deploy string `yaml:"deploy"`
}

func (c *Conf) getConf() *Conf {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	yamlFile, err := ioutil.ReadFile(dir + "/config-dev.yml")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return c
}

func (c *Conf) isGetConfigFromEnv() bool {
	deploy := os.Getenv("deploy")
	if deploy == "k8s" {
		return true
	}
	return false
}

func (c *Conf) getConfigFromEnv() *Conf {
	deploy := os.Getenv("deploy")

	postgresHost := os.Getenv("db.postgres.host")
	if postgresHost == "" {
		log.Fatal("PG host is empty")
	}
	postgresUsername := os.Getenv("db.postgres.username")
	if postgresUsername == "" {
		log.Fatal("PG username is empty")
	}
	postgresPort := os.Getenv("db.postgres.port")
	if postgresPort == "" {
		log.Fatal("PG port is empty")
	}
	postgresPassword := os.Getenv("db.postgres.password")
	if postgresPassword == "" {
		log.Fatal("PG password is empty")
	}
	postgresDbname := os.Getenv("db.postgres.dbname")
	if postgresDbname == "" {
		log.Fatal("PG dbname is empty")
	}
	serverPort := os.Getenv("server.port")
	if serverPort == "" {
		log.Fatal("Server port is empty")
	}
	serverDomain := os.Getenv("server.domain")
	if serverDomain == "" {
		log.Fatal("Server domain is empty")
	}

	postgres := postgres{
		Host:     postgresHost,
		Username: postgresUsername,
		Port:     postgresPort,
		Password: postgresPassword,
		DBName:   postgresDbname}

	db := db{Postgres: postgres}

	server := server{
		Port:   serverPort,
		Domain: serverDomain}

	c.Server = server
	c.Db = db
	c.Deploy = deploy

	return c
}
