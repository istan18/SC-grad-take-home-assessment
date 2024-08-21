package folders

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gofrs/uuid"
)

func FetchAllFoldersByOrgIDToken(orgID uuid.UUID, token string, pageSize int) ([]*Folder, string, error) {
	folders := GetSampleData()

	var resFolder []*Folder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	start := 0
	if token != "" {
		var err error
		start, err = decodeToken(token)
		if err != nil {
			return nil, "", err
		}
	}

	end := start + pageSize

	if start >= len(resFolder) {
		return []*Folder{}, "", nil
	}
	if end > len(resFolder) {
		end = len(resFolder)
	}

	nextToken := ""
	if end < len(resFolder) {
		nextToken = encodeToken(end)
	}

	return resFolder[start:end], nextToken, nil
}

func encodeToken(index int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", index)))
}

func decodeToken(token string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(decoded))
}

func GetAllFoldersToken(req *FetchFolderRequestToken) (*FetchFolderResponseToken, error) {
	pageSize := req.PageSize

	folders, nextToken, err := FetchAllFoldersByOrgIDToken(req.OrgID, req.Token, pageSize)
	if err != nil {
		return nil, err
	}

	response := &FetchFolderResponseToken{
		Folders: folders,
		Token:   nextToken,
	}

	return response, nil
}
