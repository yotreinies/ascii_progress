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

func (pb *ProgressBar) iterate() {
	pb.iteration++

	if pb.iteration == pb.Total {
		pb.iteration = pb.Total
	}
}

func (pb *ProgressBar) GetPercentage() float64 {
	return ((float64(pb.iteration) / float64(pb.Total)) * 100)
}

func (pb *ProgressBar) Print() {
	progress := math.Floor((pb.GetPercentage() / 100) * float64(pb.Length))

	switch pb.Style {
	case 0:
		// [===>------] 30.0%
		fmt.Printf(
			"\r[%s%s] %.2f",
			strings.Repeat("=", int(progress)-1)+">",
			strings.Repeat("-", pb.Length),
			pb.GetPercentage(),
		)
	}
}

func (pb *ProgressBar) Clear() {
	pb = nil
}
