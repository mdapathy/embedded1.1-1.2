package main

import (
	"fmt"
	"github.com/Equanox/gotron"
	"github.com/wcharczuk/go-chart"
	"math"
	"math/rand"
	"os"
	"path"
	"time"
)

func main() {
	//Variant specific data

	n := 6
	w := 1500
	N := 1024

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var y = make([]float64, N, N)
	var y2 = make([]float64, N, N)
	var x = make([]float64, N/2, N/2)

	var rval = make([]float64, N/2, N/2)
	var rval2 = make([]float64, N/2, N/2)

	var sum float64
	var sum2 float64

	//Generating first signal

	for i := 0; i < N; i++ {

		var ytemp float64
		for j := 0; j < n; j++ {
			ytemp += r.Float64() *
				(math.Sin(float64((w/n)*(j+1)*i) + r.Float64()))

		}
		y[i] = ytemp

		if i < N/2 {
			sum += ytemp
		}

	}

	//Generating second signal

	for i := 0; i < N; i++ {

		var ytemp2 float64
		for j := 0; j < n; j++ {

			ytemp2 += r.Float64() *
				(math.Sin(float64((w/n)*(j+1)*i) + r.Float64()))
		}
		y2[i] = ytemp2

		if i < N/2 {
			sum2 += ytemp2
		}

	}

	//Calculating math expectation for both signals

	expect := sum / float64(N/2)
	expect2 := sum2 / float64(N/2)

	var cor float64
	var cor2 float64

	for tau := 0; tau < N/2; tau++ {
		for t := 0; t < N/2; t++ {
			rval[t] = (y[t] - expect) * (y[t+tau] - expect)
			rval2[t] = (y[t] - expect) * (y2[t+tau] - expect2)

			cor += rval[t]
			cor2 += rval2[t]

			x[t] = float64(t)
		}
	}
	cor = cor / float64(N-1)
	cor2 = cor2 / float64(N-1)

	graph := chart.Chart{
		Width:  N * 2,
		Height: 400,
		XAxis: chart.XAxis{
			Name: "N",
		},
		YAxis: chart.YAxis{
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: rval,
			},
		},
	}

	graph2 := chart.Chart{
		Width:  N * 2,
		Height: 400,
		XAxis: chart.XAxis{
			Name: "N",
		},
		YAxis: chart.YAxis{
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: rval2,
			},
		},
	}

	f, _ := os.Create("output1.png")
	defer f.Close()
	
	if err := graph.Render(chart.PNG, f); err != nil {
		panic(err)
	}

	f2, _ := os.Create("output2.png")
	defer f2.Close()

	if err := graph2.Render(chart.PNG, f2); err != nil {
		panic(err)
	}

	fmt.Println("Auto correlation", cor, "Correlation", cor2)

	pathToDir, err := os.Getwd()

	// Create a new browser window instance
	window, err := gotron.New(path.Join(pathToDir, "html"))
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Lab1"

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	<-done

}
