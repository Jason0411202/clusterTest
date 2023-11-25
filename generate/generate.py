import numpy as np
import matplotlib.pyplot as plt

CLUSTER_NUM = 3 # 設定要產生的群數量
MIN_POINTS = 1 # 每個群的最小資料點數量
MAX_POINTS = 100 # 每個群的最大資料點數量
DISTRIBUTE_FACTOR = 5 # 設定資料分布的離散程度
SPACE_SIZE = 100 # 定義分布空間大小

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

