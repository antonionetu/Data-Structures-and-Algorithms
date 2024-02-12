def get_heigher_hourglass(matrix):
    heigher = float('-inf') # Initialize as the lowest possible value

    for i in range(0, 4):
        for j in range(0, 4):
            current_sum = matrix[i][j] + matrix[i][j+1] + matrix[i][j+2] + matrix[i+1][j+1] + matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]
            if current_sum > heigher:
                heigher = current_sum

    return heigher