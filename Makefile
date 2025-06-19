build:
	# go get -d golang.org/x/mobile/cmd/gomobile
	go install golang.org/x/mobile/cmd/gomobile
	gomobile init
	gomobile bind -target=macos -o Sources/ESBuildMobile.xcframework github.com/Pickleboyonline/ESBuildMobile/lib/esbuildmobile

test:
	@cd lib/esbuildmobile && go test
