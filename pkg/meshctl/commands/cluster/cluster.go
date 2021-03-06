package cluster

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/cluster/deregister"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/cluster/register"
	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Interacting with remote Kubernetes clusters registered to Gloo Mesh",
	}

	cmd.AddCommand(
		register.Command(ctx),
		deregister.Command(ctx),
	)

	return cmd
}
