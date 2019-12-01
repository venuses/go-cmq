package go_cmq

type TopicAPI interface {
	CreateTopic()
	SetTopicAttributes()
	ListTopic()
	GetTopicAttributes()
	DeleteTopic()
}
type TopicMessageAPI interface {
	PublishMessage()
	BatchPublishMessage()
}
type TopicSubscriptionAPI interface {
	ClearSubscriptionFilterTags()
	Subscribe()
	ListSubscriptionByTopic()
	SetSubscriptionAttributes()
	Unsubscribe()
	GetSubscriptionAttributes()
}
