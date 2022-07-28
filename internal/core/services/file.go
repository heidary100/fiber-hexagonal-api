package services

import (
	"fmt"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/pkg/crawler"
	"github.com/heidary100/fiber-hexagonal-api/internal/pkg/google"
)

type fileService struct {
	repository ports.FileRepository
}

func NewFileService(r ports.FileRepository) ports.FileService {
	return &fileService{
		repository: r,
	}
}

func (s *fileService) Insert(file *domain.File) (*domain.File, error) {
	return nil, nil
}

func (s *fileService) Fetch() (*[]domain.File, error) {
	return nil, nil
}

func (s *fileService) Update(file *domain.File) (*domain.File, error) {
	return nil, nil
}

func (s *fileService) Remove(ID string) error {
	return nil
}

func (s *fileService) Find(name string, kind string) ([]domain.File, error) {
	extensions := map[string][]string{
		"music": {"mp3", "m4a", "flac"},
		"film":  {"mp4", "mkv", "wav"},
	}
	prefix := map[string]string{
		"music": ` "دانلود" "رایگان" "آهنگ"`,
		"film":  ` "دانلود" "رایگان" "فیلم"`,
	}

	searchResult, err := google.Search(prefix[kind] + name)
	if err != nil {
		return nil, err
	}

	ch := make(chan []string)
	for _, eachSearchResult := range searchResult {
		go fetchFiles(eachSearchResult.Link, extensions[kind], ch)
	}

	var files []domain.File
	for range searchResult {
		urls := <-ch
		for _, url := range urls {
			file := domain.File{Url: url}
			files = append(files, file)
		}
	}

	return files, nil
}

func fetchFiles(url string, extensions []string, ch chan<- []string) {
	fileUrls, err := crawler.GetFilesFromWebPage(url, extensions)
	if err != nil {
		ch <- nil
	}
	fmt.Println("FOUND", len(fileUrls))
	ch <- fileUrls
}
