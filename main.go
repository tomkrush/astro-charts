package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Star struct {
	BV            string `json:"B-V"`
	Dm            string `json:"DM"`
	Dec           string `json:"Dec"`
	Dmag          string `json:"Dmag"`
	Fk5           string `json:"FK5"`
	Glat          string `json:"GLAT"`
	Glon          string `json:"GLON"`
	Hd            string `json:"HD"`
	Hr            string `json:"HR"`
	IRflag        string `json:"IRflag"`
	K             string `json:"K"`
	LuminosityCls string `json:"LuminosityCls"`
	MultCnt       string `json:"MultCnt"`
	Name          string `json:"Name"`
	NoteFlag      string `json:"NoteFlag"`
	Notes         []struct {
		Category string `json:"Category"`
		Remark   string `json:"Remark"`
	} `json:"Notes"`
	Parallax    string `json:"Parallax"`
	RI          string `json:"R-I"`
	Ra          string `json:"RA"`
	RadVel      string `json:"RadVel"`
	RotVel      string `json:"RotVel"`
	Sao         string `json:"SAO"`
	Sep         string `json:"Sep"`
	SpectralCls string `json:"SpectralCls"`
	UB          string `json:"U-B"`
	VarID       string `json:"VarID"`
	Vmag        string `json:"Vmag"`
	LRotVel     string `json:"l_RotVel"`
	NRadVel     string `json:"n_RadVel"`
	PmDE        string `json:"pmDE"`
	PmRA        string `json:"pmRA"`
}

func (s Star) Equatorial() EquatorialCoordinates {
	return EquatorialCoordinates{
		RightAscension: s.RightAscension(),
		Declination:    s.Declination(),
	}
}

func ParseHours(hours string) float64 {
	reg := regexp.MustCompile(`(\d+)h (\d+)m (\d+\.\d+)s`)
	matches := reg.FindStringSubmatch(hours)

	if len(matches) != 4 {
		return 0
	}

	h, _ := strconv.ParseFloat(matches[1], 64)
	m, _ := strconv.ParseFloat(matches[2], 64)
	s, _ := strconv.ParseFloat(matches[3], 64)

	return h + m/60 + s/3600
}

func ParseDegrees(degrees string) float64 {
	reg := regexp.MustCompile(`([+-])(\d+)° (\d+)′ (\d+)″`)
	matches := reg.FindStringSubmatch(degrees)

	if len(matches) != 5 {
		return 0
	}

	sign := 1.0
	if matches[1] == "-" {
		sign = -1.0
	}

	d, _ := strconv.ParseFloat(matches[2], 64)
	m, _ := strconv.ParseFloat(matches[3], 64)
	s, _ := strconv.ParseFloat(matches[4], 64)

	return sign * (d + m/60 + s/3600)
}

func (s Star) RightAscension() float64 {
	return ParseHours(s.Ra)
}

func (s Star) Declination() float64 {
	return ParseDegrees(s.Dec)
}

type EquatorialCoordinates struct {
	RightAscension float64
	Declination    float64
}

type HorizontalCoordinates struct {
	Azimuth  float64
	Altitude float64
}

type ObserverCoordinates struct {
	Latitude  float64
	Longitude float64
}

func EquatorialToHorizontal(e EquatorialCoordinates, observer ObserverCoordinates) HorizontalCoordinates {

	return HorizontalCoordinates{}
}

func main() {
	// Read the file
	file, err := os.Open("data/bsc5.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Decode the JSON
	decoder := json.NewDecoder(file)

	var stars []Star
	err = decoder.Decode(&stars)

	if err != nil {
		fmt.Println(err)
	}

	// Print the stars
	for _, star := range stars {
		e := star.Equatorial()

		fmt.Printf("RA: %.2f, DEC: %.2f, Vmag: %s\n", e.RightAscension, e.Declination, star.Vmag)
	}
}
