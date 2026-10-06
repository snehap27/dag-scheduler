package main // declaring the main package

import "fmt" // importing fmt package for formatted I/O

type Task struct {
	ID string // unique identifier for the task
	// Dependencies mean tasks that MUST finish before this task can run.
	Dependencies []string // list of task IDs that this task depends on
	Action func() // function to execute the task
}

func main(){
	// 1. Create tasks with dependencies
	// 	A → B → D
	//   ↘ C ↗
	tasks := []Task{
		Task{
			ID: "A",
			Dependencies: []string{},
			Action: func(){
				fmt.Println("Executing Task A") // action for task A
			},
		},
		Task{
			ID: "B",
			Dependencies: []string{"A"},
		},
		Task{
			ID: "C",
			Dependencies: []string{"A"},
		},
		Task{
			ID: "D",
			Dependencies: []string{"B", "C"},
		},
	}
	fmt.Println(tasks)

	// 2. Create task map for easy access
	// make means initialize a new map
	// In this case, we are creating a new map to store tasks by their ID.
	// map[string]Task means that the keys of the map are strings (the task IDs) and the values are of type Task.
	// C++ euivalent: unordered_map<string,Task> taskMap
	taskMap := make(map[string]Task) // creating a map to store tasks by their ID
	// _, t := range tasks means that we are iterating over the tasks slice, and for each task, we are ignoring the index (using _) and storing the task in variable t.
	for _, t := range tasks {
		taskMap[t.ID] = t // adding each task to the map with its ID as the key
	}
	fmt.Println(taskMap) // printing the task map

	// 3. Topo Sort (kahn's)
	// using bfs -> indegree array, queue, result array
	// Dependencies mean tasks that MUST finish before this task can run.
	// eg: 
	// Task{
    // ID: "B",
    // Dependencies: []string{"A"},
	// this means B depends on A (A->B)

	// forming adjacency list
	adjList := make(map[string][]string) // creating an adjacency list to represent the graph
	for _, t := range tasks {
		for _, dependency := range t.Dependencies {
			// Add t.ID to the list of tasks that depend on dependency. This means that if dependency is completed, t.ID can be executed next.
			adjList[dependency] = append(adjList[dependency], t.ID)
		}
	}

	// forming indegree array
	indegree := make(map[string]int) // creating a map to store the indegree of each task
	for _, t := range tasks {
		indegree[t.ID] = len(t.Dependencies) // the indegree of a task is the number of tasks it depends on
	}

	// queue for bfs
	queue := []string{} // creating a queue for storing tasks (task IDs) with indegree = 0
	for _,t := range tasks {
		if indegree[t.ID] == 0 {
			queue = append(queue, t.ID) // adding tasks with indegree 0 to the queue
		}
	}

	// result array
	result := []string{} // creating a slice to store the result of the topological sort

	// bfs
	for len(queue) > 0 {
		currentTaskID := queue[0] // get the first task ID from the queue
		queue = queue[1:] // remove the first task ID from the queue

		result = append(result, currentTaskID) // add the current task ID to the result
		
		for _, dependentTaskID := range adjList[currentTaskID] { // iterate over tasks that depend on the current task
			indegree[dependentTaskID]-- // decrement the indegree of the dependent task

			if indegree[dependentTaskID] == 0 {
				queue = append(queue, dependentTaskID)
			}
		}
	}
	if len(result) != len(tasks) {
    	fmt.Println("Cycle detected")
	}
	fmt.Println(result) // print the result of the topological sort
}