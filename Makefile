.PHONY = all buildcss watchcss tailwindcss clean

all: buildcss

buildcss: tailwindcss
	./tw -i ./tailwind.css -o ./static/styles/app.css --minify

watchcss:
	./tw -i ./tailwind.css -o ./static/styles/app.css --watch

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tw

clean:
	rm -rf tw static/
