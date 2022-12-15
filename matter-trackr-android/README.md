# Matter Trackr
An android integration
## Purpose

To create an android application for Matter supported devices to send data to Trackr for visualization.

[Link to the demo video](https://youtu.be/8ifrk2LAsCk)
## Table of Contents

- [Matter Trackr](#matter-trackr)
  - [Purpose](#purpose)
  - [Table of Contents](#table-of-contents)
    - [prerequisites](#prerequisites)
      - [To setup up a Matter device](#to-setup-up-a-matter-device)
      - [To set up Matter](#to-set-up-matter)
    - [Running Trackr](#running-trackr)
    - [Running the Application](#running-the-application)
    - [Using the Application](#using-the-application)
    - [More resources](#more-resources)
    - [Authors](#authors)

### prerequisites
- The android device prerequisites can be found [here](https://developers.home.google.com/samples/matter-app)
- The Matter SDK can only be installed on Mac and Linux machines so the following steps are roughly equivalent to what they might be on other machines. 
#### To setup up a Matter device
Using the ESP-32 M5 stack device the using a Mac
- First install the Mac pre-reqs [here](https://github.com/project-chip/connectedhomeip/blob/master/docs/guides/BUILDING.md)
- If any packages are missing they will have to installed using brew or pip
- Be sure to install gn and ninja if they are not already installed.

#### To set up Matter
To build and flash the all-clusters app onto the ESP-32 follow these guides:
- https://github.com/project-chip/connectedhomeip/blob/master/docs/guides/esp32/setup_idf_chip.md
- https://github.com/project-chip/connectedhomeip/blob/master/docs/guides/esp32/build_app_and_commission.md
- https://github.com/project-chip/connectedhomeip/blob/master/docs/guides/esp32/build_app_and_commission.md#build-flash-and-monitor-an-example
- If using the ESP32- M5 device the device type  is ESP32, and use the M5 setup config.
- After the ESP32 device is flashed the QR code can be accessed in the device or found in the logs.

### Running Trackr

1. Run the backend

```
cd backend
make run
```

2. Run the frontend
   
```
cd frontend
npm install
npm start
```

### Running the Application

1. Open the *Matter Trackr* application, select a target device and select the play button.
2. This will build and install the application to the target device.
3. This application was modified from the Google Sample Home App, using androids matter SDK, the app may have changed so to see the updated changes or to learn more about how the app works check out [google’s code lab](https://developers.home.google.com/codelabs/matter-sample-app#1)

### Using the Application

1. Upon opening the application a “trackr” login screen will open up. 
2. The home app doesn’t yet have register functionality so you need to register in the web-app first. 
3. After logging in you can connect a device using the QR code scanner. 
4. After a device is connected it will be viewable in the list of devices. 
5. To connect a device to trackr click on the device and press the “Connect to Trackr” button. 
6. This button creates a project template in trackr, and sends the current temperature value to trackr.
7. In future iterations this should create a thread that checks for temperature updates on the device or creates a subscription. 
8. It should also save all the returned metadata so that multiple Value writes can happen.

### More resources

- To learn more about Matter, please refer to this [website](https://developers.home.google.com/matter/develop)
- You can learn more about Trackr [here](https://github.com/COMP4560/trackr)
- To learn more about [android studio](https://developer.android.com/guide/start/)
- For Kotlin support refer to [this](https://kotlinlang.org/docs/basic-syntax.html)

### Authors

- [Aidan](https://github.com/HodgertA)
- [Kunal](https://github.com/KunalRajpal)
- [Kam](https://github.com/kamarabbas99)
- [Akshay](https://github.com/Akshaysharma21)
