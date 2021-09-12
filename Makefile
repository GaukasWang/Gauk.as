all:
	git pull
	sudo systemctl stop gauk.as
	/usr/local/go/bin/go build -o gauk.as ./src
	sudo systemctl start gauk.as
