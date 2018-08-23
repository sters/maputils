package maputils

func RecursiveFindKey(searchKey string, recursiveMap map[string]interface{}) []interface{} {
	targets := []map[string]interface{}{
		recursiveMap,
	}

	var results []interface{}

	for {
		target := targets[0]
		if val, ok := target[searchKey]; ok {
			results = append(results, val)
		}

		for _, v := range target {
			if child, ok := v.(map[string]interface{}); ok {
				targets = append(targets, child)
			}
		}

		targetsLength := len(targets)
		if targetsLength == 1 {
			break
		}
		targets = targets[1:targetsLength]
	}

	return results
}
