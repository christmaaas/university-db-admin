GO=go

BUILD_DIR=build
OUT_DIR=$(BUILD_DIR)/bin
APP_DIR=cmd/university_db_admin

linux_build:
	$(GO) build -o $(OUT_DIR)/app.out $(APP_DIR)/main.go

linux_run: linux_build
	./$(OUT_DIR)/app.out

win_build:
	$(GO) build -o $(OUT_DIR)/app.exe $(APP_DIR)/main.go

win_run: win_build
	./$(OUT_DIR)/app.exe
	
linux_clean:
	rm -rf $(BUILD_DIR)

win_clean:
	rd /s /q $(BUILD_DIR)