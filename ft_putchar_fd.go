package main

import "os"

/*
** Parameters:
**   c: The character to send.
**   fd: The file descriptor to write to.
 */
func ftPutcharFd(c byte, fd *os.File) {
	fd.Write([]byte{c})
}
