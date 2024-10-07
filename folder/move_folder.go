package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Find Source and Destination Folder
	srcFolder, srcErr := f.FindFolder(name)
	dstFolder, dstErr := f.FindFolder(dst)
	if srcErr != nil {
		return nil, errors.New("Source folder does not exist")
	}
	if dstErr != nil {
		return nil, errors.New("Destination folder does not exist")
	}
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("Cannot move a folder to a different organization")
	}
	if srcFolder.Paths == dstFolder.Paths {
		return nil, errors.New("Cannot move a folder to itself")
	}
	// Checking Prefix: https://www.codecademy.com/resources/docs/go/strings/hasprefix
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}

	// Example
	// Src = alpha.bravo
	// Dst = alpha.delta
	// Path1 = alpha.bravo
	// Path2 = alpha.bravo.charlie
	// Result :
	// Path1 = alpha.bravo 					(f.Paths) <- replace the "alpha.bravo" (srcFolder.Paths) with alpha.delta.bravo (dstFolder.Path + srcFolder.Name)
	// Path2 = alpha.bravo.charlie 	(f.Paths) <- replace the "alpha.bravo" (srcFolder.Paths) with alpha.delta.bravo (dstFolder.Path + srcFolder.Name)

	res := []Folder{}
	folders := f.folders
	for _, f := range folders {
		if strings.HasPrefix(f.Paths, srcFolder.Paths) {
			f.Paths = strings.Replace(f.Paths, srcFolder.Paths, dstFolder.Paths + "." + srcFolder.Name, 1)
		}
		res = append(res, f)
	}

	return res, nil
}


func (f *driver) FindFolder(name string) (Folder, error) {
	folders := f.folders
	for _,f := range folders {
		if f.Name == name {
			return f, nil
		}
	}

	return Folder{}, errors.New("Error")
}