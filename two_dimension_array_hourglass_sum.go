/*
Given a  2D Array, :

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
We define an hourglass in  to be a subset of values with indices falling in this pattern in 's graphical representation:

a b c
  d
e f g
There are  hourglasses in , and an hourglass sum is the sum of an hourglass' values. Calculate the hourglass sum for every hourglass in , then print the maximum hourglass sum.

For example, given the 2D array:

-9 -9 -9  1 1 1 
 0 -9  0  4 3 2
-9 -9 -9  1 2 3
 0  0  8  6 6 0
 0  0  0 -2 0 0
 0  0  1  2 4 0
We calculate the following  hourglass values:

-63, -34, -9, 12, 
-10, 0, 28, 23, 
-27, -11, -2, 10, 
9, 17, 25, 18
Our highest hourglass value is  from the hourglass:

0 4 3
  1
8 6 6
Note: If you have already solved the Java domain's Java 2D Array challenge, you may wish to skip this challenge.

Function Description

Complete the function hourglassSum in the editor below. It should return an integer, the maximum hourglass sum in the array.

hourglassSum has the following parameter(s):

arr: an array of integers
Input Format

Each of the  lines of inputs  contains  space-separated integers .

Constraints

Output Format

Print the largest (maximum) hourglass sum found in .

Sample Input

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
Sample Output

19
Explanation

 contains the following hourglasses:

image

The hourglass with the maximum sum () is:

2 4 4
  2
1 2 4
*/
func hourglassSumTest() {
	println("Please enter 6x6 array each row in new line and column separated by space -> ")
	arr := [6][6]int32{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	print(hourglassSum(arr))
}
func hourglassSum(arr [6][6]int32) int32 {
	var sum int32
	setFirstTime := true
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr)-2; j++ {
			var tempSum int32
			for k := 0; k < 3; k++ {
				tempSum += arr[i][j+k]
				fmt.Printf("%d ", arr[i][j+k])
				tempSum += arr[i+2][j+k]
				fmt.Printf("%d ", arr[i+2][j+k])
			}
			println()
			tempSum += arr[i+1][j+1]
			fmt.Printf(" %d", arr[i+1][j+1])
			println()
			fmt.Printf("tempSum %d", tempSum)
			println()
			if setFirstTime {
				sum = tempSum
				setFirstTime = false
			} else if tempSum > sum {
				sum = tempSum
			}
		}
	}
	return sum
}
