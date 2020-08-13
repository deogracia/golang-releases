security:
	go list -json -m all | docker run --rm -i sonatypecommunity/nancy:latest
