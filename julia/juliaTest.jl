using HTTP
using JSON
using Plots
using Clustering
using YAML
using Statistics

config = YAML.load_file("../parameter.yaml")

CLUSTER_NUM = config["CLUSTER_NUM"]
NOW_MULTIPLE = config["NOW_MULTIPLE"]

multiple = NOW_MULTIPLE

# 計時
start = time()

response = HTTP.get("http://localhost:8000?size=" * string(multiple))
if response.status == 200
    # 解析 JSON 格式的回應
    points_str = String(response.body)
    points = JSON.parse(points_str)
    points = JSON.parse(points)

    # 將 x 和 y 座標分別存成列表
    x_values = Float64[parse(Float64, point["x"]) for point in points]
    y_values = Float64[parse(Float64, point["y"]) for point in points]
    dx = hcat(x_values, y_values)

    println("size of dx", size(dx)) # size of dx(142, 2)


    # 計算 k-means
    result = kmeans(dx', CLUSTER_NUM)

    # 繪製散佈圖
    scatter(x_values, y_values, group=result.assignments, xlabel="X", ylabel="Y", title="K-means Clustering")

    # 分別計算並印出每群的 mean stddev
    labels = result.assignments
    for i in 1:CLUSTER_NUM
        cluster_points = dx[labels .== i, :]
        cluster_mean = mean(cluster_points, dims=1)
        cluster_stddev = std(cluster_points, dims=1)

        print("group $i means: ", cluster_mean)
        println("  stddev: ", cluster_stddev)
    end

    # 計時結束
    global end_time = time() - start
    println("time: ", end_time)

    if !isfile("./data/time.yaml")
        touch("./data/time.yaml")
    end

    # YAML.write_file("./data/time.yaml", "time_" * string(NOW_MULTIPLE) * ": " * string(end_time))

    # 儲存圖片
    savefig("./data/result_" * string(NOW_MULTIPLE) * ".png")
end
