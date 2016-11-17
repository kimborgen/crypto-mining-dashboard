Cryptocurrency Mining Dashboard

Tired of teamviewing to every mining rig, one at a time, to check on the status?
Tired of not having everything in one secure place?
Tired of not having an easy way to switch between algorithms fast?
Welcome to the ultimate dashboard! 

This project is in pre-alpha (not finished/unusable), and is written as a learning excercise to learn go and react.  

This project consists of three subprojects
/frontend  #A react app for the frontend, gets live data from the webserver
/backend/src/webserver  #A json api server written in go, recives data from several miningservers and sends it to the frontend
/backend/src/miningserver  #A json api server written in go, gets system and mining infoand sends the data to the webserver


