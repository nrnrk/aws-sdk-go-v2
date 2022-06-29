// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The custom terminology applied to the input text by Amazon Translate for the
// translated text response. This is optional in the response and will only be
// present if you specified terminology input in the request. Currently, only one
// terminology can be applied per TranslateText request.
type AppliedTerminology struct {

	// The name of the custom terminology applied to the input text by Amazon Translate
	// for the translated text response.
	Name *string

	// The specific terms of the custom terminology applied to the input text by Amazon
	// Translate for the translated text response. A maximum of 250 terms will be
	// returned, and the specific terms applied will be the first 250 terms in the
	// source text.
	Terms []Term

	noSmithyDocumentSerde
}

// The encryption key used to encrypt this object.
type EncryptionKey struct {

	// The Amazon Resource Name (ARN) of the encryption key being used to encrypt this
	// object.
	//
	// This member is required.
	Id *string

	// The type of encryption key used by Amazon Translate to encrypt this object.
	//
	// This member is required.
	Type EncryptionKeyType

	noSmithyDocumentSerde
}

// The input configuration properties for requesting a batch translation job.
type InputDataConfig struct {

	// Describes the format of the data that you submit to Amazon Translate as input.
	// You can specify one of the following multipurpose internet mail extension (MIME)
	// types:
	//
	// * text/html: The input data consists of one or more HTML files. Amazon
	// Translate translates only the text that resides in the html element in each
	// file.
	//
	// * text/plain: The input data consists of one or more unformatted text
	// files. Amazon Translate translates every character in this type of input.
	//
	// *
	// application/vnd.openxmlformats-officedocument.wordprocessingml.document: The
	// input data consists of one or more Word documents (.docx).
	//
	// *
	// application/vnd.openxmlformats-officedocument.presentationml.presentation: The
	// input data consists of one or more PowerPoint Presentation files (.pptx).
	//
	// *
	// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet: The input
	// data consists of one or more Excel Workbook files (.xlsx).
	//
	// *
	// application/x-xliff+xml: The input data consists of one or more XML Localization
	// Interchange File Format (XLIFF) files (.xlf). Amazon Translate supports only
	// XLIFF version 1.2.
	//
	// If you structure your input data as HTML, ensure that you
	// set this parameter to text/html. By doing so, you cut costs by limiting the
	// translation to the contents of the html element in each file. Otherwise, if you
	// set this parameter to text/plain, your costs will cover the translation of every
	// character.
	//
	// This member is required.
	ContentType *string

	// The URI of the AWS S3 folder that contains the input files. Amazon Translate
	// translates all the files in the folder. The folder must be in the same Region as
	// the API endpoint you are calling. The URI can also point to a single input
	// document, or it can provide the prefix for a collection of input documents. For
	// example. if you use the URI S3://bucketName/prefix and the prefix is a single
	// file, Amazon Translate uses that files as input. If more than one file begins
	// with the prefix, Amazon Translate uses all of them as input.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// The number of documents successfully and unsuccessfully processed during a
// translation job.
type JobDetails struct {

	// The number of documents that could not be processed during a translation job.
	DocumentsWithErrorsCount *int32

	// The number of documents used as input in a translation job.
	InputDocumentsCount *int32

	// The number of documents successfully processed during a translation job.
	TranslatedDocumentsCount *int32

	noSmithyDocumentSerde
}

// A supported language.
type Language struct {

	// Language code for the supported language.
	//
	// This member is required.
	LanguageCode *string

	// Language name of the supported language.
	//
	// This member is required.
	LanguageName *string

	noSmithyDocumentSerde
}

// The output configuration properties for a batch translation job.
type OutputDataConfig struct {

	// The URI of the S3 folder that contains a translation job's output file. The
	// folder must be in the same Region as the API endpoint that you are calling.
	//
	// This member is required.
	S3Uri *string

	// The encryption key used to encrypt this object.
	EncryptionKey *EncryptionKey

	noSmithyDocumentSerde
}

// Specifies the format and S3 location of the parallel data input file.
type ParallelDataConfig struct {

	// The format of the parallel data input file.
	//
	// This member is required.
	Format ParallelDataFormat

	// The URI of the Amazon S3 folder that contains the parallel data input file. The
	// folder must be in the same Region as the API endpoint you are calling.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// The location of the most recent parallel data input file that was successfully
// imported into Amazon Translate.
type ParallelDataDataLocation struct {

	// The Amazon S3 location of the parallel data input file. The location is returned
	// as a presigned URL to that has a 30-minute expiration. Amazon Translate doesn't
	// scan all input files for the risk of CSV injection attacks. CSV injection occurs
	// when a .csv or .tsv file is altered so that a record contains malicious code.
	// The record begins with a special character, such as =, +, -, or @. When the file
	// is opened in a spreadsheet program, the program might interpret the record as a
	// formula and run the code within it. Before you download an input file from
	// Amazon S3, ensure that you recognize the file and trust its creator.
	//
	// This member is required.
	Location *string

	// Describes the repository that contains the parallel data input file.
	//
	// This member is required.
	RepositoryType *string

	noSmithyDocumentSerde
}

// The properties of a parallel data resource.
type ParallelDataProperties struct {

	// The Amazon Resource Name (ARN) of the parallel data resource.
	Arn *string

	// The time at which the parallel data resource was created.
	CreatedAt *time.Time

	// The description assigned to the parallel data resource.
	Description *string

	// The encryption key used to encrypt this object.
	EncryptionKey *EncryptionKey

	// The number of records unsuccessfully imported from the parallel data input file.
	FailedRecordCount *int64

	// The number of UTF-8 characters that Amazon Translate imported from the parallel
	// data input file. This number includes only the characters in your translation
	// examples. It does not include characters that are used to format your file. For
	// example, if you provided a Translation Memory Exchange (.tmx) file, this number
	// does not include the tags.
	ImportedDataSize *int64

	// The number of records successfully imported from the parallel data input file.
	ImportedRecordCount *int64

	// The time at which the parallel data resource was last updated.
	LastUpdatedAt *time.Time

	// The time that the most recent update was attempted.
	LatestUpdateAttemptAt *time.Time

	// The status of the most recent update attempt for the parallel data resource.
	LatestUpdateAttemptStatus ParallelDataStatus

	// Additional information from Amazon Translate about the parallel data resource.
	Message *string

	// The custom name assigned to the parallel data resource.
	Name *string

	// Specifies the format and S3 location of the parallel data input file.
	ParallelDataConfig *ParallelDataConfig

	// The number of items in the input file that Amazon Translate skipped when you
	// created or updated the parallel data resource. For example, Amazon Translate
	// skips empty records, empty target texts, and empty lines.
	SkippedRecordCount *int64

	// The source language of the translations in the parallel data file.
	SourceLanguageCode *string

	// The status of the parallel data resource. When the parallel data is ready for
	// you to use, the status is ACTIVE.
	Status ParallelDataStatus

	// The language codes for the target languages available in the parallel data file.
	// All possible target languages are returned as an array.
	TargetLanguageCodes []string

	noSmithyDocumentSerde
}

// The term being translated by the custom terminology.
type Term struct {

	// The source text of the term being translated by the custom terminology.
	SourceText *string

	// The target text of the term being translated by the custom terminology.
	TargetText *string

	noSmithyDocumentSerde
}

// The data associated with the custom terminology. For information about the
// custom terminology file, see  Creating a Custom Terminology
// (https://docs.aws.amazon.com/translate/latest/dg/creating-custom-terminology.html).
type TerminologyData struct {

	// The file containing the custom terminology data. Your version of the AWS SDK
	// performs a Base64-encoding on this field before sending a request to the AWS
	// service. Users of the SDK should not perform Base64-encoding themselves.
	//
	// This member is required.
	File []byte

	// The data format of the custom terminology.
	//
	// This member is required.
	Format TerminologyDataFormat

	// The directionality of your terminology resource indicates whether it has one
	// source language (uni-directional) or multiple (multi-directional). UNI The
	// terminology resource has one source language (for example, the first column in a
	// CSV file), and all of its other languages are target languages. MULTI Any
	// language in the terminology resource can be the source language or a target
	// language. A single multi-directional terminology resource can be used for jobs
	// that translate different language pairs. For example, if the terminology
	// contains English and Spanish terms, it can be used for jobs that translate
	// English to Spanish and Spanish to English. When you create a custom terminology
	// resource without specifying the directionality, it behaves as uni-directional
	// terminology, although this parameter will have a null value.
	Directionality Directionality

	noSmithyDocumentSerde
}

// The location of the custom terminology data.
type TerminologyDataLocation struct {

	// The Amazon S3 location of the most recent custom terminology input file that was
	// successfully imported into Amazon Translate. The location is returned as a
	// presigned URL that has a 30-minute expiration . Amazon Translate doesn't scan
	// all input files for the risk of CSV injection attacks. CSV injection occurs when
	// a .csv or .tsv file is altered so that a record contains malicious code. The
	// record begins with a special character, such as =, +, -, or @. When the file is
	// opened in a spreadsheet program, the program might interpret the record as a
	// formula and run the code within it. Before you download an input file from
	// Amazon S3, ensure that you recognize the file and trust its creator.
	//
	// This member is required.
	Location *string

	// The repository type for the custom terminology data.
	//
	// This member is required.
	RepositoryType *string

	noSmithyDocumentSerde
}

// The properties of the custom terminology.
type TerminologyProperties struct {

	// The Amazon Resource Name (ARN) of the custom terminology.
	Arn *string

	// The time at which the custom terminology was created, based on the timestamp.
	CreatedAt *time.Time

	// The description of the custom terminology properties.
	Description *string

	// The directionality of your terminology resource indicates whether it has one
	// source language (uni-directional) or multiple (multi-directional). UNI The
	// terminology resource has one source language (the first column in a CSV file),
	// and all of its other languages are target languages. MULTI Any language in the
	// terminology resource can be the source language.
	Directionality Directionality

	// The encryption key for the custom terminology.
	EncryptionKey *EncryptionKey

	// The format of the custom terminology input file.
	Format TerminologyDataFormat

	// The time at which the custom terminology was last update, based on the
	// timestamp.
	LastUpdatedAt *time.Time

	// Additional information from Amazon Translate about the terminology resource.
	Message *string

	// The name of the custom terminology.
	Name *string

	// The size of the file used when importing a custom terminology.
	SizeBytes *int32

	// The number of terms in the input file that Amazon Translate skipped when you
	// created or updated the terminology resource.
	SkippedTermCount *int32

	// The language code for the source text of the translation request for which the
	// custom terminology is being used.
	SourceLanguageCode *string

	// The language codes for the target languages available with the custom
	// terminology resource. All possible target languages are returned in array.
	TargetLanguageCodes []string

	// The number of terms included in the custom terminology.
	TermCount *int32

	noSmithyDocumentSerde
}

// Provides information for filtering a list of translation jobs. For more
// information, see ListTextTranslationJobs.
type TextTranslationJobFilter struct {

	// Filters the list of jobs by name.
	JobName *string

	// Filters the list of jobs based by job status.
	JobStatus JobStatus

	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted after the specified time. Jobs
	// are returned in descending order, newest to oldest.
	SubmittedAfterTime *time.Time

	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted before the specified time. Jobs
	// are returned in ascending order, oldest to newest.
	SubmittedBeforeTime *time.Time

	noSmithyDocumentSerde
}

// Provides information about a translation job.
type TextTranslationJobProperties struct {

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that granted Amazon Translate read access to the job's input data.
	DataAccessRoleArn *string

	// The time at which the translation job ended.
	EndTime *time.Time

	// The input configuration properties that were specified when the job was
	// requested.
	InputDataConfig *InputDataConfig

	// The number of documents successfully and unsuccessfully processed during the
	// translation job.
	JobDetails *JobDetails

	// The ID of the translation job.
	JobId *string

	// The user-defined name of the translation job.
	JobName *string

	// The status of the translation job.
	JobStatus JobStatus

	// An explanation of any errors that may have occurred during the translation job.
	Message *string

	// The output configuration properties that were specified when the job was
	// requested.
	OutputDataConfig *OutputDataConfig

	// A list containing the names of the parallel data resources applied to the
	// translation job.
	ParallelDataNames []string

	// Settings that configure the translation output.
	Settings *TranslationSettings

	// The language code of the language of the source text. The language must be a
	// language supported by Amazon Translate.
	SourceLanguageCode *string

	// The time at which the translation job was submitted.
	SubmittedTime *time.Time

	// The language code of the language of the target text. The language must be a
	// language supported by Amazon Translate.
	TargetLanguageCodes []string

	// A list containing the names of the terminologies applied to a translation job.
	// Only one terminology can be applied per StartTextTranslationJob request at this
	// time.
	TerminologyNames []string

	noSmithyDocumentSerde
}

// Settings that configure the translation output.
type TranslationSettings struct {

	// You can optionally specify the desired level of formality for real-time
	// translations to supported target languages. The formality setting controls the
	// level of formal language usage (also known as register
	// (https://en.wikipedia.org/wiki/Register_(sociolinguistics))) in the translation
	// output. You can set the value to informal or formal. If you don't specify a
	// value for formality, or if the target language doesn't support formality, the
	// translation will ignore the formality setting. Note that asynchronous
	// translation jobs don't support formality. If you provide a value for formality,
	// the StartTextTranslationJob API throws an exception (InvalidRequestException).
	// For target languages that support formality, see Supported Languages and
	// Language Codes in the Amazon Translate Developer Guide
	// (https://docs.aws.amazon.com/translate/latest/dg/what-is.html).
	Formality Formality

	// Enable the profanity setting if you want Amazon Translate to mask profane words
	// and phrases in your translation output. To mask profane words and phrases,
	// Amazon Translate replaces them with the grawlix string “?$#@$“. This 5-character
	// sequence is used for each profane word or phrase, regardless of the length or
	// number of words. Amazon Translate doesn't detect profanity in all of its
	// supported languages. For languages that support profanity detection, see
	// Supported Languages and Language Codes in the Amazon Translate Developer Guide
	// (https://docs.aws.amazon.com/translate/latest/dg/what-is.html).
	Profanity Profanity

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
