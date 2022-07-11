package ttRequests

import (
	"errors"
	"reflect"

	models "github.com/Elfshot/tt-api-wrapper/models"
)

func GetTotalPlayers() ([]models.BaseTotalPlayer, error) {
	var newPlayers []models.BaseTotalPlayer

	for i := uint8(1); i <= uint8(len(servers)); i++ {
		getFailCount := 0
		res, err := Get_WidgetPlayers(i)
		if err != nil {
			getFailCount++
			if getFailCount >= len(servers) {
				return nil, errors.New("All get_players failed in GetTotalPlayers")
			}
			continue
		}
		players := res.Players

		for ii := 0; ii < len(players); ii++ {
			player := players[ii]
			vrp := player[2].(float64)
			if reflect.TypeOf(player[3]).Kind() == reflect.Bool {
				player[3] = ""
			}

			newPlayer := models.BaseTotalPlayer{
				Name:       player[0].(string),
				VrpId:      uint32(vrp),
				DiscordUrl: player[3].(string),
				IsStaff:    player[4].(bool),
				IsDonor:    player[6].(bool),
				Job:        player[5].(string),
				ServerId:   i,
			}
			newPlayers = append(newPlayers, newPlayer)
		}
	}
	return newPlayers, nil
}
