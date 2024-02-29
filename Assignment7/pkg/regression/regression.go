package regression

import (
	"image/color"

	"example.com/assignment3/pkg/dataloading"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotRegression(name string, data []dataloading.DataPoint) error {
	p := plot.New()

	p.Title.Text = "Linear Regression for " + name
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	pts := make(plotter.XYs, len(data))
	for i := range data {
		pts[i].X = data[i].X
		pts[i].Y = data[i].Y
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	xs := make([]float64, len(pts))
	ys := make([]float64, len(pts))
	for i, pt := range pts {
		xs[i] = pt.X
		ys[i] = pt.Y
	}

	alpha, beta := stat.LinearRegression(xs, ys, nil, false)

	line := plotter.NewFunction(func(x float64) float64 {
		return beta*x + alpha
	})
	line.Color = color.RGBA{B: 255, A: 255}
	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, name+"_regression.png"); err != nil {
		return err
	}

	return nil
}
