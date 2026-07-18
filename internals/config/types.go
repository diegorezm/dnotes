package config

type Config struct {
	Version int     `json:"version"`
	Editor  Editor  `json:"editor"`
	Cluster Cluster `json:"cluster"`
}

type Editor struct {
	Command string `json:"command"`
}

type Cluster struct {
	Active string `json:"active"`
}
