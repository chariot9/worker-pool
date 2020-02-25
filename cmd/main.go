package main

import (
	"fmt"
	"github.com/chariot9/worker-pool/pool"
	"github.com/chariot9/worker-pool/result"
)

func ResourceProcessor(resource interface{}) error {
	fmt.Printf("Resource processor got: %s\n", resource)
	return nil
}

func ResultProcessor(result result.Result) error {
	fmt.Printf("Result processor got: %d\n", result.Job.Id)
	return nil
}

func main() {
	strings := []string{"first", "second"}
	resources := make([]interface{}, len(strings))
	for i, s := range strings {
		resources[i] = s
	}

	p := pool.NewPool(3)
	p.Start(resources, ResourceProcessor, ResultProcessor)
}
