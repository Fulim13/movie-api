# To Run the /debug/var

go run ./cmd/api -limiter-enabled=false -port=4000

TotalAlloc — Cumulative bytes allocated on the heap (will not decrease).
HeapAlloc — Current number of bytes on the heap.
HeapObjects — Current number of objects on the heap.
Sys — Total bytes of memory obtained from the OS (i.e. total memory reserved by the Go runtime for the heap, stacks, and other internal data structures).
NumGC — Number of completed garbage collector cycles.
NextGC — The target heap size of the next garbage collector cycle (Go aims to keep HeapAlloc ≤ NextGC).
