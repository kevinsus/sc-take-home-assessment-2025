package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	orgID2 := uuid.FromStringOrNil("d3c847c6-8f5f-40c4-8a97-2ec06d5d4836")
	sampleData := []folder.Folder{
		{
			Name: "alpha",
			OrgId: orgID1,
			Paths: "alpha",
		},
		{
			Name: "bravo",
			OrgId: orgID1,
			Paths: "alpha.bravo",
		},
		{
			Name: "charlie",
			OrgId: orgID1,
			Paths: "alpha.bravo.charlie",
		},
		{
			Name: "delta",
			OrgId: orgID1,
			Paths: "alpha.delta",
		},
		{
			Name: "echo",
			OrgId: orgID1,
			Paths: "echo",
		},
		{
			Name: "foxtrot",
			OrgId: orgID2,
			Paths: "foxtrot",
		},
	}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "alpha",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{
				{
					Name: "alpha",
					OrgId: orgID1,
					Paths: "alpha",
				},
				{
					Name: "bravo",
					OrgId: orgID1,
					Paths: "alpha.bravo",
				},
				{
					Name: "charlie",
					OrgId: orgID1,
					Paths: "alpha.bravo.charlie",
				},
				{
					Name: "delta",
					OrgId: orgID1,
					Paths: "alpha.delta",
				},
				{
					Name: "echo",
					OrgId: orgID1,
					Paths: "echo",
				},
			},
		},
		{
			name: "foxtrot",
			orgID: orgID2,
			folders: sampleData,
			want: []folder.Folder{
				{
					Name: "foxtrot",
					OrgId: orgID2,
					Paths: "foxtrot",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t,tt.want, get, "Test Unsuccessfull: %s", tt.name)
		})
	}
}



func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	orgID2 := uuid.FromStringOrNil("d3c847c6-8f5f-40c4-8a97-2ec06d5d4836")
	sampleData := []folder.Folder{
		{
			Name: "alpha",
			OrgId: orgID1,
			Paths: "alpha",
		},
		{
			Name: "bravo",
			OrgId: orgID1,
			Paths: "alpha.bravo",
		},
		{
			Name: "charlie",
			OrgId: orgID1,
			Paths: "alpha.bravo.charlie",
		},
		{
			Name: "delta",
			OrgId: orgID1,
			Paths: "alpha.delta",
		},
		{
			Name: "echo",
			OrgId: orgID1,
			Paths: "echo",
		},
		{
			Name: "foxtrot",
			OrgId: orgID2,
			Paths: "foxtrot",
		},
	}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "alpha",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{
				{
					Name: "bravo",
					OrgId: orgID1,
					Paths: "alpha.bravo",
				},
				{
					Name: "charlie",
					OrgId: orgID1,
					Paths: "alpha.bravo.charlie",
				},
				{
					Name: "delta",
					OrgId: orgID1,
					Paths: "alpha.delta",
				},
			},
		},
		{
			name: "bravo",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{
				{
					Name: "charlie",
					OrgId: orgID1,
					Paths: "alpha.bravo.charlie",
				},
			},
		},
		{
			name: "charlie",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{},
		},
		{
			name: "echo",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{},
		},
		{
			name: "invalid_folder",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{},
		},
		{
			name: "foxtrot",
			orgID: orgID1,
			folders: sampleData,
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.name)
			assert.Equal(t,tt.want, get, "Test Unsuccessfull: %s", tt.name)
		})
	}
}

