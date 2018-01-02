package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// // TASK 1
	// var input = []int{1, 4, 6, 5, 8, 2}
	// sorting_visualization(input)
	// fmt.Println()
	// // true if descending, false if ascending
	// sorter(false, input)

	// TASK 3
	listdir()
}

func sorting_visualization(input []int) {
	max_val := 0
	for _, val := range input {
		if val > max_val {
			max_val = val
		}
	}

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

	for idx, val := range input {
		if idx < len(input)-1 {
			fmt.Print(val, " ")
		} else {
			fmt.Println(val)
		}
	}
}

func sorter(rev bool, input []int) {
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
				fmt.Println()
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
	path := "/home/hendra/Documents/go/"

	for _, f := range filelist(path) {
		if f.IsDir() {
			fmt.Println(f.Name(), "is Dir")
		}
	}
}
