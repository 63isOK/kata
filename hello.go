package main

func hello(somebody string) string {
	if len(somebody) == 0 {
		somebody = "world"
	}

	return "hello " + somebody
}
