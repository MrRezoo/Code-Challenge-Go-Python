FROM python:3.12-slim

LABEL name="PythonClient"

WORKDIR /app

COPY pyproject.toml poetry.lock ./

RUN pip install poetry

RUN poetry config virtualenvs.create false && poetry install --no-dev

COPY . .

CMD ["python", "main.py"]
