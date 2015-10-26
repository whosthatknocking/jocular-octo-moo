package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jawher/mow.cli"
)

var (
	defaultTimeout = 30
)

func defineCommands(parent *cli.Cmd) {

	parent.Command("ping", "Get docker ping response time", func(cmd *cli.Cmd) {
		cmd.Spec = "[--socket][--timeout]"

		endpoint := cmd.String(cli.StringOpt{
			Name:   "socket",
			Desc:   "Daemon socket to connect to",
			Value:  "unix:///var/run/docker.sock",
			EnvVar: "DOCKER_HOST",
		})

		timeout := cmd.Int(cli.IntOpt{
			Name:   "timeout",
			Desc:   "Connect timeout (seconds)",
			Value:  defaultTimeout,
			EnvVar: "",
		})

		cmd.Action = func() {
			start := time.Now()

			// client call
			client, err := NewClient(*endpoint, nil, time.Duration(*timeout))
			if err != nil {
				log.Fatalf("Failed to init new client, %v", err)
			}

			resp, err := client.Ping()
			if err != nil {
				log.Fatalf("Ping call failed, %v", err)
			}
			var status uint8 = 0
			if string(resp) == "OK" {
				status = 1
			}

			mm, err := NewPingMonitor(time.Since(start).Nanoseconds(), status)
			if err != nil {
				log.Fatalf("Failed to create new monitor, %v", err)
			}

			mmj, err := json.Marshal(mm)
			if err != nil {
				log.Fatalf("Failed to, %v", err)
			}

			fmt.Println(string(mmj))
			os.Exit(0)
		}
	})

	parent.Command("info", "Get docker info", func(cmd *cli.Cmd) {
		cmd.Spec = "[--socket][--timeout]"

		endpoint := cmd.String(cli.StringOpt{
			Name:   "socket",
			Desc:   "Daemon socket to connect to",
			Value:  "unix:///var/run/docker.sock",
			EnvVar: "DOCKER_HOST",
		})

		timeout := cmd.Int(cli.IntOpt{
			Name:   "timeout",
			Desc:   "Connect timeout (seconds)",
			Value:  defaultTimeout,
			EnvVar: "",
		})

		cmd.Action = func() {
			client, err := NewClient(*endpoint, nil, time.Duration(*timeout))
			if err != nil {
				log.Fatalf("Failed to init new client, %v", err)
			}

			resp, err := client.Info()
			if err != nil {
				log.Fatalf("Info failed, %v", err)
			}

			mm, err := NewInfoMonitor(resp)
			if err != nil {
				log.Fatalf("Failed to create new monitor, %v", err)
			}

			mmj, err := json.Marshal(mm)
			if err != nil {
				log.Fatalf("Failed to, %v", err)
			}

			fmt.Println(string(mmj))
			os.Exit(0)
		}
	})

	parent.Command("version", "Show version information", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			fmt.Printf("Version:\t%v\n", Version)
			fmt.Printf("API version:\t%.2f\n", APIVersion)
			fmt.Printf("Git commit:\t%v\n", BuildCommit)
			fmt.Printf("Built:\t\t%v\n", BuildDate)
			os.Exit(0)
		}
	})
}
