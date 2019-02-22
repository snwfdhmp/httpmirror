package main

func unmap(m map[interface{}]interface{}, fn func(name, value interface{}) (stop bool, returnValue interface{})) (returnValue interface{}) {
	for n, v := range m {
		if stop, retVal := fn(n, v); stop {
			return retVal
		}
	}

	return nil
}

func mapreduce(m map[interface{}]interface{}, fn func(name, value, result interface{}) (newResult interface{})) interface{} {
	var result interface{}
	for n, v := range m {
		result = fn(n, v, result)
	}

	return result
}
