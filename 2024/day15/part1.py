class Grid:
    def __init__(self, filepath):
        self.grid, self.dirs = self.parse_input_file(filepath)
        self.robot_row, self.robot_col = self.get_initial_robot_pos()
    
    def parse_input_file(self, filepath):
        with open(filepath, 'r') as f:
            content = f.read()
            parts = content.split('\n\n')
            gridStr = parts[0].split('\n')
            grid = []
            for i in range(len(gridStr)):
                grid.append(list(gridStr[i]))
            dirs = parts[1]
            dirs.replace('\n', '')
            print(type(dirs[0]))
        return grid, dirs
    
    def get_initial_robot_pos(self):
        for i in range(len(self.grid)):
            for j in range(len(self.grid[i])):
                if self.grid[i][j] == '@': 
                    return [i, j]
    def update_pixel(self, new_row, new_col, row, col):
        self.grid[new_row][new_col] = self.grid[row][col]
        self.grid[row][col] = '.'

    def update_robot_pos(self, new_row, new_col):
        self.robot_row = new_row
        self.robot_col = new_col
    
    def update(self):
        for dir in self.dirs:
            if dir not in {'^', 'v', '>', '<'}:
                continue
            new_row, new_col = self.get_new_pos(self.robot_row, self.robot_col, dir)
            if self.grid[new_row][new_col] == '#':
                continue
            if self.grid[new_row][new_col] == '.':
                self.update_pixel(new_row, new_col, self.robot_row, self.robot_col)
                self.update_robot_pos(new_row, new_col)

            if self.grid[new_row][new_col] == 'O':
                while self.grid[new_row][new_col] == 'O':
                    new_row, new_col = self.get_new_pos(new_row, new_col, dir)
                if self.grid[new_row][new_col] == '#':
                    continue
                while (new_row, new_col) != (self.robot_row, self.robot_col):
                    new_row, new_col = self.get_old_pos(new_row, new_col, dir)
                    r, c = self.get_new_pos(new_row, new_col, dir)
                    self.update_pixel(r, c, new_row, new_col)
                r, c = self.get_new_pos(self.robot_row, self.robot_col, dir)
                self.update_robot_pos(r, c)
    def get_gps_number(self, r , c):
        return (r*100) + c
    
    def sum_gps_number(self):
        res = 0
        for i in range(len(self.grid)):
            for j in range(len(self.grid[i])):
                if self.grid[i][j] == 'O':
                    res += self.get_gps_number(i, j)
        return res

    
    def get_new_pos(self, row, col, dir):
        if dir == '^':
            row -= 1
        elif dir == 'v':
            row += 1
        elif dir == '>':
            col += 1
        elif dir == '<':
            col -= 1
        return row, col
    
    def get_old_pos(self, row, col, dir):
        if dir == '^':
            row += 1
        elif dir == 'v':
            row -= 1
        elif dir == '>':
            col -= 1
        elif dir == '<':
            col += 1
        return row, col

    def print_grid(self):
        for i in range(len(self.grid)):
            print(''.join(self.grid[i]))
    def print_dirs(self):
        print(self.dirs)

g = Grid('./2024/day15/input.txt')
g.print_grid()
print("")
#g.dirs = '<<'
g.update()
print(g.sum_gps_number())
g.print_grid()