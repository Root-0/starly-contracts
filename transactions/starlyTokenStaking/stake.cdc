import FungibleToken from "../../contracts/FungibleToken.cdc"
import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import StarlyToken from "../../contracts/StarlyToken.cdc"
import StarlyTokenStaking from "../../contracts/StarlyTokenStaking.cdc"

transaction(amount: UFix64) {
    let vaultRef: &StarlyToken.Vault
    let stakeCollectionRef: &StarlyTokenStaking.Collection

    prepare(acct: AuthAccount) {
        self.vaultRef = acct.borrow<&StarlyToken.Vault>(from: StarlyToken.TokenStoragePath)
            ?? panic("Could not borrow reference to the owner's StarlyToken vault!")

        self.stakeCollectionRef = acct.borrow<&StarlyTokenStaking.Collection>(from: StarlyTokenStaking.CollectionStoragePath)
            ?? panic("Could not borrow reference to the owner's StarlyTokenStaking collection!")
    }

    execute {
        let vault <- self.vaultRef.withdraw(amount: amount) as! @StarlyToken.Vault
        self.stakeCollectionRef.stake(principalVault: <-vault)
    }
}
