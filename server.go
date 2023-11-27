package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	if c.Request().Method == "GET" {
		// 讀取檔案
		quarySize := c.QueryParam("size")
		if quarySize == "" {
			return c.String(200, "can't find size parameter")
		}

		file, err := os.Open("./generate/data/data_" + string(quarySize) + ".txt")
		if err != nil {
			fmt.Println("open file failed!, err:", err)
		}

		// 將檔案內容轉成json 格式
		var data []map[string]interface{}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			values := strings.Split(line, ",")
			x := values[0]
			y := values[1]
			// label := values[2]

			dataPoint := map[string]interface{}{
				"x": x,
				"y": y,
				// "label": label,
			}

			data = append(data, dataPoint)
		}

		// jsonData, err := json.MarshalIndent(data, "", "  ")
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("MarshalIndent failed!, err:", err)
		}

		return c.JSON(200, string(jsonData))
	}

	return c.String(200, "error")
}

func main() {
	e := echo.New()

	e.GET("/", home)

	e.Start(":8000")
}
