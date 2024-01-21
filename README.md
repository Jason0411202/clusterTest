# 檔案架構
* generate資料夾: 包含產生資料的程式跟存放產生的資料
* python資料夾: 包含python的分群程式，以及畫出來的圖、耗時
* R資料夾: 包含R的分群程式，以及畫出來的圖、耗時
* time_analysis資料夾: 包含分析耗時的程式，以及分析結果
* parameter.yaml: 包含各程式所需參數設定
* server.go: 提供 http 服務
# 使用方法
1. 移動至generate_data.py所在的資料夾中，執行該程式產生資料
2. 移動至專案根目錄中，執行server.go提供 http 服務
3. 根據不同程式語言執行對應的分群程式
    * python 的部分，移動至pythonTest.py所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)
    * R 的部分，移動至RTest.R所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)

# python
* Python version 3.11.3
* 分群演算法採用KMeans演算法
* 程式第一次呼叫Kmean function 時，耗時較長；重復呼叫時，耗時較短
* 社群最大，查資料較為容易，容易跟其他功能整合

# R
* R version 4.1.2 (2021-11-01)
* 分群演算法採用kmeans演算法
* 在資料量極大的時候，圖無法順利存檔
* 繪圖並存檔的速度慢
* 總是會跳出一些奇怪的錯誤，語法不太直觀

# julia
* julia version 1.9.4
* 分群演算法採用kmeans演算法
* 每次要執行程式的時候，都要大概先等十秒左右編譯的時間 (聽說已經優化過了，更早以前的版本要等一分鐘)
* 感覺尚未成熟，以及社群過小；許多功能都只查的到 github 上簡短的使用方法、甚至還在開發中 (例如將 julia 物件寫入 yaml 中)
* 但去除掉編譯時間，執行速度跟其他語言相比不慢

# Go
* go version go1.21.4 windows/amd64
* 分群演算法採用kmeans演算法
* Go 感覺在資料分析這塊的社群不太大，沒有像 python 有有名的 package 專門用來做資料分析
* 使用 "github.com/muesli/kmeans" 進行 kmeans 分析，這是目前找到包得最好的 package
* 使用 "gonum.org/v1/plot" 進行繪圖