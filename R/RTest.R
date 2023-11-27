library(jsonlite)
library(httr)
library(ggplot2)
library(yaml)

# 讀取參數設定
config <- yaml::yaml.load_file('../parameter.yaml') # nolint
CLUSTER_NUM <- config$CLUSTER_NUM
NOW_MULTIPLE <- config$NOW_MULTIPLE

multiple <- NOW_MULTIPLE

# 計時開始
start <- Sys.time()

response <- GET(paste0("http://localhost:8000?size=", multiple))
if (status_code(response) == 200) {
  # 解析 JSON 格式的回應
  points_str <- content(response, "text", encoding = "UTF-8") # 取得回應內容
  points <- fromJSON(points_str) # 將 JSON 格式轉換為 string 格式
  prettified_json <- prettify(points) # 將 string 格式轉換為 JSON 格式
  points <- fromJSON(prettified_json) # 將 JSON 格式轉換為 R 資料結構

  # 資料處理
  x_values <- as.numeric(points$x)
  y_values <- as.numeric(points$y)
  
  dx <- cbind(x_values, y_values) # 創建一個矩陣包含 x_values 和 y_values
  # print(dx)

  # 使用kmeans進行分群
  kmeans_result <- kmeans(dx, centers = CLUSTER_NUM, nstart = 10)
  
  # 視覺化 k-means 分群結果，標記出每群的中心點座標
  ggplot(data.frame(x_values, y_values, cluster = as.factor(kmeans_result$cluster)), aes(x = x_values, y = y_values, color = cluster)) +
    geom_point() +
    stat_ellipse(level = 0.95, type = "norm", geom = "polygon", alpha = 0.2) + # nolint: line_length_linter.
    stat_ellipse(level = 0.95, type = "norm", geom = "path", alpha = 0.2) +
    geom_point(data = data.frame(kmeans_result$centers), aes(x = x_values, y = y_values), color = "black", size = 3) +
    ggtitle(paste("K-means Clustering (", multiple, " points)", sep = "")) +
    theme(plot.title = element_text(hjust = 0.5)) +
    theme(legend.position = "none") +
    theme_bw()

  # 使用 kmeans_result$cluster 將原始資料 dx 按群集分組
  print("means: ")
  print(kmeans_result$centers)

  clusters <- kmeans_result$cluster
  centers <- kmeans_result$centers
  df <- as.data.frame(dx)
  res_sd <- NULL

  for (cl in c(unique(clusters))){
    df_part <- df[clusters == cl, ] # 取出群集cl的資料
    sd_s <- apply(df_part, 2, sd)
    names(sd_s) <- paste("sd_", colnames(df_part), sep = "")
    res_part <- c(cluster = cl, sd_s)
    res_sd <- rbind(res_sd, res_part)
  }


  res_sd <- as.data.frame(res_sd)
  rownames(res_sd) <- res_sd$cluster
  res_sd <- res_sd[order(res_sd$cluster), ]
  print("stddev: ")
  print(res_sd)

  # 計時結束
  end <- Sys.time()
  print(paste("時間: ", end - start))
  
  if (!file.exists('./data/time.yaml')) {
    file.create('./data/time.yaml')
  }

  data <- yaml::yaml.load_file('./data/time.yaml')
  data[paste0("time_", as.character(multiple))] <- end - start
  yaml::write_yaml(data, './data/time.yaml')

  # 顯示圖片
  ggsave(paste0("./data/result_", multiple, ".png"), width = 10, height = 10)
}
