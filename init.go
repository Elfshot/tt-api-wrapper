package ttRequests

func Init() {
	initAlive()
}

func initAlive() {
	if aliveToggle == true {
		return
	}
	aliveToggle = true
	servers = [2]serverType{
		{
			Name:      "NY-1",
			Id:        1,
			CfxId:     "2epova",
			DirectUrl: "http://v1.api.tycoon.community/main",
			CfxUrl:    "https://tycoon-2epova.users.cfx.re/status",
			AliveUrl:  "",
			Alive:     false,
		},
		{
			Name:      "NY-2",
			Id:        2,
			CfxId:     "njyvop",
			DirectUrl: "http://v1.api.tycoon.community/beta",
			CfxUrl:    "https://tycoon-njyvop.users.cfx.re/status",
			AliveUrl:  "",
			Alive:     false,
		},
	}
	// AliveServers.Primary_server = servers[0].AliveUrl
	AliveServers.Primary_server = ""
	AliveServers.Servers = servers
	checkServers()
	// cron.Every("5s").Do(checkServers)
	go cronAlive()
}
