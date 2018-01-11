package sysutil

/*
#include <unistd.h>

int getJiffies() {
    return sysconf(_SC_CLK_TCK);
}
*/
import "C"

func GetJiffies() int {
	return int(C.getJiffies())
}
