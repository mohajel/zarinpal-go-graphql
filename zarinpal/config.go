package zarinpal

const (
	// BaseURL is the base URL for the Zarinpal API
	BaseURL = "https://payment.zarinpal.com"
	// GraphQLURL is the GraphQL URL for the Zarinpal API
	GraphQLURL = "https://next.zarinpal.com/api/v4/graphql/"
	// SandboxBaseURL is the base URL for the Zarinpal API in sandbox mode
	SandboxBaseURL = "https://sandbox.zarinpal.com"
	// TODO: see if sandbox exists or not
	// SandboxGraphQLURL is the GraphQL URL for the Zarinpal API in sandbox mode
	SandboxGraphQLURL = "https://sandbox.zarinpal.com/api/v4/graphql/"
)

type Config struct {
	Sandbox     bool
	MerchantID  string
	AccessToken string
}
