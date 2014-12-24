RIVER_TARGET = bin/river
RIVER_SRC = src/main.c
RIVER_OBJ = $(RIVER_SRC:.c=.o)

RIVUP_TARGET = bin/rivup
RIVUP_SRC = src/update.c
RIVUP_OBJ = $(RIVUP_SRC:.c=.o)

RIVDOC_TARGET = bin/rivdoc
RIVDOC_SRC = src/doc.c
RIVDOC_OBJ = $(RIVDOC_SRC:.c=.o)

CFLAGS = -Isrc

ifeq (Windows_NT,$(OS))
RIVER_TARGET := $(RIVER_TARGET).exe
RIVER_SRC += src/ansicolor-w32.c
RIVUP_TARGET := $(RIVUP_TARGET).exe
RIVUP_SRC += src/ansicolor-w32.c
RIVDOC_TARGET := $(RIVDOC_TARGET).exe
RIVDOC_SRC += src/ansicolor-w32.c
endif

.SUFFIXES: .c.o

all: $(RIVER_TARGET) $(RIVUP_TARGET) $(RIVDOC_TARGET)

$(RIVER_TARGET) : $(RIVER_OBJ)
	gcc $(CFLAGS) -o $@ $^

$(RIVUP_TARGET) : $(RIVUP_OBJ)
	gcc $(CFLAGS) -o $@ $^

$(RIVDOC_TARGET) : $(RIVDOC_OBJ)
	gcc $(CFLAGS) -o $@ $^

.c.o:
	gcc $(CFLAGS) -o $@ -c $<

clean:
	-rm -rf bin/* src/*.o

