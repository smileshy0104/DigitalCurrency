package repo

import "job-center/internal/model"

type BtcTransactionRepo interface {
	FindByTxId(txId string) (*model.BitCoinTransaction, error)
	Save(bt *model.BitCoinTransaction) error
}
