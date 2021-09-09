package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNestingResponse is returned by InterfaceNesting on success.
type InterfaceNestingResponse struct {
	Root InterfaceNestingRootTopic `json:"root"`
}

// InterfaceNestingRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                `json:"id"`
	Children []InterfaceNestingRootTopicChildrenContent `json:"-"`
}

func (v *InterfaceNestingRootTopic) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InterfaceNestingRootTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceNestingRootTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Children
		raw := firstPass.Children
		*target = make(
			[]InterfaceNestingRootTopicChildrenContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			err = __unmarshalInterfaceNestingRootTopicChildrenContent(
				target, raw)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal InterfaceNestingRootTopic.Children: %w", err)
			}
		}
	}

	return nil
}

// InterfaceNestingRootTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenContentParentTopic `json:"parent"`
}

// InterfaceNestingRootTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNestingRootTopicChildrenContent is implemented by the following types:
// InterfaceNestingRootTopicChildrenArticle
// InterfaceNestingRootTopicChildrenVideo
// InterfaceNestingRootTopicChildrenTopic
//
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetParent returns the interface-field "parent" from its implementation.
	GetParent() InterfaceNestingRootTopicChildrenContentParentTopic
}

func (v *InterfaceNestingRootTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenArticle) GetId() testutil.ID { return v.Id }

// GetParent is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenArticle) GetParent() InterfaceNestingRootTopicChildrenContentParentTopic {
	return v.Parent
}

func (v *InterfaceNestingRootTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenVideo) GetId() testutil.ID { return v.Id }

// GetParent is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenVideo) GetParent() InterfaceNestingRootTopicChildrenContentParentTopic {
	return v.Parent
}

func (v *InterfaceNestingRootTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenTopic) GetId() testutil.ID { return v.Id }

// GetParent is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenTopic) GetParent() InterfaceNestingRootTopicChildrenContentParentTopic {
	return v.Parent
}

func __unmarshalInterfaceNestingRootTopicChildrenContent(v *InterfaceNestingRootTopicChildrenContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNestingRootTopicChildrenArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNestingRootTopicChildrenVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNestingRootTopicChildrenTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNestingRootTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNestingRootTopicChildrenContentParentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenContentParentTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                                          `json:"id"`
	Children []InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent `json:"-"`
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopic) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InterfaceNestingRootTopicChildrenContentParentTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceNestingRootTopicChildrenContentParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Children
		raw := firstPass.Children
		*target = make(
			[]InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			err = __unmarshalInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent(
				target, raw)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal InterfaceNestingRootTopicChildrenContentParentTopic.Children: %w", err)
			}
		}
	}

	return nil
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent is implemented by the following types:
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic
//
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle) GetId() testutil.ID {
	return v.Id
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo) GetId() testutil.ID {
	return v.Id
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic) GetId() testutil.ID {
	return v.Id
}

func __unmarshalInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent(v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// InterfaceNestingRootTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenContentParentTopic `json:"parent"`
}

// InterfaceNestingRootTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenContentParentTopic `json:"parent"`
}

func InterfaceNesting(
	client graphql.Client,
) (*InterfaceNestingResponse, error) {
	var err error

	var retval InterfaceNestingResponse
	err = client.MakeRequest(
		nil,
		"InterfaceNesting",
		`
query InterfaceNesting {
	root {
		id
		children {
			__typename
			id
			parent {
				id
				children {
					__typename
					id
				}
			}
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

