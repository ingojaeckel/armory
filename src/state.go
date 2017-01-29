package main

var m map[string]string

func updateState(id string, newState string) {
	if m == nil {
		m = make(map[string]string)
	}
	m[id] = newState
}

func updateStateWithDebug(id string, logMessage string, newState string) {
	if m == nil {
		m = make(map[string]string)
	}
	Log(logMessage)
	m[id] = newState
}

func getState(id string) string {
	if val, exists := m[id]; exists {
		return val
	}
	return "unknown"
}
