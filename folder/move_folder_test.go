package folder_test

import (
	"testing"
	"errors"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
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
			Paths: "alpha.delta.echo",
		},
		{
			Name: "foxtrot",
			OrgId: orgID2,
			Paths: "foxtrot",
		},
		{
			Name: "golf",
			OrgId: orgID1,
			Paths: "golf",
		},
	}
	tests := [...]struct {
		name    	string
		dst				string
		folders 	[]folder.Folder
		wantError	error
		want    	[]folder.Folder
	}{
		{
			name: "bravo",
			dst: "delta",
			folders: sampleData,
			wantError: nil,
			want: []folder.Folder{
				{
					Name: "alpha",
					OrgId: orgID1,
					Paths: "alpha",
				},
				{
					Name: "bravo",
					OrgId: orgID1,
					Paths: "alpha.delta.bravo",
				},
				{
					Name: "charlie",
					OrgId: orgID1,
					Paths: "alpha.delta.bravo.charlie",
				},
				{
					Name: "delta",
					OrgId: orgID1,
					Paths: "alpha.delta",
				},
				{
					Name: "echo",
					OrgId: orgID1,
					Paths: "alpha.delta.echo",
				},
				{
					Name: "foxtrot",
					OrgId: orgID2,
					Paths: "foxtrot",
				},
				{
					Name: "golf",
					OrgId: orgID1,
					Paths: "golf",
				},
			},
		},
		{
			name: "bravo",
			dst: "golf",
			folders: sampleData,
			wantError: nil,
			want: []folder.Folder{
				{
					Name: "alpha",
					OrgId: orgID1,
					Paths: "alpha",
				},
				{
					Name: "bravo",
					OrgId: orgID1,
					Paths: "golf.bravo",
				},
				{
					Name: "charlie",
					OrgId: orgID1,
					Paths: "golf.bravo.charlie",
				},
				{
					Name: "delta",
					OrgId: orgID1,
					Paths: "alpha.delta",
				},
				{
					Name: "echo",
					OrgId: orgID1,
					Paths: "alpha.delta.echo",
				},
				{
					Name: "foxtrot",
					OrgId: orgID2,
					Paths: "foxtrot",
				},
				{
					Name: "golf",
					OrgId: orgID1,
					Paths: "golf",
				},
			},
		},
		{
			name: "bravo",
			dst: "charlie",
			folders: sampleData,
			wantError: errors.New("Cannot move a folder to a child of itself"),
			want: []folder.Folder{},
		},
		{
			name: "bravo",
			dst: "bravo",
			folders: sampleData,
			wantError: errors.New("Cannot move a folder to itself"),
			want: []folder.Folder{},
		},
		{
			name: "bravo",
			dst: "foxtrot",
			folders: sampleData,
			wantError: errors.New("Cannot move a folder to a different organization"),
			want: []folder.Folder{},
		},
		{
			name: "invalid_folder",
			dst: "delta",
			folders: sampleData,
			wantError: errors.New("Source folder does not exist"),
			want: []folder.Folder{},
		},
		{
			name: "bravo",
			dst: "invalid_folder",
			folders: sampleData,
			wantError: errors.New("Destination folder does not exist"),
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.name, tt.dst)
			if tt.wantError != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t,tt.want, get, "Test Unsuccessfull: %s", tt.name)
			}
		})
	}
}
