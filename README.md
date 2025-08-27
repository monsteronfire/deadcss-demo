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
