# Silent Map Access in Go: Unexpected Zero Values

This repository demonstrates a common pitfall in Go: accessing a non-existent key in a map doesn't result in a runtime panic. Instead, it silently returns the zero value for the map's value type. This can lead to subtle bugs if not handled correctly.

**Problem:**
The `bug.go` file shows a simple example.  Accessing the key "foo" in an uninitialized map returns 0, without any error or indication of the missing key.  This behavior might be unexpected for programmers coming from languages that throw an exception in similar cases.

**Solution:**
The `bugSolution.go` file shows a better approach by explicitly checking if the key exists before accessing it using the `comma ok` idiom.  This allows for proper handling of missing keys and avoids unexpected results.