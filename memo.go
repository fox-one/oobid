package main

import (
	"encoding/base64"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/ugorji/go/codec"
)

// side : A (ask) or B (bid)
// priceType : L (limit) or  M (market)
func createMemo(side, price, priceType, assetId string) (string, error) {
	asset_uuid, err := uuid.FromString(assetId)
	if err != nil {
		return "", errors.Wrap(err, "invalid asset id")
	}

	action := map[string]interface{}{
		"S": side,
		"P": price,
		"T": priceType,
		"A": asset_uuid,
	}

	memo := make([]byte, 140)
	handle := new(codec.MsgpackHandle)
	encoder := codec.NewEncoderBytes(&memo, handle)
	if err := encoder.Encode(action); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(memo), nil
}
