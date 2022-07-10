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

func Get_DataAdv(playerId uint32) (data *models.UserData, Error error) {
	var res models.UserData
	_, err := primaryGet(fmt.Sprintf("/status/dataadv/%d", playerId), &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_Players() (data *models.Players, Error error) {
	var res models.Players
	_, err := primaryGet("/status/players.json", &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_FxPlayers() (data *models.FxPlayers, Error error) {
	var res models.FxPlayers
	_, err := primaryGet("/players.json", &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_Sotd() (data *models.Sotd, Error error) {
	var res models.Sotd
	_, err := primaryGet("/status/sotd.json", &res)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
