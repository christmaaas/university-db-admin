linux_build:
	go build -o build/bin/app.out cmd/university_db_admin/main.go

linux_run: linux_build
	./build/bin/app.out

win_build:
	go build -o build/bin/app.exe cmd/university_db_admin/main.go

win_run: win_build
	./build/bin/app.exe
	
linux_clean:
	rm -rf build

win_clean:
	rd /s /q build