package util

func StringSliceToSet(hosts []string) Set {
	if len(hosts) == 0 {
		return nil
	}
	hostSet := NewSet()
	for _, host := range hosts {
		hostSet.Add(host)
	}
	return hostSet
}
