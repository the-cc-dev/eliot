package printers

import (
	"io"

	pb "github.com/ernoaapa/can/pkg/api/services/pods/v1"
	"github.com/ernoaapa/can/pkg/config"
)

// ResourcePrinter is an interface that knows how to print runtime objects.
type ResourcePrinter interface {
	// PrintPodsTable receives a list of Pods, formats it as table and prints it to a writer.
	PrintPodsTable([]*pb.Pod, io.Writer) error
	// PrintPodDetails receives a list of Pods, formats it to detailed description and prints it to the writer.
	PrintPodDetails(*pb.Pod, io.Writer) error
	// PrintConfig receives config and formats it to human readable format
	PrintConfig(*config.Config, io.Writer) error
}