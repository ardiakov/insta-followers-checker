package main

import (
	"fmt"
	"insta-follower-notifier/internal/app"
)

func main() {
	fmt.Println("Run")

	config := app.InitConfig()
	context := app.InitContext(config)

	//context.MongoDbClient.CreateDocument()

	unfollowers := context.MongoDbClient.GetUnfollowers()

	//fmt.Println(unfollowers)

	/*followers := context.InstagramClient.GetFollowers()
	context.MongoDbClient.Update("followers", "followers", followers)
	followings := context.InstagramClient.GetFollowings()
	context.MongoDbClient.Update("followers", "followings", followings)

	context.MongoDbClient.UpdateUnfollowers("followers", "followings")
	newUnfollowers := context.MongoDbClient.GetUnfollowers()*/

	data := []string{"test", "test"}

	context.MongoDbClient.DiffBetweenUnfollowers(unfollowers, data)

	// ["alex_veselova","ragulina","mi_ana_","sweet_lyla","gsportarena.ru","aliina_shubiina","alenaabramchuk","elizavetta_mm","anastasiyaleon","nasty_keratin95","kseniam.313","sonya.tarasova","julia_frolova16","j.kiiese","mileeenych","prozorova_viktoriya","ma.is.me","kuzminova_life","_ekaterina_eikhler_","anastasiasmirnova17","freestyle_extremepark","_xx_sam_","bogacheva.nastua","_dsots_","vivdishaa","rineeonly","brainpicture","eliettaa","n_ase4ka","asya_rugrats","daniela_yovanovich","nudavaitak","odna_actrisa","nysheva","perfectkeller","art_ferro","zakharooooova","anitamurnieks","vikki112233","hey__adele","meshchersky","a.chesnokovva","cassette_cafe","1khappiness","lasalute_lublino","verona_tmk","velosklad","1serhyo.sneakers","yana_novikova24","xrommova","polina.zhitlova","georgemsk","s.staciia","eddy_lex005","ean777","jekakasatkin","nikstepushin","stefanovic92","chernegaaa","esin.19","raketa_k","romanzaripov","ivan_krasavin_","alina_volkusha","evjenes","dhanishgajjar","don__greko","yannaamee","megakikiboy","leshapooh","dariaalexander","borzmos","marselaaa","markovsky","kamaletdinovam","artemmrz","mickeysanta","demid_rezin","smeshlivaya","kosiner.codes","kaaatey__","akshay.code","thedevlife","xo4y_b","sashatikhomirov","maisto","dimeloper_","stasferbers","kosyachenko_ilya","memeclubru","tanyamoiseenko","136th","martinsoft","naastika","kasatov_movie","denissemenikhin","eddylyfts","chau_codes","pp_yzhin","taha_safari","kinovime","csjack9","yoparveen","boostproductivityjake","peterpandev","shmami","snowboardingss","wylsacom","danemcbeth","eric_davidich","garikkharlamov"]

}
