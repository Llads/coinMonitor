# coinMonitor
A console application with alert integration for monitoring the prices of crypto using the coingecko api.

Thank you to: https://github.com/martinlindhe for his goLang notification library // https://github.com/gookit for their awesome color library // https://github.com/sneakers for help with my issues with json and requests.

alert.png can be changed to whatever you feel like works best for you.

Program is currently setup to alert you to dips, this can be changed in around one line.

Currently will only run on windows due to hooking the cmd for changing console title, plan to either remove this which will allow it to run on mac devices, however you can just remove the setTitle() function and remove the call to it at the top.
