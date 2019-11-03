package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/caarlos0/env"
)

var config *Settings

//Settings ...
type Settings struct {
	DBBizNome     string `json:"DBBizNome" env:"DB_BIZ_NOME"`
	DBBizHost     string `json:"DBBizHost" env:"DB_BIZ_HOST"`
	DBBizPorta    int    `json:"DBBizPorta" env:"DB_BIZ_PORTA"`
	DBBizUser     string `json:"DBBizUser" env:"DB_BIZ_USER"`
	DBBizPassword string `json:"DBBizPassword" env:"DB_BIZ_PASSWORD"`
	EnableLogFile bool   `json:"enableLogFile" env:"ENABLE_LOG_FILE"`
	LogFile       string `json:"logFile" env:"LOG_FILE"`
	RedisHost     string `json:"redisHost" env:"REDIS_HOST"`
	RedisSenha    string `json:"redisSenha" env:"REDIS_SENHA"`
	Port          int    `json:"port" env:"PORT"`
	AllowedParam  string `json:"allowedParam" env:"ALLOWED_PARAM"`
}

//NewConfig ...
func NewConfig(file string) *Settings {
	var erro error

	conf := &Settings{}

	if file != "" {

		bufConf, err := ioutil.ReadFile(file)
		if err == nil {
			erro = json.Unmarshal(bufConf, conf)
			if erro != nil {
				log.Println(erro)
			}
		}
	}

	if erro = env.Parse(conf); erro != nil {
		log.Println(erro)
	}

	return conf
}

//Config ...
func Config() *Settings {
	return config
}
