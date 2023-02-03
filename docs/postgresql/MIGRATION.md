Example empty migration creation
```sh
go run runners/migration-runner.go create -n example_migration
```

Run all migrations
```sh
go run runners/migration-runner.go up  
```

Run next uncommited 2 migration
```sh
go run runners/migration-runner.go up --step=2
```

Undo last committed 1 migration
```sh
go run runners/migration-runner.go up --step=1
```

Show all migrations status
```sh
go run runners/migration-runner.go status
```