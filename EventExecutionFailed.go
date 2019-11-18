package eventago

type EventExecutionFailed struct {
	service   string
	exception string
	event     string
}
