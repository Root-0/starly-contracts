import StarlyCardStakingClaims from "../../contracts/StarlyCardStakingClaims.cdc"

pub fun main(address: Address): UFix64 {
    return StarlyCardStakingClaims.getRemainingDailyClaimAmountByAddress(address: address)
}
