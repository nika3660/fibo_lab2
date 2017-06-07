import matplotlib.pyplot as plt

file1 = open("file1.txt")
f1 = []
for i in file1.readlines():
    a = i.split(" ")
    f1.append(a[1])

file2 = open("file2.txt")
f2 = []
for i in file2.readlines():
    a = i.split(" ")
    f2.append(a[1])

fig = plt.figure(1)

plt.subplot(2, 1, 1)
plt.plot(f1)

plt.subplot(2, 1, 2)
plt.plot(f2)

plt.show()
