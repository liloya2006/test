package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const timeLayout = "15:04"

type weather struct {
	Sol        int       `json:"sol"`
	MinTemp    int       `json:"min_temp"`
	MaxTemp    int       `json:"max_temp"`
	SunsetStr  string    `json:"sunset"`
	Sunset     time.Time `json:"-"`
	SunriseStr string    `json:"sunrise"`
	Sunrise    time.Time `json:"-"`
	Pressure   int       `json:"pressure"`
}

func GetSolInfo(sol string) (weather, error) {
	url := fmt.Sprintf("https://api.maas2.apollorion.com/%v", sol)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var data weather
	if err != nil || resp.StatusCode == 404 || resp.StatusCode == 500 {
		log.Println(resp.Status)
		log.Fatal(err)
		return data, err
	}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return data, err
}

func GetData() map[int]weather {
	currentSol, err := GetSolInfo("")
	if err != nil {
		log.Fatal(err)
	}

	sol := currentSol.Sol
	startSol := sol - 6

	sols := make(map[int]weather)

	for i := startSol; i <= sol; i++ {
		previousSol, err := GetSolInfo(strconv.Itoa(i))
		if err == nil {
			sunset, err := time.Parse(timeLayout, previousSol.SunsetStr)
			if err != nil {
				log.Fatal(err)
			}
			previousSol.Sunset = sunset

			sunrise, err := time.Parse(timeLayout, previousSol.SunriseStr)
			if err != nil {
				log.Fatal(err)
			}
			previousSol.Sunrise = sunrise
			sols[i] = previousSol
		}
	}

	return sols
}

func OneDay(c *gin.Context) {
	sols := GetData()
	var arr []int

	for _, sol := range sols {
		arr = append(arr, sol.Sol)
	}

	var ranIndex = rand.Intn(len(arr))

	var ranDay = sols[arr[ranIndex]]

	//fmt.Println(fmt.Sprintf("Sol: %d, MinTemp: %d, MaxTemp: %d, Sunset: %s, Sunrise: %s, Pressure: %d",
	//	ranDay.Sol, ranDay.MinTemp, ranDay.MaxTemp, ranDay.SunsetStr, ranDay.SunriseStr, ranDay.Pressure))

	c.JSON(http.StatusOK, ranDay)
}

func Summary(c *gin.Context) {
	sols := GetData()

	c.JSON(http.StatusOK, sols)
}

func Average(c *gin.Context) {
	marsSols := GetData()
	var sumPressure int
	var sunSetMid int
	var sunRiseMid int
	for _, marsWeatherStruct := range marsSols {
		sumPressure = sumPressure + marsWeatherStruct.Pressure
		sunSetMid = sunSetMid + (marsWeatherStruct.Sunset.Hour()*60 + marsWeatherStruct.Sunset.Minute())
		sunRiseMid = sunRiseMid + (marsWeatherStruct.Sunrise.Hour()*60 + marsWeatherStruct.Sunrise.Minute())
	}

	avgPressure := sumPressure / len(marsSols)

	sunSetMidTime, _ := time.Parse(timeLayout, fmt.Sprintf("%d:%d", sunSetMid/len(marsSols)/60, sunSetMid/len(marsSols)%60))
	sunRiseMidTime, _ := time.Parse(timeLayout, fmt.Sprintf("%d:%d", sunRiseMid/len(marsSols)/60, sunRiseMid/len(marsSols)%60))

	c.JSON(http.StatusOK, gin.H{
		"avg_pressure":     avgPressure,
		"sunset_mid_time":  sunSetMidTime.Format(timeLayout),
		"sunrise_mid_time": sunRiseMidTime.Format(timeLayout),
	})
}

func main() {
	r := gin.Default()

	r.GET("/average", Average)
	r.GET("/summary", Summary)
	r.GET("/oneday", OneDay)

	r.Run()
}
