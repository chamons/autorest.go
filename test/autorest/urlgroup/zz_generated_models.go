//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import "time"

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// PathItemsGetAllWithValuesOptions contains the optional parameters for the PathItems.GetAllWithValues method.
type PathItemsGetAllWithValuesOptions struct {
	// should contain value 'localStringQuery'
	LocalStringQuery *string
	// A string value 'pathItemStringQuery' that appears as a query parameter
	PathItemStringQuery *string
}

// PathItemsGetGlobalAndLocalQueryNullOptions contains the optional parameters for the PathItems.GetGlobalAndLocalQueryNull method.
type PathItemsGetGlobalAndLocalQueryNullOptions struct {
	// should contain null value
	LocalStringQuery *string
	// A string value 'pathItemStringQuery' that appears as a query parameter
	PathItemStringQuery *string
}

// PathItemsGetGlobalQueryNullOptions contains the optional parameters for the PathItems.GetGlobalQueryNull method.
type PathItemsGetGlobalQueryNullOptions struct {
	// should contain value 'localStringQuery'
	LocalStringQuery *string
	// A string value 'pathItemStringQuery' that appears as a query parameter
	PathItemStringQuery *string
}

// PathItemsGetLocalPathItemQueryNullOptions contains the optional parameters for the PathItems.GetLocalPathItemQueryNull method.
type PathItemsGetLocalPathItemQueryNullOptions struct {
	// should contain value null
	LocalStringQuery *string
	// should contain value null
	PathItemStringQuery *string
}

// PathsArrayCSVInPathOptions contains the optional parameters for the Paths.ArrayCSVInPath method.
type PathsArrayCSVInPathOptions struct {
	// placeholder for future optional parameters
}

// PathsBase64URLOptions contains the optional parameters for the Paths.Base64URL method.
type PathsBase64URLOptions struct {
	// placeholder for future optional parameters
}

// PathsByteEmptyOptions contains the optional parameters for the Paths.ByteEmpty method.
type PathsByteEmptyOptions struct {
	// placeholder for future optional parameters
}

// PathsByteMultiByteOptions contains the optional parameters for the Paths.ByteMultiByte method.
type PathsByteMultiByteOptions struct {
	// placeholder for future optional parameters
}

// PathsByteNullOptions contains the optional parameters for the Paths.ByteNull method.
type PathsByteNullOptions struct {
	// placeholder for future optional parameters
}

// PathsDateNullOptions contains the optional parameters for the Paths.DateNull method.
type PathsDateNullOptions struct {
	// placeholder for future optional parameters
}

// PathsDateTimeNullOptions contains the optional parameters for the Paths.DateTimeNull method.
type PathsDateTimeNullOptions struct {
	// placeholder for future optional parameters
}

// PathsDateTimeValidOptions contains the optional parameters for the Paths.DateTimeValid method.
type PathsDateTimeValidOptions struct {
	// placeholder for future optional parameters
}

// PathsDateValidOptions contains the optional parameters for the Paths.DateValid method.
type PathsDateValidOptions struct {
	// placeholder for future optional parameters
}

// PathsDoubleDecimalNegativeOptions contains the optional parameters for the Paths.DoubleDecimalNegative method.
type PathsDoubleDecimalNegativeOptions struct {
	// placeholder for future optional parameters
}

// PathsDoubleDecimalPositiveOptions contains the optional parameters for the Paths.DoubleDecimalPositive method.
type PathsDoubleDecimalPositiveOptions struct {
	// placeholder for future optional parameters
}

// PathsEnumNullOptions contains the optional parameters for the Paths.EnumNull method.
type PathsEnumNullOptions struct {
	// placeholder for future optional parameters
}

// PathsEnumValidOptions contains the optional parameters for the Paths.EnumValid method.
type PathsEnumValidOptions struct {
	// placeholder for future optional parameters
}

// PathsFloatScientificNegativeOptions contains the optional parameters for the Paths.FloatScientificNegative method.
type PathsFloatScientificNegativeOptions struct {
	// placeholder for future optional parameters
}

// PathsFloatScientificPositiveOptions contains the optional parameters for the Paths.FloatScientificPositive method.
type PathsFloatScientificPositiveOptions struct {
	// placeholder for future optional parameters
}

// PathsGetBooleanFalseOptions contains the optional parameters for the Paths.GetBooleanFalse method.
type PathsGetBooleanFalseOptions struct {
	// placeholder for future optional parameters
}

// PathsGetBooleanTrueOptions contains the optional parameters for the Paths.GetBooleanTrue method.
type PathsGetBooleanTrueOptions struct {
	// placeholder for future optional parameters
}

// PathsGetIntNegativeOneMillionOptions contains the optional parameters for the Paths.GetIntNegativeOneMillion method.
type PathsGetIntNegativeOneMillionOptions struct {
	// placeholder for future optional parameters
}

// PathsGetIntOneMillionOptions contains the optional parameters for the Paths.GetIntOneMillion method.
type PathsGetIntOneMillionOptions struct {
	// placeholder for future optional parameters
}

// PathsGetNegativeTenBillionOptions contains the optional parameters for the Paths.GetNegativeTenBillion method.
type PathsGetNegativeTenBillionOptions struct {
	// placeholder for future optional parameters
}

// PathsGetTenBillionOptions contains the optional parameters for the Paths.GetTenBillion method.
type PathsGetTenBillionOptions struct {
	// placeholder for future optional parameters
}

// PathsStringEmptyOptions contains the optional parameters for the Paths.StringEmpty method.
type PathsStringEmptyOptions struct {
	// placeholder for future optional parameters
}

// PathsStringNullOptions contains the optional parameters for the Paths.StringNull method.
type PathsStringNullOptions struct {
	// placeholder for future optional parameters
}

// PathsStringURLEncodedOptions contains the optional parameters for the Paths.StringURLEncoded method.
type PathsStringURLEncodedOptions struct {
	// placeholder for future optional parameters
}

// PathsStringURLNonEncodedOptions contains the optional parameters for the Paths.StringURLNonEncoded method.
type PathsStringURLNonEncodedOptions struct {
	// placeholder for future optional parameters
}

// PathsStringUnicodeOptions contains the optional parameters for the Paths.StringUnicode method.
type PathsStringUnicodeOptions struct {
	// placeholder for future optional parameters
}

// PathsUnixTimeURLOptions contains the optional parameters for the Paths.UnixTimeURL method.
type PathsUnixTimeURLOptions struct {
	// placeholder for future optional parameters
}

// QueriesArrayStringCSVEmptyOptions contains the optional parameters for the Queries.ArrayStringCSVEmpty method.
type QueriesArrayStringCSVEmptyOptions struct {
	// an empty array [] of string using the csv-array format
	ArrayQuery []string
}

// QueriesArrayStringCSVNullOptions contains the optional parameters for the Queries.ArrayStringCSVNull method.
type QueriesArrayStringCSVNullOptions struct {
	// a null array of string using the csv-array format
	ArrayQuery []string
}

// QueriesArrayStringCSVValidOptions contains the optional parameters for the Queries.ArrayStringCSVValid method.
type QueriesArrayStringCSVValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format
	ArrayQuery []string
}

// QueriesArrayStringNoCollectionFormatEmptyOptions contains the optional parameters for the Queries.ArrayStringNoCollectionFormatEmpty method.
type QueriesArrayStringNoCollectionFormatEmptyOptions struct {
	// Array-typed query parameter. Pass in ['hello', 'nihao', 'bonjour'].
	ArrayQuery []string
}

// QueriesArrayStringPipesValidOptions contains the optional parameters for the Queries.ArrayStringPipesValid method.
type QueriesArrayStringPipesValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the pipes-array format
	ArrayQuery []string
}

// QueriesArrayStringSsvValidOptions contains the optional parameters for the Queries.ArrayStringSsvValid method.
type QueriesArrayStringSsvValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the ssv-array format
	ArrayQuery []string
}

// QueriesArrayStringTsvValidOptions contains the optional parameters for the Queries.ArrayStringTsvValid method.
type QueriesArrayStringTsvValidOptions struct {
	// an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the tsv-array format
	ArrayQuery []string
}

// QueriesByteEmptyOptions contains the optional parameters for the Queries.ByteEmpty method.
type QueriesByteEmptyOptions struct {
	// placeholder for future optional parameters
}

// QueriesByteMultiByteOptions contains the optional parameters for the Queries.ByteMultiByte method.
type QueriesByteMultiByteOptions struct {
	// '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array
	ByteQuery []byte
}

// QueriesByteNullOptions contains the optional parameters for the Queries.ByteNull method.
type QueriesByteNullOptions struct {
	// null as byte array (no query parameters in uri)
	ByteQuery []byte
}

// QueriesDateNullOptions contains the optional parameters for the Queries.DateNull method.
type QueriesDateNullOptions struct {
	// null as date (no query parameters in uri)
	DateQuery *time.Time
}

// QueriesDateTimeNullOptions contains the optional parameters for the Queries.DateTimeNull method.
type QueriesDateTimeNullOptions struct {
	// null as date-time (no query parameters)
	DateTimeQuery *time.Time
}

// QueriesDateTimeValidOptions contains the optional parameters for the Queries.DateTimeValid method.
type QueriesDateTimeValidOptions struct {
	// placeholder for future optional parameters
}

// QueriesDateValidOptions contains the optional parameters for the Queries.DateValid method.
type QueriesDateValidOptions struct {
	// placeholder for future optional parameters
}

// QueriesDoubleDecimalNegativeOptions contains the optional parameters for the Queries.DoubleDecimalNegative method.
type QueriesDoubleDecimalNegativeOptions struct {
	// placeholder for future optional parameters
}

// QueriesDoubleDecimalPositiveOptions contains the optional parameters for the Queries.DoubleDecimalPositive method.
type QueriesDoubleDecimalPositiveOptions struct {
	// placeholder for future optional parameters
}

// QueriesDoubleNullOptions contains the optional parameters for the Queries.DoubleNull method.
type QueriesDoubleNullOptions struct {
	// null numeric value
	DoubleQuery *float64
}

// QueriesEnumNullOptions contains the optional parameters for the Queries.EnumNull method.
type QueriesEnumNullOptions struct {
	// null string value
	EnumQuery *URIColor
}

// QueriesEnumValidOptions contains the optional parameters for the Queries.EnumValid method.
type QueriesEnumValidOptions struct {
	// 'green color' enum value
	EnumQuery *URIColor
}

// QueriesFloatNullOptions contains the optional parameters for the Queries.FloatNull method.
type QueriesFloatNullOptions struct {
	// null numeric value
	FloatQuery *float32
}

// QueriesFloatScientificNegativeOptions contains the optional parameters for the Queries.FloatScientificNegative method.
type QueriesFloatScientificNegativeOptions struct {
	// placeholder for future optional parameters
}

// QueriesFloatScientificPositiveOptions contains the optional parameters for the Queries.FloatScientificPositive method.
type QueriesFloatScientificPositiveOptions struct {
	// placeholder for future optional parameters
}

// QueriesGetBooleanFalseOptions contains the optional parameters for the Queries.GetBooleanFalse method.
type QueriesGetBooleanFalseOptions struct {
	// placeholder for future optional parameters
}

// QueriesGetBooleanNullOptions contains the optional parameters for the Queries.GetBooleanNull method.
type QueriesGetBooleanNullOptions struct {
	// null boolean value
	BoolQuery *bool
}

// QueriesGetBooleanTrueOptions contains the optional parameters for the Queries.GetBooleanTrue method.
type QueriesGetBooleanTrueOptions struct {
	// placeholder for future optional parameters
}

// QueriesGetIntNegativeOneMillionOptions contains the optional parameters for the Queries.GetIntNegativeOneMillion method.
type QueriesGetIntNegativeOneMillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesGetIntNullOptions contains the optional parameters for the Queries.GetIntNull method.
type QueriesGetIntNullOptions struct {
	// null integer value
	IntQuery *int32
}

// QueriesGetIntOneMillionOptions contains the optional parameters for the Queries.GetIntOneMillion method.
type QueriesGetIntOneMillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesGetLongNullOptions contains the optional parameters for the Queries.GetLongNull method.
type QueriesGetLongNullOptions struct {
	// null 64 bit integer value
	LongQuery *int64
}

// QueriesGetNegativeTenBillionOptions contains the optional parameters for the Queries.GetNegativeTenBillion method.
type QueriesGetNegativeTenBillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesGetTenBillionOptions contains the optional parameters for the Queries.GetTenBillion method.
type QueriesGetTenBillionOptions struct {
	// placeholder for future optional parameters
}

// QueriesStringEmptyOptions contains the optional parameters for the Queries.StringEmpty method.
type QueriesStringEmptyOptions struct {
	// placeholder for future optional parameters
}

// QueriesStringNullOptions contains the optional parameters for the Queries.StringNull method.
type QueriesStringNullOptions struct {
	// null string value
	StringQuery *string
}

// QueriesStringURLEncodedOptions contains the optional parameters for the Queries.StringURLEncoded method.
type QueriesStringURLEncodedOptions struct {
	// placeholder for future optional parameters
}

// QueriesStringUnicodeOptions contains the optional parameters for the Queries.StringUnicode method.
type QueriesStringUnicodeOptions struct {
	// placeholder for future optional parameters
}
