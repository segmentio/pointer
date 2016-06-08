package pointer

func Int(v int) *int {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func String(v string) *string {
	return &v
}

func Float(v float64) *float64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}
