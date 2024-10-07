package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	// "github.com/gofrs/uuid"
)

func main() {
	// Create data by calling generateData function within the "folder" files
	// Printing to terminal and Write the output into sample.json
	// res := folder.GenerateData()
	// folder.PrettyPrint(res)
	// folder.WriteSampleData(res)

	// Fetches all folders
	res := folder.GetAllFolders()
	
	// 1 - Testing Get Folder
	get orgID,  create new driver from "folder", and from the returned result, we call getGetAllChildFolders
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	folderDriver := folder.NewDriver(res)
	name := "sterling-skyrocket"
	orgFolder := folderDriver.GetAllChildFolders(orgID, name)
	folder.PrettyPrint(orgFolder)

	// 2 - Testing Move Folder
	folderDriver := folder.NewDriver(res)
	name := "sterling-skyrocket"
	dst := "harmless-falcon"
	orgFolder, err := folderDriver.MoveFolder(name, dst)
	if err != nil {
		fmt.Println(err)
	} else {
		folder.PrettyPrint(orgFolder)
	}
}




// package main

// import (
// 	"fmt"

// 	"github.com/georgechieng-sc/interns-2022/folder"
// 	"github.com/gofrs/uuid"
// )

// func main() {
// 	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

// 	res := folder.GetAllFolders()

// 	// example usage
// 	folderDriver := folder.NewDriver(res)
// 	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

// 	folder.PrettyPrint(res)
// 	fmt.Printf("\n Folders for orgID: %s", orgID)
// 	folder.PrettyPrint(orgFolder)
// }