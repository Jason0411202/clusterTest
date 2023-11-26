import numpy as np
import matplotlib.pyplot as plt
import yaml

with open('../parameter.yaml', 'r', encoding='utf-8') as file:
    config = yaml.safe_load(file)

CLUSTER_NUM = config['CLUSTER_NUM']
MIN_POINTS = config['MIN_POINTS']
MAX_POINTS = config['MAX_POINTS']
SPACE_SIZE = config['SPACE_SIZE']
DISTRIBUTE_FACTOR = config['DISTRIBUTE_FACTOR']

# 產生隨機資料
data = []
for i in range(CLUSTER_NUM):
    num = np.random.randint(MIN_POINTS, MAX_POINTS) # 每個群的資料點數量
    center = np.random.randint(SPACE_SIZE, size=2) # 生成隨機中心點
    data.append(np.random.randn(num, 2)*DISTRIBUTE_FACTOR + center) # 生成隨機資料

# 繪製資料
for i in range(CLUSTER_NUM):
    plt.scatter(data[i][:, 0], data[i][:, 1], s=1)

plt.show()

# 將資料寫入data 資料夾中
with open('generate/data/data.txt', 'w') as f:
    for i in range(CLUSTER_NUM):
        for j in range(data[i].shape[0]):
            f.write(str(data[i][j][0]) + ',' + str(data[i][j][1]) + ',' + str(i) + '\n')

