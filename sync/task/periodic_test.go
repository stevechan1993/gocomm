package task

import (
	"github.com/stevechan1993/gocomm/common"
	"log"
	"time"
)

func ExamplePeriodic() {
	count := 0
	task := NewPeriodic(time.Second*2, func() error {
		count++
		log.Println("current count:", count)
		return nil
	})
	common.Must(task.Start())
	time.Sleep(time.Second * 5)
	common.Must(task.Close())
	log.Println("Count:", count)
	common.Must(task.Start())
	time.Sleep(time.Second * 5)
	log.Println("Count:", count)
	common.Must(task.Close())
}
