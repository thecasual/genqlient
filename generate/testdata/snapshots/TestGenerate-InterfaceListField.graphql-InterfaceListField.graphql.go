// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceListFieldResponse is returned by InterfaceListField on success.
type InterfaceListFieldResponse struct {
	Root        InterfaceListFieldRootTopic         `json:"root"`
	WithPointer *InterfaceListFieldWithPointerTopic `json:"withPointer"`
}

// GetRoot returns InterfaceListFieldResponse.Root, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldResponse) GetRoot() InterfaceListFieldRootTopic { return v.Root }

// GetWithPointer returns InterfaceListFieldResponse.WithPointer, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldResponse) GetWithPointer() *InterfaceListFieldWithPointerTopic {
	return v.WithPointer
}

// InterfaceListFieldRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListFieldRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                  `json:"id"`
	Name     string                                       `json:"name"`
	Children []InterfaceListFieldRootTopicChildrenContent `json:"-"`
}

// GetId returns InterfaceListFieldRootTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldRootTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopic) GetName() string { return v.Name }

// GetChildren returns InterfaceListFieldRootTopic.Children, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopic) GetChildren() []InterfaceListFieldRootTopicChildrenContent {
	return v.Children
}

func (v *InterfaceListFieldRootTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InterfaceListFieldRootTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceListFieldRootTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Children
		src := firstPass.Children
		*dst = make(
			[]InterfaceListFieldRootTopicChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalInterfaceListFieldRootTopicChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"unable to unmarshal InterfaceListFieldRootTopic.Children: %w", err)
				}
			}
		}
	}
	return nil
}

type __premarshalInterfaceListFieldRootTopic struct {
	Id testutil.ID `json:"id"`

	Name string `json:"name"`

	Children []json.RawMessage `json:"children"`
}

func (v *InterfaceListFieldRootTopic) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InterfaceListFieldRootTopic) __premarshalJSON() (*__premarshalInterfaceListFieldRootTopic, error) {
	var retval __premarshalInterfaceListFieldRootTopic

	retval.Id = v.Id
	retval.Name = v.Name
	{

		dst := &retval.Children
		src := v.Children
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = __marshalInterfaceListFieldRootTopicChildrenContent(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to marshal InterfaceListFieldRootTopic.Children: %w", err)
			}
		}
	}
	return &retval, nil
}

// InterfaceListFieldRootTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceListFieldRootTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListFieldRootTopicChildrenArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenArticle) GetTypename() string { return v.Typename }

// GetId returns InterfaceListFieldRootTopicChildrenArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenArticle) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldRootTopicChildrenArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenArticle) GetName() string { return v.Name }

// InterfaceListFieldRootTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceListFieldRootTopicChildrenContent is implemented by the following types:
// InterfaceListFieldRootTopicChildrenArticle
// InterfaceListFieldRootTopicChildrenTopic
// InterfaceListFieldRootTopicChildrenVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListFieldRootTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceListFieldRootTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *InterfaceListFieldRootTopicChildrenArticle) implementsGraphQLInterfaceInterfaceListFieldRootTopicChildrenContent() {
}
func (v *InterfaceListFieldRootTopicChildrenTopic) implementsGraphQLInterfaceInterfaceListFieldRootTopicChildrenContent() {
}
func (v *InterfaceListFieldRootTopicChildrenVideo) implementsGraphQLInterfaceInterfaceListFieldRootTopicChildrenContent() {
}

func __unmarshalInterfaceListFieldRootTopicChildrenContent(b []byte, v *InterfaceListFieldRootTopicChildrenContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceListFieldRootTopicChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceListFieldRootTopicChildrenTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceListFieldRootTopicChildrenVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceListFieldRootTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceListFieldRootTopicChildrenContent(v *InterfaceListFieldRootTopicChildrenContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceListFieldRootTopicChildrenArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListFieldRootTopicChildrenArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListFieldRootTopicChildrenTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListFieldRootTopicChildrenTopic
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListFieldRootTopicChildrenVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListFieldRootTopicChildrenVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceListFieldRootTopicChildrenContent: "%T"`, v)
	}
}

// InterfaceListFieldRootTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListFieldRootTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListFieldRootTopicChildrenTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenTopic) GetTypename() string { return v.Typename }

// GetId returns InterfaceListFieldRootTopicChildrenTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldRootTopicChildrenTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenTopic) GetName() string { return v.Name }

// InterfaceListFieldRootTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceListFieldRootTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListFieldRootTopicChildrenVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenVideo) GetTypename() string { return v.Typename }

// GetId returns InterfaceListFieldRootTopicChildrenVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenVideo) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldRootTopicChildrenVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldRootTopicChildrenVideo) GetName() string { return v.Name }

// InterfaceListFieldWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListFieldWithPointerTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                         `json:"id"`
	Name     string                                              `json:"name"`
	Children []InterfaceListFieldWithPointerTopicChildrenContent `json:"-"`
}

// GetId returns InterfaceListFieldWithPointerTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldWithPointerTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopic) GetName() string { return v.Name }

// GetChildren returns InterfaceListFieldWithPointerTopic.Children, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopic) GetChildren() []InterfaceListFieldWithPointerTopicChildrenContent {
	return v.Children
}

func (v *InterfaceListFieldWithPointerTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InterfaceListFieldWithPointerTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceListFieldWithPointerTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Children
		src := firstPass.Children
		*dst = make(
			[]InterfaceListFieldWithPointerTopicChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalInterfaceListFieldWithPointerTopicChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"unable to unmarshal InterfaceListFieldWithPointerTopic.Children: %w", err)
				}
			}
		}
	}
	return nil
}

type __premarshalInterfaceListFieldWithPointerTopic struct {
	Id testutil.ID `json:"id"`

	Name string `json:"name"`

	Children []json.RawMessage `json:"children"`
}

func (v *InterfaceListFieldWithPointerTopic) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InterfaceListFieldWithPointerTopic) __premarshalJSON() (*__premarshalInterfaceListFieldWithPointerTopic, error) {
	var retval __premarshalInterfaceListFieldWithPointerTopic

	retval.Id = v.Id
	retval.Name = v.Name
	{

		dst := &retval.Children
		src := v.Children
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = __marshalInterfaceListFieldWithPointerTopicChildrenContent(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to marshal InterfaceListFieldWithPointerTopic.Children: %w", err)
			}
		}
	}
	return &retval, nil
}

// InterfaceListFieldWithPointerTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceListFieldWithPointerTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListFieldWithPointerTopicChildrenArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenArticle) GetTypename() string { return v.Typename }

// GetId returns InterfaceListFieldWithPointerTopicChildrenArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenArticle) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldWithPointerTopicChildrenArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenArticle) GetName() string { return v.Name }

// InterfaceListFieldWithPointerTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceListFieldWithPointerTopicChildrenContent is implemented by the following types:
// InterfaceListFieldWithPointerTopicChildrenArticle
// InterfaceListFieldWithPointerTopicChildrenTopic
// InterfaceListFieldWithPointerTopicChildrenVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListFieldWithPointerTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceListFieldWithPointerTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *InterfaceListFieldWithPointerTopicChildrenArticle) implementsGraphQLInterfaceInterfaceListFieldWithPointerTopicChildrenContent() {
}
func (v *InterfaceListFieldWithPointerTopicChildrenTopic) implementsGraphQLInterfaceInterfaceListFieldWithPointerTopicChildrenContent() {
}
func (v *InterfaceListFieldWithPointerTopicChildrenVideo) implementsGraphQLInterfaceInterfaceListFieldWithPointerTopicChildrenContent() {
}

func __unmarshalInterfaceListFieldWithPointerTopicChildrenContent(b []byte, v *InterfaceListFieldWithPointerTopicChildrenContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceListFieldWithPointerTopicChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceListFieldWithPointerTopicChildrenTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceListFieldWithPointerTopicChildrenVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceListFieldWithPointerTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceListFieldWithPointerTopicChildrenContent(v *InterfaceListFieldWithPointerTopicChildrenContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceListFieldWithPointerTopicChildrenArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListFieldWithPointerTopicChildrenArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListFieldWithPointerTopicChildrenTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListFieldWithPointerTopicChildrenTopic
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListFieldWithPointerTopicChildrenVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListFieldWithPointerTopicChildrenVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceListFieldWithPointerTopicChildrenContent: "%T"`, v)
	}
}

// InterfaceListFieldWithPointerTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListFieldWithPointerTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListFieldWithPointerTopicChildrenTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenTopic) GetTypename() string { return v.Typename }

// GetId returns InterfaceListFieldWithPointerTopicChildrenTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenTopic) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldWithPointerTopicChildrenTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenTopic) GetName() string { return v.Name }

// InterfaceListFieldWithPointerTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceListFieldWithPointerTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListFieldWithPointerTopicChildrenVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenVideo) GetTypename() string { return v.Typename }

// GetId returns InterfaceListFieldWithPointerTopicChildrenVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenVideo) GetId() testutil.ID { return v.Id }

// GetName returns InterfaceListFieldWithPointerTopicChildrenVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListFieldWithPointerTopicChildrenVideo) GetName() string { return v.Name }

func InterfaceListField(
	client graphql.Client,
) (*InterfaceListFieldResponse, error) {
	req := &graphql.Request{
		OpName: "InterfaceListField",
		Query: `
query InterfaceListField {
	root {
		id
		name
		children {
			__typename
			id
			name
		}
	}
	withPointer: root {
		id
		name
		children {
			__typename
			id
			name
		}
	}
}
`,
	}
	var err error

	var data InterfaceListFieldResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

