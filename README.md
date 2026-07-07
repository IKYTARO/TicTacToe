# TicTacToe — Веб-приложение на Go

Учебный проект: игра «Крестики-нолики» с алгоритмом Минимакс.
Клиент взаимодействует с сервером через REST API или веб-интерфейс.

## Архитектура

Проект построен по слоёной архитектуре (Clean Architecture). Каждый слой — отдельный пакет.

```
cmd/server/main.go           # точка входа
internal/
├── domain/                  # доменные модели и интерфейсы
│   ├── model/               # Board, Game, GameResult, CellState
│   ├── service/             # интерфейс GameService
│   ├── repository/          # интерфейс GameRepository
│   └── errors/              # доменные ошибки
├── application/             # бизнес-логика
│   ├── game_service.go      # оркестрация хода (ProcessGame)
│   ├── game_over.go         # проверка окончания игры
│   ├── validation.go        # валидация хода игрока
│   └── minimax.go           # алгоритм Минимакс
├── datasource/              # хранение данных (in-memory)
│   ├── storage.go           # потокобезопасное хранилище (sync.Map)
│   ├── repository.go        # реализация GameRepository
│   ├── model.go             # модель для хранения
│   └── mapper.go            # domain ↔ datasource
├── web/                     # HTTP-слой
│   ├── handler.go           # обработчики запросов
│   ├── dto.go               # структуры для JSON
│   ├── mapper.go            # domain ↔ web
│   └── static/
│       └── index.html       # веб-интерфейс
└── di/                      # внедрение зависимостей (uber/fx)
└── fx.go
```

## Зоны ответственности слоёв

### Domain

Описывает **что такое игра**, но не **как** она работает:

- `Board` — игровое поле 3×3, клетки: `Empty` (0), `Cross` (1), `Nought` (2)
- `Game` — текущая игра (UUID + Board)
- `GameResult` — результат: `InProgress`, `CrossWon`, `NoughtWon`, `Draw`
- `GameService` (интерфейс) — `ProcessGame`, `Validate`, `NextMove`, `CheckGameOver`, `CreateGame`
- `GameRepository` (интерфейс) — `Save`, `FindByID`

### Application

Реализует бизнес-логику:

- **`Validate`** — проверяет, что игрок сделал ровно один ход, поставил `Cross` в пустую клетку, не изменил предыдущие ходы
- **`CheckGameOver`** — определяет победителя или ничью
- **`NextMove`** — вычисляет лучший ход компьютера (`Nought`) через Минимакс
- **`ProcessGame`** — оркестрирует полный цикл: валидация → сохранение → ход компьютера → результат

### Datasource

Хранит игры в памяти через `sync.Map`. Потокобезопасно. Поддерживает одновременные игры (ключ — UUID).

### Web

Принимает JSON, конвертирует в доменную модель, вызывает `GameService`, возвращает JSON с обновлённой доской и статусом. Отдаёт веб-интерфейс.

Эндпоинты:
- `GET /` — веб-интерфейс игры
- `POST /game/` — создать новую игру
- `POST /game/{uuid}` — сделать ход

### DI

Собирает граф зависимостей через `uber/fx`: `Storage → Repository → Service → Handler`.

## API

### Создание игры
```
POST /game/
Response: {"id":"uuid","board":[[0,0,0],[0,0,0],[0,0,0]],"status":"in_progress"}
```

### Ход игрока
```
POST /game/{uuid}
Body: {"board":[[1,0,0],[0,0,0],[0,0,0]]}
Response: {"id":"uuid","board":[[1,0,0],[0,2,0],[0,0,0]],"status":"in_progress"}
```

Статусы: `in_progress`, `cross_won`, `nought_won`, `draw`.

Игрок всегда `Cross` (1), компьютер — `Nought` (2). Игрок ходит первым.

## Веб-интерфейс

Откройте `http://localhost:8080` в браузере после запуска сервера. Клетки кликабельны, статус отображается над доской.

## Тесты

Покрыты все слои:

| Пакет | Что тестируется |
|-------|-----------------|
| `application` | `CheckGameOver`, `Validate`, `minimax`, `bestMove`, `evaluate` |
| `datasource` | `Storage`, мапперы, `Repository` |

```bash
go test ./internal/... -v
```

## Быстрый старт

```bash
# Запустить сервер
go run cmd/server/main.go

# Открыть в браузере
http://localhost:8080
```

Демо через curl:

```bash
bash examples/demo.sh
```

Или одной командой:

```bash
make demo
```

## Зависимости

- `github.com/google/uuid` — генерация UUID
- `go.uber.org/fx` — dependency injection