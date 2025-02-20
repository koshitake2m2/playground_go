module example.com/main

go 1.22.5

replace example.com/aaa => ../aaa

replace example.com/bbb => ../bbb

replace example.com/ccc => ../ccc

require example.com/bbb v0.0.0-00010101000000-000000000000

require (
	example.com/aaa v0.0.0-00010101000000-000000000000
	example.com/ccc v0.0.0-00010101000000-000000000000 // indirect
)
