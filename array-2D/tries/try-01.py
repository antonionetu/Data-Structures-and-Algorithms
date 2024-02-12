def get_heigher_hourglass(matrix):
    heigher = {
        'hourglass': [],
        'sum': float('-inf')
    }
    step_X = 0
    step_Y = 0

    while step_Y < 4:
        if step_X < 4:
            current_hourglass = get_current_hourglass(matrix, step_X, step_Y)
            result = sum_hourglass(current_hourglass)

            if result > heigher['sum']:
                heigher = {
                    'hourglass': current_hourglass,
                    'sum': result
                }

            step_X += 1

        else:
            step_X = 0
            step_Y += 1

    return heigher


def get_current_hourglass(matrix, x, y):
    return [
        [matrix[y][x], matrix[y][x + 1], matrix[y][x + 2]],
        [matrix[y + 1][x + 1]],
        [matrix[y + 2][x], matrix[y + 2][x + 1], matrix[y + 2][x + 2]]
    ]


def sum_hourglass(matrix):
    result = 0
    for row in matrix:
        for item in row:
            result += item

    return result