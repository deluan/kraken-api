package krakenapi

import (
	"net/url"
	"strconv"
)

/*
URL: https://api.kraken.com/0/private/Withdraw

Input:
aclass = asset class (optional):
    currency (default)
asset = asset being withdrawn
key = withdrawal key name, as set up on your account
amount = amount to withdraw, including fees

Result: associative array of withdrawal transaction:
refid = reference id
*/
func (api *KrakenApi) ApiWithdraw(asset string, key string, amount float64) (string, error) {
	params := url.Values{}
	params.Set("asset", asset)
	params.Set("key", key)
	params.Set("amount", strconv.FormatFloat(amount, 'f', -1, 64))

	resp, err := api.Query(URL_PRIVATE_WITHDRAW, params, true)
	if err != nil {
		return "", err
	}

	content, err := parse(resp, nil)
	if err != nil {
		return "", err
	}

	refIds, _ := content.(map[string]string)
	return refIds["refid"], nil
}

/*
URL: https://api.kraken.com/0/private/WithdrawStatus

Input:
aclass = asset class (optional):
    currency (default)
asset = asset being withdrawn
method = withdrawal method name (optional)

Result: array of array withdrawal status information:
method = name of the withdrawal method used
aclass = asset class
asset = asset X-ISO4217-A3 code
refid = reference id
txid = method transaction id
info = method transaction information
amount = amount withdrawn
fee = fees paid
time = unix timestamp when request was made
status = status of withdrawal
status-prop = additional status properties (if available)
    cancel-pending = cancelation requested
    canceled = canceled
    cancel-denied = cancelation requested but was denied
    return = a return transaction initiated by Kraken; it cannot be canceled
    onhold = withdrawal is on hold pending review
*/
func (api *KrakenApi) ApiWithdrawStatus(asset string, method string) ([]WithdrawStatus, error) {
	params := url.Values{}
	params.Set("asset", asset)
	params.Set("method", method)

	resp, err := api.Query(URL_PRIVATE_WITHDRAW_STATUS, params, true)
	if err != nil {
		return nil, err
	}

	var status []WithdrawStatus
	_, err = parse(resp, &status)
	return status, err
}
