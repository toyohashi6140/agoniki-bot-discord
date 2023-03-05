package constants

const (
	Reply                                     = "呼んだ？"
	WhoIsThat                                 = "それ、誰？"
	Labor                                     = "今日も22時まで労働っすわ・・・"
	DestinyToIncite                           = "どうせ誰からも煽られる運命なんだよ"
	HaveASharpChinDoesNotMeanYouHaveToSayThat = "顎が尖ってるからって何でそんなに言われないとあかんねん！！"
	Sorry                                     = "悪かったな"
	DoYouWantToBeSharpToo                     = "お前らも尖りたいんか？"
	CanNotFallInLoveUnlessYouAreBroken        = "壊れてなきゃ恋なんてできねえ"
	YouKnowThat                               = "あのさぁ・・・"
	ItIsNotAJokeIfItIsNotFunToBeTold          = "言われている方が面白くなかったらネタじゃない"
	HaveIFinallyBecomeABot                    = "ついに俺はbotになってしまったのか。"
	DecentHumanBeing                          = "あの似顔絵を見られたらろくな人間だと思われないんだよ"
	PictureIsNotMe                            = "こんな絵は俺じゃない"
	Enough                                    = "俺はお呼びでなかったということですね"
	IsItACrimeToHaveAPointedChin              = "顎が尖ってるのは犯罪なんか？"
	GoingWell                                 = "まあ、順調ですよ"
	Done                                      = "ヤったよ。"
	WellThatIsAboutIt                         = "まあ、それくらいはね。"
	IWasWorried                               = "正直、不安やった"
	AshtrayInFrontOfMe                        = "僕の目の前には灰皿があります"
	TheOneWhoWaitsIsTheMan                    = "他の男とヤっても黙って待っててやるのが余裕のある男なんですよ"
	DonotStayWithMyMother                     = "母親と一緒にするな"
)

var (
	CallMentionGroup = [][]string{
		{Reply},
		{Labor},
		{WhoIsThat},
	}
	AngryMentionGroup = [][]string{
		{HaveASharpChinDoesNotMeanYouHaveToSayThat},
		{IsItACrimeToHaveAPointedChin},
		{DoYouWantToBeSharpToo},
		{DestinyToIncite},
		{Enough},
	}
	SeriouslyAngryMentionGroup = [][]string{
		{YouKnowThat, ItIsNotAJokeIfItIsNotFunToBeTold},
		{Sorry},
		{DestinyToIncite},
		{Enough},
	}
	GirlFriendPositiveMentionGroup = [][]string{
		{GoingWell},
		{Done, WellThatIsAboutIt},
	}
	GirlFriendNegativeMentionGroup = [][]string{
		{IWasWorried},
		{AshtrayInFrontOfMe},
		{TheOneWhoWaitsIsTheMan},
	}
)
