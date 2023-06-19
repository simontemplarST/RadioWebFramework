# RadioWeb Framework
` A self reliant framework that puts controls back into the hands of Ham Radio Operators. `

---

We are dependent on two or three sources for our contact confirmation, On-Air status, and web logging. 
This is more of an ideological framework than an actual software framework. We're not too awesome at development. But we like to think we have some cool ideas. 

---

## What is still needed?
- We need example databases from any and all Ham Radio logging programs. 
    - Even if we say we already have them, send them in any ways. Never hurts to have multiple things to run against.
- Development. If you are a software developer and like the idea. Follow the guidelines, and have fun! That's what the hobby is about.
- Logging software integration
    - If you are a developer of a logging application, and want to have this in your logger, By all means, implement it! 

## Things that haven't been set in stone yet. 
Log.xml

- Do we use a common domain structure? like `callsign.radio`? So all we have to do is put the callsign in a var statement then slap `.radio/log.xml` on the tail?
- Do we offload the confirming to the [Confrim](https://gist.github.com/simontemplarST/07a58a3fefc167178b58d8eebd8fb0c8) server, and put metadata in the generated files for scrapers to index? 
    - Would these servers communicate with each other and vote on the most accurate data?
- Do we say find a way to implement ActivityPub?
- Do we make everyone a part of a "WebRing" and source the data that way?

Status.xml
- I am wondering if current frequency, mode, and a freeform message should be added into the schema as well?

### Guidelines


Three files have to be generated. 
   - log.`filename here` (the operators log in its entirety) -in the demo case, it's using a markdown file to generate a hugo post- 
   - log.xml (RSS feed for confirming contacts)
   - status.xml (To show on or off-air status)

What data creates those files? 
   - The SQL or SQL like database your logger saves to. 

Where is the data saved?
- On the machine you are currently using for logging by default. You can also set it up to run on an external device, and then access the database over the network. 

How do I get this uploaded?
- You can use an FTP client, Syncthing (or similar), or if your web host has a RESTful API to upload files (Like mine does), you can use that as well.

What if I'm not too technical?
- That's the beauty of Open Source. it promotes creation of tool sets that interact with the framework. In the near future this can include full on GUI applications. 


### Schemas
status.xml
```xml
<status>
    <updated>$date-time</updated>
    <airstatus>$On/Off</airstatus>
    <mode>$SSB</mode>
    <frequency>$Frequency</frequency>
    <msg>Freeform message goes here</msg>
</status>
```
log.xml
```xml
<log>
    <entry>
        <Call>OtherStation Call</call>
        <QSO_Date>UTC Date</QSO_Date>
        <Time_On>UTC Time QSO Start</Time_On>
        <Time_Off>UTC Time QSO End</Time_Off>
        <Band>Band</Band>
        <Cont>OtherStation Continent</Cont>
        <Country>OtherStation Country</Country>
        <DXCC>OtherStation Entity</DXCC>
        <My_Cnty>My Country</My_Cnty>
        <CQz>CQ Zone</CQz>
        <Freq>frequency</Freq>
        <My_Gridsquare>My Maidenhead Grid</My_Gridsquare>
        <MY_State>My State</MY_State>
        <ITUz>OtherStation ITU Zone</ITUz>
        <Mode>Operating Mode</Mode>
        <Name>OtherStation Operator Name</Name>
        <OPERATOR>My Call Sign</OPERATOR>
        <RST_Sent>RST Sent to OtherStation</RST_Sent>
        <RST_Rcvd>RST from OtherStation</RST_Rcvd>
        <State>OtherStation State</State>
    </entry>
</log>

```

You can take the log.`filename here` And create it with what ever you'd like, as long as it's searchable, and public, showing your digital log in its entirety. You can style it to your taste, or (eventually) there will be templates you can use for all of the major platforms. 

### Config File generator
This generator is VERY basic and ugly. 
- It will have the major logging software listed in the "Logging Software" dropdown. 
    - If you would like more added please email me with your logging software's database backup or send me a snippet of the table structure. Thank you!
- You can select where the database is located with the "Open Location" button.
- Soon, there will be check boxes to select which additional columns you want displayed in your public web log aside from the required ones by the standard. 

### How to install
Pre-reqs: Go 1.20 or newer

1. Clone the repo to your local machine. IE: `git clone git@github.com:simontemplarST/RadioWebFramework.git`
2. Navigate to the repo in your terminal. IE: `cd RadioWebFramework`
3. Execute `go get`