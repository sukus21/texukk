package texukk

import (
	"errors"
	"image"
	"io/fs"
	"os"
)

func NewAtlasFromFolder(srcPath string, width int, height int, recursive bool) (*Atlas, error) {
	fsys := os.DirFS(srcPath)
	atlas := NewAtlas(width, height)
	errs := []error{}

	fs.WalkDir(fsys, ".", func(fname string, d fs.DirEntry, err error) error {
		if err != nil {
			errs = append(errs, err)
			return nil
		}

		// This is probably an image
		if !d.IsDir() {
			file, err := fsys.Open(fname)
			if err != nil {
				errs = append(errs, err)
				return nil
			}
			img, _, err := image.Decode(file)
			file.Close()
			if err != nil {
				errs = append(errs, err)
				return nil
			}
			atlas.AddNamed(img, fname)
		}

		// Directory, what do?
		if d.IsDir() {
			if !recursive && fname != "." {
				return fs.SkipDir
			}
		}

		// All good
		return nil
	})

	return atlas, errors.Join(errs...)
}

func RenderFromFolder(path string, width int, height int, recursive bool) (image.Image, map[string]image.Rectangle, error) {
	atlas, err := NewAtlasFromFolder(path, width, height, recursive)
	if err != nil {
		return nil, nil, err
	}
	return atlas.Render()
}
