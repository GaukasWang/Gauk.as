all:
	git pull
	sudo systemctl stop gauk.as
	/usr/local/go build -o gauk.as ./src
	sudo systemctl start gauk.as
