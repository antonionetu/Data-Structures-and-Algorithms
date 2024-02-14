const hourGlasses = [
    [1, 1, 1, 0, 0, 0], 
    [0, 1, 0, 0, 0, 0], 
    [1, 1, 1, 0, 0, 0], 
    [0, 0, 2, 4, 4, 0], 
    [0, 0, 0, 2, 0, 0], 
    [0, 0, 1, 2, 4, 0]
]

const higherHourglass = (matrix) => {
    let higher = Number.NEGATIVE_INFINITY
    let sum

    for(let i = 0; i < 4; i++) {
        for(let j = 0; j < 4; j++) {
            sum = matrix[i][j] + matrix[i][j+1] + matrix[i][j+2] + matrix[i+1][j+1] + matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]
            sum > higher 
                ? higher = sum 
                : higher
        }
	}
    
    return higher
}

console.log(
    higherHourglass(hourGlasses)
);