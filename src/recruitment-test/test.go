package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	param := ""
	for param != "exit" {
		fmt.Println("=======================")
		fmt.Println("1. TASK 1")
		fmt.Println("2. TASK 2")
		fmt.Println("3. TASK 3")
		fmt.Println("4. TASK 4")
		fmt.Print("Give me a number: ")
		fmt.Scanln(&param)
		fmt.Println("=======================")

		switch param {
		case "1":
			fmt.Println("TASK 1 - ASCENDING/DESCENDING SORT")
			fmt.Println()
			var input = []int{1, 4, 5, 6, 8, 2}
			reverse := false
			sorter(reverse, input)

		case "2":
			fmt.Println("TASK 2 - QUEUE")
			fmt.Println()
			fmt.Println()

		case "3":
			fmt.Println("TASK 3 - COMPARE 2 FOLDER")
			fmt.Println()
			listdir()

		case "4":
			fmt.Println("TASK 4 - CONCURRENCY")
			fmt.Println()
			fmt.Println()
		}
	}
}

func sorting_visualization(input []int) {
	max_val := 0
	for _, val := range input {
		if val > max_val {
			max_val = val
		}
	}

	// Print BarChart
	for i := max_val; i > 0; i-- {
		for _, val := range input {
			if val-i >= 0 && val-i <= val {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

	// Print list
	for idx, val := range input {
		if idx < len(input)-1 {
			fmt.Print(val, " ")
		} else {
			fmt.Println(val)
		}
	}

	fmt.Println()
}

func sorter(rev bool, input []int) {
	sorting_visualization(input)
	param := false
	for i := len(input) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if rev {
				param = input[i] > input[j]
			} else {
				param = input[i] < input[j]
			}

			if param {
				input[i] += input[j]
				input[j] = input[i] - input[j]
				input[i] -= input[j]
				sorting_visualization(input)
			}
		}
	}
}

func listdir() {
	var filelist = func(path string) []os.FileInfo {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		return files
	}
	path := "./"

	for _, f := range filelist(path + "/source") {
		if f.IsDir() {
			// fmt.Println(f.Name())

		} else {
			for _, ff := range filelist(path + "/target") {
				if ff.IsDir() {

				} else {
					fmt.Println("ff name: ", f.Name(), ff.Name())
					if f.Name() == ff.Name() {
						fmt.Println("EXIST", ff.Name())
					}
				}
			}
		}
	}
}
