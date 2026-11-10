.PHONY: run test lint fmt

run:
	uv run uvicorn evidence_desk.app:app --reload

test:
	uv run pytest

lint:
	uv run ruff check .
	uv run ruff format --check .
	uv run mypy

fmt:
	uv run ruff format .
