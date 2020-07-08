// Code generated by skv2. DO NOT EDIT.

// Definitions for Output Snapshots
package output

import (
	"context"
	"sort"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	discovery_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	discovery_smh_solo_io_v1alpha1_sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/sets"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of MeshServices with a given set of labels
	MeshServices() []LabeledMeshServiceSet
	// return the set of MeshWorkloads with a given set of labels
	MeshWorkloads() []LabeledMeshWorkloadSet
	// return the set of Meshes with a given set of labels
	Meshes() []LabeledMeshSet

	// apply the snapshot to the cluster, garbage collecting stale resources
	Apply(ctx context.Context, clusterClient client.Client) error

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client) error
}

type snapshot struct {
	name string

	meshServices  []LabeledMeshServiceSet
	meshWorkloads []LabeledMeshWorkloadSet
	meshes        []LabeledMeshSet
}

func NewSnapshot(
	name string,

	meshServices []LabeledMeshServiceSet,
	meshWorkloads []LabeledMeshWorkloadSet,
	meshes []LabeledMeshSet,

) Snapshot {
	return &snapshot{
		name: name,

		meshServices:  meshServices,
		meshWorkloads: meshWorkloads,
		meshes:        meshes,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	meshServices discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet,
	meshWorkloads discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet,
	meshes discovery_smh_solo_io_v1alpha1_sets.MeshSet,

) (Snapshot, error) {

	partitionedMeshServices, err := partitionMeshServicesByLabel(labelKey, meshServices)
	if err != nil {
		return nil, err
	}
	partitionedMeshWorkloads, err := partitionMeshWorkloadsByLabel(labelKey, meshWorkloads)
	if err != nil {
		return nil, err
	}
	partitionedMeshes, err := partitionMeshesByLabel(labelKey, meshes)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedMeshServices,
		partitionedMeshWorkloads,
		partitionedMeshes,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) Apply(ctx context.Context, cli client.Client) error {
	var genericLists []output.ResourceList

	for _, outputSet := range s.meshServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshWorkloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
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

	for _, outputSet := range s.meshServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshWorkloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
		genericLists = append(genericLists, outputSet.Generic())
	}

	return output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient)
}

func partitionMeshServicesByLabel(labelKey string, set discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet) ([]LabeledMeshServiceSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "MeshService", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "MeshService", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha1_sets.NewMeshServiceSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshServices []LabeledMeshServiceSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshServiceSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshServices = append(partitionedMeshServices, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshServices, func(i, j int) bool {
		leftLabelValue := partitionedMeshServices[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshServices[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshServices, nil
}

func partitionMeshWorkloadsByLabel(labelKey string, set discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet) ([]LabeledMeshWorkloadSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "MeshWorkload", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "MeshWorkload", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha1_sets.NewMeshWorkloadSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshWorkloads []LabeledMeshWorkloadSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshWorkloadSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshWorkloads = append(partitionedMeshWorkloads, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshWorkloads, func(i, j int) bool {
		leftLabelValue := partitionedMeshWorkloads[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshWorkloads[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshWorkloads, nil
}

func partitionMeshesByLabel(labelKey string, set discovery_smh_solo_io_v1alpha1_sets.MeshSet) ([]LabeledMeshSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha1_sets.MeshSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha1_sets.NewMeshSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshes []LabeledMeshSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshes = append(partitionedMeshes, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshes, func(i, j int) bool {
		leftLabelValue := partitionedMeshes[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshes[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshes, nil
}

func (s snapshot) MeshServices() []LabeledMeshServiceSet {
	return s.meshServices
}

func (s snapshot) MeshWorkloads() []LabeledMeshWorkloadSet {
	return s.meshWorkloads
}

func (s snapshot) Meshes() []LabeledMeshSet {
	return s.meshes
}

// LabeledMeshServiceSet represents a set of meshServices
// which share a common set of labels.
// These labels are used to find diffs between MeshServiceSets.
type LabeledMeshServiceSet interface {
	// returns the set of Labels shared by this MeshServiceSet
	Labels() map[string]string

	// returns the set of MeshServicees with the given labels
	Set() discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshServiceSet struct {
	set    discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet
	labels map[string]string
}

func NewLabeledMeshServiceSet(set discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet, labels map[string]string) (LabeledMeshServiceSet, error) {
	// validate that each MeshService contains the labels, else this is not a valid LabeledMeshServiceSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on MeshService %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshServiceSet{set: set, labels: labels}, nil
}

func (l *labeledMeshServiceSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshServiceSet) Set() discovery_smh_solo_io_v1alpha1_sets.MeshServiceSet {
	return l.set
}

func (l labeledMeshServiceSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha1.MeshServiceList
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

// LabeledMeshWorkloadSet represents a set of meshWorkloads
// which share a common set of labels.
// These labels are used to find diffs between MeshWorkloadSets.
type LabeledMeshWorkloadSet interface {
	// returns the set of Labels shared by this MeshWorkloadSet
	Labels() map[string]string

	// returns the set of MeshWorkloades with the given labels
	Set() discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshWorkloadSet struct {
	set    discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet
	labels map[string]string
}

func NewLabeledMeshWorkloadSet(set discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet, labels map[string]string) (LabeledMeshWorkloadSet, error) {
	// validate that each MeshWorkload contains the labels, else this is not a valid LabeledMeshWorkloadSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on MeshWorkload %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshWorkloadSet{set: set, labels: labels}, nil
}

func (l *labeledMeshWorkloadSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshWorkloadSet) Set() discovery_smh_solo_io_v1alpha1_sets.MeshWorkloadSet {
	return l.set
}

func (l labeledMeshWorkloadSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha1.MeshWorkloadList
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

// LabeledMeshSet represents a set of meshes
// which share a common set of labels.
// These labels are used to find diffs between MeshSets.
type LabeledMeshSet interface {
	// returns the set of Labels shared by this MeshSet
	Labels() map[string]string

	// returns the set of Meshes with the given labels
	Set() discovery_smh_solo_io_v1alpha1_sets.MeshSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshSet struct {
	set    discovery_smh_solo_io_v1alpha1_sets.MeshSet
	labels map[string]string
}

func NewLabeledMeshSet(set discovery_smh_solo_io_v1alpha1_sets.MeshSet, labels map[string]string) (LabeledMeshSet, error) {
	// validate that each Mesh contains the labels, else this is not a valid LabeledMeshSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Mesh %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshSet{set: set, labels: labels}, nil
}

func (l *labeledMeshSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshSet) Set() discovery_smh_solo_io_v1alpha1_sets.MeshSet {
	return l.set
}

func (l labeledMeshSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha1.MeshList
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
