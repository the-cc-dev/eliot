package main

import (
	"fmt"
	"os"

	"github.com/ernoaapa/can/cmd"
	pb "github.com/ernoaapa/can/pkg/api/services/pods/v1"
	"github.com/ernoaapa/can/pkg/printers"
	"github.com/ernoaapa/can/pkg/term"
	"github.com/urfave/cli"
)

var runCommandHelp = `
	# Start new container
	can-cli run my-pod docker.io/ernoaapa/hello-world:latest
`

var runCommand = cli.Command{
	Name:        "run",
	HelpName:    "run",
	Usage:       "Start container in the device",
	Description: "With run command, you can start new containers in the device",
	UsageText: `can-cli run [options] NAME

	 # Start new container
	 can-cli run --image docker.io/eaapa/hello-world:latest my-pod
`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "image",
			Usage: "The container image to start",
		},
		cli.BoolFlag{
			Name:  "detach, d",
			Usage: "Run container in background and print container information",
		},
		cli.StringSliceFlag{
			Name:  "mount",
			Usage: "Attach a filesystem mount to the container",
		},
		cli.StringSliceFlag{
			Name:  "bind",
			Usage: "Bind a directory in host to the container. Format: /source:/target:options, E.g. /var:/var:rshared",
		},
		cli.StringSliceFlag{
			Name:  "env, e",
			Usage: "Set environment variables",
		},
		cli.BoolFlag{
			Name:  "rm",
			Usage: "Automatically remove the container when it exits",
		},
		cli.BoolFlag{
			Name:  "tty, t",
			Usage: "Allocate TTY for each container in the pod",
		},
		cli.BoolFlag{
			Name:  "stdin, i",
			Usage: "Keep stdin open on the container(s) in the pod, even if nothing is attached.",
		},
		cli.StringFlag{
			Name:  "workdir, w",
			Usage: "Working directory inside the container",
		},
	},
	Action: func(clicontext *cli.Context) error {
		var (
			name    = clicontext.Args().First()
			image   = clicontext.String("image")
			detach  = clicontext.Bool("detach")
			rm      = clicontext.Bool("rm")
			tty     = clicontext.Bool("tty")
			env     = clicontext.StringSlice("env")
			workdir = clicontext.String("workdir")
			mounts  = cmd.GetMounts(clicontext)
			binds   = cmd.GetBinds(clicontext)
			args    = []string{}
			stdin   = os.Stdin
			stdout  = os.Stdout
			stderr  = os.Stderr
		)

		if clicontext.NArg() > 1 {
			args = clicontext.Args()[1:]
		}

		if name == "" {
			return fmt.Errorf("You must give NAME parameter")
		}

		if image == "" {
			return fmt.Errorf("You must define --image option")
		}
		image = cmd.ExpandToFQIN(image)

		if detach && rm {
			return fmt.Errorf("You cannot use --detach flag with --rm, it would remove right away after container started")
		}

		config := cmd.GetConfig(clicontext)
		client := cmd.GetClient(clicontext)

		pod := &pb.Pod{
			Metadata: &pb.ResourceMetadata{
				Name:      name,
				Namespace: config.GetCurrentContext().Namespace,
			},
			Spec: &pb.PodSpec{
				HostNetwork: true,
				Containers: []*pb.Container{
					&pb.Container{
						Name:       name,
						Image:      image,
						Tty:        tty,
						Args:       args,
						Env:        env,
						WorkingDir: workdir,
						Mounts:     append(mounts, binds...),
					},
				},
			},
		}

		if rm {
			defer client.DeletePod(pod)
		}

		result, err := client.CreatePod(pod)
		if err != nil {
			return err
		}

		if detach {
			writer := printers.GetNewTabWriter(os.Stdout)
			printer := cmd.GetPrinter(clicontext)
			return printer.PrintPodDetails(result, writer)
		}

		term := term.TTY{
			Out: stdout,
		}

		if clicontext.Bool("stdin") {
			term.In = stdin
			term.Raw = true
		} else {
			stdin = nil
		}

		return term.Safe(func() error {
			return client.Attach(result.Spec.Containers[0].Name, stdin, stdout, stderr)
		})
	},
}
