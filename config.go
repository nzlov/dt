package main

import "github.com/nzlov/utils"

type Config struct {
	utils.BaseConfig

	Size  int    `json:"size" yaml:"size" mapstructure:"size"`
	Bg    int    `json:"bg" yaml:"bg" mapstructure:"bg"`
	Img   string `json:"img" yaml:"img" mapstructure:"img"`
	Tmpls []Tmpl `json:"tmpls" yaml:"tmpls" mapstructure:"tmpls"`
}

type Tmpl struct {
	T string `json:"t" yaml:"t" mapstructure:"t"`
	O string `json:"o" yaml:"o" mapstructure:"o"`
	E string `json:"e" yaml:"e" mapstructure:"e"`
}
