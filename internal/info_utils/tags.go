package info_utils

import "math/rand"


var tags = []string{"#nnscs #photooftheday #fashion #beautiful #happy #tbt #like4like #followme #picoftheday #me #selfie #instadaily #friends #fun #style #smile #food #instalike #likeforlike #travel #fitness #tagsforlikes #follow4follow #amazing #instamood #instagram #photography #watchpeopledieinside #unexpectedvideos #perfectlycutscreams",
					"#nnscs #watchpeopledieinside #perfectlycutscreams #meme #tiktokmemes #memes #unexpected #memesdaily #humour#trending #photooftheday #fashion #like4like #followme #picoftheday #follow #me #selfie #summer #art #instadaily #instalike #likeforlike #travel #fitness #igers #follow4follow #life #beauty #amazing #instagram"}

func GetTag() string {
	return tags[rand.Intn(len(tags))]
}