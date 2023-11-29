// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Defines the Amazon S3 bucket where the configured audience is stored.
type AudienceDestination struct {

	// The Amazon S3 bucket and path for the configured audience.
	//
	// This member is required.
	S3Destination *S3ConfigMap

	noSmithyDocumentSerde
}

// Provides information about the audience export job.
type AudienceExportJobSummary struct {

	// The Amazon Resource Name (ARN) of the audience generation job that was exported.
	//
	// This member is required.
	AudienceGenerationJobArn *string

	// The size of the generated audience. Must match one of the sizes in the
	// configured audience model.
	//
	// This member is required.
	AudienceSize *AudienceSize

	// The time at which the audience export job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience export job.
	//
	// This member is required.
	Name *string

	// The status of the audience export job.
	//
	// This member is required.
	Status AudienceExportJobStatus

	// The most recent time at which the audience export job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the audience export job.
	Description *string

	// The Amazon S3 bucket where the audience export is stored.
	OutputLocation *string

	// Details about the status of a resource.
	StatusDetails *StatusDetails

	noSmithyDocumentSerde
}

// Defines the Amazon S3 bucket where the training data for the configured
// audience is stored.
type AudienceGenerationJobDataSource struct {

	// The Amazon S3 bucket where the training data for the configured audience is
	// stored.
	//
	// This member is required.
	DataSource *S3ConfigMap

	// The ARN of the IAM role that can read the Amazon S3 bucket where the training
	// data is stored.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

// Provides information about the configured audience generation job.
type AudienceGenerationJobSummary struct {

	// The Amazon Resource Name (ARN) of the audience generation job.
	//
	// This member is required.
	AudienceGenerationJobArn *string

	// The Amazon Resource Name (ARN) of the configured audience model that was used
	// for this audience generation job.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The time at which the audience generation job was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience generation job.
	//
	// This member is required.
	Name *string

	// The status of the audience generation job.
	//
	// This member is required.
	Status AudienceGenerationJobStatus

	// The most recent time at which the audience generation job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The identifier of the collaboration that contains this audience generation job.
	CollaborationId *string

	// The description of the audience generation job.
	Description *string

	// The AWS Account that submitted the job.
	StartedBy *string

	noSmithyDocumentSerde
}

// The audience model metrics.
type AudienceModelMetric struct {

	// The number of users that were used to generate these model metrics.
	//
	// This member is required.
	ForTopKItemPredictions *int32

	// The audience model metric.
	//
	// This member is required.
	Type AudienceModelMetricType

	// The value of the audience model metric
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// Information about the audience model.
type AudienceModelSummary struct {

	// The Amazon Resource Name (ARN) of the audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The time at which the audience model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience model.
	//
	// This member is required.
	Name *string

	// The status of the audience model.
	//
	// This member is required.
	Status AudienceModelStatus

	// The Amazon Resource Name (ARN) of the training dataset that was used for the
	// audience model.
	//
	// This member is required.
	TrainingDatasetArn *string

	// The most recent time at which the audience model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the audience model.
	Description *string

	noSmithyDocumentSerde
}

// Metrics that describe the quality of the generated audience.
type AudienceQualityMetrics struct {

	// The relevance scores of the generated audience.
	//
	// This member is required.
	RelevanceMetrics []RelevanceMetric

	noSmithyDocumentSerde
}

// The size of the generated audience. Must match one of the sizes in the
// configured audience model.
type AudienceSize struct {

	// Whether the audience size is defined in absolute terms or as a percentage. You
	// can use the ABSOLUTE AudienceSize to configure out audience sizes using the
	// count of identifiers in the output. You can use the Percentage AudienceSize to
	// configure sizes in the range 1-100 percent.
	//
	// This member is required.
	Type AudienceSizeType

	// Specify an audience size value.
	//
	// This member is required.
	Value *int32

	noSmithyDocumentSerde
}

// Configure the list of audience output sizes that can be created. A request to
// StartAudienceGenerationJob that uses this configured audience model must have an
// audienceSize selected from this list. You can use the ABSOLUTE AudienceSize to
// configure out audience sizes using the count of identifiers in the output. You
// can use the Percentage AudienceSize to configure sizes in the range 1-100
// percent.
type AudienceSizeConfig struct {

	// An array of the different audience output sizes.
	//
	// This member is required.
	AudienceSizeBins []int32

	// Whether the audience output sizes are defined as an absolute number or a
	// percentage.
	//
	// This member is required.
	AudienceSizeType AudienceSizeType

	noSmithyDocumentSerde
}

// Metadata for a column.
type ColumnSchema struct {

	// The name of a column.
	//
	// This member is required.
	ColumnName *string

	// The data type of column.
	//
	// This member is required.
	ColumnTypes []ColumnType

	noSmithyDocumentSerde
}

// Configuration information necessary for the configure audience model output.
type ConfiguredAudienceModelOutputConfig struct {

	// Defines the Amazon S3 bucket where the configured audience is stored.
	//
	// This member is required.
	Destination *AudienceDestination

	// The ARN of the IAM role that can write the Amazon S3 bucket.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

// Information about the configured audience model.
type ConfiguredAudienceModelSummary struct {

	// The Amazon Resource Name (ARN) of the audience model that was used to create
	// the configured audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The Amazon Resource Name (ARN) of the configured audience model that you are
	// interested in.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The time at which the configured audience model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the configured audience model.
	//
	// This member is required.
	Name *string

	// The output configuration of the configured audience model.
	//
	// This member is required.
	OutputConfig *ConfiguredAudienceModelOutputConfig

	// The status of the configured audience model.
	//
	// This member is required.
	Status ConfiguredAudienceModelStatus

	// The most recent time at which the configured audience model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the configured audience model.
	Description *string

	noSmithyDocumentSerde
}

// Defines where the training dataset is located, what type of data it contains,
// and how to access the data.
type Dataset struct {

	// A DatasetInputConfig object that defines the data source and schema mapping.
	//
	// This member is required.
	InputConfig *DatasetInputConfig

	// What type of information is found in the dataset.
	//
	// This member is required.
	Type DatasetType

	noSmithyDocumentSerde
}

// Defines the Glue data source and schema mapping information.
type DatasetInputConfig struct {

	// A DataSource object that specifies the Glue data source for the training data.
	//
	// This member is required.
	DataSource *DataSource

	// The schema information for the training data.
	//
	// This member is required.
	Schema []ColumnSchema

	noSmithyDocumentSerde
}

// Defines information about the Glue data source that contains the training data.
type DataSource struct {

	// A GlueDataSource object that defines the catalog ID, database name, and table
	// name for the training data.
	//
	// This member is required.
	GlueDataSource *GlueDataSource

	noSmithyDocumentSerde
}

// Defines the Glue data source that contains the training data.
type GlueDataSource struct {

	// The Glue database that contains the training data.
	//
	// This member is required.
	DatabaseName *string

	// The Glue table that contains the training data.
	//
	// This member is required.
	TableName *string

	// The Glue catalog that contains the training data.
	CatalogId *string

	noSmithyDocumentSerde
}

// The relevance score of a generated audience.
type RelevanceMetric struct {

	// The size of the generated audience. Must match one of the sizes in the
	// configured audience model.
	//
	// This member is required.
	AudienceSize *AudienceSize

	// The relevance score of the generated audience.
	Score *float64

	noSmithyDocumentSerde
}

// Provides information about an Amazon S3 bucket and path.
type S3ConfigMap struct {

	// The Amazon S3 location URI.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// Details about the status of a resource.
type StatusDetails struct {

	// The error message that was returned. The message is intended for human
	// consumption and can change at any time. Use the statusCode for programmatic
	// error handling.
	Message *string

	// The status code that was returned. The status code is intended for programmatic
	// error handling. Clean Rooms ML will not change the status code for existing
	// error conditions.
	StatusCode *string

	noSmithyDocumentSerde
}

// Provides information about the training dataset.
type TrainingDatasetSummary struct {

	// The time at which the training dataset was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the training dataset.
	//
	// This member is required.
	Name *string

	// The status of the training dataset.
	//
	// This member is required.
	Status TrainingDatasetStatus

	// The Amazon Resource Name (ARN) of the training dataset.
	//
	// This member is required.
	TrainingDatasetArn *string

	// The most recent time at which the training dataset was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the training dataset.
	Description *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
