package account

//import (
//	"testing"
//
//	"github.com/andyj29/wannabet/internal/domain/common"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestAddFunds(t *testing.T) {
//	t.Parallel()
//
//	testCases := map[string]struct {
//		inititalBalance, amountToBeAdded  int64
//		initialCurrencyCode, currencyCode string
//		expectedErr                       bool
//		err                               error
//	}{
//		"Valid - Add 100 CAD to 10 CAD": {
//			inititalBalance:     10,
//			initialCurrencyCode: "CAD",
//			amountToBeAdded:     100,
//			currencyCode:        "CAD",
//		},
//		"Invalid / Can't add 0 - Add 0 CAD": {
//			inititalBalance:     100,
//			initialCurrencyCode: "CAD",
//			currencyCode:        "CAD",
//			expectedErr:         true,
//			err:                 c.FundsNotAddable,
//		},
//		"Invalid / Incompatible currency - Add 100 VND to 50 CAD": {
//			inititalBalance:     50,
//			initialCurrencyCode: "CAD",
//			amountToBeAdded:     100,
//			currencyCode:        "VND",
//			expectedErr:         true,
//			err:                 common.FundsNotAddable,
//		},
//	}
//
//	for name, testCase := range testCases {
//		testCase := testCase
//		t.Run(name, func(t *testing.T) {
//			t.Parallel()
//
//			instance, _ := NewAccount("testID", Email("testEmail"), "testName")
//			initialBalance, _ := common.NewMoney(testCase.inititalBalance, testCase.initialCurrencyCode)
//			if err := instance.AddFunds(initialBalance); err != nil {
//				t.Errorf("Failed to initialize account with a balance of %v", initialBalance)
//			}
//
//			amountToBeAdded, _ := common.NewMoney(testCase.amountToBeAdded, testCase.currencyCode)
//			err := instance.AddFunds(amountToBeAdded)
//
//			if testCase.expectedErr {
//				assert.Equal(t, testCase.err, err)
//				assert.True(t, instance.Balance.IsEqual(initialBalance))
//				return
//			}
//
//			expectedBalance, _ := initialBalance.Add(amountToBeAdded)
//			assert.Nilf(t, err, "err should be nil")
//			assert.True(t, instance.Balance.IsEqual(expectedBalance))
//		})
//	}
//}
//
//func TestDeductFunds(t *testing.T) {
//	t.Parallel()
//
//	testCases := map[string]struct {
//		initialBalance, amountToBeDeducted int64
//		initialCurrencyCode, currencyCode  string
//		expectedErr                        bool
//		err                                error
//	}{
//		"Valid - Deduct 50 CAD from 100 CAD": {
//			initialBalance:      100,
//			initialCurrencyCode: "CAD",
//			amountToBeDeducted:  50,
//			currencyCode:        "CAD",
//		},
//		"Invalid / Insufficient funds - Deduct 100 CAD from 50 CAD": {
//			initialBalance:      50,
//			initialCurrencyCode: "CAD",
//			amountToBeDeducted:  150,
//			currencyCode:        "CAD",
//			expectedErr:         true,
//			err:                 common.FundsNotDeductible,
//		},
//		"Invalid / Incompatible currency - Deduct 1 VND from 100 CAD": {
//			initialBalance:      100,
//			initialCurrencyCode: "CAD",
//			amountToBeDeducted:  1,
//			currencyCode:        "VND",
//			expectedErr:         true,
//			err:                 common.FundsNotDeductible,
//		},
//	}
//
//	for name, testCase := range testCases {
//		testCase := testCase
//		t.Run(name, func(t *testing.T) {
//			t.Parallel()
//
//			instance, _ := NewAccount("testID", Email("testEmail"), "testName")
//			initialBalance, _ := common.NewMoney(testCase.initialBalance, testCase.initialCurrencyCode)
//			if err := instance.AddFunds(initialBalance); err != nil {
//				t.Errorf("Failed to initialize account with a balance of %v", initialBalance)
//			}
//
//			amountToBeDeducted, _ := common.NewMoney(testCase.amountToBeDeducted, testCase.currencyCode)
//			err := instance.DeductFunds(amountToBeDeducted)
//
//			if testCase.expectedErr {
//				assert.Equal(t, testCase.err, err)
//				assert.True(t, instance.Balance.IsEqual(initialBalance))
//				return
//			}
//
//			expectedBalance, _ := initialBalance.Deduct(amountToBeDeducted)
//			assert.Nilf(t, err, "err should be nil")
//			assert.True(t, instance.Balance.IsEqual(expectedBalance))
//
//		})
//	}
//}
