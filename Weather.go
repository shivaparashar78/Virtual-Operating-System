package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"net/http"
)

func WeatherApp() {
	a := app.New()

	w := a.NewWindow("Weather App")
	w.Resize(fyne.NewSize(400 , 500))

	//API Part

	res , err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=ee57e340e20c6252600d725cef70262c")
	if err != nil{	
		fmt.Print(err);
	}

	defer res.Body.Close()

	body , err := ioutil.ReadAll(res.Body)
	if err != nil{	
		fmt.Print(err);
	}

	weather , err := UnmarshalWelcome(body)
	if err != nil{
		fmt.Print(err)
	}

	img := canvas.NewImageFromFile("icon\\weatherhome.png")
	img.FillMode = canvas.ImageFillOriginal

	label1:= canvas.NewText("Weather Details:", color.White) 
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label2:= canvas.NewText(fmt.Sprintf("Country: \t%s" , weather.Sys.Country), color.White)
	label3:= canvas.NewText(fmt.Sprintf("City: \t%s" , weather.Name), color.White) 
	label4:= canvas.NewText(fmt.Sprintf("Wind Speed: \t%.2f" , weather.Wind.Speed), color.White) 
	label5:= canvas.NewText(fmt.Sprintf("Temperature: \t%.2f" , weather.Main.Temp), color.White) 
	label6:= canvas.NewText(fmt.Sprintf("Humidity: \t%.2d" , weather.Main.Humidity), color.White) 

	weatherContainer := container.NewVBox(
			img,
			label1,
			label2,
			label3,
			label4,
			label5,
			label6,
	)
    r, _ := fyne.LoadResourceFromPath("icon\\Weather.png")
	w.SetIcon(r)

	w.SetContent(container.NewBorder(nil,nil,nil,nil,weatherContainer),)
	
	w.Show()
}


func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}
