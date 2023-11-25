package main

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

func ReadConfig(filename string) (Config, error) {
    var config Config

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return config, err
    }

    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return config, err
    }

    return config, nil
}
