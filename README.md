# 🛠️ Config Validator CLI

A fast, flexible, and lightweight CLI tool built in Go to validate configuration files (YAML, JSON, TOML, and ENV). Designed to help developers and DevOps teams catch misconfigurations early in the CI/CD pipeline.

---

## 🚀 Features

- ✅ Validate configuration syntax across multiple formats
- ⚡ Blazing fast CLI performance with Go
- 🧪 Easily integrable into CI/CD workflows
- 📁 Recursive directory scanning
- 🖨️ Clear and color-coded output for quick debugging

---

## 📦 Supported Formats

- `JSON`
- `YAML`
- `TOML`
- `.env` / Key-Value `.env` files

---

## 🏁 Quick Start

### 🔧 Install

```bash
go install github.com/jordan-kasper/Config-Validator-CLIr@latest
```

Or clone locally:

```bash
git clone https://github.com/jordan-kasper/Config-Validator-CLI
cd Config-Validator-CLI
go build -o Config-Validator-CLI
```

### 🧪 Usage

```bash
Config-Validator-CLI validate config.yaml
```

Validate all config files recursively:

```bash
Config-Validator-CLI validate ./configs --recursive
```

Specify expected schema:

```bash
Config-Validator-CLI validate config.yaml --schema schema.json
```

---

## 📘 Example Schema File

```json
{
  "database": {
    "host": "string",
    "port": "int",
    "user": "string",
    "password": "string"
  },
  "logging": {
    "level": "string",
    "format": "string"
  }
}
```

---

## 🛠️ CLI Options

| Flag             | Description                              |
|------------------|------------------------------------------|
| `--recursive`     | Scan directories recursively             |
| `--schema`        | Path to schema definition file           |
| `--format`        | Output format: `json`, `yaml`, `table`  |
| `--ignore-errors` | Continue validation despite errors       |
| `--strict`        | Fail on unknown fields                   |

---

## 🧰 Tech Stack

- **Language:** Go
- **Dependencies:**
- **Testing:** Go’s built-in `testing` package

---

## 🧪 Running Tests

```bash
go test ./...
```

---

## 🎯 Use Cases

- Pre-deployment config validation
- Static analysis in CI pipelines
- Local developer tooling
- Plugin-based config policy enforcement

---

## 📂 Project Structure

```
config-validator/
├── cmd/                 # CLI commands
├── internal/            # Core validation logic
├── schemas/             # Sample schemas
├── testdata/            # Test config files
├── main.go
└── README.md
```

---

## 🧠 Lessons Learned

- Gained hands-on experience with parsing and validating structured data
- Learned how to build extensible tooling for real-world DevOps pipelines

---

## 📄 License

MIT © [Jordan Kasper](https://github.com/jordan-kasper)