package main

import (
	"encoding/json"
	"github.com/JedBeom/airq"
	"io/ioutil"
)

type Config struct {
	AirqKey string `json:"airq_key"`
}

func loadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var c Config
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}

	err = airq.SetKey(c.AirqKey)
	if err != nil {
		panic(err)
	}
}
