package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type FetchFolderRequestToken struct {
	OrgID    uuid.UUID `json:"org_id"`
	Token    string    `json:"token"`
	PageSize int       `json:"page_size"`
}

type FetchFolderResponseToken struct {
	Folders []*Folder `json:"folders"`
	Token   string    `json:"token"`
}
