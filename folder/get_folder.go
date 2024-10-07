package folder

import (
	"fmt"
	"strings"
	"errors"
	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	
	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.folders
	res := []Folder{}
	folderExist := false 
	
	// Check string contains: https://www.codecademy.com/resources/docs/go/strings/contains
	for _, f := range folders {
		// Case 1: Check if folder doesnt exist in dataset
		if f.Name == name {
			folderExist = true
			// Case 2: Check if folder exist, but not in the specified org
			if f.OrgId != orgID {
				fmt.Println("Folder does not exist in the specified organization")
				return []Folder{}
			}
			break
		}
	}

	if !folderExist {
		fmt.Println("Folder does not exist")
		return []Folder{}
	}
		
	foldersOrg := f.GetFoldersByOrgID(orgID)
	// Case 3: Normal Case
	for _, f := range foldersOrg {
		if strings.Contains(f.Paths, name + ".") {
			res = append(res, f)
		}
	}
	
	return res
}