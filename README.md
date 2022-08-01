# tinygo-devcontainer

This is an example of setting up TinyGo via VSCode's devcontainer.

## How to use (with VSCode)

First docker pull tinygo/tinygo-dev as it is used internally.

```
$ docker pull tinygo/tinygo-dev
```

Install [Visual Studio Code Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

Enable remote container with `Remote-Containers: Reopen in Container`

![](./images/reopen.png)

### examples/wasm

The examples in examples/wasm are very easy to execute.

```
$ cd ./examples/wasm/
$ make slices
$ go run ./server.go
```

![](./images/examples-wasm.png)

#### lsp not working properly

1. `TinyGo target` and select `wasm`
2. Add GOOS/GOARCH to `./.vscode/settings.json`
3. Reload window (`>Developer: Reload Window`)

```
{
    "go.toolsEnvVars": {
        "GOOS": "js",
        "GOARCH": "wasm",
        "GOROOT": "...",
        "GOFLAGS": "...",
    }
}
```

The following PRs attempt to improve the situation.

* https://github.com/tinygo-org/vscode-tinygo/pull/6

## Internals

The following are used internally

* https://golang.org/x/tools/gopls
* https://golang.org/x/tools/cmd/goimports
* https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo
    * `TinyGo target` to enable LSP  
    ![](./images/lsp.png)
* https://github.com/sago35/tinygo-autocmpl
    * You can use tinygo's bash completion  
    ![](./images/tinygo-autocmpl.gif)

The configuration file can be found at

* [./.devcontainer](./.devcontainer)

## LICENSE

MIT

## Author

sago35 - <sago35@gmail.com>
