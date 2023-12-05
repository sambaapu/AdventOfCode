fileObj = open('./data', 'r');
lines = fileObj.readlines();
res = 0;
for line in lines:
    temp = "";
    for ch in line:
        if ch.isdigit():
            temp += ch;
            break;
    for i in range(len(line)-1, -1, -1):
        if line[i].isdigit():
            temp += line[i];
            break;
    res += int(temp);
    print(temp)
print(res)