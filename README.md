# gocat

for print something pre or post cat

## how to install
```
go get -u github.com/umaumax/gocat
```

## how to use
```
gocat -prefix="\x1b[2J\x1b[1;1H\033[0m" -suffix="# END\n" | nc localhost 3939
```

