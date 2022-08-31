 to test run go programs, run 'go run <file-name>'
 to build a compiled version of the go program, run 'go build <file-name>
 to save your build under a custom name, use the '-o' flag:
	 ex: go build -o hello_world hello.go
	 ^^This saves the build to a file name 'hello_world'

 you can use the 'go install' command to install software
 you simply write 'go install' then point it toward the location of 
 the source code repositor, followed by @ to specify the version
 If you want the latest version, just write @latest.
 to update to the latest version, rerun the install command
 pointing to the appropriate endpoint and specify the newer
 version or @latest.

 To run compiled go programs, write ./<compiled-program-name>

 go uses explicit formatting. To format your file
 according to go standards, use the command 'go fmt'.

To clean up your import statements, us goimports. This can be
 installed with go install golang.org/x/tools/cmd/goimports@latest
 and can be run with goimports -l -w.

 goimports will remove unused imports, alphabetize imports,
 and fill in unspecified imports. These guesses are sometimes
 wrong, so it's better to write in all your imports yourself.

 to run this command, run goimports -l -w .

 Go requires semicolons at the end of every statement.
 However, developers don't put these semicolons in themselves
 - the compiler handles it for them.

 Code linting tools are also used to search for bugs and formatting errors
 While golint was the standard for a time, it's been deprecated.
 New tools are Staticcheck and revive. 

 To install staticcheck, run 'go install honnef.co/go/tools/cmd/staticcheck@latest'
 To run staticcheck, simply run staticcheck . This will check the current package.
 To check all packages, run staticcheck ./...

 To test new version of go before implementing them run ` go get golang.org/dl/go.1.15.6
$ go1.15.6 download` , where 1.15.6 is a stand-in for the new version.

You can then use go1.15.6 build as a command instead of go build.

To remove this new version, run 
```
go1.15.6 env GOROOT
/Users/gobook/sdk/go1.15.6
$ rm -rf $(go1.15.6 env GOROOT)
$ rm $(go env GOPATH)/bin/go1.15.6
```

You can use homebrew to update to the latest go version.

