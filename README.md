# Task Tracker

A command-line tool built in Golang to manage tasks. Allow you to add, update, delete, list, and mark the status of tasks. Tasks are stored in a JSON file for persistence.

Roadmap Project Challenge: https://roadmap.sh/projects/task-tracker

## Features

- Add, Update, and Delete tasks.
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

## Install
1. Ensure you have Golang installed.
2. Clone the project
```
git clone 
```
3. Build the project
```
cd task-tracker

go build -o task-tracker
```

## Usage
```
# add a Task
./task-tracker add "task descriptin"

# update a task
./task-tracker update {task_id} "updataed task description"

# delete a task
./task-tracker delete {task_id}

# mark a task as "in-progress" and "done"
./task-tracker mark-in-progress {task_id}

./task-tracker mark-done {task_id}

# list all tasks
./task-tracker list

# list tasks that status is todo
./task-tracker list todo

# list tasks that status is in-progress
./task-tracker list in-progress

# list tasks that status is done
./task-tracker list done
```

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/wsy0414/task-tracker/blob/master/LICENSE) file for details.
