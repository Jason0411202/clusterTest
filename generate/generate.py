import numpy as np
import matplotlib.pyplot as plt
import yaml
import os

with open('../parameter.yaml', 'r', encoding='utf-8') as file:
    config = yaml.safe_load(file)

CLUSTER_NUM = config['CLUSTER_NUM']
MIN_POINTS = config['MIN_POINTS']
MAX_POINTS = config['MAX_POINTS']
SPACE_SIZE = config['SPACE_SIZE']
DATA_MAXSIZE_MULTIPLE = config['DATA_MAXSIZE_MULTIPLE']

multiple=1
while(multiple<=DATA_MAXSIZE_MULTIPLE):
    # 產生隨機資料
    data = []
    for i in range(CLUSTER_NUM):
        num = np.random.randint(MIN_POINTS*multiple, MAX_POINTS*multiple) # 每個群的資料點數量
        center = np.random.randint(SPACE_SIZE*multiple, size=2) # 生成隨機中心點
        distribute_factor=np.random.randint(1,15) # 生成隨機分布因子
        data.append(np.random.randn(num, 2)*distribute_factor*multiple + center) # 生成隨機資料

    # 繪製資料
    for i in range(CLUSTER_NUM):
        plt.scatter(data[i][:, 0], data[i][:, 1], s=1)

    plt.show()

    # 將資料寫入data 資料夾中

    if not os.path.exists('./data/data_'+ str(multiple) +'.txt'):
        open('./data/data_'+ str(multiple) +'.txt', 'a').close()
    with open('./data/data_'+ str(multiple) +'.txt', 'w') as f:
        for i in range(CLUSTER_NUM):
            for j in range(data[i].shape[0]):
                f.write(str(data[i][j][0]) + ',' + str(data[i][j][1]) + ',' + str(i) + '\n')

    multiple*=10

