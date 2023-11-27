import requests
import numpy as np
import matplotlib.pyplot as plt
import json
from sklearn.cluster import KMeans
import yaml
import time
import os

with open('../parameter.yaml', 'r', encoding='utf-8') as file:
    config = yaml.safe_load(file)

CLUSTER_NUM = config['CLUSTER_NUM']
NOW_MULTIPLE = config['NOW_MULTIPLE']

multiple=NOW_MULTIPLE

# 計時
start = time.time()

response = requests.get("http://localhost:8000?size="+str(multiple))
if response.status_code == 200:
    # 解析 JSON 格式的回應
    points_str = response.json()
    points = json.loads(points_str)

    # 將 x 和 y 座標分別存成列表
    x_values = [float(point['x']) for point in points]
    y_values = [float(point['y']) for point in points]
    dx = np.array(list(zip(x_values, y_values))).reshape(len(x_values), 2)

    # 分k群並繪製圖表
    kmeans = KMeans(n_clusters=CLUSTER_NUM, n_init=10)
    kmeans.fit(dx)

    new_dy = kmeans.predict(dx)
    centers = kmeans.cluster_centers_  # 取得各群中心點

    plt.clf()
    plt.scatter(x_values, y_values, c=new_dy, s=1)

    # 標記每個群的平均值和標準差
    for i, center in enumerate(centers):
        cluster_points = dx[new_dy == i]  # 取得每個群的點
        cluster_mean = np.mean(cluster_points, axis=0)  # 計算群的平均值
        cluster_stddev = np.std(cluster_points, axis=0)  # 計算群的標準差
        plt.scatter(center[0], center[1], marker='x', color='red', s=100)  # 標記中心點
        plt.errorbar(
            center[0],
            center[1],
            xerr=cluster_stddev[0],
            yerr=cluster_stddev[1],
            fmt='o',
            color='black'
        )  # 標記標準差
        
        # 顯示平均值和標準差的數值
        plt.text(
            center[0] + 0.1, center[1] + 0.1,
            f'average: ({cluster_mean[0]:.2f}, {cluster_mean[1]:.2f})\nstd: ({cluster_stddev[0]:.2f}, {cluster_stddev[1]:.2f})',
            fontsize=8,
            color='blue'
        )

    # 計時結束
    end = time.time()
    print("time: ", end-start)

    if not os.path.exists('./data/time.yaml'):
        open('./data/time.yaml', 'a').close()
    with open('./data/time.yaml', 'r') as file:
        data = yaml.safe_load(file)  or {}

    data['time_'+str(multiple)] = end-start
    with open('./data/time.yaml', 'w') as file:
        yaml.dump(data, file)

    # plt.show()
    plt.savefig('./data/result_' + str(multiple) + '.png')
    