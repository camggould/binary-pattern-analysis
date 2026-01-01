package utils

import (
	"fmt"

	"github.com/NimbleMarkets/ntcharts/barchart"
	"github.com/charmbracelet/lipgloss"
)

func PPrintMap(input map[string]int) {
	var barData []barchart.BarData
	for label, count := range input {
		data := barchart.BarData{
			Label: label,
			Values: []barchart.BarValue{
				{
					Value:  float64(count),
					Style:  lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
				},
			},
		}
		barData = append(barData, data)
	}

	bc := barchart.New(len(input) * 2, 20)
	bc.PushAll(barData)
	bc.Draw()
	fmt.Println(bc.View())
}