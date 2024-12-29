Atoken = 3
Btoken = 1
import re
class Solution():
    def __init__(self):
        self.minToken = float('inf')
        self.visited = set()

    def setPrizeLoc(self, x, y):
        self.prizeLoc = (x,y)
    
    def setATravel(self, x, y):
        self.ATravel = (x,y)
    
    def setBTravel(self, x, y):
        self.BTravel = (x, y)
    
    def dfs(self, x, y, token):
        if x > self.prizeLoc[0] or y > self.prizeLoc[1]:
            return
        if x == self.prizeLoc[0] and y == self.prizeLoc[1]:
            self.minToken = min(self.minToken, token)
            return
        if (x,y) in self.visited:
            return
        
        self.dfs(x+self.ATravel[0], y+self.ATravel[1], token + Atoken)
        self.dfs(x+self.BTravel[0], y+self.BTravel[1], token + Btoken)
        self.visited.add((x,y))
    
    def driver(self):
        self.dfs(0,0,0)
        return self.minToken

def parse_input_file(filename):
    with open(filename, 'r') as file:
        lines = file.readlines()
        data = []
        for i in range(0, len(lines), 4):
            match = re.search(r'X\+(\d+), Y\+(\d+)', lines[i].strip())
            a_travel = (int(match.group(1)), int(match.group(2)))
            match = re.search(r'X\+(\d+), Y\+(\d+)', lines[i+1].strip())
            b_travel = (int(match.group(1)), int(match.group(2)))
            match = re.search(r'X\=(\d+), Y\=(\d+)', lines[i+2].strip())
            prize = (int(match.group(1)), int(match.group(2)))
            data.append((a_travel, b_travel, prize))
    return data
def part1():
    data = parse_input_file('input.txt')
    result = 0
    for a, b, p in data:
        s = Solution()
        s.setATravel(a[0], a[1])
        s.setBTravel(b[0], b[1])
        s.setPrizeLoc(p[0], p[1])
        token = s.driver()
        if token != float('inf'):
            result += token
    return result
def part2():
    data = parse_input_file('input.txt')
    result = 0

    # No need of DFS this is simple linear equation
    for a, b, p in data:
        p0 = p[0] + 10000000000000
        p1 = p[1] + 10000000000000
        den = abs(a[0]*b[1] - a[1]*b[0])
        numx = abs(p0*b[1] - p1*b[0])
        numy = abs(a[0]*p1 - a[1]*p0)
        if den == 0:
            continue
        if numx % den != 0 or numy % den != 0:
            continue
        x = numx // den
        y = numy // den
        token = x*Atoken + y*Btoken
        if token != float('inf'):
            result += token
    return result
def main():
    #print(part1())
    print(part2())

import sys
sys.setrecursionlimit(1500)
if __name__ == '__main__':
    main()