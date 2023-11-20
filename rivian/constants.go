package rivian

const (
	basepath         string = "https://rivian.com/api/gql"
	gateway          string = basepath + "/gateway/graphql"
	charging         string = basepath + "/chrg/user/graphql"
	websocket        string = "wss://api.rivian.com/gql-consumer-subscriptions/graphql"
	apolloClientName string = "com.rivian.ios.consumer-apollo-ios"
)
