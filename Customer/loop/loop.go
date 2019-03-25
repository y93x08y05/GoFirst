package loop

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)//括号如果换行，末尾加逗号，否则不用加
	printFile("abc.txt")
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
	forever()
}

func forever()  {
	for {
		fmt.Println("abc")
	}
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
