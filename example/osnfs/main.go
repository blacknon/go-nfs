package main

import (
	"fmt"
	"net"
	"os"

	nfs "github.com/blacknon/go-nfs-sshlib"
	nfshelper "github.com/blacknon/go-nfs-sshlib/helpers"
	osfs "github.com/go-git/go-billy/v5/osfs"
)

func main() {
	port := ""
	if len(os.Args) < 2 {
		fmt.Printf("Usage: osnfs </path/to/folder> [port]\n")
		return
	} else if len(os.Args) == 3 {
		port = os.Args[2]
	}

	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	fmt.Printf("osnfs server running at %s\n", listener.Addr())

	bfs := osfs.New(os.Args[1])
	bfsPlusChange := NewChangeOSFS(bfs)

	handler := nfshelper.NewNullAuthHandler(bfsPlusChange)
	cacheHelper := nfshelper.NewCachingHandler(handler, 1024)
	fmt.Printf("%v", nfs.Serve(listener, cacheHelper))
}
