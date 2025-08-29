# deadcss-demo

## Dev setup
This project uses a `devcontainer` to make `Go` and `Python` available. Once the devcontainer is built you can check the versions using the following commands

```bash
go version
python --version
pip --version
uv --version
```

### Package managers
The repo has been initialised to use `Go Modules`. The `python/` directory is using `uv` to manage the virtual environment and the dependencies (packages).

### Cobra-cli
To help with scaffolding tasks like initialising a new CLI app and adding new commands, the `cobra-cli` has been installed in the devcontainer.

To check if it has been installed correctly, run:
```bash
cobra-cli help
```

### Services
The python API needs to be running before it is accessible to the CLI or other tools. In one terminal window
```
cd python
uv run main.py
```

To run the `scan` command of the CLI, from project root
```
go run cmd/deadcss/main.go scan
```

Ollama should be installed on your system. The project uses the `codellama` model. You can download it via the `ollama` CLI tool
```
ollama run codellama
```
