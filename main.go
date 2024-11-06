package main

import (
	"fmt"
	"image"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/nzlov/utils"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("$HOME/.config/dt")
	cfg, err := utils.LoadConfig[Config]()
	if err != nil {
		panic(err)
	}

	i, err := loadImage(cfg.Img)
	if err != nil {
		panic(err)
	}

	size := 3
	if cfg.Size != 0 {
		size = cfg.Size
	}
	bg := 2
	if cfg.Bg != 0 {
		bg = cfg.Bg
	}
	if bg >= size {
		bg = size - 1
	}

	vs, err := prominentcolor.KmeansWithAll(size, i, prominentcolor.ArgumentNoCropping, uint(i.Bounds().Dx()), nil)
	if err != nil {
		panic(err)
	}

	m := map[string]string{}

	for i, v := range vs {
		m[fmt.Sprintf("C%d", i)] = v.AsString()
		m[fmt.Sprintf("RGB%d", i)] = fmt.Sprintf("%d,%d,%d", v.Color.R, v.Color.G, v.Color.B)
		m[fmt.Sprintf("OR%d", i)] = intToRGBHex(RGBToInt(v) | RGBToInt(vs[bg]))
		m[fmt.Sprintf("AND%d", i)] = intToRGBHex(RGBToInt(v) & RGBToInt(vs[bg]))
	}

	fmt.Println(m)

	for _, c := range cfg.Tmpls {
		if err := execute(c, m); err != nil {
			panic(err)
		}
	}
}

func execute(c Tmpl, m map[string]string) error {
	slog.Default().Info("Render:" + c.O)
	td, err := os.ReadFile(c.T)
	if err != nil {
		return err
	}
	tmpl := template.Must(template.New(c.T).Parse(string(td)))

	f, err := os.Create(c.O)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tmpl.Execute(f, &m); err != nil {
		return err
	}
	e := strings.TrimSpace(c.E)
	if e != "" {
		es := strings.Split(e, " ")
		return exec.Command(es[0], es[1:]...).Run()
	}
	return nil
}

func RGBToInt(c prominentcolor.ColorItem) int {
	return int(c.Color.R)<<16 | int(c.Color.G)<<8 | int(c.Color.B)
}

func intToRGBHex(color int) string {
	r := (color >> 16) & 0xff
	g := (color >> 8) & 0xff
	b := color & 0xff

	return fmt.Sprintf("%.2x%.2x%.2x", r, g, b)
}

func loadImage(url string) (image.Image, error) {
	if strings.HasPrefix(url, "http") {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		img, _, err := image.Decode(resp.Body)
		return img, err
	}
	f, err := os.Open(url)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}
