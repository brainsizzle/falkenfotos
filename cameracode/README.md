build for raspberry pi

	env GOOS=linux GOARCH=arm go build -o ./out/falconcam
    scp out/falconcam pi@192.168.3.10:.

