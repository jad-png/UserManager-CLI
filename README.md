# UserManager-CLI
## Project Structure

```sh
    user-manager/
├── cmd/
│   └── cli/           # CLI entry point
├── internal/
│   ├── auth/          # Authentication
│   ├── commands/      # Command handlers
│   ├── storage/       # Memory storage
│   ├── models/        # User data structures
│   └── concurrent/    # Channel/goroutine features
├── pkg/
│   └── utils/         # Shared utilities
└── docs/              # Learning documentation
```