package util

/*
#include <unistd.h>

int getJiffies() {
    return sysconf(_SC_CLK_TCK);
}
*/
import "C"

func getJiffies() int {
	return int(C.getJiffies())
}
