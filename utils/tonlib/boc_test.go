package tonlib

import (
	"fmt"

	"testing"
)

func Test_boc(t *testing.T) {

	t.Run("boc", func(t *testing.T) {

		boc, _ := CreateTransaction(JettonTransferOption{
			Jetton:              "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS",
			Destination:         "UQAfB7KjPFWxD5GpnvQ6s2yhMaxig7seoSe8URS_o3vCw-DI",
			ResponseDestination: "UQDUewtDjeb4WwiSutRkXXTcne5jxL1QiUJt1WEy12Zz2Qpu",
			Message:             "order_000000010",
			Amount:              100e9,
		})

		fmt.Println("Generated BOC:", boc)
	})

}
