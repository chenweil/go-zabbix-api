package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	zabbix "github.com/tpretz/go-zabbix-api"
)

func main() {
	var (
		url      = flag.String("url", "http://localhost/api_jsonrpc.php", "Zabbix API URL")
		user     = flag.String("user", "Admin", "Zabbix username")
		password = flag.String("password", "zabbix", "Zabbix password")
		action   = flag.String("action", "version", "Action: version, hosts, groups")
	)
	flag.Parse()

	fmt.Printf("Connecting to Zabbix API: %s\n", *url)

	config := zabbix.Config{
		Url: *url,
	}
	api := zabbix.NewAPI(config)

	_, err := api.Login(*user, *password)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	fmt.Println("Login successful!")

	switch *action {
	case "version":
		version, err := api.Version()
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		fmt.Printf("Zabbix API Version: %s\n", version)

	case "hosts":
		hosts, err := api.HostsGet(zabbix.Params{})
		if err != nil {
			log.Fatalf("Failed to get hosts: %v", err)
		}
		fmt.Printf("Found %d hosts:\n", len(hosts))
		for _, h := range hosts {
			fmt.Printf("  - %s (ID: %s, Status: %s)\n", h.Host, h.HostID, h.Status)
		}

	case "groups":
		groups, err := api.HostGroupsGet(zabbix.Params{})
		if err != nil {
			log.Fatalf("Failed to get host groups: %v", err)
		}
		fmt.Printf("Found %d host groups:\n", len(groups))
		for _, g := range groups {
			fmt.Printf("  - %s (ID: %s)\n", g.Name, g.GroupID)
		}

	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: version, hosts, groups")
		os.Exit(1)
	}
}
