package zarinpal

type zarinpal struct {
	merchantID  string
	accessToken string
	baseUrl     string
	graphQlURL  string
}

func New(merchantID, accessToken string, sandbox bool) *zarinpal {
	var baseUrl, graphQlURL string

	if sandbox {
		baseUrl = SandboxBaseURL
		graphQlURL = SandboxGraphQLURL
	} else {
		baseUrl = BaseURL
		graphQlURL = GraphQLURL
	}

	return &zarinpal{
		merchantID:  merchantID,
		accessToken: accessToken,
		baseUrl:     baseUrl,
		graphQlURL:  graphQlURL,
	}
}


