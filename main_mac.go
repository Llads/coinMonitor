package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/martinlindhe/notify"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	color.Yellow.Println("           _                          _ _             \n  ___ ___ (_)_ __   /\\/\\   ___  _ __ (_) |_ ___  _ __ \n / __/ _ \\| | '_ \\ /    \\ / _ \\| '_ \\| | __/ _ \\| '__|\n| (_| (_) | | | | / /\\/\\ \\ (_) | | | | | || (_) | |   \n \\___\\___/|_|_| |_\\/    \\/\\___/|_| |_|_|\\__\\___/|_|   \n                                                      ")
	var coinNameInput string
	var coinValueInput float64
	color.Yellow.Println("Enter coin name(Not ticker):")
	fmt.Scanln(&coinNameInput)
	color.Yellow.Println("Enter coin alert price:")
	fmt.Scanln(&coinValueInput)
	fmt.Println("Monitoring...")

	var coinName = strings.ToLower(coinNameInput)
	var coinValue = coinValueInput

	for {
		req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/simple/price?ids="+coinName+"&vs_currencies=usd", nil)
		req.Header.Set("accept", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		var jsonResp map[string]map[string]float64
		json.Unmarshal(body, &jsonResp)

		var face = jsonResp[coinName]["usd"]
		
		if coinValue < face {
			coinString := fmt.Sprint(" ", face)
			path, _ := os.Getwd()
			notify.Notify("coinMonitor", " ", coinName+" has reached"+"\n$"+coinString, path + "/alert.png")
			os.Exit(1)
		}

		time.Sleep(1000*time.Millisecond)
	}
}

