package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"wsy0414/task-tracker/service"
)

func main() {
	arguments := os.Args

	serv := service.NewTaskService("tasks.json")

	switch arguments[1] {
	case "add":
		err := serv.Add(arguments[2])
		if err != nil {
			fmt.Print("add error:", err.Error())
		}
	case "update":
		if len(arguments) < 4 {
			fmt.Print("please enter arguments")
			return
		}
		idx, err := strconv.ParseInt(arguments[2], 10, 64)
		if err != nil {
			fmt.Print("invalid task id")
			return
		}
		err = serv.Update(int(idx), arguments[3])
		if err != nil {
			fmt.Print("update error:", err.Error())
		}

	case "delete":
		if len(arguments) < 3 {
			fmt.Print("please enter arguments")
			return
		}
		idx, err := strconv.ParseInt(arguments[2], 10, 64)
		if err != nil {
			fmt.Print("invalid task id")
			return
		}
		err = serv.Delete(int(idx))
		if err != nil {
			fmt.Print("delete error:", err.Error())
		}

	case "list":
		if len(arguments) == 2 {
			err := serv.List("")
			if err != nil {
				fmt.Print("list error:", err.Error())
			}
		} else if len(arguments) == 3 {
			err := serv.List(arguments[2])
			if err != nil {
				fmt.Print("list error:", err.Error())
			}
		} else {
			fmt.Print("invalid arguments")
		}

	case "mark-in-progress", "mark-done":
		if len(arguments) < 3 {
			fmt.Print("please enter arguments")
			return
		}
		idx, err := strconv.ParseInt(arguments[2], 10, 64)
		if err != nil {
			fmt.Print("invalid task id")
			return
		}

		status, _ := strings.CutPrefix(arguments[1], "mark-")
		err = serv.Mark(int(idx), status)
		if err != nil {
			fmt.Print("mark error:", err.Error())
		}

	default:
		print("invalid command")
	}
}
