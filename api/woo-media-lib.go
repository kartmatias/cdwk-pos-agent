package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
	"github.com/kartmatias/cdwk-pos-agent/cfg"
	"go.uber.org/zap"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func UploadImageToWordPressMedia(imageName string, logger *zap.Logger) (string, error) {
	var imagePath string
	// Read the image file
	if imageName == "" {
		return "", nil
	}

	myCfg := cfg.GetInstance()
	imagePath = filepath.Join(myCfg.ImagePath, imageName)
	if doesFileNotExist(imagePath) {
		return "", nil
	}

	err := UploadToFTP(myCfg.FtpServer, myCfg.FtpUser, myCfg.FtpPassword, imagePath, myCfg.FtpRemotePath+imageName)
	if err != nil {
		return "", nil
	}

	return myCfg.FtpUrlBase + imageName, nil

}

// UploadToFTP send file to ftp server
func UploadToFTP(ftpServer, ftpUser, ftpPassword, localFilePath, remoteFilePath string) error {
	// Connect to FTP server
	c, err := ftp.Dial(ftpServer)
	if err != nil {
		return err
	}

	// Login
	err = c.Login(ftpUser, ftpPassword)
	if err != nil {
		return err
	}
	defer c.Quit()

	// Open local file
	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload to FTP
	err = c.Stor(remoteFilePath, file)
	if err != nil {
		return err
	}

	fmt.Println("File uploaded successfully")
	return nil
}

// UploadImageV0 Send image file to WordPress media library
func UploadImageV0(imagePath, wordpressURL, username, password string) error {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the image file
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Create a new multipart form file
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return err
	}
	part.Write(fileContents)

	// Close the writer
	err = writer.Close()
	if err != nil {
		return err
	}

	// Prepare the request
	request, err := http.NewRequest("POST", wordpressURL+"/wp-json/wp/v2/media", body)
	if err != nil {
		return err
	}

	// Set headers
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Content-Disposition", "attachment; filename="+file.Name())
	request.Header.Set("Authorization", "Basic "+basicAuth(username, password))

	// Make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check response
	if response.StatusCode != http.StatusCreated {
		return errors.New("failed to upload image, status code: " + strconv.Itoa(response.StatusCode))
	}

	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// UploadImageV1 Send image file to WordPress media library
func UploadImageV1(productID int, imageName string, logger *zap.Logger) {
	var imagePath string
	// Read the image file
	if imageName == "" {
		return
	}
	myCfg := cfg.GetInstance()

	imagePath = filepath.Join(myCfg.ImagePath, imageName)

	imageData, err := os.Open(imagePath)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to load image file: %s, Error: %v\n", imagePath, err.Error()))
		return
	}
	defer imageData.Close()

	// Create a Resty client
	client := resty.New()

	// Set basic authentication credentials
	client.SetBasicAuth(myCfg.WPKey, myCfg.WPSecret)

	// Define the endpoint for product images
	imageEndpoint := fmt.Sprintf("%s/wp-json/wp/v2/media/", myCfg.BaseUrl)

	// Create a POST request to upload the image
	imageName = strings.ToLower(imageName)
	response, err := client.R().
		SetFileReader("image", imageName, io.NopCloser(imageData)).
		Post(imageEndpoint)

	if err != nil {
		return
	}

	// Check for a successful response
	if response.IsSuccess() {
		logger.Info("Image uploaded successfully.", zap.Int("Product", productID))
		return
	} else {
		logger.Error(fmt.Sprintf("Failed to upload image. Status code: %d, Error: %v\n", response.StatusCode(), response.Error()))
		return
	}
	// Print error details if the request was not successful
}

// UploadImageV2 Send image file to WordPress media library with content detection
func UploadImageV2(imageName string, logger *zap.Logger) (string, error) {
	var imagePath string
	// Read the image file
	if imageName == "" {
		return "", nil
	}

	myCfg := cfg.GetInstance()
	imagePath = filepath.Join(myCfg.ImagePath, imageName)
	if doesFileNotExist(imagePath) {
		return "", nil
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(imageData)
	switch contentType {
	case "image/png":
		logger.Info("Image type is already PNG")
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageData))
		if err != nil {
			return "", err
		}
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return "", err
		}
		imageData = buf.Bytes()
	default:
		logger.Error(fmt.Sprintf("unsupported content type: %s", contentType))
		return "", fmt.Errorf("unsupported content type: %s", contentType)
	}

	imgBase64 := base64.StdEncoding.EncodeToString(imageData)
	// Create a Resty client
	client := resty.New()

	// Set the base URL for the WordPress REST API
	wpAPIEndpoint := fmt.Sprintf("%s/wp-json/wp/v2/media", myCfg.BaseUrl)

	// Create a POST request to upload the image
	response, err := client.R().
		//SetBasicAuth(myCfg.WPKey, myCfg.WPSecret).
		SetHeader("Authorization", "Basic "+basicAuth(myCfg.WPKey, myCfg.WPSecret)).
		SetHeader("Content-disposition", "attachment; filename="+imageName).
		SetHeader("Content-type", "image/jpeg").
		SetFormData(map[string]string{
			"file": "data:image/jpeg;base64," + string(imgBase64),
		}).
		Post(wpAPIEndpoint)

	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	// Check for a successful response
	if response.IsSuccess() {
		// Parse the response body to get the image URL
		var resultMap map[string]interface{}
		err := json.Unmarshal(response.Body(), &resultMap)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		// Retrieve the image URL from the response
		imageURL, ok := resultMap["guid"].(map[string]interface{})["rendered"].(string)
		if !ok {
			logger.Error("Failed to parse image URL from response", zap.Any("result", resultMap))
			return "", fmt.Errorf("failed to parse image url from response")
		}

		fmt.Println("Image uploaded successfully.")
		return imageURL, nil
	}

	// Print error details if the request was not successful
	fmt.Printf("Failed to upload image. Status code: %d, Error: %v\n", response.StatusCode(), response.Error())
	return "", fmt.Errorf("failed to upload image")
}

// UploadImageV3 Send image file to WordPress media library with file existent check
func UploadImageV3(imageName string, logger *zap.Logger) (string, error) {
	var imagePath string
	// Read the image file
	if imageName == "" {
		return "", nil
	}

	myCfg := cfg.GetInstance()
	imagePath = filepath.Join(myCfg.ImagePath, imageName)
	if doesFileNotExist(imagePath) {
		return "", nil
	}
	// Imagem a ser enviada
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		logger.Error(err.Error())
		return "", nil
	}

	// Set the base URL for the WordPress REST API
	wpAPIEndpoint := fmt.Sprintf("%s/wp-json/wp/v2/media", myCfg.BaseUrl)
	// Faz o upload da imagem
	url, err := UploadImageV4(wpAPIEndpoint, imageData)
	if err != nil {
		logger.Error(err.Error())
		return "", nil
	}
	return url, nil
}

func doesFileNotExist(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	return os.IsNotExist(error)
}

// UploadImageV4 Send image file to WordPress media library with file existent check
func UploadImageV4(url string, imageData []byte) (string, error) {

	// Cria um cliente HTTP
	client := &http.Client{}

	// Cria uma requisição POST para o endpoint de upload de mídia do WordPress
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(imageData))
	if err != nil {
		return "", err
	}

	// Autentica a requisição com as credenciais do WordPress
	myCfg := cfg.GetInstance()
	request.SetBasicAuth(myCfg.WPKey, myCfg.WPSecret)

	// Envia a requisição
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	// Verifica o status da resposta
	if response.StatusCode != http.StatusCreated {
		return "", errors.New(response.Status)
	}

	// Converte a resposta em JSON
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	// Retorna a URL da imagem
	return data["url"].(string), nil
}
