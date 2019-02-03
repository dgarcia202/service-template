APP_NAME = customers
BUILD_DIR = build/
SCRIPTS_DIR = scripts/
TARGET = $(BUILD_DIR)$(APP_NAME)
ENTRY_POINT = cmd/app/main.go
SOURCES = $(wildcard internal/*/*.go) $(wildcard cmd/*/*.go) $(wildcard pkg/*/*.go)

ifeq ($(OS), Windows_NT)
	TARGET = $(BUILD_DIR)$(APP_NAME).exe
endif

all: $(TARGET)

$(TARGET): $(SOURCES)
	@$(SCRIPTS_DIR)build.sh $(TARGET) $(ENTRY_POINT)

.PHONY: clean

clean:
	@rm -f $(TARGET)