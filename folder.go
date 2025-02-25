package texukk

import (
	"errors"
	"image"
	"os"
	"path"
)

func NewAtlasFromFolder(srcPath string, width int, height int, recursive bool) (*Atlas, error) {
	folders := []string{srcPath}
	atlas := NewAtlas(width, height)
	errs := []error{}

	for i := range folders {
		folderName := folders[i]
		folder, err := os.Open(folderName)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		content, err := folder.ReadDir(0)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		for _, v := range content {
			srcPath := path.Join(folderName, v.Name())
			if v.IsDir() {
				if recursive {
					folders = append(folders, srcPath)
				}
			} else {
				file, err := os.Open(srcPath)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				img, _, err := image.Decode(file)
				file.Close()
				if err != nil {
					errs = append(errs, err)
					continue
				}
				atlas.AddNamed(img, srcPath)
			}
		}
		folder.Close()
	}

	return atlas, errors.Join(errs...)
}

func RenderFromFolder(path string, width int, height int, recursive bool) (image.Image, map[string]image.Rectangle, error) {
	atlas, err := NewAtlasFromFolder(path, width, height, recursive)
	if err != nil {
		return nil, nil, err
	}
	return atlas.Render()
}
