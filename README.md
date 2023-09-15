# pidcheck

`go test`は複数のプロセスにタスクを分散させるが、[daichitakahashi/rsmap](https://github.com/daichitakahashi/rsmap)の開発なかで`go test`全体の実行IDになるようなデータが必要になった。

macOS, Linux, Windowsの全てで`os.Getppid()`を実行IDとして使うことができそうかどうかを検証する。

# macOS
```shell
$ echo $$
6144
$ go test -v -count 1 ./...
=== RUN   Test
    pid_test.go:9: PID: 8547
    pid_test.go:10: PID of parent: 8506
--- PASS: Test (0.00s)
PASS
ok      pidcheck        0.316s
=== RUN   Test
    pid_test.go:9: PID: 8548
    pid_test.go:10: PID of parent: 8506
--- PASS: Test (0.00s)
PASS
ok      pidcheck/sub    0.508s
=== RUN   Test
    pid_test.go:9: PID: 8549
    pid_test.go:10: PID of parent: 8506
--- PASS: Test (0.00s)
PASS
ok      pidcheck/sub2   0.668s
```

# Linux
```shell
$ echo $$
1
$ go test -v -count 1 ./...
=== RUN   Test
    pid_test.go:9: PID: 1357
    pid_test.go:10: PID of parent: 7
--- PASS: Test (0.00s)
PASS
ok      pidcheck        0.002s
=== RUN   Test
    pid_test.go:9: PID: 1361
    pid_test.go:10: PID of parent: 7
--- PASS: Test (0.00s)
PASS
ok      pidcheck/sub    0.002s
=== RUN   Test
    pid_test.go:9: PID: 1367
    pid_test.go:10: PID of parent: 7
--- PASS: Test (0.00s)
PASS
ok      pidcheck/sub2   0.002s
```
