package models

import (
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for listing transactions
type ListTransactionsResponseInterface interface {
	GetData() types.Optional[[]GetTransactionResponseInterface]
	GetPaging() types.Optional[PagingResponse]
}

// Response object for getting a charge
type GetChargeResponseInterface interface {
	GetId() types.Optional[string]
	GetCode() types.Optional[string]
	GetGatewayId() types.Optional[string]
	GetAmount() types.Optional[int]
	GetStatus() types.Optional[string]
	GetCurrency() types.Optional[string]
	GetPaymentMethod() types.Optional[string]
	GetDueAt() types.Optional[time.Time]
	GetCreatedAt() types.Optional[time.Time]
	GetUpdatedAt() types.Optional[time.Time]
	GetLastTransaction() types.Optional[GetTransactionResponseInterface]
	GetInvoice() types.Optional[GetInvoiceResponse]
	GetOrder() types.Optional[GetOrderResponse]
	GetCustomer() types.Optional[GetCustomerResponse]
	GetMetadata() types.Optional[map[string]string]
	GetPaidAt() types.Optional[time.Time]
	GetCanceledAt() types.Optional[time.Time]
	GetCanceledAmount() types.Optional[int]
	GetPaidAmount() types.Optional[int]
	GetInterestAndFinePaid() types.Optional[int]
	GetRecurrencyCycle() types.Optional[string]
}

// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetUrl() *string
	GetBankTid() *string
	GetBank() *string
	GetPaidAt() *time.Time
	GetPaidAmount() *int
}

// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetUrl() types.Optional[string]
	GetBankTid() types.Optional[string]
	GetPaidAt() types.Optional[time.Time]
	GetPaidAmount() types.Optional[int]
}

// Response for voucher transactions
type GetVoucherTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() types.Optional[string]
	GetAcquirerName() types.Optional[string]
	GetAcquirerAffiliationCode() types.Optional[string]
	GetAcquirerTid() types.Optional[string]
	GetAcquirerNsu() types.Optional[string]
	GetAcquirerAuthCode() types.Optional[string]
	GetAcquirerMessage() types.Optional[string]
	GetAcquirerReturnCode() types.Optional[string]
	GetOperationType() types.Optional[string]
	GetCard() types.Optional[GetCardResponse]
}

// Response object for getting a boleto transaction
type GetBoletoTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetUrl() types.Optional[string]
	GetBarcode() types.Optional[string]
	GetNossoNumero() types.Optional[string]
	GetBank() types.Optional[string]
	GetDocumentNumber() types.Optional[string]
	GetInstructions() types.Optional[string]
	GetBillingAddress() types.Optional[GetBillingAddressResponse]
	GetDueAt() types.Optional[time.Time]
	GetQrCode() types.Optional[string]
	GetLine() types.Optional[string]
	GetPdfPassword() types.Optional[string]
	GetPdf() types.Optional[string]
	GetPaidAt() types.Optional[time.Time]
	GetPaidAmount() types.Optional[string]
	GetType() types.Optional[string]
	GetCreditAt() types.Optional[time.Time]
	GetStatementDescriptor() types.Optional[string]
}

// Response object for getting a debit card transaction
type GetDebitCardTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() types.Optional[string]
	GetAcquirerName() types.Optional[string]
	GetAcquirerAffiliationCode() types.Optional[string]
	GetAcquirerTid() types.Optional[string]
	GetAcquirerNsu() types.Optional[string]
	GetAcquirerAuthCode() types.Optional[string]
	GetOperationType() types.Optional[string]
	GetCard() types.Optional[GetCardResponse]
	GetAcquirerMessage() types.Optional[string]
	GetAcquirerReturnCode() types.Optional[string]
	GetMpi() types.Optional[string]
	GetEci() types.Optional[string]
	GetAuthenticationType() types.Optional[string]
	GetThreedAuthenticationUrl() types.Optional[string]
	GetFundingSource() types.Optional[string]
}

// Generic response object for getting a transaction.
type GetTransactionResponseInterface interface {
	GetGatewayId() types.Optional[string]
	GetAmount() types.Optional[int]
	GetStatus() types.Optional[string]
	GetSuccess() types.Optional[bool]
	GetCreatedAt() types.Optional[time.Time]
	GetUpdatedAt() types.Optional[time.Time]
	GetAttemptCount() types.Optional[int]
	GetMaxAttempts() types.Optional[int]
	GetSplits() types.Optional[[]GetSplitResponse]
	GetNextAttempt() types.Optional[time.Time]
	GetTransactionType() *string
	GetId() types.Optional[string]
	GetGatewayResponse() types.Optional[GetGatewayResponseResponse]
	GetAntifraudResponse() types.Optional[GetAntifraudResponse]
	GetMetadata() types.Optional[map[string]string]
	GetSplit() types.Optional[[]GetSplitResponse]
	GetInterest() types.Optional[GetInterestResponse]
	GetFine() types.Optional[GetFineResponse]
	GetMaxDaysToPayPastDue() types.Optional[int]
}

// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() types.Optional[string]
	GetAcquirerName() types.Optional[string]
	GetAcquirerAffiliationCode() types.Optional[string]
	GetAcquirerTid() types.Optional[string]
	GetAcquirerNsu() types.Optional[string]
	GetAcquirerAuthCode() types.Optional[string]
	GetOperationType() types.Optional[string]
	GetCard() types.Optional[GetCardResponse]
	GetAcquirerMessage() types.Optional[string]
	GetAcquirerReturnCode() types.Optional[string]
	GetInstallments() types.Optional[int]
}

// Response object for getting a cash transaction
type GetCashTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetDescription() types.Optional[string]
}

// Response object for getting a credit card transaction
type GetCreditCardTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetStatementDescriptor() types.Optional[string]
	GetAcquirerName() *string
	GetAcquirerAffiliationCode() types.Optional[string]
	GetAcquirerTid() *string
	GetAcquirerNsu() *string
	GetAcquirerAuthCode() types.Optional[string]
	GetOperationType() types.Optional[string]
	GetCard() types.Optional[GetCardResponse]
	GetAcquirerMessage() types.Optional[string]
	GetAcquirerReturnCode() types.Optional[string]
	GetInstallments() types.Optional[int]
	GetThreedAuthenticationUrl() types.Optional[string]
	GetFundingSource() types.Optional[string]
}

// Response object for listing charge transactions
type ListChargeTransactionsResponseInterface interface {
	GetData() types.Optional[[]GetTransactionResponseInterface]
	GetPaging() types.Optional[PagingResponse]
}

// Response object when getting a pix transaction
type GetPixTransactionResponseInterface interface {
	GetTransactionResponseInterface
	GetQrCode() types.Optional[string]
	GetQrCodeUrl() types.Optional[string]
	GetExpiresAt() types.Optional[time.Time]
	GetAdditionalInformation() types.Optional[[]PixAdditionalInformation]
	GetEndToEndId() types.Optional[string]
	GetPayer() types.Optional[GetPixPayerResponse]
	GetPixProviderTid() types.Optional[string]
}

// Generic response object for getting a BalanceOperation.
type GetBalanceOperationResponseInterface interface {
	GetId() types.Optional[string]
	GetStatus() types.Optional[string]
	GetBalanceAmount() types.Optional[string]
	GetBalanceOldAmount() types.Optional[string]
	GetType() types.Optional[string]
	GetAmount() types.Optional[string]
	GetFee() types.Optional[string]
	GetCreatedAt() types.Optional[string]
	GetMovementObject() *GetMovementObjectBaseResponseInterface
}

// Generic response object for getting a MovementObjectBase.
type GetMovementObjectBaseResponseInterface interface {
	GetObject() *string
	GetId() types.Optional[string]
	GetStatus() types.Optional[string]
	GetAmount() types.Optional[string]
	GetCreatedAt() types.Optional[string]
	GetType() types.Optional[string]
	GetChargeId() types.Optional[string]
	GetGatewayId() types.Optional[string]
}

// Generic response object for getting a MovementObjectRefund.
type GetMovementObjectRefundResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetFraudCoverageFee() types.Optional[string]
	GetChargeFeeRecipientId() types.Optional[string]
	GetBankAccountId() types.Optional[string]
	GetLocalTransactionId() types.Optional[string]
	GetUpdatedAt() types.Optional[string]
}

// Generic response object for getting a MovementObjectFeeCollection.
type GetMovementObjectFeeCollectionResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetDescription() types.Optional[string]
	GetPaymentDate() types.Optional[string]
	GetRecipientId() types.Optional[string]
}

type GetMovementObjectPayableResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetFee() types.Optional[string]
	GetAnticipationFee() string
	GetFraudCoverageFee() string
	GetInstallment() string
	GetSplitId() string
	GetBulkAnticipationId() string
	GetAnticipationId() string
	GetRecipientId() string
	GetOriginatorModel() string
	GetOriginatorModelId() string
	GetPaymentDate() string
	GetOriginalPaymentDate() string
	GetPaymentMethod() string
	GetAccrualAt() string
	GetLiquidationArrangementId() string
}

type GetMovementObjectTransferResponseInterface interface {
	GetMovementObjectBaseResponseInterface
	GetSourceType() types.Optional[string]
	GetSourceId() types.Optional[string]
	GetTargetType() types.Optional[string]
	GetTargetId() types.Optional[string]
	GetFee() types.Optional[string]
	GetFundingDate() types.Optional[string]
	GetFundingEstimatedDate() types.Optional[string]
	GetBankAccount() types.Optional[string]
}
