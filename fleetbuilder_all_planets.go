/* DESCRIPTION
 * Automatically build ships on all planets periodically
 * Discord and Telegram logs if prefer
 * min and max time is adjustable
 */



//############################### Settings ############################### 

// #### time settings ####
minutes_min = 60
minutes_max = 90

// #### Log settings ####
logTelegram         = true
logDiscord          = true

toBuild = {
	DESTROYER: 20,
	
/* example 	
   LIGHTFIGHTER: 6,
	HEAVYFIGHTER: 4,
*/

}



//############################### Functions ############################### 
//            (don't change if you don't know what you doing

planetsToBuildOn = GetPlanets()

func millisecondsToTime(milliseconds) {
    minutes = (milliseconds / 1000) / 60;
    return minutes;
}

func doLogging(msg, loglevel) {
    switch loglevel {
        case "error":
            LogError(msg)
        case "warn":
            LogWarn(msg)
        case "info":
            LogInfo(msg)
    }
    if logTelegram {
        SendTelegram(TELEGRAM_CHAT_ID, msg)
    }   
    if logDiscord {
        SendDiscord(DISCORD_WEBHOOK, msg)
    }
}

for {
    for planet in planetsToBuildOn {
        // Get planet celestial
        planetCelestial, _ = GetCachedCelestial(planet.Coordinate)
        // Iterate over ships to build
        for unitID, nbrToBuild in toBuild {
            // See how many ships we can build
            unitPrice = GetPrice(unitID, 1)
            resources, _ = planetCelestial.GetResources()
            canBuild = resources.Div(unitPrice)
            if canBuild > 0 {
                if canBuild > nbrToBuild {
                    canBuild = nbrToBuild
                }
                doLogging("Building: "+canBuild+"/"+nbrToBuild+" "+ID2Str(unitID)+" on "+planet, "info")
                Build(planetCelestial.GetID(), unitID, canBuild)
            } else {
                doLogging("Not enough resources on "+planet+" to build at least 1 of "+nbrToBuild+" "+ID2Str(unitID), "warn")
            }
        }
    }
    interval = Random(minutes_min*60*1000, minutes_max*60*1000) // 60-90min
    timeleft = (interval / 1000)
    doLogging("All ships built. Resume in " + ShortDur(timeleft), "info")
    Sleep(interval)
}