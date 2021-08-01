package main

import "github.com/vartanbeno/go-reddit/reddit"

func main() {
	client, _ := reddit.NewClient(reddit.Credentials{}, reddit.FromEnv)
}
