package main

import "log"

type configuration struct {
	port    int
	env     string
	version string
}

type application struct {
	config configuration
	logger *log.Logger
}
