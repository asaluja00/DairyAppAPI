# DairyAppAPI
Develop a HTTP JSON Go API for Diary application which helps the
user to manage daily activities. 
It allows users to list down each event of the current day and review past
days entries. 


This API will have following routes:
● /login : This route will allow a user to login with unique secret code and if user is found, it will
return details of the user

● /register : This is for creating a user with name and email address and will return complete details
of newly created user(secret code, user id)

● /addEntry : This route will allow a user to add a log to the present day.

● /updateEntry : This route will allow a user to update an existing log for the present day only.

● /deleteEntry : This route will allow users to delete any log of a particular day.

● /showEntry : This route will return all the logs of a particular day.
