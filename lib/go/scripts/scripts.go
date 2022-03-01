package scripts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../scripts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../scripts/...

import (
	"fmt"
	"github.com/StarlyIO/starly-contracts/lib/go/scripts/internal/assets"
	"strings"
)

const (

    // Flow
	readFlowBalanceFilename = "flow/read_balance.cdc"

	// FUSD
	readFUSDBalanceFilename = "fusd/read_balance.cdc"

	// StarlyCard
	readStarlyCardCollectionIdsFilename    = "starlyCard/read_collection_ids.cdc"
	readStarlyCardCollectionLengthFilename = "starlyCard/read_collection_length.cdc"
	readStarlyCardSupplyFilename           = "starlyCard/read_supply.cdc"
	readStarlyIdFilename                   = "starlyCard/read_starly_id.cdc"

	// StarlyCardMarket
	readStarlyCardMarketCollectionIdsFilename    = "starlyCardMarket/read_collection_ids.cdc"
	readStarlyCardMarketCollectionLengthFilename = "starlyCardMarket/read_collection_length.cdc"

	defaultFungibleTokenAddress    = "../../contracts/FungibleToken.cdc"
	defaultNonFungibleTokenAddress = "../../contracts/NonFungibleToken.cdc"
	defaultFlowTokenAddress        = "../../contracts/FlowToken.cdc"
	defaultFUSDAddress             = "../../contracts/FUSD.cdcADDRESS"
	defaultStarlyCardAddress       = "../../contracts/StarlyCard.cdc"
	defaultStarlyCardMarketAddress = "../../contracts/StarlyCardMarket.cdc"
)

func withHexPrefix(address string) string {
	if address == "" {
		return ""
	}

	if address[0:2] == "0x" {
		return address
	}

	return fmt.Sprintf("0x%s", address)
}

func ReadFlowBalanceScript(
	fungibleTokenAddress string,
	flowTokenAddress string,
) []byte {
	code := assets.MustAssetString(readFlowBalanceFilename)

	code = strings.ReplaceAll(
		code,
		defaultFungibleTokenAddress,
		withHexPrefix(fungibleTokenAddress),
	)
	code = strings.ReplaceAll(
		code,
		defaultFlowTokenAddress,
		withHexPrefix(flowTokenAddress),
	)

	return []byte(code)
}

func ReadFUSDBalanceScript(
	fungibleTokenAddress string,
	fusdAddress string,
) []byte {
	code := assets.MustAssetString(readFUSDBalanceFilename)

	code = strings.ReplaceAll(
		code,
		defaultFungibleTokenAddress,
		withHexPrefix(fungibleTokenAddress),
	)
	code = strings.ReplaceAll(
		code,
		defaultFUSDAddress,
		withHexPrefix(fusdAddress),
	)

	return []byte(code)
}

func ReadStarlyCardCollectionIdsScript(
	nonFungibleTokenAddress string,
	starlyCardAddress string,
) []byte {
	code := assets.MustAssetString(readStarlyCardCollectionIdsFilename)

	code = strings.ReplaceAll(
		code,
		defaultNonFungibleTokenAddress,
		withHexPrefix(nonFungibleTokenAddress),
	)
	code = strings.ReplaceAll(
		code,
		defaultStarlyCardAddress,
		withHexPrefix(starlyCardAddress),
	)

	return []byte(code)
}

func ReadStarlyCardCollectionLengthScript(
	nonFungibleTokenAddress string,
	starlyCardAddress string,
) []byte {
	code := assets.MustAssetString(readStarlyCardCollectionLengthFilename)

	code = strings.ReplaceAll(
		code,
		defaultNonFungibleTokenAddress,
		withHexPrefix(nonFungibleTokenAddress),
	)
	code = strings.ReplaceAll(
		code,
		defaultStarlyCardAddress,
		withHexPrefix(starlyCardAddress),
	)

	return []byte(code)
}

func ReadStarlyIdScript(
	nonFungibleTokenAddress string,
	starlyCardAddress string,
) []byte {
	code := assets.MustAssetString(readStarlyIdFilename)

	code = strings.ReplaceAll(
		code,
		defaultNonFungibleTokenAddress,
		withHexPrefix(nonFungibleTokenAddress),
	)
	code = strings.ReplaceAll(
		code,
		defaultStarlyCardAddress,
		withHexPrefix(starlyCardAddress),
	)

	return []byte(code)
}

func ReadStarlyCardSupplyScript(
	starlyCardAddress string,
) []byte {
	code := assets.MustAssetString(readStarlyCardSupplyFilename)

	code = strings.ReplaceAll(
		code,
		defaultStarlyCardAddress,
		withHexPrefix(starlyCardAddress),
	)

	return []byte(code)
}

func ReadStarlyCardMarketCollectionIdsScript(
	starlyCardMarketAddress string,
) []byte {
	code := assets.MustAssetString(readStarlyCardMarketCollectionIdsFilename)

	code = strings.ReplaceAll(
		code,
		defaultStarlyCardMarketAddress,
		withHexPrefix(starlyCardMarketAddress),
	)

	return []byte(code)
}

func ReadStarlyCardMarketCollectionLengthScript(
	starlyCardMarketAddress string,
) []byte {
	code := assets.MustAssetString(readStarlyCardMarketCollectionLengthFilename)

	code = strings.ReplaceAll(
		code,
		defaultStarlyCardMarketAddress,
		withHexPrefix(starlyCardMarketAddress),
	)

	return []byte(code)
}
