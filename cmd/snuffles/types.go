package main

type Monitor struct {
	Timestamp   int64  `json:"timestamp"`
	Application string `json:"application"`
	Dimensions  struct {
		Host string `json:"host"`
	} `json:"dimensions"`
}

type Ping struct {
	Latency int64 `json:"latency"`
	Status  uint8 `json:"status"`
}

type Info struct {
	Images     int `json:"images"`
	Containers int `json:"containers"`
	// storage driver status
	DataUsed          int64 `json:"data.used"`
	DataTotal         int64 `json:"data.total"`
	DataAvailable     int64 `json:"data.available"`
	MetadataUsed      int64 `json:"metadata.used"`
	MetadataTotal     int64 `json:"metadata.total"`
	MetadataAvailable int64 `json:"metadata.available"`
}

type PingMonitor struct {
	Monitor
	Metrics Ping `json:"metrics"`
}

type InfoMonitor struct {
	Monitor
	Metrics Info `json:"metrics"`
}

type InfoRaw struct {
	Containers   int         `json:"Containers"`
	Images       int         `json:"Images"`
	Driver       string      `json:"Driver"`
	DriverStatus [][2]string `json:"DriverStatus"`
}
