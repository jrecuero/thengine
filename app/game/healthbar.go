package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type HealthBar struct {
	*widgets.Gauge
}

func NewHealthBar(name string, position *api.Point, size *api.Size, total int) *HealthBar {
	return &HealthBar{
		Gauge: widgets.NewGauge(name, position, size, &theStyleGreenOverBlack, total),
	}
}

func (g *HealthBar) UpdateStyle(completed int) {
	if completed < g.GetTotal()/2 {
		g.SetStyle(&theStyleRedOverBlack)
		//tools.Logger.WithField("module", "healthBar").
		//    WithField("method", "UpdateCanvas").
		//    Debugf("%s", tools.StyleToString(g.GetStyle()))
	} else {
		g.SetStyle(&theStyleGreenOverBlack)
	}
}
