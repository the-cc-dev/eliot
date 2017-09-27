package printers

import (
	"bytes"
	"testing"

	pb "github.com/ernoaapa/can/pkg/api/services/pods/v1"
	"github.com/ernoaapa/can/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPrintTable(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var buffer bytes.Buffer
	printer := NewHumanReadablePrinter()

	data := []*pb.Pod{
		&pb.Pod{
			Metadata: &pb.ResourceMetadata{
				Name:      "foo",
				Namespace: "cand",
			},
			Spec: &pb.PodSpec{
				Containers: []*pb.Container{
					&pb.Container{},
					&pb.Container{},
				},
			},
		},
	}

	err := printer.PrintPodsTable(data, &buffer)
	assert.NoError(t, err, "Printing pods table should not return error")

	result := buffer.String()

	assert.True(t, len(result) > 0, "Should write something to the writer")
}

func TestPrintDetails(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var buffer bytes.Buffer
	printer := NewHumanReadablePrinter()

	data := &pb.Pod{
		Metadata: &pb.ResourceMetadata{
			Name:      "foo",
			Namespace: "cand",
		},
		Spec: &pb.PodSpec{
			Containers: []*pb.Container{
				&pb.Container{},
				&pb.Container{},
			},
		},
	}

	err := printer.PrintPodDetails(data, &buffer)
	assert.NoError(t, err, "Printing pod details should not return error")

	result := buffer.String()

	assert.True(t, len(result) > 0, "Should write something to the writer")
}
func TestPrintConfig(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var buffer bytes.Buffer
	printer := NewHumanReadablePrinter()

	data := &config.Config{
		Endpoints: []config.Endpoint{
			config.Endpoint{Name: "localhost", URL: "localhost:5000"},
		},
		Users: []config.User{
			config.User{Name: "tester"},
		},
		Contexts: []config.Context{
			config.Context{Name: "local-dev", Endpoint: "localhost", User: "tester", Namespace: "default"},
		},
		CurrentContext: "local-dev",
	}

	err := printer.PrintConfig(data, &buffer)
	assert.NoError(t, err, "Printing config should not return error")

	result := buffer.String()

	assert.True(t, len(result) > 0, "Should write something to the writer")
}