package ParquetType

import (
	"fmt"
	"github.com/xitongsys/parquet-go/parquet"
	"log"
	"reflect"
)

//base type
type BOOLEAN bool
type INT32 int32
type INT64 int64
type INT96 string // length=96
type FLOAT float32
type DOUBLE float64
type BYTE_ARRAY string
type FIXED_LEN_BYTE_ARRAY string

//logical type
type UTF8 string
type INT_8 int32
type INT_16 int32
type INT_32 int32
type INT_64 int64
type UINT_8 uint32
type UINT_16 uint32
type UINT_32 uint32
type UINT_64 uint64
type DATE int32
type TIME_MILLIS int32
type TIME_MICROS int64
type TIMESTAMP_MILLIS int64
type TIMESTAMP_MICROS int64
type INTERVAL string // length=12
type DECIMAL string

func ParquetTypeToGoType(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	typeName := reflect.TypeOf(value).Name()
	switch typeName {
	case "BOOLEAN":
		return bool(value.(BOOLEAN))
	case "INT32":
		return int32(value.(INT32))
	case "INT64":
		return int64(value.(INT64))
	case "INT96":
		return string(value.(INT96))
	case "FLOAT":
		return float32(value.(FLOAT))
	case "DOUBLE":
		return float64(value.(DOUBLE))
	case "BYTE_ARRAY":
		return string(value.(BYTE_ARRAY))
	case "FIXED_LEN_BYTE_ARRAY":
		return string(value.(FIXED_LEN_BYTE_ARRAY))
	case "UTF8":
		return string(value.(UTF8))
	case "INT_8":
		return int32(value.(INT_8))
	case "INT_16":
		return int32(value.(INT_16))
	case "INT_32":
		return int32(value.(INT_32))
	case "INT_64":
		return int64(value.(INT_64))
	case "UINT_8":
		return uint32(value.(UINT_8))
	case "UINT_16":
		return uint32(value.(UINT_16))
	case "UINT_32":
		return uint32(value.(UINT_32))
	case "UINT_64":
		return uint64(value.(UINT_64))
	case "DATE":
		return int32(value.(DATE))
	case "TIME_MILLIS":
		return int32(value.(TIME_MILLIS))
	case "TIME_MICROS":
		return int64(value.(TIME_MICROS))
	case "TIMESTAMP_MILLIS":
		return int64(value.(TIMESTAMP_MILLIS))
	case "TIMESTAMP_MICROS":
		return int64(value.(TIMESTAMP_MICROS))
	case "INTERVAL":
		return string(value.(INTERVAL))
	case "DECIMAL":
		return string(value.(DECIMAL))
	}
	return nil

}

//Scan a string to parquet value
func StrToParquetType(s string, typeName string) interface{} {
	if typeName == "BOOLEAN" {
		var v BOOLEAN
		fmt.Sscanf(s, "%t", &v)
		return v
	} else if typeName == "INT32" {
		var v INT32
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "INT64" {
		var v INT64
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "INT96" {
		var v INT96
		v = INT96(s)
		return v
	} else if typeName == "FLOAT" {
		var v FLOAT
		fmt.Sscanf(s, "%f", &v)
		return v
	} else if typeName == "DOUBLE" {
		var v DOUBLE
		fmt.Sscanf(s, "%f", &v)
		return v
	} else if typeName == "BYTE_ARRAY" {
		var v BYTE_ARRAY
		v = BYTE_ARRAY(s)
		return v
	} else if typeName == "FIXED_LEN_BYTE_ARRAY" {
		var v FIXED_LEN_BYTE_ARRAY
		v = FIXED_LEN_BYTE_ARRAY(s)
		return v
	} else if typeName == "UTF8" {
		var v UTF8
		v = UTF8(s)
		return v
	} else if typeName == "INT_8" {
		var v INT_8
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "INT_16" {
		var v INT_16
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "INT_32" {
		var v INT_32
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "INT_64" {
		var v INT_64
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "UINT_8" {
		var v UINT_8
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "UINT_16" {
		var v UINT_16
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "UINT_32" {
		var v UINT_32
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "UINT_64" {
		var v UINT_64
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "DATE" {
		var v DATE
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "TIME_MILLIS" {
		var v TIME_MILLIS
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "TIME_MICROS" {
		var v TIME_MICROS
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "TIMESTAMP_MILLIS" {
		var v TIMESTAMP_MILLIS
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "TIMESTAMP_MICROS" {
		var v TIMESTAMP_MICROS
		fmt.Sscanf(s, "%d", &v)
		return v
	} else if typeName == "INTERVAL" {
		var v INTERVAL
		v = INTERVAL(s)
		return v
	} else if typeName == "DECIMAL" {
		var v DECIMAL
		v = DECIMAL(s)
		return v
	} else {
		log.Printf("Type Error: %v ", typeName)
		return nil
	}
}

func GoTypeToParquetType(src interface{}, pT *parquet.Type, cT *parquet.ConvertedType) interface{} {
	if cT == nil {
		if *pT == parquet.Type_BOOLEAN {
			return BOOLEAN(src.(bool))
		} else if *pT == parquet.Type_INT32 {
			return INT32(src.(int32))
		} else if *pT == parquet.Type_INT64 {
			return INT64(src.(int64))
		} else if *pT == parquet.Type_INT96 {
			return INT96(src.(string))
		} else if *pT == parquet.Type_FLOAT {
			return FLOAT(src.(float32))
		} else if *pT == parquet.Type_DOUBLE {
			return DOUBLE(src.(float64))
		} else if *pT == parquet.Type_BYTE_ARRAY {
			return BYTE_ARRAY(src.(string))
		} else if *pT == parquet.Type_FIXED_LEN_BYTE_ARRAY {
			return FIXED_LEN_BYTE_ARRAY(src.(string))
		}
		return nil
	}

	if *cT == parquet.ConvertedType_UTF8 {
		return UTF8(src.(string))
	} else if *cT == parquet.ConvertedType_INT_8 {
		return INT_8(src.(int32))
	} else if *cT == parquet.ConvertedType_INT_16 {
		return INT_16(src.(int32))
	} else if *cT == parquet.ConvertedType_INT_32 {
		return INT_32(src.(int32))
	} else if *cT == parquet.ConvertedType_INT_64 {
		return INT_64(src.(int64))
	} else if *cT == parquet.ConvertedType_UINT_8 {
		return UINT_8(src.(uint32))
	} else if *cT == parquet.ConvertedType_UINT_16 {
		return UINT_16(src.(uint32))
	} else if *cT == parquet.ConvertedType_UINT_32 {
		return UINT_32(src.(uint32))
	} else if *cT == parquet.ConvertedType_UINT_64 {
		return UINT_64(src.(uint64))
	} else if *cT == parquet.ConvertedType_DATE {
		return DATE(src.(int32))
	} else if *cT == parquet.ConvertedType_TIME_MILLIS {
		return TIME_MILLIS(src.(int32))
	} else if *cT == parquet.ConvertedType_TIME_MICROS {
		return TIME_MICROS(src.(int64))
	} else if *cT == parquet.ConvertedType_TIMESTAMP_MILLIS {
		return TIMESTAMP_MILLIS(src.(int64))
	} else if *cT == parquet.ConvertedType_TIMESTAMP_MICROS {
		return TIMESTAMP_MICROS(src.(int64))
	} else if *cT == parquet.ConvertedType_INTERVAL {
		return INTERVAL(src.(string))
	} else if *cT == parquet.ConvertedType_DECIMAL {
		return DECIMAL(src.(string))
	} else {
		return nil
	}
}
