package translation

import (
	"bytes"
	"context"
	"fmt"

	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/appmesh"

	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	appmeshoutput "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/appmesh"
	istiooutput "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	localoutput "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/local"
	smioutput "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/smi"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/osm"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type OutputSnapshots struct {
	istio   istiooutput.Snapshot
	appmesh appmeshoutput.Snapshot
	smi     smioutput.Snapshot
	local   localoutput.Snapshot
}

func (t OutputSnapshots) MarshalJSON() ([]byte, error) {
	istioByt, err := t.istio.MarshalJSON()
	if err != nil {
		return nil, err
	}
	appmeshByt, err := t.appmesh.MarshalJSON()
	if err != nil {
		return nil, err
	}
	smiByt, err := t.smi.MarshalJSON()
	if err != nil {
		return nil, err
	}
	localByt, err := t.local.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return bytes.Join([][]byte{istioByt, appmeshByt, smiByt, localByt}, []byte("\n")), nil
}

func (t OutputSnapshots) Apply(
	ctx context.Context,
	clusterClient client.Client,
	multiClusterClient multicluster.Client,
	errHandler output.ErrorHandler,
) {
	// Apply mesh resources to registered clusters
	t.istio.ApplyMultiCluster(ctx, multiClusterClient, errHandler)
	t.appmesh.ApplyMultiCluster(ctx, multiClusterClient, errHandler)
	t.smi.ApplyMultiCluster(ctx, multiClusterClient, errHandler)
	// Apply local resources only to management cluster
	t.local.ApplyLocalCluster(ctx, clusterClient, errHandler)
}

// the networking translator translates an istio input networking snapshot to an istiooutput snapshot of mesh config resources
type Translator interface {
	// errors reflect an internal translation error and should never happen
	Translate(
		ctx context.Context,
		in input.Snapshot,
		reporter reporting.Reporter,
	) (OutputSnapshots, error)
}

type translator struct {
	totalTranslates   int // TODO(ilackarms): metric
	istioTranslator   istio.Translator
	appmeshTranslator appmesh.Translator
	osmTranslator     osm.Translator
}

func NewTranslator(
	istioTranslator istio.Translator,
	appmeshTranslator appmesh.Translator,
	osmTranslator osm.Translator,
) Translator {
	return &translator{
		istioTranslator:   istioTranslator,
		appmeshTranslator: appmeshTranslator,
		osmTranslator:     osmTranslator,
	}
}

func (t *translator) Translate(
	ctx context.Context,
	in input.Snapshot,
	reporter reporting.Reporter,
) (OutputSnapshots, error) {
	t.totalTranslates++
	ctx = contextutils.WithLogger(ctx, fmt.Sprintf("translation-%v", t.totalTranslates))

	istioOutputs := istiooutput.NewBuilder(ctx, fmt.Sprintf("networking-istio-%v", t.totalTranslates))
	appmeshOutputs := appmeshoutput.NewBuilder(ctx, fmt.Sprintf("networking-appmesh-%v", t.totalTranslates))
	smiOutputs := smioutput.NewBuilder(ctx, fmt.Sprintf("networking-smi-%v", t.totalTranslates))
	localOutputs := localoutput.NewBuilder(ctx, fmt.Sprintf("networking-local-%v", t.totalTranslates))

	t.istioTranslator.Translate(ctx, in, istioOutputs, localOutputs, reporter)

	t.appmeshTranslator.Translate(ctx, in, appmeshOutputs, reporter)

	t.osmTranslator.Translate(ctx, in, smiOutputs, reporter)

	istioSnapshot, err := istioOutputs.BuildSinglePartitionedSnapshot(metautils.TranslatedObjectLabels())
	if err != nil {
		return OutputSnapshots{}, err
	}

	appmeshSnapshot, err := appmeshOutputs.BuildSinglePartitionedSnapshot(metautils.TranslatedObjectLabels())
	if err != nil {
		return OutputSnapshots{}, err
	}

	smiSnapshot, err := smiOutputs.BuildSinglePartitionedSnapshot(metautils.TranslatedObjectLabels())
	if err != nil {
		return OutputSnapshots{}, err
	}

	localSnapshot, err := localOutputs.BuildSinglePartitionedSnapshot(metautils.TranslatedObjectLabels())
	if err != nil {
		return OutputSnapshots{}, err
	}

	return OutputSnapshots{
		istio:   istioSnapshot,
		appmesh: appmeshSnapshot,
		smi:     smiSnapshot,
		local:   localSnapshot,
	}, nil
}
