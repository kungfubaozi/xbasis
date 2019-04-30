package workflow

type byPriority func(f1, f2 *sequenceFlow) bool

type sequenceFlowSorter struct {
	flows []*sequenceFlow
	by    byPriority
}

func (f *sequenceFlowSorter) Len() int {
	return len(f.flows)
}

func (f *sequenceFlowSorter) Less(i, j int) bool {
	return f.by(f.flows[i], f.flows[j])
}

func (f *sequenceFlowSorter) Swap(i, j int) {
	f.flows[i], f.flows[j] = f.flows[j], f.flows[i]
}
