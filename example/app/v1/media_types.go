//************************************************************************//
// API "congo" version v1: Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "github.com/raphael/goa"

// A response to a CFP
// Identifier: application/vnd.proposal+json
type Proposal struct {
	// Response abstract
	Abstract *string
	// Response detail
	Detail *string
	// API href of user
	Href *string
	// ID of user
	ID *int
	// Response title
	Title *string
}

// Proposal views
type ProposalViewEnum string

const (
	// Proposal default view
	ProposalDefaultView ProposalViewEnum = "default"
	// Proposal link view
	ProposalLinkView ProposalViewEnum = "link"
)

// LoadProposal loads raw data into an instance of Proposal
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposal(raw interface{}) (res *Proposal, err error) {
	res, err = UnmarshalProposal(raw, err)
	return
}

// Dump produces raw data from an instance of Proposal running all the
// validations. See LoadProposal for the definition of raw data.
func (mt *Proposal) Dump(view ProposalViewEnum) (res map[string]interface{}, err error) {
	if view == ProposalDefaultView {
		res, err = MarshalProposal(mt, err)
	}
	if view == ProposalLinkView {
		res, err = MarshalProposalLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Proposal) Validate() (err error) {
	if mt.Abstract != nil {
		if len(*mt.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 50, true, err)
		}
	}
	if mt.Abstract != nil {
		if len(*mt.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 500, false, err)
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 100, true, err)
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 2000, false, err)
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 10, true, err)
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 200, false, err)
		}
	}
	return
}

// MarshalProposal validates and renders an instance of Proposal into a interface{}
// using view "default".
func MarshalProposal(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp129 := map[string]interface{}{
		"abstract": source.Abstract,
		"detail":   source.Detail,
		"href":     source.Href,
		"id":       source.ID,
		"title":    source.Title,
	}
	target = tmp129
	return
}

// MarshalProposalLink validates and renders an instance of Proposal into a interface{}
// using view "link".
func MarshalProposalLink(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp130 := map[string]interface{}{
		"href":  source.Href,
		"id":    source.ID,
		"title": source.Title,
	}
	target = tmp130
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp131 string
			if val, ok := v.(string); ok {
				tmp131 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp131
		}
		if v, ok := val["detail"]; ok {
			var tmp132 string
			if val, ok := v.(string); ok {
				tmp132 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			target.Detail = &tmp132
		}
		if v, ok := val["href"]; ok {
			var tmp133 string
			if val, ok := v.(string); ok {
				tmp133 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp133
		}
		if v, ok := val["id"]; ok {
			var tmp134 int
			if f, ok := v.(float64); ok {
				tmp134 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp134
		}
		if v, ok := val["title"]; ok {
			var tmp135 string
			if val, ok := v.(string); ok {
				tmp135 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			target.Title = &tmp135
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// ProposalCollection media type
// Identifier: application/vnd.proposal+json; type=collection
type ProposalCollection []*Proposal

// LoadProposalCollection loads raw data into an instance of ProposalCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposalCollection(raw interface{}) (res ProposalCollection, err error) {
	res, err = UnmarshalProposalCollection(raw, err)
	return
}

// Dump produces raw data from an instance of ProposalCollection running all the
// validations. See LoadProposalCollection for the definition of raw data.
func (mt ProposalCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp136 := range mt {
		var tmp137 map[string]interface{}
		tmp137, err = MarshalProposal(tmp136, err)
		res[i] = tmp137
	}
	return
}

// Validate validates the media type instance.
func (mt ProposalCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Abstract != nil {
			if len(*e.Abstract) < 50 {
				err = goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 50, true, err)
			}
		}
		if e.Abstract != nil {
			if len(*e.Abstract) > 500 {
				err = goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 500, false, err)
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) < 100 {
				err = goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 100, true, err)
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) > 2000 {
				err = goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 2000, false, err)
			}
		}
		if e.Title != nil {
			if len(*e.Title) < 10 {
				err = goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 10, true, err)
			}
		}
		if e.Title != nil {
			if len(*e.Title) > 200 {
				err = goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 200, false, err)
			}
		}
	}
	return
}

// MarshalProposalCollection validates and renders an instance of ProposalCollection into a interface{}
// using view "default".
func MarshalProposalCollection(source ProposalCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalProposal(res, err)
	}
	return
}

// UnmarshalProposalCollection unmarshals and validates a raw interface{} into an instance of ProposalCollection
func UnmarshalProposalCollection(source interface{}, inErr error) (target ProposalCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Proposal, len(val))
		for tmp138, v := range val {
			target[tmp138], err = UnmarshalProposal(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// A review is submitted by a reviewer
// Identifier: application/vnd.review+json
type Review struct {
	// Review comments
	Comment *string
	// API href of user
	Href *string
	// ID of user
	ID *int
	// Rating of proposal, from 1-5
	Rating *int
}

// Review views
type ReviewViewEnum string

const (
	// Review default view
	ReviewDefaultView ReviewViewEnum = "default"
	// Review link view
	ReviewLinkView ReviewViewEnum = "link"
)

// LoadReview loads raw data into an instance of Review
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReview(raw interface{}) (res *Review, err error) {
	res, err = UnmarshalReview(raw, err)
	return
}

// Dump produces raw data from an instance of Review running all the
// validations. See LoadReview for the definition of raw data.
func (mt *Review) Dump(view ReviewViewEnum) (res map[string]interface{}, err error) {
	if view == ReviewDefaultView {
		res, err = MarshalReview(mt, err)
	}
	if view == ReviewLinkView {
		res, err = MarshalReviewLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Review) Validate() (err error) {
	if mt.Comment != nil {
		if len(*mt.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 10, true, err)
		}
	}
	if mt.Comment != nil {
		if len(*mt.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 200, false, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	return
}

// MarshalReview validates and renders an instance of Review into a interface{}
// using view "default".
func MarshalReview(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp139 := map[string]interface{}{
		"comment": source.Comment,
		"href":    source.Href,
		"id":      source.ID,
		"rating":  source.Rating,
	}
	target = tmp139
	return
}

// MarshalReviewLink validates and renders an instance of Review into a interface{}
// using view "link".
func MarshalReviewLink(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp140 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp140
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["comment"]; ok {
			var tmp141 string
			if val, ok := v.(string); ok {
				tmp141 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			target.Comment = &tmp141
		}
		if v, ok := val["href"]; ok {
			var tmp142 string
			if val, ok := v.(string); ok {
				tmp142 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp142
		}
		if v, ok := val["id"]; ok {
			var tmp143 int
			if f, ok := v.(float64); ok {
				tmp143 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp143
		}
		if v, ok := val["rating"]; ok {
			var tmp144 int
			if f, ok := v.(float64); ok {
				tmp144 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			target.Rating = &tmp144
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// ReviewCollection media type
// Identifier: application/vnd.review+json; type=collection
type ReviewCollection []*Review

// LoadReviewCollection loads raw data into an instance of ReviewCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReviewCollection(raw interface{}) (res ReviewCollection, err error) {
	res, err = UnmarshalReviewCollection(raw, err)
	return
}

// Dump produces raw data from an instance of ReviewCollection running all the
// validations. See LoadReviewCollection for the definition of raw data.
func (mt ReviewCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp145 := range mt {
		var tmp146 map[string]interface{}
		tmp146, err = MarshalReview(tmp145, err)
		res[i] = tmp146
	}
	return
}

// Validate validates the media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
	}
	return
}

// MarshalReviewCollection validates and renders an instance of ReviewCollection into a interface{}
// using view "default".
func MarshalReviewCollection(source ReviewCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalReview(res, err)
	}
	return
}

// UnmarshalReviewCollection unmarshals and validates a raw interface{} into an instance of ReviewCollection
func UnmarshalReviewCollection(source interface{}, inErr error) (target ReviewCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Review, len(val))
		for tmp147, v := range val {
			target[tmp147], err = UnmarshalReview(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// A user belonging to a tenant account
// Identifier: application/vnd.user+json
type User struct {
	// Biography of user
	Bio *string
	// City of residence
	City *string
	// Country of residence
	Country *string
	// Email address of user
	Email *string
	// First name of user
	Firstname *string
	// API href of user
	Href *string
	// ID of user
	ID *int
	// Last name of user
	Lastname *string
	// Role of user
	Role *string
	// State of residence
	State *string
}

// User views
type UserViewEnum string

const (
	// User default view
	UserDefaultView UserViewEnum = "default"
	// User link view
	UserLinkView UserViewEnum = "link"
)

// LoadUser loads raw data into an instance of User
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUser(raw interface{}) (res *User, err error) {
	res, err = UnmarshalUser(raw, err)
	return
}

// Dump produces raw data from an instance of User running all the
// validations. See LoadUser for the definition of raw data.
func (mt *User) Dump(view UserViewEnum) (res map[string]interface{}, err error) {
	if view == UserDefaultView {
		res, err = MarshalUser(mt, err)
	}
	if view == UserLinkView {
		res, err = MarshalUserLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if mt.Bio != nil {
		if len(*mt.Bio) > 500 {
			err = goa.InvalidLengthError(`response.bio`, *mt.Bio, len(*mt.Bio), 500, false, err)
		}
	}
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2, err)
		}
	}
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
// using view "default".
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp148 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"href":      source.Href,
		"id":        source.ID,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp148
	return
}

// MarshalUserLink validates and renders an instance of User into a interface{}
// using view "link".
func MarshalUserLink(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp149 := map[string]interface{}{
		"email": source.Email,
		"href":  source.Href,
		"id":    source.ID,
	}
	target = tmp149
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["bio"]; ok {
			var tmp150 string
			if val, ok := v.(string); ok {
				tmp150 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			target.Bio = &tmp150
		}
		if v, ok := val["city"]; ok {
			var tmp151 string
			if val, ok := v.(string); ok {
				tmp151 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = &tmp151
		}
		if v, ok := val["country"]; ok {
			var tmp152 string
			if val, ok := v.(string); ok {
				tmp152 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp152
		}
		if v, ok := val["email"]; ok {
			var tmp153 string
			if val, ok := v.(string); ok {
				tmp153 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp153
		}
		if v, ok := val["firstname"]; ok {
			var tmp154 string
			if val, ok := v.(string); ok {
				tmp154 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp154
		}
		if v, ok := val["href"]; ok {
			var tmp155 string
			if val, ok := v.(string); ok {
				tmp155 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp155
		}
		if v, ok := val["id"]; ok {
			var tmp156 int
			if f, ok := v.(float64); ok {
				tmp156 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp156
		}
		if v, ok := val["lastname"]; ok {
			var tmp157 string
			if val, ok := v.(string); ok {
				tmp157 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp157
		}
		if v, ok := val["role"]; ok {
			var tmp158 string
			if val, ok := v.(string); ok {
				tmp158 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = &tmp158
		}
		if v, ok := val["state"]; ok {
			var tmp159 string
			if val, ok := v.(string); ok {
				tmp159 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = &tmp159
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// UserCollection media type
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// LoadUserCollection loads raw data into an instance of UserCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUserCollection(raw interface{}) (res UserCollection, err error) {
	res, err = UnmarshalUserCollection(raw, err)
	return
}

// Dump produces raw data from an instance of UserCollection running all the
// validations. See LoadUserCollection for the definition of raw data.
func (mt UserCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp160 := range mt {
		var tmp161 map[string]interface{}
		tmp161, err = MarshalUser(tmp160, err)
		res[i] = tmp161
	}
	return
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Bio != nil {
			if len(*e.Bio) > 500 {
				err = goa.InvalidLengthError(`response[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
	}
	return
}

// MarshalUserCollection validates and renders an instance of UserCollection into a interface{}
// using view "default".
func MarshalUserCollection(source UserCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalUser(res, err)
	}
	return
}

// UnmarshalUserCollection unmarshals and validates a raw interface{} into an instance of UserCollection
func UnmarshalUserCollection(source interface{}, inErr error) (target UserCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*User, len(val))
		for tmp162, v := range val {
			target[tmp162], err = UnmarshalUser(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}