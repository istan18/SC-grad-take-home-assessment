package folders

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllFolders(t *testing.T) {
	req := &FetchFolderRequest{
		OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
	}

	resp, err := GetAllFolders(req)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	expectedIDs := []uuid.UUID{
		uuid.FromStringOrNil("7ee73e98-b5a7-4ff5-a710-bfd8077ac0a9"),
		uuid.FromStringOrNil("c03cfef0-3256-4e46-8ec9-280a913c8592"),
		uuid.FromStringOrNil("c0d25887-f71a-4176-ac3b-6a9bcea92ba4"),
		uuid.FromStringOrNil("aca43bf2-11c0-4c61-937c-c1b8f752da69"),
	}

	for _, expectedID := range expectedIDs {
		found := false
		for _, folder := range resp.Folders {
			if folder.Id == expectedID {
				found = true
				break
			}
		}
		assert.True(t, found, "Folder with ID %s not found", expectedID)
	}
}

func TestGetAllFolders_NoMatchingOrgID(t *testing.T) {
	req := &FetchFolderRequest{
		OrgID: uuid.Must(uuid.NewV4()),
	}

	resp, err := GetAllFolders(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 0, len(resp.Folders))
}

func TestFetchAllFoldersByOrgIDToken_FirstPage(t *testing.T) {
	orgID := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	pageSize := 2
	folders, nextToken, err := FetchAllFoldersByOrgIDToken(orgID, "", pageSize)
	assert.NoError(t, err)

	assert.Equal(t, pageSize, len(folders))
	assert.NotEmpty(t, nextToken)
}

func TestFetchAllFoldersByOrgIDToken_LastPage(t *testing.T) {
	orgID := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	pageSize := 2

	var allFolders []*Folder
	var token string
	var err error

	for {
		var folders []*Folder
		folders, token, err = FetchAllFoldersByOrgIDToken(orgID, token, pageSize)
		assert.NoError(t, err)
		allFolders = append(allFolders, folders...)
		if token == "" {
			break
		}
	}

	assert.Equal(t, 666, len(allFolders))
}

func TestGetAllFoldersToken(t *testing.T) {
	req := &FetchFolderRequestToken{
		OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		PageSize: 2,
	}

	resp, err := GetAllFoldersToken(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.PageSize, len(resp.Folders))
	assert.NotEmpty(t, resp.Token)

	req.Token = resp.Token
	resp, err = GetAllFoldersToken(req)
	assert.NoError(t, err)
	assert.Equal(t, req.PageSize, len(resp.Folders))
	assert.NotEqual(t, "", resp.Token)
}

func TestGetAllFoldersToken_NoMoreData(t *testing.T) {
	req := &FetchFolderRequestToken{
		OrgID:    uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		PageSize: 5000,
	}

	resp, err := GetAllFoldersToken(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Empty(t, resp.Token)
}
