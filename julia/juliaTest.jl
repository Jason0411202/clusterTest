using HTTP
using JSON
using Plots
using Clustering
using YAML

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
    # points_str="[{\"x\":\"90.75090985203879\",\"y\":\"21.853916258697804\"}]"
    points = JSON.parse(points_str)
    points = JSON.parse(points)

    # 將 x 和 y 座標分別存成列表
    x_values = Float64[parse(Float64, point["x"]) for point in points]
    y_values = Float64[parse(Float64, point["y"]) for point in points]
    dx = hcat(x_values, y_values)

    println("size of dx", size(dx)) # size of dx(142, 2)


    # # 分k群並繪製圖表
    # kmeans = kmeans(dx', CLUSTER_NUM; init = :kmpp, n_init = 10)
    # new_dy = assignments(kmeans)
    # centers = [mean(dx[findall(new_dy .== i), :], dims=1) for i in 1:CLUSTER_NUM]

    # plt = scatter(x_values, y_values, c = new_dy, markersize = 1)

    # # 標記每個群的平均值和標準差
    # for i in 1:CLUSTER_NUM
    #     cluster_points = dx[findall(new_dy .== i), :]
    #     cluster_mean = mean(cluster_points, dims=1)
    #     cluster_stddev = std(cluster_points, dims=1)
    #     scatter!([centers[i][1]], [centers[i][2]], marker = (:x, 8, :red), label = "")
    #     errorbar!(
    #         [centers[i][1]],
    #         [centers[i][2]],
    #         xerr = cluster_stddev[1],
    #         yerr = cluster_stddev[2],
    #         marker = (:circle, 4, 0.2, :black),
    #         label = "",
    #     )

    #     # 顯示平均值和標準差的數值
    #     plt = annotate!(
    #         centers[i][1] + 0.1,
    #         centers[i][2] + 0.1,
    #         text("average: \($(round(cluster_mean[1], digits=2)), $(round(cluster_mean[2], digits=2)))\nstd: \($(round(cluster_stddev[1], digits=2)), $(round(cluster_stddev[2], digits=2)))", :blue),
    #     )
    # end

    # # 計時結束
    # global end_time = time() - start
    # println("time: ", end_time)

    # if !isfile("./data/time.yaml")
    #     touch("./data/time.yaml")
    # end

    # data = YAML.load_file("./data/time.yaml") |> x -> x == nothing ? Dict() : x
    # data["time_$(multiple)"] = end_time
    # YAML.write("./data/time.yaml", data)

    # # savefig("./data/result_$(multiple).png")
    # png("./data/result_$(multiple).png")
    # plt
end
