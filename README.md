# Haiku LSP

## Dev setup
This project uses a `devcontainer` to make `Go` and `Nodejs` available. Once the devcontainer is built you can check the versions using the following commands

```bash
go version
node --version
```

### Package managers
The repo has been initialised to use `Go Modules`. The `vscode-extension/` directory is using `npm` to manage the virtual environment and the dependencies (packages).


### Ollama
Ollama should be installed on your system. The project uses the `codellama` model. You can download it via the `ollama` CLI tool
```
ollama run codellama
```

As a side note, some models come with variants. For instance, if you want to download the python variant, you can run the following command
```
ollama pull codellama:python
ollama run codellama:python
```

This project runs within a devcontainer, so the `ollama` CLI tool cannot be used from within the devcontainer. Instead, you can use the API to access the host machine's ollama installation. For instance, to see the version, you can run
```
curl http://host.docker.internal:11434/api/version
```

## Building server into binary

Inside the `lsp-server/` directory, build the binary

```bash
go build -o lsp-server /workspaces/haiku-lsp/lsp-server
```

Then create `.local/bin` directory
```bash
 mkdir -p $HOME/.local/bin
 ```

 Move binary into `.local/bin`
 ```bash
 mv lsp-server $HOME/.local/bin/
 ```

 Verify that it's available

```bash
which lsp-server
```

## Debugging

You can view the logs in the `Output`. To open the `Output` panel, ensure that you're in the `Extension Development Host` and go to `View > Output` in the top navbar. You should be able to select `Funchaiku Language Server` in the dropdown. Throw a bunch of `log.Println()` in the places you need to debug.

You will need to rebuild the go binary whenever you make changes, then move it to the appropriate folder using the instructions above.