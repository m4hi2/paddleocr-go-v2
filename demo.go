package main

import (
	"log"

	"github.com/m4hi2/paddleocr-go/ocr"
)

var (
	confFile     string
	image        string
	imageDir     string
	useServering bool
	port         string
)

//func init() {
//	flag.StringVar(&confFile, "config", "config/conf.yaml", "config from ocr system. If not given, will use default config.")
//	flag.StringVar(&image, "image", "", "image to predict. if not given, will use image_dir")
//	flag.StringVar(&imageDir, "image_dir", "", "imgs in dir to be predicted. if not given, will check servering")
//	flag.BoolVar(&useServering, "use_servering", false, "whether to use ocr server. [default: false]")
//	flag.StringVar(&port, "port", "18600", "which port to serve ocr server. [default: 18600].")
//}

func main() {
	//flag.Parse()
	confFile = "config/conf.yaml"
	image = "images/test.jpg"

	sys := ocr.NewOCRSystem(confFile, nil)

	if image != "" {
		img := ocr.ReadImage(image)
		results := sys.PredictOneImage(img)
		for _, res := range results {
			log.Println(res)
		}
		return
	}

	if imageDir != "" {
		results := sys.PredictDirImages(imageDir)
		for k, vs := range results {
			log.Printf("======== image: %v =======\n", k)
			for _, res := range vs {
				log.Println(res)
			}
		}
	}

	if useServering {
		sys.StartServer(port)
	}
}
