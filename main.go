package main

import (
	"fmt"

	requests "github.com/Elfshot/tt-api-wrapper/requests"
)

func errNull(err error) {
	if err != nil {
		fmt.Printf("There was an error: %s\n", err.Error())
	}
}

func main() {
	sotd, err := requests.Get_Sotd()
	errNull(err)
	data, err := requests.Get_DataAdv(59504)
	errNull(err)
	players, err := requests.Get_Players()
	errNull(err)
	fxplayers, err := requests.Get_FxPlayers()
	errNull(err)

	fmt.Printf("%+v\n", (*sotd))
	fmt.Printf("%+v\n", data.Data.Gaptitudes_v)
	fmt.Printf("%+v\n", players.Players[0])
	fmt.Printf("%+v\n", (*fxplayers)[0])
}
