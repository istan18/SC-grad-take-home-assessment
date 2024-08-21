package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders What the code does:
// Declares an empty array of folders using empty slice declaration using literal (could use nil slice declaration instead for consistency - small semantic differences regarding initialisation)
// Fetches all folders associated with the requests organisation ID and appends to the array of folders
// Declares an empty array of pointers to folders
// Creates a folder response which is a defined struct which contains an array of pointers to folders
// Returns the response with no errors
// IMPROVEMENT: Variable names are not descriptive, rename to something more meaningful for easier understanding
// IMPROVEMENT: Delete unused variables
// IMPROVEMENT: Remove unnecessary loops to simplify complexity
// IMPROVEMENT: Handle error cases
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	response := &FetchFolderResponse{
		Folders: folders,
	}

	return response, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// Gets some sample folder data
	folders := GetSampleData()

	var resFolder []*Folder
	// Filters the folders by the organisation ID
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	// Returns the filtered folders with no errors
	return resFolder, nil
}
