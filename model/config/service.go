package config

import "net"

type (
	UI struct {
		Port    int    `yaml:"port"`    //env:"UI_Port"
		UTC     string `yaml:"UTC"`     //env:"UI_UTC"
		NTP     net.IP `yaml:"NTP"`     //env:"UI_NTP"
		NTPPort int    `yaml:"NTPport"` //env:"UI_Port"
	}

	Service struct {
		Ip   net.IP `yaml:"serip"`   // env:"{Name}_serip"
		Port int    `yaml:"serport"` // env:"{Name}_serport"
		Type string `yaml:"sertype"` // env:"{Name}_sertype"
		Ac   string `yaml:"serac"`   // env:"{Name}_serac"
		PW   string `yaml:"serpw"`   // env:"{Name}_serpw"
		//Name string `yaml:"sername"`
	}
)
