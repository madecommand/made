package main
//
//import "fmt"
//
//const root = "  root!"
//
//type DepTree struct {
//	deps map[string][]string
//}
//
//func (dt *DepTree) Add(source string, dest ...string) error {
//	_, ok := dt.deps[source]
//	if ok {
//		return fmt.Errorf("task %q was already added", source)
//	}
//	dt.deps[source] = dest
//
//	if _, ok := dt.deps[root]; !ok {
//		dt.deps[root] = []string{}
//	}
//	dt.deps[root] = append(dt.deps[root], source)
//
//	return nil
//}
//
//func (dt *DepTree) Flatten(task ...string) ([]string, error) {
//	if len(task) == 0 {
//		return []string{}, nil
//	}
//	r, ok := dt.deps[root]
//	if !ok {
//		return []string{}, fmt.Errorf("no tasks defined")
//	}
//
//	for _, deps := 
//
//	return []string{}
//}
//