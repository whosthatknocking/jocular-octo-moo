package main

import (
	"encoding/json"
	"os"
	"time"
)

func NewMonitor(app string) (*Monitor, error) {
	var err error
	mon := &Monitor{
		Timestamp:   time.Now().Unix(),
		Application: app,
	}
	mon.Dimensions.Host, err = os.Hostname()
	if err != nil {
		return nil, err
	}
	return mon, nil
}

func NewPingMonitor(latency int64, status uint8) (*PingMonitor, error) {
	mon, err := NewMonitor("docker.ping")
	if err != nil {
		return nil, err
	}

	return &PingMonitor{
		Monitor: *mon,
		Metrics: Ping{
			Latency: latency,
			Status:  status,
		},
	}, nil
}

// Export images and containers. If used `Driver` is devicemapper export
// `DriverStatus` storage metrics.
func NewInfoMonitor(data []byte) (*InfoMonitor, error) {
	mon, err := NewMonitor("docker.info")
	if err != nil {
		return nil, err
	}

	raw := &InfoRaw{}
	err = json.Unmarshal(data, raw)
	if err != nil {
		return nil, err
	}

	info := &InfoMonitor{
		Monitor: *mon,
		Metrics: Info{
			Images:     raw.Images,
			Containers: raw.Containers,
		},
	}

	if raw.Driver == "devicemapper" && raw.DriverStatus != nil {
		for _, pair := range raw.DriverStatus {
			switch pair[0] {
			case "Data Space Used":
				info.Metrics.DataUsed, err = FromHumanSize(pair[1])
				if err != nil {
					return nil, err
				}
			case "Data Space Total":
				info.Metrics.DataTotal, err = FromHumanSize(pair[1])
				if err != nil {
					return nil, err
				}
			case "Data Space Available":
				info.Metrics.DataAvailable, err = FromHumanSize(pair[1])
				if err != nil {
					return nil, err
				}
			case "Metadata Space Used":
				info.Metrics.MetadataUsed, err = FromHumanSize(pair[1])
				if err != nil {
					return nil, err
				}
			case "Metadata Space Total":
				info.Metrics.MetadataTotal, err = FromHumanSize(pair[1])
				if err != nil {
					return nil, err
				}
			case "Metadata Space Available":
				info.Metrics.MetadataAvailable, err = FromHumanSize(pair[1])
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return info, nil
}
