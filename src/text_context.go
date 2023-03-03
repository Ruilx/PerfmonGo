package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const WiredTime = "2006-01-02 15:04:05"

type ReqTask struct {
	TaskID string
}

type Task struct {
	status string
	cancel context.CancelFunc
}

func generate_taskid() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)

}

func runtask(ctx context.Context) {
	var tmptaskid interface{}
	if tmptaskid = ctx.Value("TASKID"); tmptaskid != nil {
		fmt.Println("found value:", tmptaskid)
	}
	var taskid = fmt.Sprintf("%v", tmptaskid)

	for {
		select {
		default:
			fmt.Printf("Task %s is running \n", taskid)
			time.Sleep(time.Duration(1) * time.Hour)
			fmt.Printf("Task %s is ok \n", taskid)
		case <-ctx.Done():
			fmt.Printf("Task %s is stoped \n", taskid)
			return
		}
	}
}

func now() string {
	return time.Now().Format(WiredTime)
}

func main() {
	alltaskinfo := make(map[string]*Task)

	var task Task
	var ctx context.Context
	fmt.Printf("%s startTask...\n", now())
	taskId := strconv.Itoa(generate_taskid())
	task.status = "run"
	ctx, task.cancel = context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "TASKID", taskId)
	alltaskinfo[taskId] = &task
	go runtask(ctx)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("%s query...\n", now())
	task2, ok := alltaskinfo[taskId]
	if !ok {
		fmt.Printf("%s no task found.\n", now())
		panic("PANIC")
	}
	fmt.Printf("%s task %s status %s\n", now(), taskId, task2.status)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("%s stopping...\n", now())
	task3, ok := alltaskinfo[taskId]
	if !ok {
		fmt.Printf("%s no task found..\n", now())
		panic("PANIC")
	}
	task.cancel()
	task.status = "stopped"
	fmt.Printf("%s task %s status %s\n", now(), taskId, task3.status)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("%s turn start.\n", now())
	task.status = "run"
	go runtask(ctx)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("%s query...\n", now())
	task4, ok := alltaskinfo[taskId]
	if !ok {
		fmt.Printf("%s no task found.\n", now())
		panic("PANIC")
	}
	fmt.Printf("%s task %s status %s\n", now(), taskId, task4.status)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("%s stopping...\n", now())
	task5, ok := alltaskinfo[taskId]
	if !ok {
		fmt.Printf("%s no task found..\n", now())
		panic("PANIC")
	}
	task.cancel()
	task.status = "stopped"
	fmt.Printf("%s task %s status %s\n", now(), taskId, task5.status)

}
