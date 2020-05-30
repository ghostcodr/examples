/*
Emma is playing a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. She can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus  or . She must avoid the thunderheads. Determine the minimum number of jumps it will take Emma to jump from her starting postion to the last cloud. It is always possible to win the game.

For each game, Emma will get an array of clouds numbered  if they are safe or  if they must be avoided. For example,  indexed from . The number on each cloud is its index in the list so she must avoid the clouds at indexes  and . She could follow the following two paths:  or . The first path takes  jumps while the second takes .

Function Description

Complete the jumpingOnClouds function in the editor below. It should return the minimum number of jumps required, as an integer.

jumpingOnClouds has the following parameter(s):

c: an array of binary integers
Input Format

The first line contains an integer , the total number of clouds. The second line contains  space-separated binary integers describing clouds  where .

Constraints

Output Format

Print the minimum number of jumps needed to win the game.

Sample Input 0

7
0 0 1 0 0 1 0
Sample Output 0

4
Explanation 0:
Emma must avoid  and . She can win the game with a minimum of  jumps:

jump(2).png

Sample Input 1

6
0 0 0 0 1 0
Sample Output 1

3
Explanation 1:
The only thundercloud to avoid is . Emma can win the game in  jumps:

jump(5).png
*/
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(cn []int32) int32 {
    // var noOfClouds int32s
    // println("Please enter the number of clouds -> ")
    // fmt.Scan(&noOfClouds)
    // println("Please enter the space separated cloud numbers 0|1 -> ")
    // cn := make([]int, noOfClouds)
    for i := 0; i < len(cn); i++ {
        fmt.Scan(&cn[i])
    }
    jumps := 0
    for idx := 0; idx < len(cn); idx++ {
        if idx+2 < len(cn) && cn[idx+2] == 0 {
            idx = idx + 1
            jumps++
        } else if idx+1 < len(cn) && cn[idx+1] == 0 {
            jumps++
        }
    }
    return int32(jumps)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    cTemp := strings.Split(readLine(reader), " ")

    var c []int32

    for i := 0; i < int(n); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }

    result := jumpingOnClouds(c)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
