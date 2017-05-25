package pgerror

import (
	"github.com/lib/pq"
)

// Class 00 - Successful Completion

// SuccessfulCompletion checks if the error is of code 00000
func SuccessfulCompletion(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "00000"
	}
	return false
}

// Class 01 - Warning

// Warning checks if the error is of code 01000
func Warning(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01000"
	}
	return false
}

// DynamicResultSetsReturned checks if the error is of code 0100C
func DynamicResultSetsReturned(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0100C"
	}
	return false
}

// ImplicitZeroBitPadding checks if the error is of code 01008
func ImplicitZeroBitPadding(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01008"
	}
	return false
}

// NullValueEliminatedInSetFunction checks if the error is of code 01003
func NullValueEliminatedInSetFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01003"
	}
	return false
}

// PrivilegeNotGranted checks if the error is of code 01007
func PrivilegeNotGranted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01007"
	}
	return false
}

// PrivilegeNotRevoked checks if the error is of code 01006
func PrivilegeNotRevoked(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01006"
	}
	return false
}

// StringDataRightTruncation checks if the error is of code 01004
func StringDataRightTruncation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01004" || pgerr.Code == "22001"
	}
	return false
}

// DeprecatedFeature checks if the error is of code 01P01
func DeprecatedFeature(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "01P01"
	}
	return false
}

// Class 02 - No Data (this is also a warning class per the SQL standard)

// NoData checks if the error is of code 02000
func NoData(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "02000"
	}
	return false
}

// NoAdditionalDynamicResultSetsReturned checks if the error is of code 02001
func NoAdditionalDynamicResultSetsReturned(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "02001"
	}
	return false
}

// Class 03 - SQL Statement Not Yet Complete

// SQLStatementNotYetComplete checks if the error is of code 03000
func SQLStatementNotYetComplete(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "03000"
	}
	return false
}

// Class 08 - Connection Exception

// ConnectionException checks if the error is of code 08000
func ConnectionException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08000"
	}
	return false
}

// ConnectionDoesNotExist checks if the error is of code 08003
func ConnectionDoesNotExist(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08003"
	}
	return false
}

// ConnectionFailure checks if the error is of code 08006
func ConnectionFailure(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08006"
	}
	return false
}

// SQLclientUnableToEstablishSQLconnection checks if the error is of code 08001
func SQLclientUnableToEstablishSQLconnection(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08001"
	}
	return false
}

// SQLserverRejectedEstablishmentOfSQLconnection checks if the error is of code 08004
func SQLserverRejectedEstablishmentOfSQLconnection(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08004"
	}
	return false
}

// TransactionResolutionUnknown checks if the error is of code 08007
func TransactionResolutionUnknown(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08007"
	}
	return false
}

// ProtocolViolation checks if the error is of code 08P01
func ProtocolViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "08P01"
	}
	return false
}

// Class 09 - Triggered Action Exception

// TriggeredActionException checks if the error is of code 09000
func TriggeredActionException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "09000"
	}
	return false
}

// Class 0A - Feature Not Supported

// FeatureNotSupported checks if the error is of code 0A000
func FeatureNotSupported(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0A000"
	}
	return false
}

// Class 0B - Invalid Transaction Initiation

// InvalidTransactionInitiation checks if the error is of code 0B000
func InvalidTransactionInitiation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0B000"
	}
	return false
}

// Class 0F - Locator Exception

// LocatorException checks if the error is of code 0F000
func LocatorException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0F000"
	}
	return false
}

// InvalidLocatorSpecification checks if the error is of code 0F001
func InvalidLocatorSpecification(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0F001"
	}
	return false
}

// Class 0L - Invalid Grantor

// InvalidGrantor checks if the error is of code 0L000
func InvalidGrantor(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0L000"
	}
	return false
}

// InvalidGrantOperation checks if the error is of code 0LP01
func InvalidGrantOperation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0LP01"
	}
	return false
}

// Class 0P - Invalid Role Specification

// InvalidRoleSpecification checks if the error is of code 0P000
func InvalidRoleSpecification(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0P000"
	}
	return false
}

// Class 0Z - Diagnostics Exception

// DiagnosticsException checks if the error is of code 0Z000
func DiagnosticsException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0Z000"
	}
	return false
}

// StackedDiagnosticsAccessedWithoutActiveHandler checks if the error is of code 0Z002
func StackedDiagnosticsAccessedWithoutActiveHandler(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "0Z002"
	}
	return false
}

// Class 20 - Case Not Found

// CaseNotFound checks if the error is of code 20000
func CaseNotFound(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "20000"
	}
	return false
}

// Class 21 - Cardinality Violation

// CardinalityViolation checks if the error is of code 21000
func CardinalityViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "21000"
	}
	return false
}

// Class 22 - Data Exception

// DataException checks if the error is of code 22000
func DataException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22000"
	}
	return false
}

// ArraySubscriptError checks if the error is of code 2202E
func ArraySubscriptError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2202E"
	}
	return false
}

// CharacterNotInRepertoire checks if the error is of code 22021
func CharacterNotInRepertoire(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22021"
	}
	return false
}

// DatetimeFieldOverflow checks if the error is of code 22008
func DatetimeFieldOverflow(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22008"
	}
	return false
}

// DivisionByZero checks if the error is of code 22012
func DivisionByZero(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22012"
	}
	return false
}

// ErrorInAssignment checks if the error is of code 22005
func ErrorInAssignment(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22005"
	}
	return false
}

// EscapeCharacterConflict checks if the error is of code 2200B
func EscapeCharacterConflict(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200B"
	}
	return false
}

// IndicatorOverflow checks if the error is of code 22022
func IndicatorOverflow(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22022"
	}
	return false
}

// IntervalFieldOverflow checks if the error is of code 22015
func IntervalFieldOverflow(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22015"
	}
	return false
}

// InvalidArgumentForLogarithm checks if the error is of code 2201E
func InvalidArgumentForLogarithm(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2201E"
	}
	return false
}

// InvalidArgumentForNtileFunction checks if the error is of code 22014
func InvalidArgumentForNtileFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22014"
	}
	return false
}

// InvalidArgumentForNthValueFunction checks if the error is of code 22016
func InvalidArgumentForNthValueFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22016"
	}
	return false
}

// InvalidArgumentForPowerFunction checks if the error is of code 2201F
func InvalidArgumentForPowerFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2201F"
	}
	return false
}

// InvalidArgumentForWidthBucketFunction checks if the error is of code 2201G
func InvalidArgumentForWidthBucketFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2201G"
	}
	return false
}

// InvalidCharacterValueForCast checks if the error is of code 22018
func InvalidCharacterValueForCast(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22018"
	}
	return false
}

// InvalidDatetimeFormat checks if the error is of code 22007
func InvalidDatetimeFormat(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22007"
	}
	return false
}

// InvalidEscapeCharacter checks if the error is of code 22019
func InvalidEscapeCharacter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22019"
	}
	return false
}

// InvalidEscapeOctet checks if the error is of code 2200D
func InvalidEscapeOctet(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200D"
	}
	return false
}

// InvalidEscapeSequence checks if the error is of code 22025
func InvalidEscapeSequence(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22025"
	}
	return false
}

// NonstandardUseOfEscapeCharacter checks if the error is of code 22P06
func NonstandardUseOfEscapeCharacter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22P06"
	}
	return false
}

// InvalidIndicatorParameterValue checks if the error is of code 22010
func InvalidIndicatorParameterValue(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22010"
	}
	return false
}

// InvalidParameterValue checks if the error is of code 22023
func InvalidParameterValue(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22023"
	}
	return false
}

// InvalidRegularExpression checks if the error is of code 2201B
func InvalidRegularExpression(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2201B"
	}
	return false
}

// InvalidRowCountInLimitClause checks if the error is of code 2201W
func InvalidRowCountInLimitClause(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2201W"
	}
	return false
}

// InvalidRowCountInResultOffsetClause checks if the error is of code 2201X
func InvalidRowCountInResultOffsetClause(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2201X"
	}
	return false
}

// InvalidTimeZoneDisplacementValue checks if the error is of code 22009
func InvalidTimeZoneDisplacementValue(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22009"
	}
	return false
}

// InvalidUseOfEscapeCharacter checks if the error is of code 2200C
func InvalidUseOfEscapeCharacter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200C"
	}
	return false
}

// MostSpecificTypeMismatch checks if the error is of code 2200G
func MostSpecificTypeMismatch(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200G"
	}
	return false
}

// NullValueNotAllowed checks if the error is of code 22004
func NullValueNotAllowed(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22004" || pgerr.Code == "39004"
	}
	return false
}

// NullValueNoIndicatorParameter checks if the error is of code 22002
func NullValueNoIndicatorParameter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22002"
	}
	return false
}

// NumericValueOutOfRange checks if the error is of code 22003
func NumericValueOutOfRange(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22003"
	}
	return false
}

// StringDataLengthMismatch checks if the error is of code 22026
func StringDataLengthMismatch(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22026"
	}
	return false
}

// SubstringError checks if the error is of code 22011
func SubstringError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22011"
	}
	return false
}

// TrimError checks if the error is of code 22027
func TrimError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22027"
	}
	return false
}

// UnterminatedCString checks if the error is of code 22024
func UnterminatedCString(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22024"
	}
	return false
}

// ZeroLengthCharacterString checks if the error is of code 2200F
func ZeroLengthCharacterString(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200F"
	}
	return false
}

// FloatingPointException checks if the error is of code 22P01
func FloatingPointException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22P01"
	}
	return false
}

// InvalidTextRepresentation checks if the error is of code 22P02
func InvalidTextRepresentation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22P02"
	}
	return false
}

// InvalidBinaryRepresentation checks if the error is of code 22P03
func InvalidBinaryRepresentation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22P03"
	}
	return false
}

// BadCopyFileFormat checks if the error is of code 22P04
func BadCopyFileFormat(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22P04"
	}
	return false
}

// UntranslatableCharacter checks if the error is of code 22P05
func UntranslatableCharacter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "22P05"
	}
	return false
}

// NotAnXMLDocument checks if the error is of code 2200L
func NotAnXMLDocument(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200L"
	}
	return false
}

// InvalidXMLDocument checks if the error is of code 2200M
func InvalidXMLDocument(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200M"
	}
	return false
}

// InvalidXMLContent checks if the error is of code 2200N
func InvalidXMLContent(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200N"
	}
	return false
}

// InvalidXMLComment checks if the error is of code 2200S
func InvalidXMLComment(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200S"
	}
	return false
}

// InvalidXMLProcessingInstruction checks if the error is of code 2200T
func InvalidXMLProcessingInstruction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2200T"
	}
	return false
}

// Class 23 - Integrity Constraint Violation

// IntegrityConstraintViolation checks if the error is of code 23000
func IntegrityConstraintViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23000"
	}
	return false
}

// RestrictViolation checks if the error is of code 23001
func RestrictViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23001"
	}
	return false
}

// NotNullViolation checks if the error is of code 23502
func NotNullViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23502"
	}
	return false
}

// ForeignKeyViolation checks if the error is of code 23503
func ForeignKeyViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23503"
	}
	return false
}

// UniqueViolation checks if the error is of code 23505
func UniqueViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23505"
	}
	return false
}

// CheckViolation checks if the error is of code 23514
func CheckViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23514"
	}
	return false
}

// ExclusionViolation checks if the error is of code 23P01
func ExclusionViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "23P01"
	}
	return false
}

// Class 24 - Invalid Cursor State

// InvalidCursorState checks if the error is of code 24000
func InvalidCursorState(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "24000"
	}
	return false
}

// Class 25 - Invalid Transaction State

// InvalidTransactionState checks if the error is of code 25000
func InvalidTransactionState(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25000"
	}
	return false
}

// ActiveSQLTransaction checks if the error is of code 25001
func ActiveSQLTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25001"
	}
	return false
}

// BranchTransactionAlreadyActive checks if the error is of code 25002
func BranchTransactionAlreadyActive(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25002"
	}
	return false
}

// HeldCursorRequiresSameIsolationLevel checks if the error is of code 25008
func HeldCursorRequiresSameIsolationLevel(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25008"
	}
	return false
}

// InappropriateAccessModeForBranchTransaction checks if the error is of code 25003
func InappropriateAccessModeForBranchTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25003"
	}
	return false
}

// InappropriateIsolationLevelForBranchTransaction checks if the error is of code 25004
func InappropriateIsolationLevelForBranchTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25004"
	}
	return false
}

// NoActiveSQLTransactionForBranchTransaction checks if the error is of code 25005
func NoActiveSQLTransactionForBranchTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25005"
	}
	return false
}

// ReadOnlySQLTransaction checks if the error is of code 25006
func ReadOnlySQLTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25006"
	}
	return false
}

// SchemaAndDataStatementMixingNotSupported checks if the error is of code 25007
func SchemaAndDataStatementMixingNotSupported(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25007"
	}
	return false
}

// NoActiveSQLTransaction checks if the error is of code 25P01
func NoActiveSQLTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25P01"
	}
	return false
}

// InFailedSQLTransaction checks if the error is of code 25P02
func InFailedSQLTransaction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "25P02"
	}
	return false
}

// Class 26 - Invalid SQL Statement Name

// InvalidSQLStatementName checks if the error is of code 26000
func InvalidSQLStatementName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "26000"
	}
	return false
}

// Class 27 - Triggered Data Change Violation

// TriggeredDataChangeViolation checks if the error is of code 27000
func TriggeredDataChangeViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "27000"
	}
	return false
}

// Class 28 - Invalid Authorization Specification

// InvalidAuthorizationSpecification checks if the error is of code 28000
func InvalidAuthorizationSpecification(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "28000"
	}
	return false
}

// InvalidPassword checks if the error is of code 28P01
func InvalidPassword(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "28P01"
	}
	return false
}

// Class 2B - Dependent Privilege Descriptors Still Exist

// DependentPrivilegeDescriptorsStillExist checks if the error is of code 2B000
func DependentPrivilegeDescriptorsStillExist(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2B000"
	}
	return false
}

// DependentObjectsStillExist checks if the error is of code 2BP01
func DependentObjectsStillExist(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2BP01"
	}
	return false
}

// Class 2D - Invalid Transaction Termination

// InvalidTransactionTermination checks if the error is of code 2D000
func InvalidTransactionTermination(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2D000"
	}
	return false
}

// Class 2F - SQL Routine Exception

// SQLRoutineException checks if the error is of code 2F000
func SQLRoutineException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2F000"
	}
	return false
}

// FunctionExecutedNoReturnStatement checks if the error is of code 2F005
func FunctionExecutedNoReturnStatement(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2F005"
	}
	return false
}

// ModifyingSQLDataNotPermitted checks if the error is of code 2F002
func ModifyingSQLDataNotPermitted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2F002" || pgerr.Code == "38002"
	}
	return false
}

// ProhibitedSQLStatementAttempted checks if the error is of code 2F003
func ProhibitedSQLStatementAttempted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2F003" || pgerr.Code == "38003"
	}
	return false
}

// ReadingSQLDataNotPermitted checks if the error is of code 2F004
func ReadingSQLDataNotPermitted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "2F004" || pgerr.Code == "38004"
	}
	return false
}

// Class 34 - Invalid Cursor Name

// InvalidCursorName checks if the error is of code 34000
func InvalidCursorName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "34000"
	}
	return false
}

// Class 38 - External Routine Exception

// ExternalRoutineException checks if the error is of code 38000
func ExternalRoutineException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "38000"
	}
	return false
}

// ContainingSQLNotPermitted checks if the error is of code 38001
func ContainingSQLNotPermitted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "38001"
	}
	return false
}

// Class 39 - External Routine Invocation Exception

// ExternalRoutineInvocationException checks if the error is of code 39000
func ExternalRoutineInvocationException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "39000"
	}
	return false
}

// InvalidSQLstateReturned checks if the error is of code 39001
func InvalidSQLstateReturned(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "39001"
	}
	return false
}

// TriggerProtocolViolated checks if the error is of code 39P01
func TriggerProtocolViolated(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "39P01"
	}
	return false
}

// SrfProtocolViolated checks if the error is of code 39P02
func SrfProtocolViolated(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "39P02"
	}
	return false
}

// Class 3B - Savepoint Exception

// SavepointException checks if the error is of code 3B000
func SavepointException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "3B000"
	}
	return false
}

// InvalidSavepointSpecification checks if the error is of code 3B001
func InvalidSavepointSpecification(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "3B001"
	}
	return false
}

// Class 3D - Invalid Catalog Name

// InvalidCatalogName checks if the error is of code 3D000
func InvalidCatalogName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "3D000"
	}
	return false
}

// Class 3F - Invalid Schema Name

// InvalidSchemaName checks if the error is of code 3F000
func InvalidSchemaName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "3F000"
	}
	return false
}

// Class 40 - Transaction Rollback

// TransactionRollback checks if the error is of code 40000
func TransactionRollback(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "40000"
	}
	return false
}

// TransactionIntegrityConstraintViolation checks if the error is of code 40002
func TransactionIntegrityConstraintViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "40002"
	}
	return false
}

// SerializationFailure checks if the error is of code 40001
func SerializationFailure(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "40001"
	}
	return false
}

// StatementCompletionUnknown checks if the error is of code 40003
func StatementCompletionUnknown(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "40003"
	}
	return false
}

// DeadlockDetected checks if the error is of code 40P01
func DeadlockDetected(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "40P01"
	}
	return false
}

// Class 42 - Syntax Error or Access Rule Violation

// SyntaxErrorOrAccessRuleViolation checks if the error is of code 42000
func SyntaxErrorOrAccessRuleViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42000"
	}
	return false
}

// SyntaxError checks if the error is of code 42601
func SyntaxError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42601"
	}
	return false
}

// InsufficientPrivilege checks if the error is of code 42501
func InsufficientPrivilege(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42501"
	}
	return false
}

// CannotCoerce checks if the error is of code 42846
func CannotCoerce(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42846"
	}
	return false
}

// GroupingError checks if the error is of code 42803
func GroupingError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42803"
	}
	return false
}

// WindowingError checks if the error is of code 42P20
func WindowingError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P20"
	}
	return false
}

// InvalidRecursion checks if the error is of code 42P19
func InvalidRecursion(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P19"
	}
	return false
}

// InvalidForeignKey checks if the error is of code 42830
func InvalidForeignKey(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42830"
	}
	return false
}

// InvalidName checks if the error is of code 42602
func InvalidName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42602"
	}
	return false
}

// NameTooLong checks if the error is of code 42622
func NameTooLong(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42622"
	}
	return false
}

// ReservedName checks if the error is of code 42939
func ReservedName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42939"
	}
	return false
}

// DatatypeMismatch checks if the error is of code 42804
func DatatypeMismatch(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42804"
	}
	return false
}

// IndeterminateDatatype checks if the error is of code 42P18
func IndeterminateDatatype(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P18"
	}
	return false
}

// CollationMismatch checks if the error is of code 42P21
func CollationMismatch(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P21"
	}
	return false
}

// IndeterminateCollation checks if the error is of code 42P22
func IndeterminateCollation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P22"
	}
	return false
}

// WrongObjectType checks if the error is of code 42809
func WrongObjectType(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42809"
	}
	return false
}

// UndefinedColumn checks if the error is of code 42703
func UndefinedColumn(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42703"
	}
	return false
}

// UndefinedFunction checks if the error is of code 42883
func UndefinedFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42883"
	}
	return false
}

// UndefinedTable checks if the error is of code 42P01
func UndefinedTable(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P01"
	}
	return false
}

// UndefinedParameter checks if the error is of code 42P02
func UndefinedParameter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P02"
	}
	return false
}

// UndefinedObject checks if the error is of code 42704
func UndefinedObject(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42704"
	}
	return false
}

// DuplicateColumn checks if the error is of code 42701
func DuplicateColumn(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42701"
	}
	return false
}

// DuplicateCursor checks if the error is of code 42P03
func DuplicateCursor(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P03"
	}
	return false
}

// DuplicateDatabase checks if the error is of code 42P04
func DuplicateDatabase(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P04"
	}
	return false
}

// DuplicateFunction checks if the error is of code 42723
func DuplicateFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42723"
	}
	return false
}

// DuplicatePreparedStatement checks if the error is of code 42P05
func DuplicatePreparedStatement(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P05"
	}
	return false
}

// DuplicateSchema checks if the error is of code 42P06
func DuplicateSchema(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P06"
	}
	return false
}

// DuplicateTable checks if the error is of code 42P07
func DuplicateTable(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P07"
	}
	return false
}

// DuplicateAlias checks if the error is of code 42712
func DuplicateAlias(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42712"
	}
	return false
}

// DuplicateObject checks if the error is of code 42710
func DuplicateObject(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42710"
	}
	return false
}

// AmbiguousColumn checks if the error is of code 42702
func AmbiguousColumn(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42702"
	}
	return false
}

// AmbiguousFunction checks if the error is of code 42725
func AmbiguousFunction(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42725"
	}
	return false
}

// AmbiguousParameter checks if the error is of code 42P08
func AmbiguousParameter(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P08"
	}
	return false
}

// AmbiguousAlias checks if the error is of code 42P09
func AmbiguousAlias(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P09"
	}
	return false
}

// InvalidColumnReference checks if the error is of code 42P10
func InvalidColumnReference(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P10"
	}
	return false
}

// InvalidColumnDefinition checks if the error is of code 42611
func InvalidColumnDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42611"
	}
	return false
}

// InvalidCursorDefinition checks if the error is of code 42P11
func InvalidCursorDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P11"
	}
	return false
}

// InvalidDatabaseDefinition checks if the error is of code 42P12
func InvalidDatabaseDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P12"
	}
	return false
}

// InvalidFunctionDefinition checks if the error is of code 42P13
func InvalidFunctionDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P13"
	}
	return false
}

// InvalidPreparedStatementDefinition checks if the error is of code 42P14
func InvalidPreparedStatementDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P14"
	}
	return false
}

// InvalidSchemaDefinition checks if the error is of code 42P15
func InvalidSchemaDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P15"
	}
	return false
}

// InvalidTableDefinition checks if the error is of code 42P16
func InvalidTableDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P16"
	}
	return false
}

// InvalidObjectDefinition checks if the error is of code 42P17
func InvalidObjectDefinition(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "42P17"
	}
	return false
}

// Class 44 - WITH CHECK OPTION Violation

// WithCheckOptionViolation checks if the error is of code 44000
func WithCheckOptionViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "44000"
	}
	return false
}

// Class 53 - Insufficient Resources

// InsufficientResources checks if the error is of code 53000
func InsufficientResources(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "53000"
	}
	return false
}

// DiskFull checks if the error is of code 53100
func DiskFull(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "53100"
	}
	return false
}

// OutOfMemory checks if the error is of code 53200
func OutOfMemory(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "53200"
	}
	return false
}

// TooManyConnections checks if the error is of code 53300
func TooManyConnections(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "53300"
	}
	return false
}

// ConfigurationLimitExceeded checks if the error is of code 53400
func ConfigurationLimitExceeded(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "53400"
	}
	return false
}

// Class 54 - Program Limit Exceeded

// ProgramLimitExceeded checks if the error is of code 54000
func ProgramLimitExceeded(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "54000"
	}
	return false
}

// StatementTooComplex checks if the error is of code 54001
func StatementTooComplex(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "54001"
	}
	return false
}

// TooManyColumns checks if the error is of code 54011
func TooManyColumns(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "54011"
	}
	return false
}

// TooManyArguments checks if the error is of code 54023
func TooManyArguments(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "54023"
	}
	return false
}

// Class 55 - Object Not In Prerequisite State

// ObjectNotInPrerequisiteState checks if the error is of code 55000
func ObjectNotInPrerequisiteState(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "55000"
	}
	return false
}

// ObjectInUse checks if the error is of code 55006
func ObjectInUse(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "55006"
	}
	return false
}

// CantChangeRuntimeParam checks if the error is of code 55P02
func CantChangeRuntimeParam(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "55P02"
	}
	return false
}

// LockNotAvailable checks if the error is of code 55P03
func LockNotAvailable(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "55P03"
	}
	return false
}

// Class 57 - Operator Intervention

// OperatorIntervention checks if the error is of code 57000
func OperatorIntervention(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "57000"
	}
	return false
}

// QueryCanceled checks if the error is of code 57014
func QueryCanceled(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "57014"
	}
	return false
}

// AdminShutdown checks if the error is of code 57P01
func AdminShutdown(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "57P01"
	}
	return false
}

// CrashShutdown checks if the error is of code 57P02
func CrashShutdown(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "57P02"
	}
	return false
}

// CannotConnectNow checks if the error is of code 57P03
func CannotConnectNow(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "57P03"
	}
	return false
}

// DatabaseDropped checks if the error is of code 57P04
func DatabaseDropped(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "57P04"
	}
	return false
}

// Class 58 - System Error (errors external to PostgreSQL itself)

// SystemError checks if the error is of code 58000
func SystemError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "58000"
	}
	return false
}

// IoError checks if the error is of code 58030
func IoError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "58030"
	}
	return false
}

// UndefinedFile checks if the error is of code 58P01
func UndefinedFile(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "58P01"
	}
	return false
}

// DuplicateFile checks if the error is of code 58P02
func DuplicateFile(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "58P02"
	}
	return false
}

// Class F0 - Configuration File Error

// ConfigFileError checks if the error is of code F0000
func ConfigFileError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "F0000"
	}
	return false
}

// LockFileExists checks if the error is of code F0001
func LockFileExists(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "F0001"
	}
	return false
}

// Class HV - Foreign Data Wrapper Error (SQL/MED)

// FdwError checks if the error is of code HV000
func FdwError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV000"
	}
	return false
}

// FdwColumnNameNotFound checks if the error is of code HV005
func FdwColumnNameNotFound(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV005"
	}
	return false
}

// FdwDynamicParameterValueNeeded checks if the error is of code HV002
func FdwDynamicParameterValueNeeded(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV002"
	}
	return false
}

// FdwFunctionSequenceError checks if the error is of code HV010
func FdwFunctionSequenceError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV010"
	}
	return false
}

// FdwInconsistentDescriptorInformation checks if the error is of code HV021
func FdwInconsistentDescriptorInformation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV021"
	}
	return false
}

// FdwInvalidAttributeValue checks if the error is of code HV024
func FdwInvalidAttributeValue(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV024"
	}
	return false
}

// FdwInvalidColumnName checks if the error is of code HV007
func FdwInvalidColumnName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV007"
	}
	return false
}

// FdwInvalidColumnNumber checks if the error is of code HV008
func FdwInvalidColumnNumber(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV008"
	}
	return false
}

// FdwInvalidDataType checks if the error is of code HV004
func FdwInvalidDataType(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV004"
	}
	return false
}

// FdwInvalidDataTypeDescriptors checks if the error is of code HV006
func FdwInvalidDataTypeDescriptors(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV006"
	}
	return false
}

// FdwInvalidDescriptorFieldIdentifier checks if the error is of code HV091
func FdwInvalidDescriptorFieldIdentifier(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV091"
	}
	return false
}

// FdwInvalidHandle checks if the error is of code HV00B
func FdwInvalidHandle(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00B"
	}
	return false
}

// FdwInvalidOptionIndex checks if the error is of code HV00C
func FdwInvalidOptionIndex(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00C"
	}
	return false
}

// FdwInvalidOptionName checks if the error is of code HV00D
func FdwInvalidOptionName(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00D"
	}
	return false
}

// FdwInvalidStringLengthOrBufferLength checks if the error is of code HV090
func FdwInvalidStringLengthOrBufferLength(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV090"
	}
	return false
}

// FdwInvalidStringFormat checks if the error is of code HV00A
func FdwInvalidStringFormat(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00A"
	}
	return false
}

// FdwInvalidUseOfNullPointer checks if the error is of code HV009
func FdwInvalidUseOfNullPointer(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV009"
	}
	return false
}

// FdwTooManyHandles checks if the error is of code HV014
func FdwTooManyHandles(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV014"
	}
	return false
}

// FdwOutOfMemory checks if the error is of code HV001
func FdwOutOfMemory(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV001"
	}
	return false
}

// FdwNoSchemas checks if the error is of code HV00P
func FdwNoSchemas(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00P"
	}
	return false
}

// FdwOptionNameNotFound checks if the error is of code HV00J
func FdwOptionNameNotFound(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00J"
	}
	return false
}

// FdwReplyHandle checks if the error is of code HV00K
func FdwReplyHandle(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00K"
	}
	return false
}

// FdwSchemaNotFound checks if the error is of code HV00Q
func FdwSchemaNotFound(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00Q"
	}
	return false
}

// FdwTableNotFound checks if the error is of code HV00R
func FdwTableNotFound(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00R"
	}
	return false
}

// FdwUnableToCreateExecution checks if the error is of code HV00L
func FdwUnableToCreateExecution(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00L"
	}
	return false
}

// FdwUnableToCreateReply checks if the error is of code HV00M
func FdwUnableToCreateReply(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00M"
	}
	return false
}

// FdwUnableToEstablishConnection checks if the error is of code HV00N
func FdwUnableToEstablishConnection(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "HV00N"
	}
	return false
}

// Class P0 - PL/pgSQL Error

// PlpgsqlError checks if the error is of code P0000
func PlpgsqlError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "P0000"
	}
	return false
}

// RaiseException checks if the error is of code P0001
func RaiseException(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "P0001"
	}
	return false
}

// NoDataFound checks if the error is of code P0002
func NoDataFound(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "P0002"
	}
	return false
}

// TooManyRows checks if the error is of code P0003
func TooManyRows(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "P0003"
	}
	return false
}

// Class XX - Internal Error

// InternalError checks if the error is of code XX000
func InternalError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "XX000"
	}
	return false
}

// DataCorrupted checks if the error is of code XX001
func DataCorrupted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "XX001"
	}
	return false
}

// IndexCorrupted checks if the error is of code XX002
func IndexCorrupted(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == "XX002"
	}
	return false
}
