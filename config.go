package main

type configT struct {
	SidecarPort int
	Smtp struct {
		Host     string
		Port     int
		User     string
		Password string
	}
}

var config *configT
