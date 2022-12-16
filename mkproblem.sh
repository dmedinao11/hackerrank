GO_MOD_BASE="github.com/dmedinao11/hackerrank"
GO_MAIN_PACKAGE="package main"

mkdir "$1"

if [ $? -eq 1 ]
then
  exit 1
fi

cd "$1" || exit

echo "$GO_MAIN_PACKAGE" > main.go

go mod init $GO_MOD_BASE/"$1"
go mod tidy