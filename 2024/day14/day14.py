import re
from pathlib import Path
TILE_ROWS = 103
TILE_COLS = 101
def parse_input_file(filename):
    with open(filename, 'r') as file:
        lines = file.readlines()
        robots = []
        for i in range(0, len(lines)):
            match = re.search(r'p\=(\d+),(\d+) v\=(-?\d+),(-?\d+)', lines[i].strip())
            pos = (int(match.group(2)), int(match.group(1)))
            vel = (int(match.group(4)), int(match.group(3)))
            
            robots.append((pos, vel))
    return robots

def move_robot(robot, seconds):
    pos = robot[0]
    vel = robot[1]
    new_pos = ((pos[0] + seconds * vel[0])%TILE_ROWS, (pos[1] + seconds * vel[1])%TILE_COLS)
    return new_pos

def get_quadrant(pos):
    if pos[0] == TILE_ROWS//2 or pos[1] == TILE_COLS//2:
        return 0
    elif pos[0] < TILE_ROWS//2 and pos[1] < TILE_COLS//2:
        return 1
    elif pos[0] < TILE_ROWS//2 and pos[1] > TILE_COLS//2:
        return 2
    elif pos[0] > TILE_ROWS//2 and pos[1] < TILE_COLS//2:
        return 3
    else:
        return 4

def part1():
    import sys
    args = sys.argv
    if len(args) > 2:
        print("Usage: invalid number of args")
        return
    if len(args)==2 and args[1] == '1':
        robots = parse_input_file('input1.txt')
        TILE_COLS = 11
        TILE_ROWS = 7
    else:
        robots = parse_input_file('input.txt')
    quad = [0] * 5
    safety_factor = 1
    for robot in robots:
        new_pos = move_robot(robot, 100)
        quad[get_quadrant(new_pos)] += 1
    for i in range(1, 5):
        safety_factor = safety_factor * quad[i]
    print(safety_factor)

def part2(steps=100):
    robots = parse_input_file('input.txt')
    grid = [['.' for i in range(TILE_COLS)] for j in range(TILE_ROWS)]
    images = []
    for robot in robots:
        grid[robot[0][0]][robot[0][1]] = '#'
    for s in range(1,steps):
        for i in range(len(robots)):
            robots[i] = (move_robot(robots[i], 1), robots[i][1])
        grid = [['.' for i in range(TILE_COLS)] for j in range(TILE_ROWS)]
        for robot in robots:
            grid[robot[0][0]][robot[0][1]] = '#'
        images.append(f"After {s} seconds:\n")
        for row in grid:
            images.append(''.join(row))
            images.extend("\n")
        found = False 
        for imRow in images:
            if "#########" in imRow:
                found = True
                break
        if found:
            break
        else:
            images = []
    Path(__file__, "..", "output.txt").resolve().write_text("\n".join(images))
part1()
part2(10000)