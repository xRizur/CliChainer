package main

type Command struct {
    Name   string   `yaml:"command"`
    Args   []string `yaml:"args"`
}

type AliasConfig struct {
    Commands []Command `yaml:"commands"`
}

type Config struct {
    Variables map[string]map[string]string `yaml:"variables"`
    Aliases   map[string]AliasConfig       `yaml:"cli_aliases"`
}
