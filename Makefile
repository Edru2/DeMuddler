# Makefile for building and packaging de-muddler for multiple platforms

# Define variables
APP_NAME := de-muddler
VERSION := 1.0.0
ZIP_DIR := zip

# Define platform-specific variables
LINUX_BIN := $(APP_NAME)
MACOS_BIN := $(APP_NAME)
WINDOWS_BIN := $(APP_NAME).exe

# Default target
all: clean linux macos windows

# Clean build directories
clean:
	rm -rf $(ZIP_DIR)

# Build for Linux
linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BIN) .
	chmod +x $(LINUX_BIN)
	mkdir -p $(ZIP_DIR)
	zip -j $(ZIP_DIR)/$(APP_NAME)-linux-amd64.zip $(LINUX_BIN)
	rm $(LINUX_BIN)

# Build for macOS
macos:
	GOOS=darwin GOARCH=amd64 go build -o $(MACOS_BIN) .
	chmod +x $(MACOS_BIN)
	mkdir -p $(ZIP_DIR)
	zip -j $(ZIP_DIR)/$(APP_NAME)-macos-amd64.zip $(MACOS_BIN)
	rm $(MACOS_BIN)

# Build for Windows
windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BIN) .
	mkdir -p $(ZIP_DIR)
	zip -j $(ZIP_DIR)/$(APP_NAME)-windows-amd64.zip $(WINDOWS_BIN)
	rm $(WINDOWS_BIN)

# Target to print file sizes
check-size:
	@echo "Sizes of the generated zip files:"
	@du -h $(ZIP_DIR)/*.zip

# Clean and re-build everything
rebuild: clean all

# Phony targets
.PHONY: all clean linux macos windows check-size rebuild
