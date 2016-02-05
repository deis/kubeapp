package pod

import (
	"bufio"

	"github.com/jchauncey/kubeclient/api"
)

// LogStreamer is just the GetLogs func from github.com/jchauncey/kubeclient/api.PodInterface
type LogStreamer interface {
	// GetLogs gets the logs for the given pod name, according to the given options
	GetLogs(podName string, opts *api.PodLogOptions) *api.Request
}

// Stream streams logs for the given pod from the given streamer, according to the given opts.
//
// The stream will start in a goroutine and the returned channel will receive all logs until they stop. If there was an error starting the stream, the returned channel will be closed and a non-nil error will be returned.
func Stream(streamer LogStreamer, podName string, opts *api.PodLogOptions) (<-chan string, <-chan error) {
	retCh := make(chan string)
	errCh := make(chan error)
	rc, err := streamer.GetLogs().Stream()
	if err != nil {
		go func() {
			errCh <- err
		}()
		return retCh, errCh
	}
	go func() {
		defer func() {
			// TODO: don't swallow the error returned by Close
			rc.Close()
		}()
		scanner := bufio.NewScanner(rc)
		for scanner.Scan() {
			retCh <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			errCh <- err
		}
	}()
	return retCh, errCh
}
