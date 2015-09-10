package util

func StringSliceToSet(slice []string) Set {
	if len(slice) == 0 {
		return nil
	}
	sliceSet := NewSet()
	for _, elem := range slice {
		sliceSet.Add(elem)
	}
	return sliceSet
}
