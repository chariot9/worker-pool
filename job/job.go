package job

type Job struct {
	Id       int
	Resource interface{}
}

type ProcessJob func(resource interface{}) error
