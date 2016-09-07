package main

import "fmt"

func help() {
	fmt.Println(`
Usage:
$ kubedeploy get [-n namespace]
$ kubedeploy replace -p pod -i image -n namespace
$ kubedeploy deploy -s service -n namespace
$ kubedeploy list -i image
	`)
}
