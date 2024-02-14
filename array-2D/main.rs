#[warn(dropping_copy_types)]

const HOURGLASSES: [[i32; 6]; 6] = [
    [1, 1, 1, 0, 0, 0],
    [0, 1, 0, 0, 0, 0],
    [1, 1, 1, 0, 0, 0],
    [0, 0, 2, 4, 4, 0],
    [0, 0, 0, 2, 0, 0],
    [0, 0, 1, 2, 4, 0],
];


fn main() {
    println!("{}", get_higher_hourglass(&HOURGLASSES));
}

fn get_higher_hourglass(matrix: &[[i32; 6]; 6]) -> i32 {
    let mut higher = i32::MIN;
    
    for i in 0..=3 {
        for j in 0..=3 {
            {
                let sum = matrix[i][j] + matrix[i][j + 1] + matrix[i][j + 2] +
                        matrix[i + 1][j + 1] +
                        matrix[i + 2][j] + matrix[i + 2][j + 1] + matrix[i + 2][j + 2];
                
                if sum > higher {
                    higher = sum
                }
            }
        }
    }

    return higher
}