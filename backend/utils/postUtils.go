package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

// ______________________fonction utilitaire
func FormatTimeAgo(creationDate time.Time) string {
	now := time.Now()
	diff := now.Sub(creationDate)

	var res string

	switch {
	case diff.Hours() >= 24*365:
		years := int(diff.Hours() / (24 * 365))
		res = fmt.Sprintf("%d year", years)
		if years > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 24*30:
		months := int(diff.Hours() / (24 * 30))
		res = fmt.Sprintf("%d month", months)
		if months > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 24*7:
		weeks := int(diff.Hours() / (24 * 7))
		res = fmt.Sprintf("%d week", weeks)
		if weeks > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 24:
		days := int(diff.Hours() / 24)
		res = fmt.Sprintf("%d day", days)
		if days > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 1:
		hours := int(diff.Hours())
		res = fmt.Sprintf("%d hour", hours)
		if hours > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Minutes() >= 1:
		minutes := int(diff.Minutes())
		res = fmt.Sprintf("%d minute", minutes)
		if minutes > 1 {
			res += "s"
		}
		res += " ago"
	default:
		seconds := int(diff.Seconds())
		res = fmt.Sprintf("%d second", seconds)
		if seconds != 1 {
			res += "s"
		}
		res += " ago"
	}

	return res
}

func IsValidPost(content string, privacy string, photoURL string, allowedUsers []string) (bool, string) {
	if len(content) == 0 || utf8.RuneCountInString(content) > 500 {
		return false, "Bad Request: Invalid content length"
	}

	if privacy != "public" && privacy != "private" && privacy != "almost_private" {
		return false, "Bad Request: Invalid privacy setting"
	}

	if privacy == "almost_private" && len(allowedUsers) == 0 {
		return false, "Bad Request: No users specified for almost_private post"
	}

	if photoURL != "" {
		if photoURL == "err400" {
			return false, "Bad Request: Invalid image upload"
		} else if photoURL == "err500" {
			return false, "Internal Server Error: Image upload failed"
		}
	}

	return true, ""
}

func IsValidImage(file multipart.File, handler *multipart.FileHeader) bool {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return false
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buff)

	switch contentType {
	case "image/svg+xml", "image/jpeg", "image/gif", "image/png":
	default:
		fmt.Println(contentType)
		fmt.Println("type")
		return false
	}
	return true
}

func UploadImage(w http.ResponseWriter, r *http.Request, origin string) string {
	var photoURL string

	err := r.ParseMultipartForm(10 << 10)
	if err != nil {
		return "err400"
	}
	file, handler, err := r.FormFile("file")
	if file != nil {
		if err != nil {
			return "err400"
		}
		defer file.Close()
		if IsValidImage(file, handler) {
			return "err400"
		}
		if handler.Size > 20<<20 {
			return "err408"
		}
		// path temporaire
		dirPath := "../frontend/public/uploads/"
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err := os.MkdirAll(dirPath, 0755)
			if err != nil {
				return "err500"
			}
		}
		tempFile, err := os.CreateTemp(dirPath, "upload-*"+filepath.Ext(handler.Filename))
		if err != nil {
			return "err500"
		}
		defer tempFile.Close()
		_, err = io.Copy(tempFile, file)
		if err != nil {
			return "err500"
		}
		photoURL = tempFile.Name()
	}
	if origin == "login" && photoURL == "" {
		return "lien_avatar_par_defaut_en_attendant_qu'on_trouve_image.jpg"
	}
	return photoURL
}

func GenerateUuid() string {
	return uuid.New().String()
}

func IsValidComment(content string, photoURL string) (bool, string) {
	if len(content) == 0 || utf8.RuneCountInString(content) > 500 {
		return false, "Bad Request: Invalid content length"
	}

	if photoURL != "" {
		if photoURL == "err400" {
			return false, "Bad Request: Invalid image upload"
		} else if photoURL == "err500" {
			return false, "Internal Server Error: Image upload failed"
		}
	}

	return true, ""
}
