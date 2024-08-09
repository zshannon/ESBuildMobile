build:
	go get -d golang.org/x/mobile/cmd/gomobile
	gomobile bind -target=macos -o Sources/ESBuildMobile.xcframework 
