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

# R
* R version 4.1.2 (2021-11-01)
* 分群演算法採用kmeans演算法
* 在資料量極大的時候，圖無法順利存檔
* 繪圖並存檔的速度慢

# Go
* go version go1.21.4 windows/amd64