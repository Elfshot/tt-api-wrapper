package ttRequests_models

// Dunno how I'd implement the mixed json array so...
//
//[string, uint32 uint32]
//
//[Name, Src, VrpId]
type WidgetPlayers struct {
	Server  widgetPlayersServer  `json:"server"`
	Players widgetPlayersPlayers `json:"players"`
}

type widgetPlayersServer struct {
	Dxp    widgetPlayersDxp `json:"dxp"`
	Beta   string           `json:"beta"`
	Limit  uint16           `json:"limit"`
	Name   string           `json:"name"`
	Number string           `json:"number"`
	Motd   string           `json:"motd"`
	Uptime string           `json:"uptime"`
	Region string           `json:"region"`
}

// [bool, string, uint32, uint32, uint32]
//
// [IsDxpActive, HostName, time_remaining, AdditionalTime?, TimePast?]
type widgetPlayersDxp []interface{}

type widgetPlayersPlayers []widgetPlayersPlayer

// [string, uint16 uint32, string, bool, string, bool]
//
// [Name, Src, VrpId, DiscordAvatar, IsStaff, Job, IsDonor]
type widgetPlayersPlayer []interface{}
