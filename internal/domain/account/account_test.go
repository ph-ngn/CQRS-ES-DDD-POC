package account

import (
	"testing"

	"github.com/andyj29/wannabet/internal/domain/common"
	"github.com/stretchr/testify/assert"
)

func TestAddFunds(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inititalBalance, amountToBeAdded  int64
		initialCurrencyCode, currencyCode string
		errExpected                       bool
		expectedErr                       error
	}{
		"Valid - Add 100 CAD": {
			inititalBalance:     10,
			initialCurrencyCode: "CAD",
			amountToBeAdded:     100,
			currencyCode:        "CAD",
			errExpected:         false,
		},
		"Invalid / Can't add 0 - Add 0 CAD": {
			inititalBalance:     100,
			initialCurrencyCode: "CAD",
			currencyCode:        "CAD",
			errExpected:         true,
			expectedErr:         common.FundsNotAddable,
		},
		"Invalid / Incompatible currency - Add 100 VND": {
			inititalBalance:     50,
			initialCurrencyCode: "CAD",
			amountToBeAdded:     100,
			currencyCode:        "VND",
			errExpected:         true,
			expectedErr:         common.FundsNotAddable,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			instance := NewAccount("testID", Email("testEmail"), "testName")
			initialBalance := common.NewMoney(testCase.inititalBalance, testCase.initialCurrencyCode)
			if err := instance.AddFunds(initialBalance); err != nil {
				t.Errorf("Failed to initialize account with a balance of %v", initialBalance)
			}

			amountToBeAdded := common.NewMoney(testCase.amountToBeAdded, testCase.currencyCode)
			err := instance.AddFunds(amountToBeAdded)

			if testCase.errExpected {
				assert.Equal(t, testCase.expectedErr, err)
				assert.True(t, instance.Balance.IsEqual(initialBalance))
				return
			}

			expectedBalance, err := initialBalance.Add(amountToBeAdded)
			if err != nil {
				t.Errorf("Failed to compute expected balance after success add")
			}
			assert.Nilf(t, err, "err should be nil")
			assert.True(t, instance.Balance.IsEqual(expectedBalance))
		})
	}
}

func TestDeductFunds(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		initialBalance, amountToBeDeducted int64
		initialCurrencyCode, currencyCode  string
		errExpected                        bool
		expectedErr                        error
	}{
		"Valid - Deduct 50 CAD from 100 CAD": {
			initialBalance:      100,
			initialCurrencyCode: "CAD",
			amountToBeDeducted:  50,
			currencyCode:        "CAD",
			errExpected:         false,
			expectedErr:         common.FundsNotDeductible,
		},
		"Invalid / Insufficient funds - Deduct 100 CAD": {
			initialBalance:      50,
			initialCurrencyCode: "CAD",
			amountToBeDeducted:  150,
			currencyCode:        "CAD",
			errExpected:         true,
			expectedErr:         common.FundsNotDeductible,
		},
		"Invalid / Incompatible currency - Deduct 1 VND": {
			initialBalance:      100,
			initialCurrencyCode: "CAD",
			amountToBeDeducted:  1,
			currencyCode:        "VND",
			errExpected:         true,
			expectedErr:         common.FundsNotDeductible,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			instance := NewAccount("testID", Email("testEmail"), "testName")
			initialBalance := common.NewMoney(testCase.initialBalance, testCase.initialCurrencyCode)
			if err := instance.AddFunds(initialBalance); err != nil {
				t.Errorf("Failed to initialize account with a balance of %v", initialBalance)
			}

			amountToBeDeducted := common.NewMoney(testCase.amountToBeDeducted, testCase.currencyCode)
			err := instance.DeductFunds(amountToBeDeducted)

			if testCase.errExpected {
				assert.Equal(t, testCase.expectedErr, err)
				assert.True(t, instance.Balance.IsEqual(initialBalance))
				return
			}

			expectedBalance, err := initialBalance.Deduct(amountToBeDeducted)
			if err != nil {
				t.Errorf("Failed to compute expected balance after success deduct")
			}
			assert.Nilf(t, err, "err should be nil")
			assert.True(t, instance.Balance.IsEqual(expectedBalance))

		})
	}
}
