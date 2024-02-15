import yaml
import matplotlib.pyplot as plt

filePath=['../python/data/time.yaml','../R/data/time.yaml', '../Go/data/time.yaml', '../julia/data/time.yaml']
language=['python','R', 'Go', 'julia']
plt.figure(figsize=(10, 6))

for i in range(len(filePath)):
    with open(filePath[i], 'r') as file:
        data = yaml.safe_load(file)
    times, values = zip(*data.items())
    times=[]
    for j in range(5):
        times.append(str(10**j)+" time")
    values = [float(val) for val in values]
    plt.plot(times, values, label=language[i])

plt.title('time analyze')
plt.xlabel('Data Set Magnitude')
plt.ylabel('Processing Time')
plt.legend()
plt.xticks(rotation=45)
plt.grid(True)
plt.tight_layout()

# plt.show()

# 寫檔
plt.savefig('result.png')
