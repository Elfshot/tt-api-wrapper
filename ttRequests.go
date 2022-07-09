package tt_requests

import (
	"fmt"

	models "github.com/Elfshot/tt-api-wrapper/models"
)

func Get_DataAdv(playerId uint32) (data *models.UserData, Error error) {
	var res models.UserData
	_, err := _Get(fmt.Sprintf("https://tycoon-2epova.users.cfx.re/status/dataadv/%d", playerId), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_Players() (data *models.Players, Error error) {
	var res models.Players
	_, err := _Get("https://tycoon-2epova.users.cfx.re/status/players.json", &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_FxPlayers() (data *models.FxPlayers, Error error) {
	var res models.FxPlayers
	_, err := _Get("https://tycoon-2epova.users.cfx.re/players.json", &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Get_Sotd() (data *models.Sotd, Error error) {
	var res models.Sotd
	_, err := _Get("https://tycoon-2epova.users.cfx.re/status/sotd.json", &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
