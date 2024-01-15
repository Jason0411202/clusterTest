package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"
)

type DataPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
	// 如果有其他屬性，可以在這裡添加
}

func main() {
	// 讀取配置文件
	config, err := readConfig("../parameter.yaml")
	if err != nil {
		log.Fatal(err)
	}

	//clusterNum := config.ClusterNum
	nowMultiple := config.NowMultiple

	multiple := nowMultiple

	// 計時開始
	//start := time.Now()

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

		fmt.Println("rawPoints:", rawPoints)

		// // 轉換數據為 Golearn 需要的格式
		// instData := base.NewDenseInstances()
		// instData.AddAttribute(base.NewFloatAttribute("x"))
		// instData.AddAttribute(base.NewFloatAttribute("y"))

		// 	for i := range xValues {
		// 		instData.AddInstance(base.NewDenseInstance([]float64{xValues[i], yValues[i]}))
		// 	}

		// 	// 使用 K-Means 算法進行聚類
		// 	cls := cluster.NewKMeans(CLUSTER_NUM)
		// 	cls.Fit(instData)

		// 	// 繪製圖表
		// 	plotClusters(instData, cls, "./data/result_"+strconv.Itoa(multiple)+".png")

		// 	// 計時結束
		// 	elapsed := time.Since(start)
		// 	fmt.Println("time:", elapsed)

		// 	// 寫入時間數據到 YAML 文件
		// 	writeTimeData(elapsed, multiple)
		// } else {
		// 	fmt.Println("HTTP request failed with status:", response.StatusCode)
		// }
	}
}

// readConfig 讀取配置文件

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

// Config 包含配置信息的結構體
type Config struct {
	ClusterNum  int `yaml:"CLUSTER_NUM"`
	NowMultiple int `yaml:"NOW_MULTIPLE"`
}

// plotClusters 繪製散點圖和聚類中心
// func plotClusters(instData base.FixedDataGrid, cls *cluster.KMeans, filePath string) {
// 	plt, err := plot.New()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 提取 x 和 y 軸數據
// 	xy, err := base.InstancesToRows(instData)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 提取每個實例所屬的聚類
// 	predictions, err := cls.Predict(xy)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 設置散點圖屬性
// 	scatter, err := plotter.NewScatter(xy)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scatter.GlyphStyle.Radius = vg.Points(1)
// 	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
// 	scatter.Color = color.RGBA{B: 255, A: 255}

// 	// 將散點圖添加到繪圖區域
// 	plt.Add(scatter)

// 	// 繪製聚類中心
// 	centroids := cls.Centroids
// 	for _, centroid := range centroids {
// 		centerX := centroid[0]
// 		centerY := centroid[1]

// 		// 將聚類中心添加到繪圖區域
// 		point := plotter.NewGlyphPoints(draw.CircleGlyph{})
// 		point.Color = color.RGBA{R: 255, A: 255}
// 		point.Add(plotter.XY{X: centerX, Y: centerY})
// 		plt.Add(point)

// 		// 繪製標準差
// 		// ...

// 		// 顯示平均值和標準差的數值
// 		// ...
// 	}

// 	// 保存圖表為文件
// 	err = plt.Save(4*vg.Inch, 4*vg.Inch, filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// writeTimeData 寫入時間數據到 YAML 文件
func writeTimeData(elapsed time.Duration, multiple int) {
	filePath := "./data/time.yaml"

	// 讀取 YAML 文件
	data := make(map[string]interface{})
	yamlData, err := ioutil.ReadFile(filePath)
	if err == nil {
		err = yaml.Unmarshal(yamlData, &data)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 更新時間數據
	data["time_"+strconv.Itoa(multiple)] = elapsed.Seconds()

	// 寫入 YAML 文件
	yamlBytes, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filePath, yamlBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
