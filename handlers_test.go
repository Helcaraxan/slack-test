package slacktest

import (
	"testing"

	slack "github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestPostMessageHandler(t *testing.T) {}
func TestServerListChannels(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	channels, err := client.GetChannels(true)
	assert.NoError(t, err)
	assert.Len(t, channels, 2)
	assert.Equal(t, "C024BE91L", channels[0].ID)
	assert.Equal(t, "C024BE92L", channels[1].ID)
	for _, channel := range channels {
		assert.Equal(t, "W012A3CDE", channel.Creator)
	}
}

func TestUserInfoHandler(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	user, err := client.GetUserInfo("123456")
	assert.NoError(t, err)
	assert.Equal(t, "W012A3CDE", user.ID)
	assert.Equal(t, "spengler", user.Name)
	assert.True(t, user.IsAdmin)
}

func TestBotInfoHandler(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	bot, err := client.GetBotInfo(s.BotID)
	assert.NoError(t, err)
	assert.Equal(t, s.BotID, bot.ID)
	assert.Equal(t, s.BotName, bot.Name)
	assert.False(t, bot.Deleted)
}

func TestListGroupsHandler(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	groups, err := client.GetGroups(true)
	assert.NoError(t, err)
	if !assert.Len(t, groups, 1, "should have one group") {
		t.FailNow()
	}
	mygroup := groups[0]
	assert.Equal(t, "G024BE91L", mygroup.ID, "id should match")
	assert.Equal(t, "secretplans", mygroup.Name, "name should match")
	assert.True(t, mygroup.IsGroup, "should be a group")
}

func TestListChannelsHandler(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	channels, err := client.GetChannels(true)
	assert.NoError(t, err)
	if !assert.Len(t, channels, 2, "should have two channels") {
		t.FailNow()
	}
	generalChan := channels[0]
	otherChan := channels[1]
	assert.Equal(t, "C024BE91L", generalChan.ID, "id should match")
	assert.Equal(t, "general", generalChan.Name, "name should match")
	assert.Equal(t, "Fun times", generalChan.Topic.Value)
	assert.True(t, generalChan.IsMember, "should be in channel")
	assert.Equal(t, "C024BE92L", otherChan.ID, "id should match")
	assert.Equal(t, "bot-playground", otherChan.Name, "name should match")
	assert.Equal(t, "Fun times", otherChan.Topic.Value)
	assert.True(t, otherChan.IsMember, "should be in channel")
}