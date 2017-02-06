package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"whats-that-resolution/resolution"
)

var imageStyle = `image-rendering: -webkit-optimize-contrast;
image-rendering: pixelated;
image-rendering: crisp-edges; 
image-rendering: -moz-crisp-edges; 
-ms-interpolation-mode: nearest-neighbor;`

func doImage(w http.ResponseWriter, theUrl string) {
	if theUrl != "" {
		resp, err := http.Get(theUrl)
		if err != nil {
			w.Write([]byte("error getting url"))
			return
		}
		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	w.Write([]byte("error getting image"))
		//	return
		//}

		//contentType := resp.Header["Content-Type"]
		//if contentType == "image/gif" {
		//}
		img, mime, err := image.Decode(resp.Body)
		if err != nil {
			w.Write([]byte("error getting image"))
			w.Write([]byte(err.Error()))
			return
		}

		width, height, origWidth, origHeight, err := resolution.CheckResolution(img)
		if err != nil {
			w.Write([]byte("error checking resolution"))
		}
		//w.Write([]byte(mime + ": " + w + "x" + h))
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "%s: %dx%d (original %dx%d)<br>", mime, width, height, origWidth, origHeight)

		w.Write([]byte(`<img style="` + imageStyle + `" src="` + theUrl + `">`))
	} else {
		w.Header().Set("Content-Type", "text/html")
	}
}

func doTwitterImage(w http.ResponseWriter, twitterURL string) {
	resp, err := http.Get(theUrl)
	if err != nil {
		w.Write([]byte("error getting url"))
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte("error getting twitter url"))
		return
	}
	body := string(bodyBytes)
}

func StringBetween(subject, left, right string) {

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		form := r.Form
		whichForm := form.Get("form")
		twitterURL := form.Get("twitter_url")
		theUrl := form.Get("url")
		if whichForm == "image" {
			doImage(w, theUrl)
		} else if whichForm == "twitter" {
			doTwitterImage(w, twitterURL)
		}
		w.Write([]byte(`
<br>
Enter url for pixel art image: 
<form action="/" method="GET">
<input name="url" value="` + theUrl + `"type="text">
<input name="form" type="submit" value="image"/>
</form>

Enter Twitter url:
<form action="/" method="GET">
<input name="twitter_url" value="` + twitterURL + `"type="text">
<input name="form" type="submit" value="twitter" />
</form>
`))
	})

	port := "8082"
	fmt.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
