package main

import (
	"encoding/json"
	"fmt"
	"log"

	mbconnect "github.com/nsvirk/gomoneybotslib/pkg/connect"
)

type APITest struct {
	cfg      Config
	mbClient *mbconnect.Client
}

func New(cfg *Config) *APITest {

	mbClient := mbconnect.New(cfg.KiteUserId)

	return &APITest{
		cfg:      *cfg,
		mbClient: mbClient,
	}
}

func (t *APITest) TestAPIEndpoints() {
	// session
	t.TestSessionEndpoints()
	// instruments
	t.TestInstrumentsQueryEndpoints()
	t.TestInstrumentsFNOEndpoints()
	t.TestInstrumentsOCEndpoints()
	t.TestIndicesEndpoints()
}

func (t *APITest) TestSessionEndpoints() {
	t.GenerateUserSession()
	t.GenerateTotpValue()
	// Uncomment the following line to test session deletion
	// t.DeleteUserSession()
	printSectionFooter()
}

func (t *APITest) TestInstrumentsQueryEndpoints() {
	t.InstrumentsBySymbols()
	t.InstrumentsByTokens()
	t.InstrumentsByQuery()
	t.SymbolsByQuery()
	t.TokensByQuery()
	printSectionFooter()
}

func (t *APITest) TestInstrumentsOCEndpoints() {
	t.OptionchainInstruments()
	t.OptionchainSymbols()
	t.OptionchainTokens()
	printSectionFooter()
}

func (t *APITest) TestInstrumentsFNOEndpoints() {
	t.FNOSegmentExpiries()
	t.FNOSegmentNames()
	printSectionFooter()
}

func (t *APITest) TestIndicesEndpoints() {
	t.IndicesAll()
	t.IndicesByExchange()
	t.IndexNames()
	t.IndexTokens()
	t.IndexSymbols()
	t.IndexInstruments()
	printSectionFooter()
}

func main() {
	cfg, err := LoadEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	// Print the configuration
	fmt.Println(cfg.String())

	// Initialize the API test
	apiTest := New(cfg)
	// apiTest.mbClient.SetDebug(true)
	// apiTest.mbClient.SetBaseURI(apiTest.cfg.APIDevUrl)
	apiTest.TestAPIEndpoints()

}

// --------------------------------------------------------
// Helper functions
// --------------------------------------------------------

func PrettyPrint(title string, data interface{}) {
	fmt.Printf("\n===================================\n")
	fmt.Printf("%s:\n", title)
	fmt.Printf(" type: %T\n", data)
	fmt.Printf("-----------------------------------\n")

	count := 0
	maxCount := 2
	outerCount := 0
	maxOuterCount := 1

	// check if no data
	if data == nil {
		fmt.Printf("  %s\n", "No data")
		return
	}

	switch v := data.(type) {
	case *mbconnect.UserSession:
		prettyPrintJSON(v)

	case string:
		fmt.Printf("  %s\n", v)

	case uint32:
		fmt.Printf("  %d\n", v)

	case []string:
		for i, s := range v {
			if count > maxCount {
				break
			}
			fmt.Printf("  [%d]: %s\n", i, s)
			count++
		}

	case []uint32:
		for i, n := range v {
			if count > maxCount {
				break
			}
			fmt.Printf("  [%d]: %d\n", i, n)
			count++
		}

	case map[string]string:
		for k, v := range v {
			if count > maxCount {
				break
			}
			fmt.Printf("  %s: %s\n", k, v)
			count++
		}

	case map[string]uint32:
		for k, v := range v {
			if count > maxCount {
				break
			}
			fmt.Printf("  %s: %d\n", k, v)
			count++
		}

	case map[uint32]string:
		for k, v := range v {
			if count > maxCount {
				break
			}
			fmt.Printf("  %d: %s\n", k, v)
			count++
		}

	case map[string][]string:
		for k, v := range v {
			if outerCount > maxOuterCount {
				break
			}
			fmt.Printf("  %s:\n", k)
			for i, s := range v {
				if count > maxCount {
					break
				}
				fmt.Printf("    [%d]: %s\n", i, s)
				count++
			}
			outerCount++
		}

	case map[string]mbconnect.Instrument:
		for _, v := range v {
			if count > maxCount {
				break
			}
			prettyPrintJSON(v)
			count++
		}

	case map[uint32]mbconnect.Instrument:
		for _, v := range v {
			if count > maxCount {
				break
			}
			prettyPrintJSON(v)
			count++
		}

	case map[string][]mbconnect.Instrument:
		for k, v := range v {
			if outerCount > maxOuterCount {
				break
			}
			fmt.Printf("  %s:\n", k)
			for _, model := range v {
				if count > maxCount {
					break
				}
				prettyPrintJSON(model)
				count++
			}
			outerCount++
		}

	case []mbconnect.Index:
		for _, model := range v {
			if count > maxCount {
				break
			}
			prettyPrintJSON(model)
			count++
		}

	case map[string]mbconnect.Index:
		for k, v := range v {
			if outerCount > maxOuterCount {
				break
			}
			fmt.Printf("  %s:\n", k)
			prettyPrintJSON(v)
			outerCount++
		}

	case map[string][]mbconnect.Index:
		for k, v := range v {
			if outerCount > maxOuterCount {
				break
			}
			fmt.Printf("  %s:\n", k)
			for _, model := range v {
				if count > maxCount {
					break
				}
				prettyPrintJSON(model)
				count++
			}
			outerCount++
		}

	case map[uint32][]mbconnect.Index:
		for k, v := range v {
			if outerCount > maxOuterCount {
				break
			}
			fmt.Printf("  %d:\n", k)
			for _, model := range v {
				if count > maxCount {
					break
				}
				prettyPrintJSON(model)
				count++
			}
			outerCount++
		}

	default:
		// fmt.Printf("Unsupported type: %T\n", v)
		log.Fatalf("Unsupported type: %T\n", v)
	}
}

func printSectionFooter() {
	fmt.Printf("-----------------------------------\n")
}

func prettyPrintJSON(v interface{}) {
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate json: %v", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
