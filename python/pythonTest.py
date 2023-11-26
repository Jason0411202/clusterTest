import requests
import numpy as np
import matplotlib.pyplot as plt
import json
from sklearn.cluster import KMeans
import yaml

with open('../parameter.yaml', 'r', encoding='utf-8') as file:
    config = yaml.safe_load(file)

CLUSTER_NUM = config['CLUSTER_NUM']


response = requests.get("http://localhost:8000")
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
    plt.scatter(x_values, y_values, c=new_dy, s=1)
    plt.show()
