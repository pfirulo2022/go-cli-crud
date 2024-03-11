package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/pfirulo2022/go-cli-crud.git/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{}

	}

	if len(os.Args) < 2 {
		printUsage()
		return
	} else {
		switch os.Args[1] {
		case "list":
			task.ListTask(tasks)
		case "add":
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("\n¿Cual es tu tarea?")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			tasks = task.AddTask(tasks, name)
			task.SaveTask(file, tasks)

		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Debes proporcionar un ID.")
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("El id debe ser un número.")
			}

			tasks = task.DeleteTask(tasks, id)
			task.SaveTask(file, tasks)

		case "complete":
			if len(os.Args) < 3 {
				fmt.Println("Debes proporcionar un ID.")

			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("El id debe ser un número.")

			}
			tasks = task.CompleteTask(tasks, id)
			task.SaveTask(file, tasks)

		}

	}

}

func printUsage() {
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}
