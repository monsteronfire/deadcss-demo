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

To download the python variant, you can run the following command
```
ollama pull codellama:python
ollama run codellama:python
```

This project runs within a devcontainer, so the `ollama` CLI tool cannot be used from within the devcontainer. Instead, you can use the API to access the host machine's ollama installation. For instance, to see the version, you can run
```
curl http://host.docker.internal:11434/api/version
```
