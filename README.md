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

# Windows
```shell
$ Get-Process | Out-String -Stream | Select-String shell

    821      32    60672      76484       2.39    608   4 powershell
    756      33    23804      82556       2.48   5732   4 ShellExperienceHost


$ go test -v -count 1 ./...
=== RUN   Test
    pid_test.go:9: PID: 7000
    pid_test.go:10: PID of parent: 4204
--- PASS: Test (0.01s)
PASS
ok      pidcheck        0.188s
=== RUN   Test
    pid_test.go:9: PID: 7228
    pid_test.go:10: PID of parent: 4204
--- PASS: Test (0.01s)
PASS
ok      pidcheck/sub    0.187s
=== RUN   Test
    pid_test.go:9: PID: 11452
    pid_test.go:10: PID of parent: 4204
--- PASS: Test (0.01s)
PASS
ok      pidcheck/sub2   0.195s
```
