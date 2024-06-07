package asciiprogress

import (
	"fmt"
	"math"
	"strings"
)

type ProgressBar struct {
	Name      string
	Total     int
	Length    int
	iteration int
	Style     int
}

type ProgressBarStyle int

const (
	CLASSIC ProgressBarStyle = iota
)

func DefineProgress(name string, total int, length int, style int) ProgressBar {
	return ProgressBar{
		Name:      name,
		Total:     total,
		Length:    length,
		Style:     style,
		iteration: 0,
	}
}

func (pb *ProgressBar) Iterate() {
	pb.iteration++

	if pb.iteration == pb.Total {
		pb.iteration = pb.Total
	}
}

func (pb *ProgressBar) GetPercentage() float64 {
	return ((float64(pb.iteration) / float64(pb.Total)) * 100)
}

func (pb *ProgressBar) Print() {
	blank := "-"
	bar, tip := "", ""
	progress := math.Floor((pb.GetPercentage() / 100) * float64(pb.Length))

	switch pb.Style {
	case 0:
		// [===>------] 30.0%
		bar, tip = "=", ">"
	case 1:
		// [▇▇▇-------] 30.0%
		bar, tip = "▇", "▇"
	case 2:
		// [▇▇▇       ] 30.0%
		blank = " "
		bar, tip = "▇", "▇"
	}

	fmt.Printf(
		"\r%s: [%s%s] %.2f%%",
		pb.Name,
		strings.Repeat(bar, int(progress))+tip,
		strings.Repeat(blank, pb.Length-int(progress)),
		pb.GetPercentage(),
	)
}

func (pb *ProgressBar) Clear() {
	pb = nil
	fmt.Println()
}
