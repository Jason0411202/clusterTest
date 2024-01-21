package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type DataPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func main() {
	// 讀取參數
	config, err := readConfig("../parameter.yaml")
	if err != nil {
		log.Fatal(err)
	}

	clusterNum := config.ClusterNum
	multiple := config.NowMultiple

	// 計時開始
	start := time.Now()

	// 發送 HTTP request
	url := fmt.Sprintf("http://localhost:8000?size=%d", multiple)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var Datas interface{} // 定義一個空的 interface{}，用來接回傳的 JSON 字串
		err := json.NewDecoder(response.Body).Decode(&Datas)
		if err != nil {
			log.Fatal(err)
		}

		// 此時 data 是 interface{} 型態，實際存的是字串
		myString, ok := Datas.(string) // 將 interface{} 轉換為字串
		if !ok {
			log.Fatal("convert to string failed")
		}

		// 解析 JSON 字串
		var rawPoints []map[string]interface{}
		err = json.Unmarshal([]byte(myString), &rawPoints)
		if err != nil {
			fmt.Println("解析 JSON 時發生錯誤:", err)
			return
		}

		// 將資料轉換為 kmeans 套件需要的格式
		var d clusters.Observations
		for _, point := range rawPoints {
			x_string, _ := point["x"].(string)
			x, _ := strconv.ParseFloat(x_string, 64)
			y_string, _ := point["y"].(string)
			y, _ := strconv.ParseFloat(y_string, 64)

			d = append(d, clusters.Coordinates{
				x,
				y,
			})
		}

		// 執行 k-means 演算法
		km := kmeans.New()
		clusters, err := km.Partition(d, clusterNum)

		// 計算每個群組的平均值和標準差
		for i, c := range clusters {
			pts := make(plotter.XYs, len(c.Observations))
			for j, obs := range c.Observations {
				pts[j].X = obs.Coordinates()[0]
				pts[j].Y = obs.Coordinates()[1]
			}

			meanX, meanY, stddevX, stddevY := calculateStatistics(pts)

			fmt.Printf("Group %d:\n", i+1)
			fmt.Printf("Mean X: %.2f, Mean Y: %.2f\n", meanX, meanY)
			fmt.Printf("Stddev X: %.2f, Stddev Y: %.2f\n\n", stddevX, stddevY)
		}

		// 繪圖
		p := plot.New()
		if err != nil {
			log.Fatal(err)
		}

		for i, c := range clusters {
			pts := make(plotter.XYs, len(c.Observations))
			for j, obs := range c.Observations {
				pts[j].X = obs.Coordinates()[0]
				pts[j].Y = obs.Coordinates()[1]
			}

			s, err := plotter.NewScatter(pts)
			if err != nil {
				log.Fatal(err)
			}
			s.GlyphStyle.Color = plotutil.Color(i) // 每個分群使用不同顏色

			p.Add(s)
		}

		// 新增圖片 label
		p.Title.Text = "K-Means Clustering"
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"

		// 計時結束
		end := time.Now()

		// 儲存時間資訊
		filePath := "./data/time.yaml"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.Create(filePath)
		}
		fileContent, _ := os.ReadFile(filePath)
		data := make(map[string]interface{})
		yaml.Unmarshal(fileContent, &data)
		key := fmt.Sprintf("time_%d", multiple)
		data[key] = (end.Sub(start)).Seconds()
		yamlData, _ := yaml.Marshal(data)
		os.WriteFile(filePath, yamlData, 0644)

		imgPath := fmt.Sprintf("data/result_%d.png", multiple)
		// 儲存圖片
		if err := p.Save(4*vg.Inch, 4*vg.Inch, imgPath); err != nil {
			log.Fatal(err)
		}

	}
}

// 讀取參數的函式
func readConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// Config struct
type Config struct {
	ClusterNum  int `yaml:"CLUSTER_NUM"`
	NowMultiple int `yaml:"NOW_MULTIPLE"`
}

// 計算統計數據的函式
func calculateStatistics(pts plotter.XYs) (float64, float64, float64, float64) {
	var sumX, sumY float64

	for _, pt := range pts {
		sumX += pt.X
		sumY += pt.Y
	}

	count := float64(len(pts))
	meanX := sumX / count
	meanY := sumY / count

	// 計算標準差
	var sumSqX, sumSqY float64
	for _, pt := range pts {
		sumSqX += (pt.X - meanX) * (pt.X - meanX)
		sumSqY += (pt.Y - meanY) * (pt.Y - meanY)
	}

	stddevX := math.Sqrt(sumSqX / count)
	stddevY := math.Sqrt(sumSqY / count)

	return meanX, meanY, stddevX, stddevY
}
