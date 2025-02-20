module example.com/ccc

go 1.22.5

// You don't need to use aaa module, but you should require it. Because bbb module requires it.
replace example.com/aaa => ../aaa

replace example.com/bbb => ../bbb

require example.com/bbb v0.0.0-00010101000000-000000000000

require example.com/aaa v0.0.0-00010101000000-000000000000 // indirect
