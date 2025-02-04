A minimal http web server just using [net/http](https://pkg.go.dev/net/http) and [fmt](https://pkg.go.dev/fmt) from Go standard library.

For dynamic reloading install [modd](https://github.com/cortesi/modd/)

`go install github.com/cortesi/modd/cmd/modd@latest`

And launch it within this directory

`modd`

Note: 
- Beware that if your code doesn't compile `modd` will stick with the previous version.
- It can be confusing. If you have any doubt just restart it.
