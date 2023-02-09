Example empty migration creation
```sh
go run cmd/cmd.go psql:migration-create -n dummy_one_create
```

Migrate specific migration
```sh
go run cmd/cmd.go psql:migrate -n 20230210014853_dummy_one_create  
```

Undo specific migration
```sh
go run cmd/cmd.go psql:migrate-undo -n 20230210014853_dummy_one_create
```
