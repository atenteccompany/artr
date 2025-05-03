# Aten Remote Task Runner (ARTR)

`artr` is a tool to run remote tasks remotely without exposing remote machine to caller.

**Aten Remote Task Runner (ARTR)** is a minimal, secure agent-based tool to execute controlled scripts on remote Linux servers. Designed for modern DevOps and on-premise automation, it helps you schedule and run predefined system tasks (like database backups, log rotations, and status reports) from a centralized management system — while maintaining complete execution control and operational security.

- 🛡️ Root-safe architecture
- 🔒 Encrypted secret handling (no password leakage)
- 📦 Lightweight, single-binary deployment (`artr`)
- 📜 MIT-licensed and easy to integrate

`art` is useful in situations where remote access to servers are not allowed or not preferred. It is also useful to run automated tasks on a machine where tasks are initiated from a remote 
machine using any automation mechanism like `cron` or `systemd.timer` or `at`.

---

## 🧩 Use Case

ARTR was developed as part of the **Anchor MMS** platform to manage periodic on-premise database backups and remote system tasks without exposing credentials or APIs. It runs under `systemd` for stability and is designed to execute only pre-approved scripts — ensuring zero shell injection or unsafe commands.

---

## ⚙️ Features

- Agent runs as a systemd service (`artr.service`)
- Task definitions are plain bash scripts (can be pulled securely from remote orchestrator)
- Native support for:
  - mTLS for cli/agent daemon communication
  - File transfer from agen to CLI
  - Post-execution report upload (coming soon)
- Optional counter/progress support for large files

---

## 🚀 Getting Started

### 1. Install

Download the binary:

```bash
wget https://github.com/atenteccompany/artr/releases/latest/download/artr
chmod +x artr
sudo mv artr /usr/local/bin/
```

```ini
# /etc/systemd/system/artr.service
[Unit]
Description=Aten Remote Task Runner
After=network.target

[Service]
ExecStart=/usr/local/bin/artr server --port 9443 --dir /opt/artr/scripts/
Restart=always
User=root

[Install]
WantedBy=multi-user.target
```

enable and start daemon

```ini
sudo systemctl daemon-reexec
sudo systemctl enable artr
sudo systemctl start artr
```

## Scripts Layout (tasks)

`arts` daemon will look for executable scripts inside a directory defined using `--dir` flag for `server` verb (e.g. /opt/artr/scripts/).

Each script must:

- Be executable (chmod +x)
- Follow a predefined naming pattern
- Include necessary GPG logic if applicable
- use `#::ARTR::` key to add headers for scripts, available headers are:
    - `#::ARTR::title=<task title>` : will display title on CLI output
    - `#::ARTR::result-type=<type>` : defines result type for CLI to render, supported types are:
        - `metric`: expects single value return only 
        - `table`: expects tabulated result 
        - `file`: defines task will generate a file that must be transferred to CLI once script ends
    - `file-name`: the file that `artr` daemon will send to CLI after script (Task) execution finishes. This works only when `result-type=file` is defined in script.

You may mount these scripts from containers or manage them via GitOps.

## Verbs

- `version`: show version
- `server`: run artr daemon
- `run`: runs a remote task on server and displays the result
- `inspect`: runs a remote systask for inspection purposes
- `list`: lists available remote tasks

### Example

```bash
# list available remote tasks for server 127.0.0.1 where artr daemon listens on port 9443
./artr list --addr 127.0.0.1 --port 9443

# run a remote task named backup at remote server of IP 192.168.1.2 where artr daemon listens on port 9443
./artr run backup --addr 192.168.1.2 --port 9443

# run a remote task named get-backup at remote server IP 192.168.1.2 where artr daemon 
# listens on port 9443, and task is expecting to generate file that must be transferred 
# to directory ./backups
./arts run get-backup --addr 192.168.1.2 --port 9443 --outdir ./backups
```

## Architecture

Aten Remote Task Runner (artr) are split to 3 components:

1. user CLI: a binary CLI tool at user machine.
2. management server: `arts` daemon running on a remote machine(s) using `systemd`:
    - responsible for running task on-demand (bash scripts)
    - responsible for running internal `artr` system tasks (built-in functionality)
3. task scripts: Bash scripts on each remote server for each functional requirement:
    - each task will be a single bash script 
    - bash scripts must be able to chaining (follow Linux philosophy)

## Contributing

We welcome contributions! Please read our [Contributing Guide](./CONTRIBUTING.md) and [Code of Conduct](./CODE_OF_CONDUCT.md) to get started.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE.md) file for details.

## 🧠 Credits

ARTR is developed and maintained by [AtenTEC](https://www.atentec.com), under the [Anchor](https://anchor.anchornize.com) product ecosystem.

If you find this tool helpful, please consider starring the repo ⭐️ and spreading the word.
