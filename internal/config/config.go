package config

import "time"

var (
	Host             = ":8080"
	MaxInputRequests = 100
	ParallelRequests = 4
	RequestTimeout   = time.Second
	MaxUrls          = 20
)
