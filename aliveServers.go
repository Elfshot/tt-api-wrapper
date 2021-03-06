package ttRequests

import (
	"errors"
	"time"

	gocron "github.com/go-co-op/gocron"
)

type AliveServer struct {
	Primary_server string
	Servers        [2]serverType
}

type serverType struct {
	Name      string
	Id        uint8
	CfxId     string
	DirectUrl string
	CfxUrl    string
	AliveUrl  string
	Alive     bool
}

var servers [2]serverType
var aliveToggle bool = false
var AliveServers AliveServer

func cronAlive() {
	cron := gocron.NewScheduler(time.UTC)

	cron.Every("45s").Do(checkServers)
	cron.StartBlocking()
}

func checkServers() {
	aliveDown := 0

	for i := len(&AliveServers.Servers) - 1; i >= 0; i-- {
		server := &AliveServers.Servers[i]
		aliveCheck(server)
		if server.Alive {
			AliveServers.Primary_server = server.AliveUrl
		} else {
			aliveDown++
		}
	}
	if aliveDown == len(&AliveServers.Servers) {
		AliveServers.Primary_server = ""
	}
	// fmt.Printf("AliveServers: %+v\n", AliveServers)
}

func aliveCheck(server *serverType) {
	first, err := GetNoParse_(server.DirectUrl + "/players.json")
	if err != nil || (first.StatusCode != 200 && first.StatusCode != 204) {
		second, err := GetNoParse_(server.CfxUrl + "/players.json")

		if err != nil || (second.StatusCode != 200 && second.StatusCode != 204) {
			server.AliveUrl = ""
			server.Alive = false
		} else {
			server.AliveUrl = server.CfxUrl
			server.Alive = true
		}
	} else {
		server.AliveUrl = server.DirectUrl
		server.Alive = true
	}
}

func PrimaryServerUrl() (string, error) {
	if AliveServers.Primary_server != "" {
		return AliveServers.Primary_server, nil
	} else {
		return "", errors.New("No online servers")
	}
}

func ServerSpec(id uint8) (*serverType, error) {
	if id < 1 || id > 2 {
		return nil, errors.New("Server ID is out of range")
	}
	return &AliveServers.Servers[id-1], nil
}

func ServerSpecAlive(id uint8) (*serverType, error) {
	server := &AliveServers.Servers[id-1]
	if server.Alive {
		return server, nil
	} else {
		return nil, errors.New("Requested server is not alive")
	}
}
