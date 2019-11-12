# Go mockgen generate mock failure

This project demonstrates the failure of mockgen when re-generating the mocks for a changed interface.
When the unit tests were written and the mocks originally gneerated, everything works as expected.
But once the mocks exists, and the interface changes, then mockgen fails to generate new mocks because the tests fail to compile.


You can check this in this project using the tags "pre-change", "post-change", and "post-remove". I use the command `mockgen -source some-interface.go -package main -destination some-interface_mock.go` to generate the mock.

At the pre-change tag, the code compiles and runs the 1 unit test.

At the post-change tag, the code fails to compile, and mockgen fails to generate new mock code.

```bash
go-mock-gen-failure]$ go build
go-mock-gen-failure]$ 
go-mock-gen-failure]$ mockgen -source some-interface.go -package main -destination some-interface_mock.go
go-mock-gen-failure/some_test.go:19:34: too many arguments
go-mock-gen-failure/some_test.go:24:34: too many arguments
2019/11/11 21:59:05 Loading input failed: loading package failed
go-mock-gen-failure]$
```

When I remove the mock file to regenerate it...

```bash
go-mock-gen-failure]$ rm *_mock.go
go-mock-gen-failure]$
go-mock-gen-failure]$ mockgen -source some-interface.go -package main -destination some-interface_mock.go
go-mock-gen-failure/some_test.go:14:13: undeclared name: NewMockSome
2019/11/11 22:03:51 Loading input failed: loading package failed
go-mock-gen-failure]$
```