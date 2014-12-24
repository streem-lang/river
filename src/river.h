#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#define NORMAL    "\x1B[0m"
#define RED       "\x1B[31m"
#define GREEN     "\x1B[32m"
#define YELLOW    "\x1B[33m"
#define BLUE      "\x1B[34m"
#define MAGENTA  "\x1B[35m"
#define CYAN      "\x1B[36m"
#define WHITE     "\x1B[37m"

#ifdef _WIN32
# include "ansicolor-w32.h"
#endif
