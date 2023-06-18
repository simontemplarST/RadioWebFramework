# RadioWeb Framework
` A self reliant framework that puts controls back into the hands of Ham Radio Operators. `

---

We are dependent on two or three sources for our contact confirmation, On-Air status, and web logging. 
This basic framework and starter project kind of exemplify putting the control back in the operators hands if they so choose. 

---

## What is still needed?
- We need example databases from any and all Ham Radio logging programs. 
    - Even if we say we already have them, send them in anyway. Never hurts to have multiple things to run against.
- Development. If you are a software developer and like the idea. Follow the guidelines, and have fun! That's what the hobby is about.
- Logging software integration
    - If you are a developer of a logging application, and want to have this in your logger, By all means, implement it! 

## Things that haven't been set in stone yet. 
Log.xml
- Do we use a common domain structure? like `callsign.radio`? So all we have to do is put the callsign in a var statement then slap `.radio/log.xml` on the tail?
- Do we offload the confirming to the "Confrim" server, and put metadata in the generated files for scrapers to index? 
    - Would these servers communicate with each other and vote on the most accurate data?
- Do we say find a way to implement ActivityPub?
- Do we make everyone a part of a "WebRing" and source the data that way?

Status.xml
- I am wondering if current frequency, mode, and a freeform message should be added into the schema as well?

### Guidelines


Three files have to be generated. 
   - log.`filename here` (the operators log in it's entirety) -in the demo case, it's using a markdown file to generate a hugo post- 
   - log.xml (RSS feed for confirming contacts)
   - status.xml (To show on or off air status)

What data creates those files? 
   - The QSL or SQL like database your logger saves to. 

Where is the data saved?
- On the machine you are currently using for logging by default. You can also set it up to run on an external device, and then access the database over the network. 

How do I get this uploaded?
- You can use an FTP client, Syncthing (or similar), or if your web host has a RESTful API to upload files (Like mine does), you can use that as well.

What if I'm not too technical?
- That's the beauty of Open Source. it promotes creation of tool sets that interact with the framework. In the near future this can include full on GUI applicaions. 


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
        <Call></call>
        <QSO_Date></QSO_Date>
        <Time_On></Time_On>
        <Time_Off></Time_Off>
        <Band></Band>
        <Cont></Cont>
        <Country></Country>
        <DXCC></DXCC>
        <My_Cnty></My_Cnty>
        <CQz></CQz>
        <Freq></Freq>
        <My_Gridsquare></My_Gridsquare>
        <MY_State></MY_State>
        <ITUz></ITUz>
        <Mode></Mode>
        <Name></Name>
        <OPERATOR></OPERATOR>
        <RST_Sent></RST_Sent>
        <RST_Rcvd></RST_Rcvd>
        <State></State>
    </entry>
</log>

```
You can take the log.`filename here` And create it with what ever you'd like, as long as it's searachable, and public, showing your digital log in it's entirety. You can style it to your taste, or (eventuially) there will be templates you can use for all of the major platforms. 

### Config File generator
This generator is VERY basic and ugly. 
- It will have the major logging software listed in the "Logging Software" dropdown. 
    - If you would like more added please send me an email with your logging software's database backup or send me a snippet of the table structure
- You can select where the database is located with the "Open Location" button.
- Soon, there will be check boxes to select which additional columns you want displayed in your public web log aside from the required ones by the standard. 