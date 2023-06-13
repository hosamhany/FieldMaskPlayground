package datastoragepoc_test

import (
	"fmt"
	"regexp"

	"github.com/hosamhany/datastoragepoc/testproto"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
)

var reSpaces = regexp.MustCompile(`\s+`)

// ExampleFilter_update_request illustrates an API endpoint that updates an existing entity.
// The request to that endpoint provides a field mask that should be used to update the entity.
func ExampleFilter_update_request() {
	// Assuming the profile entity is loaded from a database.
	profile := &testproto.Profile{
		User: &testproto.User{
			UserId: 64,
			Name:   "user name",
		},
		Photo: &testproto.Photo{
			PhotoId: 2,
			Path:    "photo path",
			Dimensions: &testproto.Dimensions{
				Width:  100,
				Height: 120,
			},
		},
		LoginTimestamps: []int64{1, 2, 3},
	}
	// An API request from an API user.
	updateProfileRequest := &testproto.UpdateProfileRequest{
		Profile: &testproto.Profile{
			User: &testproto.User{
				Name: "new user name",
			},
			Photo: &testproto.Photo{
				Path: "new photo path",
				Dimensions: &testproto.Dimensions{
					Width: 50,
				},
			},
			LoginTimestamps: []int64{4, 5}},
		Fieldmask: &field_mask.FieldMask{
			Paths: []string{"user.name", "photo.path", "photo.dimensions.width", "login_timestamps"}},
	}
	// Normalize and validate the field mask before using it.
	updateProfileRequest.Fieldmask.Normalize()
	if !updateProfileRequest.Fieldmask.IsValid(profile) {
		// Return an error.
		panic("invalid field mask")
	}
	// Now that the request is vetted we can merge it with the profile entity.
	proto.Merge(profile, updateProfileRequest.GetProfile())
	// The profile can now be saved in a database.
	fmt.Println(reSpaces.ReplaceAllString(profile.String(), " "))
	// Output: user:{user_id:64 name:"new user name"} photo:{photo_id:2 path:"new photo path" dimensions:{width:50 height:120}} login_timestamps:1 login_timestamps:2 login_timestamps:3 login_timestamps:4 login_timestamps:5
}
