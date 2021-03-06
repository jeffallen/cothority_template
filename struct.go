package template

/*
This holds the messages used to communicate with the service over the network.
*/

import (
	"github.com/dedis/onet"
	"github.com/dedis/onet/network"
)

// We need to register all messages so the network knows how to handle them.
func init() {
	network.RegisterMessages(
		CountRequest{}, CountResponse{},
		ClockRequest{}, ClockResponse{},
	)
}

const (
	// ErrorParse indicates an error while parsing the protobuf-file.
	ErrorParse = iota + 4000
)

// ClockRequest will run the tepmlate-protocol on the roster and return
// the time spent doing so.
type ClockRequest struct {
	Roster *onet.Roster
}

// ClockResponse returns the time spent for the protocol-run.
type ClockResponse struct {
	Time     float64
	Children int
}

// CountRequest will return how many times the protocol has been run.
type CountRequest struct {
}

// CountResponse returns the number of protocol-runs
type CountResponse struct {
	Count int
}
