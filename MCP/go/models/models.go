package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DescribeStreamInput represents the DescribeStreamInput schema from the OpenAPI specification
type DescribeStreamInput struct {
	Streamarn interface{} `json:"StreamArn"`
	Exclusivestartshardid interface{} `json:"ExclusiveStartShardId,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
}

// SequenceNumberRange represents the SequenceNumberRange schema from the OpenAPI specification
type SequenceNumberRange struct {
	Endingsequencenumber interface{} `json:"EndingSequenceNumber,omitempty"`
	Startingsequencenumber interface{} `json:"StartingSequenceNumber,omitempty"`
}

// ListStreamsInput represents the ListStreamsInput schema from the OpenAPI specification
type ListStreamsInput struct {
	Tablename interface{} `json:"TableName,omitempty"`
	Exclusivestartstreamarn interface{} `json:"ExclusiveStartStreamArn,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
}

// Stream represents the Stream schema from the OpenAPI specification
type Stream struct {
	Streamarn interface{} `json:"StreamArn,omitempty"`
	Streamlabel interface{} `json:"StreamLabel,omitempty"`
	Tablename interface{} `json:"TableName,omitempty"`
}

// GetRecordsInput represents the GetRecordsInput schema from the OpenAPI specification
type GetRecordsInput struct {
	Limit interface{} `json:"Limit,omitempty"`
	Sharditerator interface{} `json:"ShardIterator"`
}

// KeySchemaElement represents the KeySchemaElement schema from the OpenAPI specification
type KeySchemaElement struct {
	Attributename interface{} `json:"AttributeName"`
	Keytype interface{} `json:"KeyType"`
}

// GetRecordsOutput represents the GetRecordsOutput schema from the OpenAPI specification
type GetRecordsOutput struct {
	Records interface{} `json:"Records,omitempty"`
	Nextsharditerator interface{} `json:"NextShardIterator,omitempty"`
}

// MapAttributeValue represents the MapAttributeValue schema from the OpenAPI specification
type MapAttributeValue struct {
}

// StreamDescription represents the StreamDescription schema from the OpenAPI specification
type StreamDescription struct {
	Streamarn interface{} `json:"StreamArn,omitempty"`
	Keyschema interface{} `json:"KeySchema,omitempty"`
	Creationrequestdatetime interface{} `json:"CreationRequestDateTime,omitempty"`
	Lastevaluatedshardid interface{} `json:"LastEvaluatedShardId,omitempty"`
	Streamlabel interface{} `json:"StreamLabel,omitempty"`
	Tablename interface{} `json:"TableName,omitempty"`
	Streamstatus interface{} `json:"StreamStatus,omitempty"`
	Streamviewtype interface{} `json:"StreamViewType,omitempty"`
	Shards interface{} `json:"Shards,omitempty"`
}

// StreamRecord represents the StreamRecord schema from the OpenAPI specification
type StreamRecord struct {
	Streamviewtype interface{} `json:"StreamViewType,omitempty"`
	Approximatecreationdatetime interface{} `json:"ApproximateCreationDateTime,omitempty"`
	Keys interface{} `json:"Keys,omitempty"`
	Newimage interface{} `json:"NewImage,omitempty"`
	Oldimage interface{} `json:"OldImage,omitempty"`
	Sequencenumber interface{} `json:"SequenceNumber,omitempty"`
	Sizebytes interface{} `json:"SizeBytes,omitempty"`
}

// GetShardIteratorInput represents the GetShardIteratorInput schema from the OpenAPI specification
type GetShardIteratorInput struct {
	Streamarn interface{} `json:"StreamArn"`
	Sequencenumber interface{} `json:"SequenceNumber,omitempty"`
	Shardid interface{} `json:"ShardId"`
	Sharditeratortype interface{} `json:"ShardIteratorType"`
}

// Identity represents the Identity schema from the OpenAPI specification
type Identity struct {
	TypeField interface{} `json:"Type,omitempty"`
	Principalid interface{} `json:"PrincipalId,omitempty"`
}

// ListStreamsOutput represents the ListStreamsOutput schema from the OpenAPI specification
type ListStreamsOutput struct {
	Lastevaluatedstreamarn interface{} `json:"LastEvaluatedStreamArn,omitempty"`
	Streams interface{} `json:"Streams,omitempty"`
}

// DescribeStreamOutput represents the DescribeStreamOutput schema from the OpenAPI specification
type DescribeStreamOutput struct {
	Streamdescription interface{} `json:"StreamDescription,omitempty"`
}

// AttributeValue represents the AttributeValue schema from the OpenAPI specification
type AttributeValue struct {
	BoolField interface{} `json:"BOOL,omitempty"`
	Bs interface{} `json:"BS,omitempty"`
	L interface{} `json:"L,omitempty"`
	M interface{} `json:"M,omitempty"`
	Null interface{} `json:"NULL,omitempty"`
	S interface{} `json:"S,omitempty"`
	Ss interface{} `json:"SS,omitempty"`
	Ns interface{} `json:"NS,omitempty"`
	B interface{} `json:"B,omitempty"`
	N interface{} `json:"N,omitempty"`
}

// Shard represents the Shard schema from the OpenAPI specification
type Shard struct {
	Shardid interface{} `json:"ShardId,omitempty"`
	Parentshardid interface{} `json:"ParentShardId,omitempty"`
	Sequencenumberrange interface{} `json:"SequenceNumberRange,omitempty"`
}

// Record represents the Record schema from the OpenAPI specification
type Record struct {
	Useridentity interface{} `json:"userIdentity,omitempty"`
	Awsregion interface{} `json:"awsRegion,omitempty"`
	Dynamodb interface{} `json:"dynamodb,omitempty"`
	Eventid interface{} `json:"eventID,omitempty"`
	Eventname interface{} `json:"eventName,omitempty"`
	Eventsource interface{} `json:"eventSource,omitempty"`
	Eventversion interface{} `json:"eventVersion,omitempty"`
}

// AttributeMap represents the AttributeMap schema from the OpenAPI specification
type AttributeMap struct {
}

// GetShardIteratorOutput represents the GetShardIteratorOutput schema from the OpenAPI specification
type GetShardIteratorOutput struct {
	Sharditerator interface{} `json:"ShardIterator,omitempty"`
}
