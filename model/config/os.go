package config

import "net"

type (
	OS struct {
		Type              string           `yaml:"ostype"` //env:"OS_TYPE"
		CPU               string           `yaml:"cpu"`    //env:"CPU_TYPE"
		SelectedInterFAce int8             `yaml:"SelectedInterFAce"`
		Gateway           net.IP           `yaml:"netgateway"` // env:"{Name}_GATEWAY"
		Interfaces        map[int]*Network `yaml:"interfaces"`
	}

	Network struct {
		IsOnline bool      `yaml:"netonline"`
		IsDHCP   bool      `yaml:"netisdhcp"` // env:"{Name}_DHCP"
		Ip       net.IPNet `yaml:"netip"`     // env:"{Name}_IP"
		Name     string    `yaml:"netname"`
	}
)
