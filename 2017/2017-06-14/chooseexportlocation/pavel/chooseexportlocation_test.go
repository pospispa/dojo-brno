package chooseexportlocation

import (
	"reflect"
	"testing"
)

const (
	validPath              = "ip://directory"
	preferredPath          = "ip://preferred/directory"
	emptyPath              = ""
	spacesOnlyPath         = "  	  "
	shareExportLocationID1 = "123456-1"
	shareExportLocationID2 = "1234567-1"
	shareExportLocationID3 = "1234567-2"
	shareExportLocationID4 = "7654321-1"
	shareID1               = "123456"
	shareID2               = "1234567"
	shareID3               = "765321"
	shareID4               = "654321"
)

func TestChooseExportLocationSuccess(t *testing.T) {
	tests := []struct {
		testCaseName string
		locs         []ExportLocation
		shareID      string
		want         ExportLocation
	}{
		{
			testCaseName: "Match first item:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       false,
				},
			},
			shareID: shareID1,
			want: ExportLocation{
				Path:            validPath,
				ShareInstanceID: shareID1,
				IsAdminOnly:     false,
				ID:              shareExportLocationID1,
				Preferred:       false,
			},
		},
		{
			testCaseName: "Match second item:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID3,
					IsAdminOnly:     false,
					ID:              shareExportLocationID4,
					Preferred:       false,
				},
				{
					Path:            validPath,
					ShareInstanceID: shareID2,
					IsAdminOnly:     false,
					ID:              shareExportLocationID2,
					Preferred:       false,
				},
			},
			shareID: shareID2,
			want: ExportLocation{
				Path:            validPath,
				ShareInstanceID: shareID2,
				IsAdminOnly:     false,
				ID:              shareExportLocationID2,
				Preferred:       false,
			},
		},
		{
			testCaseName: "Match preferred location:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID3,
					IsAdminOnly:     false,
					ID:              shareExportLocationID4,
					Preferred:       false,
				},
				{
					Path:            validPath,
					ShareInstanceID: shareID2,
					IsAdminOnly:     false,
					ID:              shareExportLocationID2,
					Preferred:       false,
				},
				{
					Path:            preferredPath,
					ShareInstanceID: shareID2,
					IsAdminOnly:     false,
					ID:              shareExportLocationID3,
					Preferred:       true,
				},
			},
			shareID: shareID2,
			want: ExportLocation{
				Path:            preferredPath,
				ShareInstanceID: shareID2,
				IsAdminOnly:     false,
				ID:              shareExportLocationID3,
				Preferred:       true,
			},
		},
		{
			testCaseName: "Match first not-preferred location that matches shareID:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID3,
					IsAdminOnly:     false,
					ID:              shareExportLocationID4,
					Preferred:       false,
				},
				{
					Path:            validPath,
					ShareInstanceID: shareID2,
					IsAdminOnly:     false,
					ID:              shareExportLocationID2,
					Preferred:       false,
				},
				{
					Path:            preferredPath,
					ShareInstanceID: shareID2,
					IsAdminOnly:     false,
					ID:              shareExportLocationID3,
					Preferred:       false,
				},
			},
			shareID: shareID2,
			want: ExportLocation{
				Path:            validPath,
				ShareInstanceID: shareID2,
				IsAdminOnly:     false,
				ID:              shareExportLocationID2,
				Preferred:       false,
			},
		},
	}

	for _, tt := range tests {
		if got, err := ChooseExportLocation(tt.locs, tt.shareID); err != nil {
			t.Errorf("%q ChooseExportLocation(%v, %q) = (%v, %q) want (%v, nil)", tt.testCaseName, tt.locs, tt.shareID, got, err.Error(), tt.want)
		} else if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("%q ChooseExportLocation(%v, %q) = (%v, nil) want (%v, nil)", tt.testCaseName, tt.locs, tt.shareID, got, tt.want)
		}
	}
}

func TestChooseExportLocationNotFound(t *testing.T) {
	tests := []struct {
		testCaseName string
		locs         []ExportLocation
		shareID      string
	}{
		{
			testCaseName: "Empty slice:",
			locs:         []ExportLocation{},
			shareID:      shareID1,
		},
		{
			testCaseName: "shareID doesn't match:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       false,
				},
			},
			shareID: shareID4,
		},
		{
			testCaseName: "Locations for admins only:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     true,
					ID:              shareExportLocationID1,
					Preferred:       false,
				},
			},
			shareID: shareID1,
		},
		{
			testCaseName: "Preferred locations for admins only:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     true,
					ID:              shareExportLocationID1,
					Preferred:       true,
				},
			},
			shareID: shareID1,
		},
		{
			testCaseName: "Preferred location only, but shareID doesn't match:",
			locs: []ExportLocation{
				{
					Path:            validPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       true,
				},
			},
			shareID: shareID4,
		},
		{
			testCaseName: "Empty path:",
			locs: []ExportLocation{
				{
					Path:            emptyPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       false,
				},
			},
			shareID: shareID1,
		},
		{
			testCaseName: "Empty path in preferred location:",
			locs: []ExportLocation{
				{
					Path:            emptyPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       true,
				},
			},
			shareID: shareID1,
		},
		{
			testCaseName: "Path containing spaces only:",
			locs: []ExportLocation{
				{
					Path:            spacesOnlyPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       false,
				},
			},
			shareID: shareID1,
		},
		{
			testCaseName: "Preferred path containing spaces only:",
			locs: []ExportLocation{
				{
					Path:            spacesOnlyPath,
					ShareInstanceID: shareID1,
					IsAdminOnly:     false,
					ID:              shareExportLocationID1,
					Preferred:       true,
				},
			},
			shareID: shareID1,
		},
	}
	for _, tt := range tests {
		if got, err := ChooseExportLocation(tt.locs, tt.shareID); err == nil {
			t.Errorf("%q ChooseExportLocation(%v, %q) = (%v, nil) want (\"N/A\", \"an error\")", tt.testCaseName, tt.locs, tt.shareID, got)
		}
	}
}
