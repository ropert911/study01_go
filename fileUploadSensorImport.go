package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	//"bytes"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"net/http"
	"net/http"
	"study01_go/clients/types"
)

const (
	ContentType = "Content-Type"
	ContentJson = "application/json"
	ContentYaml = "application/x-yaml"
)

// Helper method to get the body from the response after making the request
func getBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// Helper method to make the request and return the response
func makeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func UploadFileRequest(url string, filePath string) (string, error) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Create multipart/form-data request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFileWriter, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(formFileWriter, bytes.NewReader(fileContents))
	if err != nil {
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add(ContentType, writer.FormDataContentType())

	resp, err := makeRequest(req)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", types.ErrResponseNil{}
	}
	defer resp.Body.Close()

	bodyBytes, err := getBody(resp)
	if err != nil {
		return "", err
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusAccepted) {
		return "", types.NewErrServiceClient(resp.StatusCode, bodyBytes)
	}

	bodyString := string(bodyBytes)
	return bodyString, nil
}

func main() {
	str, err := UploadFileRequest("http://192.168.20.56:18004/itc/api/device/import?lang=zh-cn&token=r9bak6bvhtfzbpy8lw36fzce79cwlfp9", "C:\\Users\\sk-qianxiao\\Desktop\\Area0.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
