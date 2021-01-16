0. 
source  ~/.gvm/scripts/gvm
gvm use go1.10.3 --default
1. Export the $GOPATH in $PATH
```
export GOPATH=$HOME/Dev/go
export PATH=$PATH:$GOROOT/bin
```
2. Install the dependances.
go get github.com/rivo/tview

3. echo $(ls | ./selector_tui.go )
