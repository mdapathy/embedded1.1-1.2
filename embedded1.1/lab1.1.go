package main

import (
	"fmt"
	"github.com/Equanox/gotron"
	"github.com/wcharczuk/go-chart"
	"log"
	"math"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var write strings.Builder

func main() {
	//Variant specific data
	n := 6
	w := 1500
	N := 1024


	var expect float64 // math expectation
	var dispersion float64
	var sum float64

	var x = make([]float64, N, N)
	var y = make([]float64, N, N)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Generating signal

	for i := 0; i < N; i++ {

		var ytemp float64
		for j := 0; j < n; j++ {
			ytemp += r.Float64() *
				(math.Sin(float64((w/n)*(j+1)*i) + r.Float64()))
		}
		sum += ytemp

		y[i] = ytemp
		x[i] = float64(i)

	}

	expect = sum / float64(N)

	for i := 0; i < len(y); i++ {
		dispersion += math.Pow(y[i]-expect, 2)
	}

	dispersion = dispersion / float64(N-1)

	writeToFile(expect, dispersion, N)

	graph := chart.Chart{
		Width:  N * 2,
		Height: 500,
		XAxis: chart.XAxis{
			Name: "Time",
		},
		YAxis: chart.YAxis{
			Name: "X value",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()

	graph.Render(chart.PNG, f)

	pwd, _ := os.Getwd()
	pathToDir := filepath.Dir(pwd)
	// Create a new window instance

	window, err := gotron.New(path.Join(pathToDir, "embedded1.1", "html"))
	if err != nil {
		panic(err)
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Lab1.1"

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Wait for the application to close
	<-done

}

func writeToFile(expect, dispersion float64, N int) {

	var outfile *os.File
	var err error

	if outfile, err = os.Create("info.txt"); err != nil {
		log.Fatalf("Error in the go routine : %s", err)
	}

	defer outfile.Close()

	write.WriteString(fmt.Sprintf("N: %d", N))
	write.WriteString(fmt.Sprintf("\nExpectation: %f", expect))
	write.WriteString(fmt.Sprintf("\nDispersion:%f", dispersion))

	if _, err := outfile.WriteString(write.String()); err != nil {
		log.Fatalf("Error in the go routine : %s", err)
	}

}
