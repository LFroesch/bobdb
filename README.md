# bobdb

Keyboard-first database TUI for SQLite, Postgres, and MongoDB. `bobdb` is built for the quick browse, query, inspect, and copy/export work that usually pulls you into a heavier GUI.

![bobdb hero screenshot](assets/screenshots/hero.png)

**Live demo:** [froesch.dev](https://froesch.dev)

## Release Status

Developed for WSL2/Linux first. Cross-platform testing and bug fixing for macOS and native Windows are still in progress.

Linux, WSL2, and macOS are the primary targets today. Windows binaries and installer entrypoints are available, but native Windows should still be treated as experimental.

## Install

Quick install:

```bash
curl -fsSL https://raw.githubusercontent.com/LFroesch/bobdb/main/install.sh | bash
```

Experimental native Windows install:

```powershell
irm https://raw.githubusercontent.com/LFroesch/bobdb/main/install.ps1 | iex
```

Direct installers: [`install.sh`](https://raw.githubusercontent.com/LFroesch/bobdb/main/install.sh), [`install.ps1`](https://raw.githubusercontent.com/LFroesch/bobdb/main/install.ps1)

If you cloned the repo already:

```powershell
./install.ps1
```

```bat
install.cmd
```

Other options:

```bash
go install github.com/LFroesch/bobdb@latest
make install
```

Run it with any of:

```bash
bobdb
bob
bdb
```

CLI helpers:

```bash
bobdb --help
bobdb --version
```

## Media

![bobdb query screenshot](assets/screenshots/query.png)

## Supported Backends

- SQLite using a local `.sqlite` or `.db` file
- Postgres using a DSN like `postgres://user:pass@host:5432/dbname`
- MongoDB using a DSN like `mongodb://user:pass@host:27017/dbname`

Saved connections live in `~/.config/bobdb/config.json`. The config directory is created with owner-only access and the file is written with `0600` permissions because DSNs often include credentials.

## Layout

| Tab | Purpose |
|-----|---------|
| Connections | Save, edit, delete, and open database connections |
| Browse | Inspect schema and preview table or collection data |
| Query | Write SQL or Mongo shell queries with completion, examples, and templates |
| Results | Review the latest result set and copy rows or values |

## Features

- One app for relational and document databases
- Schema browsing plus data preview
- Backend-aware query autocomplete
- Templates, examples, saved queries, and recent history overlays
- Structured row view plus quick copy/export as JSON or CSV
- Confirmation for saved-connection deletion and obvious write queries
- Optional Ollama query generation with `ctrl+g`
- Password masking in connection detail views so screenshots do not expose DSNs

## Query Workflow

- SQL and Mongo use the same query tab
- `tab` accepts the current completion
- `ctrl+t` opens templates
- `ctrl+e` opens examples for the current backend
- `ctrl+o` opens recent query history
- `ctrl+u` opens saved queries
- `ctrl+r` runs the current query

Mongo queries use normal shell-style commands such as:

```javascript
db.users.find({})
db.users.aggregate([{ $match: { active: true } }])
db.users.updateOne({ _id: ... }, { $set: { name: "..." } })
```

The completion flow is contextual rather than keyword-dump based. SQL suggestions follow common clause order, and Mongo suggestions follow command, collection, field, operator, and value positions.

## Ollama

`ctrl+g` sends a natural-language prompt plus schema metadata to an Ollama server and inserts the generated query.

Config:

- `BOBDB_OLLAMA_HOST`
- `BOBDB_OLLAMA_MODEL`

Defaults:

- host: `http://localhost:11434`
- model: `qwen2.5:7b`

Only schema names are sent, not row data. If you point `BOBDB_OLLAMA_HOST` at a remote server, treat schema names as data leaving your machine.

## Controls

| Key | Action |
|-----|--------|
| `1-4` | Switch tabs |
| `/` | Jump to Query |
| `tab` | Switch pane focus, or accept completion inside Query |
| `n` | New connection |
| `e` | Edit connection or build contextual edit query |
| `d` | Delete selected connection |
| `r` | Refresh tables |
| `ctrl+r` | Run query |
| `ctrl+l` | Clear query |
| `ctrl+o` | Open recent queries |
| `ctrl+t` | Open templates |
| `ctrl+e` | Open examples |
| `ctrl+u` | Open saved queries |
| `c` | Copy the focused item |
| `C` | Copy/export in another format |
| `v` | Open structured detail view |
| `?` | Help |
| `q` | Quit |

## License

[AGPL-3.0](LICENSE)
