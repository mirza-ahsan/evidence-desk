# Evidence Desk

Server-rendered Python app for case evidence. Python 3.12, FastAPI, and SQLite. No Docker.

## Setup in VS Code

Install Python 3.12 and the VS Code Python extension, then open this folder. On Windows, `py -3.12` is the Python 3.12 launcher; elsewhere the command is `python3.12`.

### uv

Install [uv](https://docs.astral.sh/uv/getting-started/installation/), then:

```bash
uv sync --all-groups
```

Command Palette → **Python: Select Interpreter** → `.venv/bin/python` (`.venv\Scripts\python.exe` on Windows).

```bash
uv run uvicorn evidence_desk.app:app --reload
```

### pip

```bash
python3.12 -m venv .venv
```

Activate it: `source .venv/bin/activate` on Linux and macOS, `.venv\Scripts\activate` on Windows. Then:

```bash
pip install -r requirements-dev.txt
pip install -e .
uvicorn evidence_desk.app:app --reload
```

Select the `.venv` interpreter the same way. `requirements.txt` is the runtime set; `requirements-dev.txt` adds the dev tools.

## Make targets

- `make run` starts the app
- `make test` runs pytest
- `make lint` runs ruff and mypy
- `make fmt` formats with ruff

`GET /healthz` returns `{"status": "ok"}`.
