package params

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

const (
	// Distribution locking parameteres

	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 5000000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 0
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
)

var (
	// Transaction verification parameters

	// UserVerifyTxn transaction verification parameters for user-created transactions
	UserVerifyTxn = VerifyTxn{
		// BurnFactor can be overriden with `USER_BURN_FACTOR` env var
		BurnFactor: 2,
		// MaxTransactionSize can be overriden with `USER_MAX_TXN_SIZE` env var
		MaxTransactionSize: 32768, // in bytes
		// MaxDropletPrecision can be overriden with `USER_MAX_DECIMALS` env var
		MaxDropletPrecision: 6,
	}
)

// distributionAddresses are addresses that received coins from the genesis address in the first block,
// used to calculate current and max supply and do distribution timelocking
var distributionAddresses = [DistributionAddressesTotal]string{
	"R6aHqKWSQfvpdo2fGSrq4F1RYXkBWR9HHJ",
	"2EYM4WFHe4Dgz6kjAdUkM6Etep7ruz2ia6h",
	"25aGyzypSA3T9K6rgPUv1ouR13efNPtWP5m",
	"ix44h3cojvN6nqGcdpy62X7Rw6Ahnr3Thk",
	"AYV8KEBEAPCg8a59cHgqHMqYHP9nVgQDyW",
	"2Nu5Jv5Wp3RYGJU1EkjWFFHnebxMx1GjfkF",
	"2THDupTBEo7UqB6dsVizkYUvkKq82Qn4gjf",
	"tWZ11Nvor9parjg4FkwxNVcby59WVTw2iL",
	"m2joQiJRZnj3jN6NsoKNxaxzUTijkdRoSR",
	"8yf8PAQqU2cDj8Yzgz3LgBEyDqjvCh2xR7",
	"sgB3n11ZPUYHToju6TWMpUZTUcKvQnoFMJ",
	"2UYPbDBnHUEc67e7qD4eXtQQ6zfU2cyvAvk",
	"wybwGC9rhm8ZssBuzpy5goXrAdE31MPdsj",
	"JbM25o7kY7hqJZt3WGYu9pHZFCpA9TCR6t",
	"2efrft5Lnwjtk7F1p9d7BnPd72zko2hQWNi",
	"Syzmb3MiMoiNVpqFdQ38hWgffHg86D2J4e",
	"2g3GUmTQooLrNHaRDhKtLU8rWLz36Beow7F",
	"D3phtGr9iv6238b3zYXq6VgwrzwvfRzWZQ",
	"gpqsFSuMCZmsjPc6Rtgy1FmLx424tH86My",
	"2EUF3GPEUmfocnUc1w6YPtqXVCy3UZA4rAq",
	"TtAaxB3qGz5zEAhhiGkBY9VPV7cekhvRYS",
	"2fM5gVpi7XaiMPm4i29zddTNkmrKe6TzhVZ",
	"ix3NDKgxfYYANKAb5kbmwBYXPrkAsha7uG",
	"2RkPshpFFrkuaP98GprLtgHFTGvPY5e6wCK",
	"Ak1qCDNudRxZVvcW6YDAdD9jpYNNStAVqm",
	"2eZYSbzBKJ7QCL4kd5LSqV478rJQGb4UNkf",
	"KPfqM6S96WtRLMuSy4XLfVwymVqivdcDoM",
	"5B98bU1nsedGJBdRD5wLtq7Z8t8ZXio8u5",
	"2iZWk5tmBynWxj2PpAFyiZzEws9qSnG3a6n",
	"XUGdPaVnMh7jtzPe3zkrf9FKh5nztFnQU5",
	"hSNgHgewJme8uaHrEuKubHYtYSDckD6hpf",
	"2DeK765jLgnMweYrMp1NaYHfzxumfR1PaQN",
	"orrAssY5V2HuQAbW9K6WktFrGieq2m23pr",
	"4Ebf4PkG9QEnQTm4MVvaZvJV6Y9av3jhgb",
	"7Uf5xJ3GkiEKaLxC2WmJ1t6SeekJeBdJfu",
	"oz4ytDKbCqpgjW3LPc52pW2CaK2gxCcWmL",
	"2ex5Z7TufQ5Z8xv5mXe53fSQRfUr35SSo7Q",
	"WV2ap7ZubTxeDdmEZ1Xo7ufGMkekLWikJu",
	"ckCTV4r1pNuz6j2VBRHhaJN9HsCLY7muLV",
	"MXJx96ZJVSjktgeYZpVK8vn1H3xWP8ooq5",
	"wyQVmno9aBJZmQ99nDSLoYWwp7YDJCWsrH",
	"2cc9wKxCsFNRkoAQDAoHke3ZoyL1mSV14cj",
	"29k9g3F5AYfVaa1joE1PpZjBED6hQXes8Mm",
	"2XPLzz4ZLf1A9ykyTCjW5gEmVjnWa8CuatH",
	"iH7DqqojTgUn2JxmY9hgFp165Nk7wKfan9",
	"RJzzwUs3c9C8Y7NFYzNfFoqiUKeBhBfPki",
	"2W2cGyiCRM4nwmmiGPgMuGaPGeBzEm7VZPn",
	"ALJVNKYL7WGxFBSriiZuwZKWD4b7fbV1od",
	"tBaeg9zE2sgmw5ZQENaPPYd6jfwpVpGTzS",
	"2hdTw5Hk3rsgpZjvk8TyKcCZoRVXU5QVrUt",
	"A1QU6jKq8YgTP79M8fwZNHUZc7hConFKmy",
	"q9RkXoty3X1fuaypDDRUi78rWgJWYJMmpJ",
	"2Xvm6is5cAPA85xnSYXDuAqiRyoXiky5RaD",
	"4CW2CPJEzxhn2PS4JoSLoWGL5QQ7dL2eji",
	"24EG6uTzL7DHNzcwsygYGRR1nfu5kco7AZ1",
	"KghGnWw5fppTrqHSERXZf61yf7GkuQdCnV",
	"2WojewRA3LbpyXTP9ANy8CZqJMgmyNm3MDr",
	"2BsMfywmGV3M2CoDA112Rs7ZBkiMHfy9X11",
	"kK1Q4gPyYfVVMzQtAPRzL8qXMqJ67Y7tKs",
	"28J4mx8xfUtM92DbQ6i2Jmqw5J7dNivfroN",
	"gQvgyG1djgtftoCVrSZmsRxr7okD4LheKw",
	"3iFGBKapAWWzbiGFSr5ScbhrEPm6Esyvia",
	"NFW2akQH2vu7AqkQXxFz2P5vkXTWkSqrSm",
	"2MQJjLnWRp9eHh6MpCwpiUeshhtmri12mci",
	"2QjRQUMyL6iodtHP9zKmxCNYZ7k3jxtk49C",
	"USdfKy7B6oFNoauHWMmoCA7ND9rHqYw2Mf",
	"cA49et9WtptYHf6wA1F8qqVgH3kS5jJ9vK",
	"qaJT9TjcMi46sTKcgwRQU8o5Lw2Ea1gC4N",
	"22pyn5RyhqtTQu4obYjuWYRNNw4i54L8xVr",
	"22dkmukC6iH4FFLBmHne6modJZZQ3MC9BAT",
	"z6CJZfYLvmd41GRVE8HASjRcy5hqbpHZvE",
	"GEBWJ2KpRQDBTCCtvnaAJV2cYurgXS8pta",
	"oS8fbEm82cprmAeineBeDkaKd7QownDZQh",
	"rQpAs1LVQdphyj9ipEAuukAoj9kNpSP8cM",
	"6NSJKsPxmqipGAfFFhUKbkopjrvEESTX3j",
	"cuC68ycVXmD2EBzYFNYQ6akhKGrh3FGjSf",
	"bw4wtYU8toepomrhWP2p8UFYfHBbvEV425",
	"HvgNmDz5jD39Gwmi9VfDY1iYMhZUpZ8GKz",
	"SbApuZAYquWP3Q6iD51BcMBQjuApYEkRVf",
	"2Ugii5yxJgLzC59jV1vF8GK7UBZdvxwobeJ",
	"21N2iJ1qnQRiJWcEqNRxXwfNp8QcmiyhtPy",
	"9TC4RGs6AtFUsbcVWnSoCdoCpSfM66ALAc",
	"oQzn55UWG4iMcY9bTNb27aTnRdfiGHAwbD",
	"2GCdwsRpQhcf8SQcynFrMVDM26Bbj6sgv9M",
	"2NRFe7REtSmaM2qAgZeG45hC8EtVGV2QjeB",
	"25RGnhN7VojHUTvQBJA9nBT5y1qTQGULMzR",
	"26uCBDfF8E2PJU2Dzz2ysgKwv9m4BhodTz9",
	"Wkvima5cF7DDFdmJQqcdq8Syaq9DuAJJRD",
	"286hSoJYxvENFSHwG51ZbmKaochLJyq4ERQ",
	"FEGxF3HPoM2HCWHn82tyeh9o7vEQq5ySGE",
	"h38DxNxGhWGTq9p5tJnN5r4Fwnn85Krrb6",
	"2c1UU8J6Y3kL4cmQh21Tj8wkzidCiZxwdwd",
	"2bJ32KuGmjmwKyAtzWdLFpXNM6t83CCPLq5",
	"2fi8oLC9zfVVGnzzQtu3Y3rffS65Hiz6QHo",
	"TKD93RxFr2Am44TntLiJQus4qcEwTtvEEQ",
	"zMDywYdGEDtTSvWnCyc3qsYHWwj9ogws74",
	"25NbotTka7TwtbXUpSCQD8RMgHKspyDubXJ",
	"2ayCELBERubQWH5QxUr3cTxrYpidvUAzsSw",
	"RMTCwLiYDKEAiJu5ekHL1NQ8UKHi5ozCPg",
	"ejJjiCwp86ykmFr5iTJ8LxQXJ2wJPTYmkm",
}
