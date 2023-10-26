.PHONY = all buildcss watchcss tailwindcss clean

all: tailwindcss buildcss

buildcss:
	./tw -i ./tailwind.css -o ./static/styles/app.css --minify

watchcss:
	./tw -i ./tailwind.css -o ./static/styles/app.css --watch

watchserver:
	~/go/bin/air

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tw

clean:
	rm -rf static/
