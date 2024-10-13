package main 

import (
     _ "embed"
     "strings"
     "fmt"
     "regexp"
     "strconv"
     "flag"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans, err := part1(input)
        if err != nil {
            panic("Error in part 1")
        }
		fmt.Println("Output:", ans)
	} 
}

func part1(input string) (int, error) {
    input = strings.TrimRight(input, "\n")
    lines := strings.Split(input, "\n")
    sum := 0
    for _, line := range lines {
        reg := regexp.MustCompile("[0-9]")
        numbers := reg.FindAll([]byte(line), -1)
        calValues := strings.Join([]string{string(numbers[0]), string(numbers[len(numbers)-1])}, "")
        numberToAdd, err := strconv.Atoi(calValues)
        if err != nil {
            return sum, err
        }
        sum += numberToAdd
    }
    return sum, nil

}
