Example empty migration creation
```sh
go run ./cmd psql:migration-create -n user_create
```

Migrate specific migration
```sh
go run ./cmd psql:migrate -n 20230210014853_user_create  
```

Undo specific migration
```sh
go run ./cmd psql:migrate-undo -n 20230210014853_user_create
```
