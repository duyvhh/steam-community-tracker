# steam-community-tracker
To track a certain item's price in Steam Community market

## Requirements

As this project is based on Google App Engine (GAE), you must have a Google Developer account to be able to run. 

## Usage

1. Clone the repository
2. Modify the URL in tracker.go to point to your favorite item on Steam community market you wish to track its price
3. Modify the SENDER const with your registered developer's email address
4. Modify the receiver's address at line 76
5. Modify your threshold, below which an email notification will be sent to notify you

Follow this guide to upload your application:
https://cloud.google.com/appengine/docs/go/gettingstarted/uploading

By default, the tracker will run every 15 minutes. You can modify this value in cron.yaml file, following Google Cron Format

https://cloud.google.com/appengine/docs/python/config/cron#Python_app_yaml_The_schedule_format

## Known limitations

1. Only USD currency is supported
2. Can only track 1 item at the moment. A workaround is to deploy multiple GAE applications.

## Credits
This project is using the following library to extract price information from steam community market

- GoQuery (github.com/PuerkitoBio/goquery)
