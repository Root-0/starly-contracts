import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import StarlyCard from "../../contracts/StarlyCard.cdc"

// This script returns the size of an account's StarlyCard collection.

pub fun main(address: Address): Int {
    let account = getAccount(address)

    let collectionRef = account.getCapability(StarlyCard.CollectionPublicPath)!
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")

    return collectionRef.getIDs().length
}
