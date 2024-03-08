package tasks

import "fmt"

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTask(tasks []Task) {
	if len(tasks) == 0 {
		println("No hay tareas.")
		return
	}

	for i, task := range tasks {
		fmt.Println("%d %s", i, task.Name)
	}
}
