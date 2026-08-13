package file

import "github.com/rvxt21/bucket-inventory/pkg/dto"

// makeFileResponse converts an internal file record into its transport shape.
// StorageKey is deliberately dropped: the client only ever sees the link.
func makeFileResponse(file dto.File, link string) dto.FileResponse {
	return dto.FileResponse{
		ID:          file.ID,
		Name:        file.Name,
		ContentType: file.ContentType,
		Size:        file.Size,
		CreatedAt:   file.CreatedAt,
		UpdatedAt:   file.UpdatedAt,
		Link:        link,
	}
}
