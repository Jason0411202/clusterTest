# English
## File Structure
* generate folder: Contains programs for generating data and storing generated data
* Go folder: Contains Go clustering programs, along with generated graphs and time analysis
* julia folder: Contains julia clustering programs, along with generated graphs and time analysis
* python folder: Contains Python clustering programs, along with generated graphs and time analysis
* R folder: Contains R clustering programs, along with generated graphs and time analysis
* time_analysis folder: Contains programs for analyzing processing time, as well as analysis results
* parameter.yaml: Contains parameters required for each program
* server.go: Provides HTTP service

## Usage
1. Navigate to the folder where generate_data.py is located and execute the program to generate data
2. Navigate to the root directory of the project and execute server.go to provide HTTP service
3. Execute the corresponding clustering program based on different programming languages
    *  For the Go part, navigate to the folder where GoTest.go is located and execute the program for clustering (to ensure consistent conditions for each execution, modify NOW_MULTIPLE in parameter.yaml for different sizes of data sets)
    *  For the julia part, navigate to the folder where juliaTest.jl is located and execute the program for clustering (to ensure consistent conditions for each execution, modify NOW_MULTIPLE in parameter.yaml for different sizes of data sets)
    * For the Python part, navigate to the folder where pythonTest.py is located and execute the program for clustering (to ensure consistent conditions for each execution, modify NOW_MULTIPLE in parameter.yaml for different sizes of data sets)
    * For the R part, navigate to the folder where RTest.R is located and execute the program for clustering (to ensure consistent conditions for each execution, modify NOW_MULTIPLE in parameter.yaml for different sizes of data sets)

## Python
* Python version 3.11.3
* KMeans algorithm is used for clustering
* The first call to the Kmean function takes longer; subsequent calls are shorter
* Largest community, easier to find information, and easy integration with other functionalities

## R
* R version 4.1.2 (2021-11-01)
* kmeans algorithm is used for clustering
* Unable to save graphs smoothly when data volume is extremely large
* Slow speed in plotting and saving graphs
* Always encounter some strange errors, and the syntax is not very intuitive

## Julia
* Julia version 1.9.4
* kmeans algorithm is used for clustering
* Every time the program is executed, it takes about ten seconds to compile (it is said that it has been optimized, but in earlier versions, it took about a minute)
* Feels immature, and the community is small; many functionalities can only be found on GitHub with brief usage methods, or are still under development (such as writing Julia objects into YAML)
* But excluding the compilation time, the execution speed is not slow compared to other languages

## Go
* go version go1.21.4 windows/amd64
* kmeans algorithm is used for clustering
* The community of Go in the field of data analysis is not very large, and there is no famous package like Python specifically for data analysis
* Using "github.com/muesli/kmeans" for kmeans analysis, this is currently the best package found
* Using "gonum.org/v1/plot" for plotting

## Summary
julia > python > R > Go
* Among these four languages, julia is the most recommended. julia was initially designed for data analysis, so it performs excellently in terms of speed and syntactic conciseness. However, julia also has a fatal flaw, which is the small community. Many functionalities are still under development, and there is little information available, making it inconvenient to search for information. If more people use julia in the future and it continues to be maintained, I believe it will be the preferred programming language.
* The second recommendation is Python. Python has many advantages that other languages do not have, such as a large community, easy integration with other functionalities, and high completeness of data analysis-related modules. It is suitable for this project. The only drawback is that its speed is not as fast as other languages (especially when dealing with big data). Overall, its advantages outweigh its disadvantages. If speed is not of utmost importance, Python is a good choice.
* The third recommendation is R. Although R is also a language specifically designed for data analysis, it has been around for 30 years without gaining a decisive advantage in data analysis. This may indirectly reflect some problems with this language. R is one of the languages I find the most difficult to learn, with non-intuitive syntax and strange errors appearing constantly during development, which gives me a not-so-good impression of the language. However, it is still a mature language with a relatively large community, and its speed is not slow. If someone is familiar with R, it is also a good choice.
* The last recommendation is Go. Go is one of my favorite languages at the moment, especially suitable for handling backend services, and its inherent support for parallel processing features also improves the efficiency of handling transactions. However, in the field of data analysis, the community is extremely small, and it is noticeable that the discussion level is low when searching for information. There is no powerful data analysis module like Python that includes all common functionalities, and many functionalities need to be implemented by oneself. The biggest advantage is that, after testing, Go has the best efficiency in handling big data. If efficiency is highly valued, it can be considered for use.

# 中文
## 檔案架構
* generate資料夾: 包含產生資料的程式跟存放產生的資料
* Go資料夾: 包含Go的分群程式，以及畫出來的圖、耗時
* julia資料夾: 包含julia的分群程式，以及畫出來的圖、耗時
* python資料夾: 包含python的分群程式，以及畫出來的圖、耗時
* R資料夾: 包含R的分群程式，以及畫出來的圖、耗時
* time_analysis資料夾: 包含分析耗時的程式，以及分析結果
* parameter.yaml: 包含各程式所需參數設定
* server.go: 提供 http 服務

## 使用方法
1. 移動至generate_data.py所在的資料夾中，執行該程式產生資料
2. 移動至專案根目錄中，執行server.go提供 http 服務
3. 根據不同程式語言執行對應的分群程式
    * Go 的部分，移動至GoTest.go所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)
    * julia 的部分，移動至juliaTest.jl所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)
    * python 的部分，移動至pythonTest.py所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)
    * R 的部分，移動至RTest.R所在的資料夾中，執行該程式進行分群 (針對不同大小的資料集，為確保每次執行的條件相同，需修改parameter.yaml中的NOW_MULTIPLE)

## python
* Python version 3.11.3
* 分群演算法採用KMeans演算法
* 程式第一次呼叫Kmean function 時，耗時較長；重復呼叫時，耗時較短
* 社群最大，查資料較為容易，容易跟其他功能整合

## R
* R version 4.1.2 (2021-11-01)
* 分群演算法採用kmeans演算法
* 在資料量極大的時候，圖無法順利存檔
* 繪圖並存檔的速度慢
* 總是會跳出一些奇怪的錯誤，語法不太直觀

## julia
* julia version 1.9.4
* 分群演算法採用kmeans演算法
* 每次要執行程式的時候，都要大概先等十秒左右編譯的時間 (聽說已經優化過了，更早以前的版本要等一分鐘)
* 感覺尚未成熟，以及社群過小；許多功能都只查的到 github 上簡短的使用方法、甚至還在開發中 (例如將 julia 物件寫入 yaml 中)
* 但去除掉編譯時間，執行速度跟其他語言相比不慢

## Go
* go version go1.21.4 windows/amd64
* 分群演算法採用kmeans演算法
* Go 感覺在資料分析這塊的社群不太大，沒有像 python 有有名的 package 專門用來做資料分析
* 使用 "github.com/muesli/kmeans" 進行 kmeans 分析，這是目前找到包得最好的 package
* 使用 "gonum.org/v1/plot" 進行繪圖

## 總結
julia > python > R > Go
* 這四種語言中，最推薦 julia。julia 起初設計的目的便是為了資料分析，故 julia 在這塊表現上，無論是速度還是語法的簡潔性都非常優異，且作為一個新興的現代語言，也吸收了過去各語言的優點。但 julia 也有個致命的缺點，便是社群過小，不但許多功能都還在開發中，搜尋到的資料也很少，因此在查資料時有許多不便。要是未來有更多人使用 julia，且不斷維護的話，這個程式語言我覺得將會是首選。
* 第二推薦的是 python，python 有許多其他語言所沒有的優點，龐大的社群、容易與其他功能整合、資料分析相關模組完成度非常高等等，很適合拿來做這次專案，最可惜的地方是速度比不過其他語言 (尤其是面對大數據時)，整體來說瑕不掩瑜，如果沒有對速度有極致追求的話，python 是個很好的選擇。
* 第三推薦的是 R，R 雖然也是專門用於資料分析的語言，但已經 30 年了，沒有在資料分析這塊取得決定性的優勢，可能間接反映出這個語言的一些問題。R 語言感覺上是我學的最差的語言之一，語法不太直觀、開發時奇怪的錯誤不斷出現，都讓我對 R 這個語言的印象不太好。但畢竟也算是個成熟且社群不小的語言，且速度不慢，如果有人熟悉 R 的話，也是個不錯的選擇。
* 最後推薦的是 Go，Go 是我目前最喜歡的語言之一，拿來處理後端服務非常適合，且本身可以輕鬆支援平行處理的特性，也使得 Go 處理事務的效率有所提升。但在資料分析這塊，社群異常小，查資料的時候能感覺到討論度很低，也沒有像 python 有包含所有常見功能的強大資料分析模組可以使用，甚至許多功能還需要自己寫，最大的優點是經過測試，Go 在處理大數據時效率是最好的，如果非常重視效率，可以考慮一用。