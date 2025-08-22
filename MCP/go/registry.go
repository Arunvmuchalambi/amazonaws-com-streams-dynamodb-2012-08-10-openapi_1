package main

import (
	"github.com/amazon-dynamodb-streams/mcp-server/config"
	"github.com/amazon-dynamodb-streams/mcp-server/models"
	tools_x_amz_target_dynamodbstreams_20120810_getrecords "github.com/amazon-dynamodb-streams/mcp-server/tools/x_amz_target_dynamodbstreams_20120810_getrecords"
	tools_x_amz_target_dynamodbstreams_20120810_getsharditerator "github.com/amazon-dynamodb-streams/mcp-server/tools/x_amz_target_dynamodbstreams_20120810_getsharditerator"
	tools_x_amz_target_dynamodbstreams_20120810_liststreams "github.com/amazon-dynamodb-streams/mcp-server/tools/x_amz_target_dynamodbstreams_20120810_liststreams"
	tools_x_amz_target_dynamodbstreams_20120810_describestream "github.com/amazon-dynamodb-streams/mcp-server/tools/x_amz_target_dynamodbstreams_20120810_describestream"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_x_amz_target_dynamodbstreams_20120810_getrecords.CreateGetrecordsTool(cfg),
		tools_x_amz_target_dynamodbstreams_20120810_getsharditerator.CreateGetsharditeratorTool(cfg),
		tools_x_amz_target_dynamodbstreams_20120810_liststreams.CreateListstreamsTool(cfg),
		tools_x_amz_target_dynamodbstreams_20120810_describestream.CreateDescribestreamTool(cfg),
	}
}
