// Code generated by skv2. DO NOT EDIT.

// Definitions for Output Snapshots
package istio

import (
	"context"
	"sort"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networking_istio_io_v1alpha3_sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of DestinationRules with a given set of labels
	DestinationRules() []LabeledDestinationRuleSet
	// return the set of EnvoyFilters with a given set of labels
	EnvoyFilters() []LabeledEnvoyFilterSet
	// return the set of VirtualServices with a given set of labels
	VirtualServices() []LabeledVirtualServiceSet

	// apply the snapshot to the cluster, garbage collecting stale resources
	Apply(ctx context.Context, clusterClient client.Client) error

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client) error
}

type snapshot struct {
	name string

	destinationRules []LabeledDestinationRuleSet
	envoyFilters     []LabeledEnvoyFilterSet
	virtualServices  []LabeledVirtualServiceSet
}

func NewSnapshot(
	name string,

	destinationRules []LabeledDestinationRuleSet,
	envoyFilters []LabeledEnvoyFilterSet,
	virtualServices []LabeledVirtualServiceSet,

) Snapshot {
	return &snapshot{
		name: name,

		destinationRules: destinationRules,
		envoyFilters:     envoyFilters,
		virtualServices:  virtualServices,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	destinationRules networking_istio_io_v1alpha3_sets.DestinationRuleSet,
	envoyFilters networking_istio_io_v1alpha3_sets.EnvoyFilterSet,
	virtualServices networking_istio_io_v1alpha3_sets.VirtualServiceSet,

) (Snapshot, error) {

	partitionedDestinationRules, err := partitionDestinationRulesByLabel(labelKey, destinationRules)
	if err != nil {
		return nil, err
	}
	partitionedEnvoyFilters, err := partitionEnvoyFiltersByLabel(labelKey, envoyFilters)
	if err != nil {
		return nil, err
	}
	partitionedVirtualServices, err := partitionVirtualServicesByLabel(labelKey, virtualServices)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedDestinationRules,
		partitionedEnvoyFilters,
		partitionedVirtualServices,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) Apply(ctx context.Context, cli client.Client) error {
	var genericLists []output.ResourceList

	for _, outputSet := range s.destinationRules {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.envoyFilters {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.virtualServices {
		genericLists = append(genericLists, outputSet.Generic())
	}

	return output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.Sync(ctx, cli)
}

// apply the desired resources to multiple cluster states; remove stale resources where necessary
func (s *snapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client) error {
	var genericLists []output.ResourceList

	for _, outputSet := range s.destinationRules {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.envoyFilters {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.virtualServices {
		genericLists = append(genericLists, outputSet.Generic())
	}

	return output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient)
}

func partitionDestinationRulesByLabel(labelKey string, set networking_istio_io_v1alpha3_sets.DestinationRuleSet) ([]LabeledDestinationRuleSet, error) {
	setsByLabel := map[string]networking_istio_io_v1alpha3_sets.DestinationRuleSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "DestinationRule", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "DestinationRule", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = networking_istio_io_v1alpha3_sets.NewDestinationRuleSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedDestinationRules []LabeledDestinationRuleSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledDestinationRuleSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedDestinationRules = append(partitionedDestinationRules, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedDestinationRules, func(i, j int) bool {
		leftLabelValue := partitionedDestinationRules[i].Labels()[labelKey]
		rightLabelValue := partitionedDestinationRules[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedDestinationRules, nil
}

func partitionEnvoyFiltersByLabel(labelKey string, set networking_istio_io_v1alpha3_sets.EnvoyFilterSet) ([]LabeledEnvoyFilterSet, error) {
	setsByLabel := map[string]networking_istio_io_v1alpha3_sets.EnvoyFilterSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "EnvoyFilter", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "EnvoyFilter", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = networking_istio_io_v1alpha3_sets.NewEnvoyFilterSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedEnvoyFilters []LabeledEnvoyFilterSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledEnvoyFilterSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedEnvoyFilters = append(partitionedEnvoyFilters, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedEnvoyFilters, func(i, j int) bool {
		leftLabelValue := partitionedEnvoyFilters[i].Labels()[labelKey]
		rightLabelValue := partitionedEnvoyFilters[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedEnvoyFilters, nil
}

func partitionVirtualServicesByLabel(labelKey string, set networking_istio_io_v1alpha3_sets.VirtualServiceSet) ([]LabeledVirtualServiceSet, error) {
	setsByLabel := map[string]networking_istio_io_v1alpha3_sets.VirtualServiceSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "VirtualService", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "VirtualService", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = networking_istio_io_v1alpha3_sets.NewVirtualServiceSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedVirtualServices []LabeledVirtualServiceSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledVirtualServiceSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedVirtualServices = append(partitionedVirtualServices, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedVirtualServices, func(i, j int) bool {
		leftLabelValue := partitionedVirtualServices[i].Labels()[labelKey]
		rightLabelValue := partitionedVirtualServices[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedVirtualServices, nil
}

func (s snapshot) DestinationRules() []LabeledDestinationRuleSet {
	return s.destinationRules
}

func (s snapshot) EnvoyFilters() []LabeledEnvoyFilterSet {
	return s.envoyFilters
}

func (s snapshot) VirtualServices() []LabeledVirtualServiceSet {
	return s.virtualServices
}

// LabeledDestinationRuleSet represents a set of destinationRules
// which share a common set of labels.
// These labels are used to find diffs between DestinationRuleSets.
type LabeledDestinationRuleSet interface {
	// returns the set of Labels shared by this DestinationRuleSet
	Labels() map[string]string

	// returns the set of DestinationRulees with the given labels
	Set() networking_istio_io_v1alpha3_sets.DestinationRuleSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledDestinationRuleSet struct {
	set    networking_istio_io_v1alpha3_sets.DestinationRuleSet
	labels map[string]string
}

func NewLabeledDestinationRuleSet(set networking_istio_io_v1alpha3_sets.DestinationRuleSet, labels map[string]string) (LabeledDestinationRuleSet, error) {
	// validate that each DestinationRule contains the labels, else this is not a valid LabeledDestinationRuleSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on DestinationRule %v", k, v, item.Name)
			}
		}
	}

	return &labeledDestinationRuleSet{set: set, labels: labels}, nil
}

func (l *labeledDestinationRuleSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledDestinationRuleSet) Set() networking_istio_io_v1alpha3_sets.DestinationRuleSet {
	return l.set
}

func (l labeledDestinationRuleSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list networking_istio_io_v1alpha3.DestinationRuleList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources: desiredResources,
		ListFunc:  listFunc,
	}
}

// LabeledEnvoyFilterSet represents a set of envoyFilters
// which share a common set of labels.
// These labels are used to find diffs between EnvoyFilterSets.
type LabeledEnvoyFilterSet interface {
	// returns the set of Labels shared by this EnvoyFilterSet
	Labels() map[string]string

	// returns the set of EnvoyFilteres with the given labels
	Set() networking_istio_io_v1alpha3_sets.EnvoyFilterSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledEnvoyFilterSet struct {
	set    networking_istio_io_v1alpha3_sets.EnvoyFilterSet
	labels map[string]string
}

func NewLabeledEnvoyFilterSet(set networking_istio_io_v1alpha3_sets.EnvoyFilterSet, labels map[string]string) (LabeledEnvoyFilterSet, error) {
	// validate that each EnvoyFilter contains the labels, else this is not a valid LabeledEnvoyFilterSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on EnvoyFilter %v", k, v, item.Name)
			}
		}
	}

	return &labeledEnvoyFilterSet{set: set, labels: labels}, nil
}

func (l *labeledEnvoyFilterSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledEnvoyFilterSet) Set() networking_istio_io_v1alpha3_sets.EnvoyFilterSet {
	return l.set
}

func (l labeledEnvoyFilterSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list networking_istio_io_v1alpha3.EnvoyFilterList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources: desiredResources,
		ListFunc:  listFunc,
	}
}

// LabeledVirtualServiceSet represents a set of virtualServices
// which share a common set of labels.
// These labels are used to find diffs between VirtualServiceSets.
type LabeledVirtualServiceSet interface {
	// returns the set of Labels shared by this VirtualServiceSet
	Labels() map[string]string

	// returns the set of VirtualServicees with the given labels
	Set() networking_istio_io_v1alpha3_sets.VirtualServiceSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledVirtualServiceSet struct {
	set    networking_istio_io_v1alpha3_sets.VirtualServiceSet
	labels map[string]string
}

func NewLabeledVirtualServiceSet(set networking_istio_io_v1alpha3_sets.VirtualServiceSet, labels map[string]string) (LabeledVirtualServiceSet, error) {
	// validate that each VirtualService contains the labels, else this is not a valid LabeledVirtualServiceSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on VirtualService %v", k, v, item.Name)
			}
		}
	}

	return &labeledVirtualServiceSet{set: set, labels: labels}, nil
}

func (l *labeledVirtualServiceSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledVirtualServiceSet) Set() networking_istio_io_v1alpha3_sets.VirtualServiceSet {
	return l.set
}

func (l labeledVirtualServiceSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list networking_istio_io_v1alpha3.VirtualServiceList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources: desiredResources,
		ListFunc:  listFunc,
	}
}
