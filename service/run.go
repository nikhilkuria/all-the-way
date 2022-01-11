package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type run struct {
	Distance float64 `json:distance`
	Duration int32   `json:duration`
	Place    string  `json.place`
}

func postRun(c *gin.Context) {
	print("Inside postRun")
	var newRun run
	runner_id := c.Param("runner_id")
	if err := c.BindJSON(&newRun); err != nil {
		print("error here")
		log.Fatal(err)
	}
	print("binding is success")
	// TODO: Improve logging with levels
	log.Output(1, fmt.Sprintf("Posting a run for user %v ", runner_id))
	fmt.Printf("%+v\n", newRun)
}
