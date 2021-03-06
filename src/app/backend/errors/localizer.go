package errors

import (
	"strings"

	"gopkg.in/square/go-jose.v2"
)

const (
	MsgDeployNamespaceMismatchError    = "MSG_DEPLOY_NAMESPACE_MISMATCH_ERROR"
	MsgDeployEmptyNamespaceError       = "MSG_DEPLOY_EMPTY_NAMESPACE_ERROR"
	MsgLoginUnauthorizedError          = "MSG_LOGIN_UNAUTHORIZED_ERROR"
	MsgEncryptionKeyChanged            = "MSG_ENCRYPTION_KEY_CHANGED"
	MsgDashboardExclusiveResourceError = "MSG_DASHBOARD_EXCLUSIVE_RESOURCE_ERROR"
	MsgTokenExpiredError               = "MSG_TOKEN_EXPIRED_ERROR"
)

var partialsToErrorsMap = map[string]string{
	"does not match the namespace":                               MsgDeployNamespaceMismatchError,
	"empty namespace may not be set":                             MsgDeployEmptyNamespaceError,
	"the server has asked for the client to provide credentials": MsgLoginUnauthorizedError,
	jose.ErrCryptoFailure.Error():                                MsgEncryptionKeyChanged,
}

func LocalizeError(err error) error {
	if err == nil {
		return nil
	}

	for partial, errString := range partialsToErrorsMap {
		if strings.Contains(err.Error(), partial) {
			if IsUnauthorized(err) {
				return NewUnauthorized(errString)
			}

			return NewBadRequest(errString)
		}
	}

	return err
}
