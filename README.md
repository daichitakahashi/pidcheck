# pidcheck

`go test`は複数のプロセスにタスクを分散させるが、[daichitakahashi/rsmap](https://github.com/daichitakahashi/rsmap)の開発なかで`go test`全体の実行IDになるようなデータが必要になった。

macOS, Linux, Windowsの全てで`os.Getppid()`を実行IDとして使うことができそうかどうかを検証する。

# macOS
```shell
$ go test -v -count 1 ./...
=== RUN   Test
    pid_test.go:9: PID: 7503
    pid_test.go:10: PID of parent: 7463
--- PASS: Test (0.00s)
PASS
ok      pidcheck        0.208s
=== RUN   Test
    pid_test.go:9: PID: 7504
    pid_test.go:10: PID of parent: 7463
--- PASS: Test (0.00s)
PASS
ok      pidcheck/sub    0.147s
=== RUN   Test
    pid_test.go:9: PID: 7505
    pid_test.go:10: PID of parent: 7463
--- PASS: Test (0.00s)
PASS
ok      pidcheck/sub2   0.381s
```
