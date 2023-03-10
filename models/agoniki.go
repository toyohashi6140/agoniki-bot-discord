package models

import (
	"strings"

	"github.com/toyohashi6140/discord-bot/config"
	"github.com/toyohashi6140/discord-bot/constants"
)

const (
	Angry          = 1
	SeriouslyAngry = 2
	Positive       = 1
	Negative       = -1
)

type Agoniki struct {
	called         bool
	Angry          bool
	SeriouslyAngry bool
	emotion        int
}

func (a *Agoniki) Call(msg string) bool {
	for _, callText := range constants.Calls {
		if strings.Contains(msg, callText) {
			a.called = true
		}
	}
	return a.called
}

// Anger 怒りフラグ設定
func (a *Agoniki) Anger(msg string) {
	if !a.called {
		angerCount := 0
		for _, swearing := range config.DiscordEnv.Swearings {
			if strings.Contains(msg, swearing) {
				angerCount++
			}
		}
		for _, swearing := range config.DiscordEnv.Swearings2 {
			if strings.Contains(msg, swearing) {
				angerCount += 2
			}
		}
		if angerCount == Angry {
			a.Angry = true
		} else if angerCount >= SeriouslyAngry {
			a.SeriouslyAngry = true
		}
	}
}

// NegativeOrPositive -1（ネガティブ）, 0（ニュートラル）, 1（ポジティブ）で判定
func (a *Agoniki) NegativeOrPositive(msg string) int {
	hasPositiveWord := false
	for _, girlfriend := range constants.Girlfriends {
		if strings.Contains(msg, girlfriend) {
			hasPositiveWord = true
			a.emotion = 1
		}
	}
	if hasPositiveWord {
		// ネガティブワードを同時に拾うとネガティブになる
		for _, cheating := range constants.Cheatings {
			if strings.Contains(msg, cheating) {
				a.emotion = -1
			}
		}
	}
	return a.emotion
}
