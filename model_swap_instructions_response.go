/*
 * Jupiter API v6
 *
 * The core of [jup.ag](https://jup.ag). Easily get a quote and swap through Jupiter API.  ### Rate Limit We update our rate limit from time to time depending on the load of our servers. We recommend running your own instance of the API if you want to have high rate limit, here to learn how to run the [self-hosted API](https://station.jup.ag/docs/apis/self-hosted).  ### API Wrapper - Typescript [@jup-ag/api](https://github.com/jup-ag/jupiter-quote-api-node)  ### Data types - Public keys are base58 encoded strings - raw data such as Vec<u8\\> are base64 encoded strings
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package jupiter

type SwapInstructionsResponse struct {
	TokenLedgerInstruction *Instruction `json:"tokenLedgerInstruction,omitempty"`
	// The necessary instructions to setup the compute budget.
	ComputeBudgetInstructions []Instruction `json:"computeBudgetInstructions"`
	// Setup missing ATA for the users.
	SetupInstructions  []Instruction `json:"setupInstructions"`
	SwapInstruction    *Instruction  `json:"swapInstruction"`
	CleanupInstruction *Instruction  `json:"cleanupInstruction,omitempty"`
	// The lookup table addresses that you can use if you are using versioned transaction.
	AddressLookupTableAddresses []string `json:"addressLookupTableAddresses"`
}
