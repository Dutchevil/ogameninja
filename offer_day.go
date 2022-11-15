/* DESCRIPTION
 * Automatically buys the offer of the day at 13:10 + random 10 - 50 delay
 * Discord and Telegram logs if prefer
 */

//############################### Settings ############################### 


// #### Log settings ####
logTelegram         = true
logDiscord          = true


//############################### Functions ############################### 
//            (don't change if you don't know what you doing

planetsToBuildOn = GetPlanets()
b = GetBotByID(BotID)

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



func callback() {
    SleepRandMin(10, 50) // Pause execution for a random duration between 10min to 50min
    err = BuyOfferOfTheDay()
    Print("Bought offer of the day", err)
    doLogging("<"+b.GetUniverse()+"-"+b.GetLang()+"|"+b.GetPlayerName()+"> Bought offer of the day.", "info");
    
}
CronExec("@13h10", callback) // Execute callback every day at 13:10 + random 10 - 50 delay
<-OnQuitCh