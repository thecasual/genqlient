package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNoFragmentsQueryRandomItemArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNoFragmentsQueryRandomItemContent is implemented by the following types:
// InterfaceNoFragmentsQueryRandomItemArticle
// InterfaceNoFragmentsQueryRandomItemVideo
// InterfaceNoFragmentsQueryRandomItemTopic
//
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryRandomItemContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent()
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

func (v *InterfaceNoFragmentsQueryRandomItemArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemArticle) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemArticle) GetName() string { return v.Name }

func (v *InterfaceNoFragmentsQueryRandomItemVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemVideo) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemVideo) GetName() string { return v.Name }

func (v *InterfaceNoFragmentsQueryRandomItemTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemTopic) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemContent.
func (v *InterfaceNoFragmentsQueryRandomItemTopic) GetName() string { return v.Name }

func __unmarshalInterfaceNoFragmentsQueryRandomItemContent(v *InterfaceNoFragmentsQueryRandomItemContent, m json.RawMessage) error {
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
		*v = new(InterfaceNoFragmentsQueryRandomItemArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryRandomItemVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryRandomItemTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNoFragmentsQueryRandomItemContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNoFragmentsQueryRandomItemTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemWithTypeNameContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNoFragmentsQueryRandomItemWithTypeNameContent is implemented by the following types:
// InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle
// InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo
// InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic
//
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent()
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

func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle) GetName() string { return v.Name }

func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo) GetName() string { return v.Name }

func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRandomItemWithTypeNameContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryRandomItemWithTypeNameContent.
func (v *InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic) GetName() string { return v.Name }

func __unmarshalInterfaceNoFragmentsQueryRandomItemWithTypeNameContent(v *InterfaceNoFragmentsQueryRandomItemWithTypeNameContent, m json.RawMessage) error {
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
		*v = new(InterfaceNoFragmentsQueryRandomItemWithTypeNameArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNoFragmentsQueryRandomItemWithTypeNameContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryRandomItemWithTypeNameVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryResponse is returned by InterfaceNoFragmentsQuery on success.
type InterfaceNoFragmentsQueryResponse struct {
	Root                   InterfaceNoFragmentsQueryRootTopic                     `json:"root"`
	RandomItem             InterfaceNoFragmentsQueryRandomItemContent             `json:"-"`
	RandomItemWithTypeName InterfaceNoFragmentsQueryRandomItemWithTypeNameContent `json:"-"`
	WithPointer            *InterfaceNoFragmentsQueryWithPointerContent           `json:"-"`
}

func (v *InterfaceNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*InterfaceNoFragmentsQueryResponse
		RandomItem             json.RawMessage `json:"randomItem"`
		RandomItemWithTypeName json.RawMessage `json:"randomItemWithTypeName"`
		WithPointer            json.RawMessage `json:"withPointer"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceNoFragmentsQueryResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.RandomItem
		raw := firstPass.RandomItem
		err = __unmarshalInterfaceNoFragmentsQueryRandomItemContent(
			target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal InterfaceNoFragmentsQueryResponse.RandomItem: %w", err)
		}
	}

	{
		target := &v.RandomItemWithTypeName
		raw := firstPass.RandomItemWithTypeName
		err = __unmarshalInterfaceNoFragmentsQueryRandomItemWithTypeNameContent(
			target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal InterfaceNoFragmentsQueryResponse.RandomItemWithTypeName: %w", err)
		}
	}

	{
		target := &v.WithPointer
		raw := firstPass.WithPointer
		*target = new(InterfaceNoFragmentsQueryWithPointerContent)
		err = __unmarshalInterfaceNoFragmentsQueryWithPointerContent(
			*target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal InterfaceNoFragmentsQueryResponse.WithPointer: %w", err)
		}
	}

	return nil
}

// InterfaceNoFragmentsQueryRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRootTopic struct {
	// ID is documented in the Content interface.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryWithPointerArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryWithPointerArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceNoFragmentsQueryWithPointerContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNoFragmentsQueryWithPointerContent is implemented by the following types:
// InterfaceNoFragmentsQueryWithPointerArticle
// InterfaceNoFragmentsQueryWithPointerVideo
// InterfaceNoFragmentsQueryWithPointerTopic
//
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryWithPointerContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() *testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() *string
}

func (v *InterfaceNoFragmentsQueryWithPointerArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerArticle) GetId() *testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerArticle) GetName() *string { return v.Name }

func (v *InterfaceNoFragmentsQueryWithPointerVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerVideo) GetId() *testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerVideo) GetName() *string { return v.Name }

func (v *InterfaceNoFragmentsQueryWithPointerTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryWithPointerContent() {
}

// GetTypename is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerTopic) GetId() *testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface InterfaceNoFragmentsQueryWithPointerContent.
func (v *InterfaceNoFragmentsQueryWithPointerTopic) GetName() *string { return v.Name }

func __unmarshalInterfaceNoFragmentsQueryWithPointerContent(v *InterfaceNoFragmentsQueryWithPointerContent, m json.RawMessage) error {
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
		*v = new(InterfaceNoFragmentsQueryWithPointerArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceNoFragmentsQueryWithPointerVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceNoFragmentsQueryWithPointerTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceNoFragmentsQueryWithPointerContent: "%v"`, tn.TypeName)
	}
}

// InterfaceNoFragmentsQueryWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryWithPointerTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceNoFragmentsQueryWithPointerVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryWithPointerVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

func InterfaceNoFragmentsQuery(
	client graphql.Client,
) (*InterfaceNoFragmentsQueryResponse, error) {
	var err error

	var retval InterfaceNoFragmentsQueryResponse
	err = client.MakeRequest(
		nil,
		"InterfaceNoFragmentsQuery",
		`
query InterfaceNoFragmentsQuery {
	root {
		id
		name
	}
	randomItem {
		__typename
		id
		name
	}
	randomItemWithTypeName: randomItem {
		__typename
		id
		name
	}
	withPointer: randomItem {
		__typename
		id
		name
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

