# 使用方法
1. 移動至generate_data.py所在的資料夾中，執行該程式產生資料
2. 移動至專案根目錄中，執行server.go提供 http 服務
3. 根據不同程式語言執行對應的分群程式
    * python 的部分，移動至pythonTest.py所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)

# python
* 分群演算法採用sklearn套件中的KMeans演算法
* 程式第一次呼叫Kmean function 時，耗時較長