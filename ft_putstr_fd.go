package main

import "os"

/*
** Parameters:
**   s: The string to send.
**   fd: The file descriptor to write to.
 */
func ftPutstrFd(s string, fd *os.File) {
	fd.Write([]byte(s))
}
