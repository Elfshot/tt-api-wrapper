package ttRequests

import (
	"fmt"

	models "github.com/Elfshot/tt-api-wrapper/models"
)

func primaryGet(append string, struc any) (*any, error) {
	server, servErr := PrimaryServerUrl()
	if servErr != nil {
		return nil, servErr
	}

	_, err := Get_(server+append, struc)
	if err != nil {
		return nil, err
	}

	return &struc, nil
}

func serverGet(serverId uint8, append string, struc any) (*any, error) {
	server, servErr := ServerSpecAlive(serverId)
	if servErr != nil {
		return nil, servErr
	}

	_, err := Get_(server.AliveUrl+append, struc)
	if err != nil {
		return nil, err
	}

	return &struc, nil
}

func Get_DataAdv(playerId uint32) (data *models.UserData, Error error) {
	var res models.UserData
	_, err := primaryGet(fmt.Sprintf("/dataadv/%d", playerId), &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_Players(serverId uint8) (data *models.Players, Error error) {
	var res models.Players
	_, err := serverGet(serverId, "/players.json", &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_WidgetPlayers(serverId uint8) (data *models.WidgetPlayers, Error error) {
	var res models.WidgetPlayers
	_, err := serverGet(serverId, "/widget/players.json", &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

// func Get_FxPlayers() (data *models.FxPlayers, Error error) {
// 	var res models.FxPlayers
// 	_, err := primaryGet("/players.json", &res)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &res, nil
// }

func Get_Sotd() (data *models.Sotd, Error error) {
	var res models.Sotd
	_, err := primaryGet("/sotd.json", &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
