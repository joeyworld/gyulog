package post

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/russross/blackfriday"
)

// Post - Overall posts parsed by parts
type Post struct {
	Title         string
	PublishedDate string
	Summary       string
	Body          string
}

// 1. markdown으로 post 작성후 submit 누르면 (파일 업로드?)
// 2. post 받아와서 string으로 변환
// 3. Post struct 형식에 맞게 파싱
// 4. DB 에 저장

func GetPost(w http.ResponseWriter, r *http.Request) {
	// 업로드된 파일 받아오기
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["post"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("failed to open file")
		log.Fatalln(err)
	}

	// post 받아와서 string으로 변환
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	post := generatePost(data)
	post.insert()
}

func generatePost(data []byte) Post {
	rendered := string(blackfriday.MarkdownCommon(data))
	result := Post{}

	// 제목 찾는 인덱스
	startIndex := strings.Index(rendered, "<h1>")
	endIndex := strings.Index(rendered, "</h1>") + 5
	result.Title = rendered[startIndex:endIndex]

	// 요약 찾는 인덱스
	startIndex = strings.Index(rendered, "<p>")
	endIndex = strings.Index(rendered, "</p>") + 4
	result.Summary = rendered[startIndex:endIndex]

	startIndex = endIndex
	result.Body = rendered[startIndex:]

	return result
}
