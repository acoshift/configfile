test:
	env \
		DATA1=true \
		DATA2=false \
		DATA3=9 \
		DATA4=0 \
		DATA5=3m5s \
		DATA6="aGVsbG8=" \
		EMPTY="" \
		ONLYENV=1 \
		go test ./...
