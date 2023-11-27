import yaml
import matplotlib.pyplot as plt

# 讀取第一個 time.yaml 檔案
with open('../python/data/time.yaml', 'r') as file1:
    data1 = yaml.safe_load(file1)

# 讀取第二個 time.yaml 檔案
with open('../R/data/time.yaml', 'r') as file2:
    data2 = yaml.safe_load(file2)

# 分離資料中的時間和數值
times1, values1 = zip(*data1.items())
times2, values2 = zip(*data2.items())

# 將字串形式的數值轉換成浮點數
values1 = [float(val) for val in values1]
values2 = [float(val) for val in values2]

# 繪製折線圖
plt.figure(figsize=(10, 6))

plt.plot(times1, values1, label='python')
plt.plot(times2, values2, label='R')

plt.title('time analyze')
plt.xlabel('time')
plt.ylabel('value')
plt.legend()
plt.xticks(rotation=45)
plt.grid(True)
plt.tight_layout()

plt.show()
