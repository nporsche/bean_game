package storage

import (
	"model"
)

var Ele *Elements

type Elements struct {
	Players []model.Player
	Beans   []model.Bean
}

func init() {
	Ele = new(Elements)
	Ele.Players = make([]model.Player, 0)
	Ele.Beans = make([]model.Bean, 0)
}

func (this *Elements) PlayerReport(id uint64, longitude float64, latitude float64) {
	//TODO:
}

func (this *Elements) UpdateBean(id uint64, state uint8, longitude float64, latitude float64) {
	//TODO:
}
