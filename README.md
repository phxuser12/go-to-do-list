# Go To-Do List

A minimal to-do list app for learning Go.

- **Backend:** Go standard library only (`net/http`, `html/template`)
- **UI:** server-rendered HTML + [htmx](https://htmx.org) + [Pico.css](https://picocss.com) (both loaded from CDNs, no build step)
- **Storage:** in-memory (lost on restart) behind a `TodoStore` interface
- **Packaging:** templates embedded with `go:embed`, so it builds to a single binary

## Run

Requires Go 1.24+.

```sh
go run .
```

Then open http://localhost:8080.

## Test

```sh
go test ./...
```

## Layout

| File | Purpose |
| --- | --- |
| `main.go` | HTTP server, routes, handlers |
| `store.go` | `Todo` type, `TodoStore` interface, `MemoryStore` |
| `templates/index.html` | Page and the `list` fragment htmx swaps in |
| `main_test.go` | Handler tests using `httptest` |

## Routes

| Method | Path | Action |
| --- | --- | --- |
| GET | `/` | Full page |
| POST | `/todos` | Add a todo (form field `text`) |
| POST | `/todos/{id}/toggle` | Toggle done |
| DELETE | `/todos/{id}` | Delete |

The last three return the updated list fragment, which htmx swaps into the page.

## Ideas for extending

- Add a SQLite-backed `TodoStore` (e.g. `database/sql` with `modernc.org/sqlite`)
- Edit todo text, due dates, filtering
- JSON API with a React front end
