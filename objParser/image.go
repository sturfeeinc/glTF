package obj

import (
	"strings"
)

func getChannels(colorType int) int {
	switch (colorType) {
	case 0: // greyscale
		return 1;
	case 2: // RGB
		return 3;
	case 4: // greyscale + alpha
		return 2;
	case 6: // RGB + alpha
		return 4;
	default:
		return 3;
	}
}

func getUriType(extension string) string {
	switch (extension) {
	case "png":
		return "data:image/png";
	case "jpg":
		return "data:image/jpeg";
	case "jpeg":
		return "data:image/jpeg";
	case "gif":
		return "data:image/gif";
	default:
		return "data:image/" + extension;
	}
}

type images map[string]interface{}
type info struct {
	transparent bool
	channels    int
	data        []byte
	uri         string
}

func Parse2(imagePath string) info {
	sArr := strings.Split(imagePath, ".")
	extension := sArr[len(sArr) - 1]
	uriType := getUriType(extension);
	uri := uriType + ";base64,";

	Info := info{
		transparent: false,
		channels: 3,
		data: []byte{},
		uri: uri,
	};

	if extension == "png" {
		// Color type is encoded in the 25th bit of the png
		colorType := 1;
		channels := getChannels(colorType);
		Info.channels = channels;
		Info.transparent = (channels == 4);
	}
	return Info
}
