//############################### Settings ############################### 

// #### Galaxy and system range to watch ####
galaxy = 4
fromSystem = 1
toSystem = 3

// #### time settings ####
minutes_min = 60
minutes_max = 90

// #### Log settings ####
logTelegram         = true
logDiscord          = true



//############################### Functions ############################### 
//            (don't change if you don't know what you doing

interval = Random(minutes_min*60*1000, minutes_max*60*1000) // 60-90min
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

data = {}

for {
    for system = fromSystem; system <= toSystem; system++ {
        systemInfos, err = GalaxyInfos(galaxy, system)
        if err != nil {
            Print(err)
            continue
        }
        arr = []
        systemInfos.Each(func(planetInfos) {
            arr += planetInfos == nil ? 0 : 1
        })
        key = galaxy+":"+system
        if data[key] != nil && data[key] != arr {
            doLogging(b.GetUniverse()+"-"+b.GetLang()+"|"+b.GetPlayerName()+ " New/Removed planets in "+key, "info")
        }
        data[key] = arr
    }
    Sleep(interval)
}