package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.", time.Now())
}

func superWang() {
	fmt.Println("I am running superWang.", time.Now())
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	//c1()
	c2()
}

func c1() {
	//每隔1秒执行一个任务，每隔4秒执行另一个任务
	s := gocron.NewScheduler()
	s.Every(1).Seconds().Do(task)
	s.Every(4).Seconds().Do(superWang)

	sc := s.Start() // keep the channel
	go test(s, sc)  // wait
	<-sc            // it will happen if the channel is closed
}

func test(s *gocron.Scheduler, sc chan bool) {
	time.Sleep(8 * time.Second)
	s.Remove(task) //remove task
	time.Sleep(8 * time.Second)
	s.Clear()
	fmt.Println("All task removed")
	close(sc) // close the channel
}

func c2() {
	// Do jobs without params
	gocron.Every(1).Second().Do(task)
	gocron.Every(2).Seconds().Do(task)
	gocron.Every(1).Minute().Do(task)
	gocron.Every(2).Minutes().Do(task)
	gocron.Every(1).Hour().Do(task)
	gocron.Every(2).Hours().Do(task)
	gocron.Every(1).Day().Do(task)
	gocron.Every(2).Days().Do(task)
	gocron.Every(1).Week().Do(task)
	gocron.Every(2).Weeks().Do(task)

	// Do jobs with params
	gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs on specific weekday
	gocron.Every(1).Monday().Do(task)
	gocron.Every(1).Thursday().Do(task)

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	gocron.Every(1).Day().At("10:30").Do(task)
	gocron.Every(1).Monday().At("18:30").Do(task)
	gocron.Every(1).Tuesday().At("18:30:59").Do(task)

	// Begin job immediately upon start
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)

	// Begin job at a specific date/time
	t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	gocron.Every(1).Hour().From(&t).Do(task)

	// NextRun gets the next running time
	_, time := gocron.NextRun()
	fmt.Println(time)

	// Remove a specific job
	gocron.Remove(task)

	// Clear all scheduled jobs
	gocron.Clear()

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<-s.Start()
}
