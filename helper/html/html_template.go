package html

var (
	BoardingPassHTML = `
	<html> 
    <head></head>
    <body>
        <style type="text/css">
            @import url("https://fonts.googleapis.com/css?family=Poppins&display=swap");

            body {
                font-family: "Poppins", sans-serif;
                background-color: #ffffff;
                margin: 0;
                padding: 0;
            }

            .container {
                margin: 0 auto;
                position: relative;
                width: auto;
            }

            @media screen and (min-width: 1024px) {
                .container {
                    max-width: 960px;
                }
            }

            /* CARD HEADER SECTION */
            .header {
                width: 100%;
                margin-bottom: 0;
                height: 225px;
                background-color: #04847D;
                color: white;
            }

            .header > h2 {
                font-size: 24px;
                padding-left: 7%;
                padding-top: 2%;
            }

            .card-bp {
                float: left;
                margin: -17% 6% 0 6%;
                background-color: #ffffff;
                width: 88%;
                min-height: 200px;
                max-height: 270px;
                border-radius: 15px;
                border: 2px solid #c4c4c4;
            }

            .card-bp-header {
                min-height: 30px;
                max-height: 70px;
                padding: 10px 3%;
                border-bottom: 2.5px solid #c4c4c4;
            }

            .icon-lion {
                vertical-align: middle;
                height: 26px;
            }

            .icon-barcode {
                vertical-align: middle;
                height: 39.5px;
                margin: 0 0 0 70%;
            }

            .column-barcode {
                float: left;
                min-height: 20%;
                max-height: 30%;
                width: 17%;
            }

            .icon-body-barcode {
                margin: 0px 24px;
                /* min-width: 50px; */
                /* max-width: 70px; */
                width: 70px;
                height: 190px;
            }

            #column-information-boarding {
                width: 74%;
                display: inline;
                float: left;
            }

            .card-bp-body {
                min-width: 90%;
                max-width: 100%;
                float: left;
                margin: 1% 4% 1% 0;
            }

            .content-inside-information {
                width: 83%;
                display: inline-block;
            }

            #content-info-header {
                width: 115%;
                min-height: 50px;
                max-height: 70px;
                padding: 1% 1% 2% 0;
                border-bottom: 2.5px solid #c4c4c4;
            }
            
            .content-info-header-field {
                float: left;
                width: 30%;
            }

            .second-content-inside-information {
                display: flex;
                flex-direction: row;
            }

            .column-inside-information {
                float: left;
                width: 33.3%;
                /* min-width: 33.3%; */
                /* max-width: 38%; */
            }

            .title-main-boarding {
                color: #72a092;
                font-weight: 600;
                margin: 0px;
                font-size: 15px;
                font-family: "Poppins", sans-serif;
            }

            .second-title-main-boarding {
                color: #72a092;
                font-weight: 600;
                margin: 0px;
                font-size: 10px;
                font-family: "Poppins", sans-serif;
            }

            .subtitle-main-boarding {
                color: #454545;
                margin: 0px;
                font-weight: 600;
                font-size: 15px;
                font-family: "Poppins", sans-serif;
            }

            .second-subtitle-main-boarding {
                color: #454545;
                margin: 0px;
                margin-top: 5%;
                font-weight: 600;
                max-width: 50%;
                font-size: 10px;
                font-family: "Poppins", sans-serif;
            }

            .column-sub-inside-information {
                float: left;
                width: 50%;
            }

            .content-title-boarding {
                float: left;
                margin: 20px 0;
                height: 31px;
                width: 100%;
                text-align: center;
                /* background-color: blue; */
            }

            .border-gray-title {
                /* float: left; */
                background-color: #dddddd;
                height: 1px;
                margin-top: 15px;
                /* width: 165px; */
                width: 100%;
                /* min-width: 50px;
                max-width: 80px; */
            }

            .title-boarding {
                /* float: left; */
                font-size: 19px;
                color: #454545;
                width: 600px;
                height: 31px;
                margin: -15px auto 0 auto;
                /* margin-top: -15px; */
                font-weight: 600;
                z-index: 10;
                background-color: #fff;
            }

            .column-card-boarding {
                float: left;
                width: 50%;
            }

            .column-card-boarding-four {
                margin: 12px 20px 10px 0;
                float: left;
                width: 76%;
                height: 280px;
            }

            .column-card-boarding-five {
                float: left;
                margin-top: 12px;
                width: 175px;
                height: 280px;
            }

            .row {
                margin: 0 6% 0 6%;
            }

            .card-boarding-one {
                width: 95%;
                min-height: 100px;
                max-height: 300px;
                border: 1px solid #c4c4c4;
                border-radius: 15px;
            }

            .card-boarding-two {
                width: 95%;
                margin-top: 3%;
                background-color: #ffffff;
                border-color: #c4c4c4;
                border-style: solid;
                border-width: 1px;
                border-radius: 15px;
            }

            .card-boarding-three {
                width: 100%;
                background-color: #ffffff;
                border: 1px solid #c4c4c4;
                border-radius: 15px;
            }

            .card-boarding-four {
                max-width: 100%;
                height: 100%;
                background-color: #ffffff;
                border-color: #c4c4c4;
                border-style: solid;
                border-width: 1px;
                border-radius: 15px;
            }

            .card-boarding-five {
                max-width: 100%;
                height: 100%;
                background-color: #ffffff;
                border-color: #c4c4c4;
                border-style: solid;
                border-width: 1px;
                border-radius: 15px;
            }


            .rectangle-one {
                width: 100%;
                height: 36px;
                background-size: contain;
                background-repeat: no-repeat;
                /* background-image: url("https://storage.googleapis.com/cabeen-dev/rectangle_one.svg"); */
            }

            .rectangle-two {
                width: 100%;
                height: 36px;
                background-size: contain;
                background-repeat: no-repeat;
                /* background-image: url("https://storage.googleapis.com/cabeen-dev/rectangle_two.svg"); */
            }

            .rectangle-three {
                width: 120%;
                height: 36px;
                background-size: contain;
                background-repeat: no-repeat;
                /* background-image: url("https://storage.googleapis.com/cabeen-dev/rectangle_three.svg"); */
            }

            .rectangle-four {
                width: 100%;
                height: 96px;
                /* margin-top: -1%; */
                /* background-size: contain; */
                /* background-repeat: no-repeat; */
                /* background-image: url("https://storage.googleapis.com/cabeen-dev/rectangle_four.svg"); */
            }

            .rectangle-five {
                width: 100%;
                height: 35px;
                background-size: contain;
                background-repeat: no-repeat;
                /* background-image: url("https://storage.googleapis.com/cabeen-dev/rectangle_five.svg"); */
            }

            .title-boarding-card {
                font-family: "Poppins", sans-serif;
                font-size: 14px;
                width: 100%;
                color: #E5F2F2;
                margin: 0px;
                position: relative;
                padding-top: 2.5%;
                padding-left: 3%;
                font-weight: bold;
            }

            .title-boarding-card-last {
                font-family: "Poppins", sans-serif;
                font-size: 14px;
                width: 100%;
                margin: 0px;
                position: relative;
                padding-top: 1%;
                padding-left: 3%;
                font-weight: bold;
                color: #ffffff;
            }

            .row-content-travel {
                margin-top: 2% 0 5% 3%;
                max-width: 100%;
                font-size: 10px;
                font-weight: 600;
            }

            .row-content-travel ul {
                /* font-size: 10px; */
                float: left;
                list-style: none;
                margin-left: -20px;
                width: 200px;
            }

            .row-content-travel ul li::before {
                content: "\2022";
                color:  #04847D;
                font-size: 18px;
                /* font-weight: bold; If you want it to be bold */
                display: inline-block; /* Needed to add space between the bullet and the text */
                width: 1em; /* Also needed for space (tweak if needed) */
                margin-left: -1em; /* Also needed for space (tweak if needed) */
            }

            .card-boarding-four ul {
                margin: 10px 0 0 3px;
                width: 89%;
                /* margin: 1em; */
            }

            .icon-boarding {
                margin: 20px 13px;
                width: 75px;
                height: 75px;
                float: left;
            }

            .stamp-here {
                font-size: 15px;
                font-weight: bold;
                font-family: "Poppins", sans-serif;
                text-align: center;
                color: #04847d;
            }

            .only-international {
                font-size: 11px;
                font-weight: 600;
                font-family: "Poppins", sans-serif;
                text-align: center;
                padding-top: 92%;
                color: #454545;
            }

            .content-stamp-here {
                display: flex;
                flex-direction: column;
                justify-content: space-between;
            }


            .row-content-travel::after {
                content: "";
                display: table;
                clear: both;
            }

            .gap-desc-travel-one {
                margin-left: 4%;
                margin-right: 4%;
            }

            .ellipse-point {
                width: 5px;
                height: 5px;
                margin-top: 4%;
                margin-right: 5%;
                border-radius: 2.5px;
                background-color: #04847d;
            }

            .ellipse-point-second {
                width: 5px;
                height: 5px;
                margin-top: 2.5%;
                margin-right: 2.5%;
                border-radius: 2.5px;
                background-color: #04847d;
            }

            .content-point {
                display: flex;
                flex-direction: row;
                align-items: flex-start;
            }

            /* Float four columns side by side */
            .column-content-img {
                float: left;
                width: 22%;
            }

            /* Responsive columns */
            @media screen and (max-width: 600px) {
                .column-content-img {
                    width: 100%;
                    display: block;
                    margin-bottom: 20px;
                }
            }

            .column-desc-boarding {
                float: left;
                margin-top: 20px;
                min-width: 50px;
                max-width: 200px;
                font-size: 10px;
                color: #454545;
                font-weight: 600;
            }

            @media screen and (max-width: 600px) {
                .column-desc-boarding {
                    width: 100%;
                    display: block;
                    margin-bottom: 20px;
                }
            }

            /* Clear floats after the columns */
            p.understand-text {
                font-size: 10px;
                color: #04847d;
                margin-top: -9px;
                font-weight: 600;
                width: 80%;
                margin-bottom: 27px;
                margin-left: 4%;
                font-weight: 600;
            }

            .content-backpack img {
                margin: 10px 0 30px 70px;
                height: 130px;
                object-fit: contain;
            }

            .bg-thanks {
                background-color: #ddefe9;
                width: 100%;
                height: 55px;
                margin-top: 10px;
                float: left;
            }

            .bg-thanks h2 {
                font-size: 18px;
                color: #04847d;
                font-weight: 600;
                text-align: center;
            }

            .footer-boarding {
                float: left;
                width: 100%;
                height:98px;
                background-color: #f0f9f6;
            }

            .icon-cabeen-boarding {
                float: left;
                height: 28px;
                margin-left: 6%;
                margin-top: 35px;
                object-fit: contain;
            }

            .download-cabeen-text {
                font-size: 12px;
                margin-top: 7px;
                margin-bottom: 7px;
                color: #99b1ab;
                font-weight: 600;
            }

            .content-download-footer {
                float: left;
                margin: 0 6% 0 45%;
            }

            .content-store img {
                height: 40px;
            }
        </style>
        <div class="container">
            <div class="header">
                <h2>Your Boarding Pass</h2>
            </div>
            <div class="card-bp">
                    <div class="card-bp-header">
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_lion.svg" class="icon-lion">
                        <img src="{{.BarcodeLandscape}}" class="icon-barcode">
                    </div>
                    <div class="card-bp-body">
                        <div class="column-barcode">
                            <img src="{{.BarcodeRotate}}" class="icon-body-barcode">
                        </div>
                        <div id="column-information-boarding">
                            <div id="content-info-header">
                                <div class="content-info-header-field">
                                    <p class="title-main-boarding">Name</p>
                                    <p class="subtitle-main-boarding">{{.FullName}}</p>
                                </div>
                                <div class="content-info-header-field">
                                    <p class="title-main-boarding">Frequent Flyer No.</p>
                                    <p class="subtitle-main-boarding">{{.FFNumber}}</p>
                                </div>
                                <div class="content-info-header-field">
                                    <p class="title-main-boarding">Security No.</p>
                                    <p class="subtitle-main-boarding">{{.SecurityNumber}}</p>
                                </div>
                            </div>
                            <div class="content-inside-information">
                                <div class="column-inside-information">
                                    <div class="second-content-inside-information">
                                        <div class="column-sub-inside-information">
                                            <p class="second-title-main-boarding">From</p>
                                            <p class="second-subtitle-main-boarding">{{.OriginCity}}</p>
                                            <br>
                                            <p class="second-title-main-boarding">Terminal</p>
                                            <p class="second-subtitle-main-boarding">{{.OriginTerminal}}</p>
                                        </div>
                                        <div class="column-sub-inside-information">
                                            <p class="second-title-main-boarding">To</p>
                                            <p class="second-subtitle-main-boarding">{{.DestinationCity}}</p>
                                            <br>
                                            <p class="second-title-main-boarding">Boarding Time</p>
                                            <p class="second-subtitle-main-boarding">{{.DestinationTerminal}}</p>
                                        </div>
                                    </div>
                                </div>
                                <div class="column-inside-information">
                                    <div class="second-content-inside-information">
                                        <div class="column-sub-inside-information">
                                            <p class="second-title-main-boarding">SSR</p>
                                            <p class="second-subtitle-main-boarding">{{.FlightSSR}}</p>
                                            <br>
                                            <p class="second-title-main-boarding">Class</p>
                                            <p class="second-subtitle-main-boarding">{{.FlightClass}}</p>
                                        </div>
                                        <div class="column-sub-inside-information">
                                            <p class="second-title-main-boarding">Flight No.</p>
                                            <p class="second-subtitle-main-boarding">{{.FlightNumber}}</p>
                                            <br>
                                            <p class="second-title-main-boarding">PNR</p>
                                            <p class="second-subtitle-main-boarding">{{.FlightPNR}}</p>
                                        </div>
                                    </div>
                                </div>
                                <div class="column-inside-information">
                                    <div class="second-content-inside-information">
                                        <div class="column-sub-inside-information">
                                            <p class="second-title-main-boarding">Date</p>
                                            <p class="second-subtitle-main-boarding">{{.DepartureDate}}</p>
                                            <br>
                                            <p class="second-title-main-boarding">Seat</p>
                                            <p class="second-subtitle-main-boarding">{{.Seat}}</p>
                                        </div>
                                        <div class="column-sub-inside-information">
                                            <p class="second-title-main-boarding">Departure Time</p>
                                            <p class="second-subtitle-main-boarding">{{.DepartureTime}}</p>
                                            <br>
                                            <p class="second-title-main-boarding">E-Ticket No.</p>
                                            <p class="second-subtitle-main-boarding">{{.ETicketNumber}}</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                <!-- </div> -->
            </div>


            <!-- TITLE PRESENT BOARDING PASS -->
            <div class="content-title-boarding">
                <div class="border-gray-title"></div>
                <div class="title-boarding">Please present this boarding pass to the airport for flight</div>
                <!-- <h3>Please present this boarding pass to the airport for flight</h3> -->
                <!-- <div class="border-gray-title"></div> -->
            </div>

            <!-- LINE ONE -->
            <div class="row">
                <div class="column-card-boarding">
                    <div class="card-boarding-one">
                        <div class="rectangle-one">
                            <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;width: 22%;">
                            <p class="title-boarding-card">Check-in with luggage</p>
                        </div>
                        <div class="row-content-travel">
                            <img src="https://storage.googleapis.com/cabeen-dev/travel_one.svg" class="icon-boarding">
                            <p class="column-desc-boarding">
                                Please head to the check-in counter no later than 60 minutes before departure
                            </p>
                        </div>
                    </div>
                    <div class="card-boarding-two">
                        <div class="rectangle-two">
                            <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;width: 22%;">
                            <p class="title-boarding-card">Safety Regulations</p>
                        </div>
                        <div class="content-card-travel">
                            <div class="row-content-travel">
                                <img src="https://storage.googleapis.com/cabeen-dev/safety.svg" class="icon-boarding">
                                <ul>
                                    <li>Make sure your items always stay with you</li>
                                    <li>Make sure you know the contents of your luggage</li>
                                    <li>Pay attention to items that are not allowed to be carried</li>
                                </ul>
                            </div>
                        </div>
                        <br>
                        <p class="understand-text">Make sure you understand, know and agree to the applicable security rules</p>
                    </div>
                </div>
                <div class="column-card-boarding">
                    <div class="card-boarding-three">
                        <div class="rectangle-three">
                            <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;width: 22%;">
                            <p class="title-boarding-card">Carry-on bag only?</p>
                        </div>
                        <div class="content-card-travel">
                            <div class="row-content-travel">
                                <img src="https://storage.googleapis.com/cabeen-dev/backpack.png" class="icon-boarding">
                                <ul>
                                    <li>Domestic destination - please head to the departure gate</li>
                                    <li>International destinations - please go to the check-in counter for verification of travel documents</li>
                                </ul>
                            </div>
                        </div>
                        <div class="rectangle-four">
                            <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;float: left;width: 22%;">
                            <p class="title-boarding-card">Size of carry-on bag</p>
                        </div>
                        <div class="content-backpack">
                            <img src="https://storage.googleapis.com/cabeen-dev/group_backpack.png">
                        </div>
                    </div>
                </div>
                <div class="column-card-boarding-four">
                    <div class="card-boarding-four">
                        <div class="rectangle-five">
                            <img src="https://storage.googleapis.com/cabeen-dev/rectangle_five.svg" style="position: absolute;width: 22%;">
                            <p class="title-boarding-card-last">Important Information</p>
                        </div>
                        <div class="content-card-travel">
                            <div class="row-content-travel">
                                <ul>
                                    <li>Make sure your travel documents are still valid</li>
                                    <li>Allowance for baggage as stipulated in your E-Ticket regulations</li>
                                    <li>Receiving boarding passes is at the discretion of the airport</li>
                                    <li>For operational, safety or security reasons we can change our seats at any time even after you board the aircraft</li>
                                    <li>You are expected to be at the Boarding Gate no later than 40 minutes before departure time to avoid being offroad from the aircraft and from being charged a no-show fee</li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="column-card-boarding-five">
                    <div class="card-boarding-five">
                        <div class="content-stamp-here">
                            <p class="stamp-here">Stamp here</p>
                            <p class="only-international">Only for international <br> destination services</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="bg-thanks">
                <h2>Thank you for choosing to fly with Lion Air</h2>
            </div>
            
            <div class="footer-boarding">
                    <img src="https://storage.googleapis.com/cabeen-dev/icon_footer.svg" class="icon-cabeen-boarding">
                    <div class="content-download-footer">
                        <p class="download-cabeen-text">Download Flapp Apps</p>
                        <div class="content-store">
                            <img src="https://storage.googleapis.com/cabeen-dev/playstore.png" onclick="location.href='https://play.google.com/store'">
                            <img src="https://storage.googleapis.com/cabeen-dev/appstore.png" onclick="location.href='https://apps.apple.com/'">
                        </div>
                    </div>
            </div>
    </body>
    </html>
    `

	FlightHTML = `
    <html>
    <head></head>
    <body>
        <style type="text/css">
            @import url("https://fonts.googleapis.com/css?family=Poppins&display=swap");

            body {
                margin: 0;
                padding: 0;
                font-family: "Poppins", sans-serif;
            }

            .container {
                margin: 0 auto;
                position: relative;
                width: auto;
            }

            @media screen and (min-width: 1024px) {
                .container {
                    max-width: 960px;
                }
            }

            .bg-banner-cabeen {
                width: 100%;
                height: 170px;
            }

            .bg-banner-cabeen>img {
                width: 100%;
                z-index: -1;
                position: relative;
            }

            .content-logo-cabeen {
                margin-top: -155px;
                padding: 30px;
            }

            .content-logo-cabeen>img {
                width: 225px;
            }

            .content-logo-cabeen>p {
                font-size: 21px;
                font-weight: normal;
                margin-left: 7%;
                margin-top: 0%;
                color: #575757;
            }

            .bg-booking-code {
                background-color: #f8f8f8;
                min-width: 80%;
                max-width: 100%;
                min-height: 58px;
                max-height: 70px;
                padding: 10px 30px;
                margin-top: -2%;
            }

            .bg-booking-code>p {
                float: left;
                font-size: 14px;
                text-align: left;
                color: #575757;
                margin: 0;
            }

            .bg-booking-code>p:nth-child(2) {
                margin-left: 55%;
                font-weight: bold;
            }

            .bg-booking-code>p:nth-child(2) b {
                font-weight: bold;
                font-size: 16px;
            }

            .bg-booking-code>p:nth-child(3) {
                font-size: 20px;
                font-weight: 900;
                color: #04847d;
            }

            .footer-boarding {
                float: left;
                width: 93%;
                background-color: #f0f9f6;
                padding: 17px 30px;
                color: #99b1ab;
            }

            .footer-boarding div {
                float: left;
                width: 50%;
            }

            .footer-boarding>div:first-child p {
                font-size: 12px;
                margin: 0;
                font-weight: 600;
            }

            .footer-boarding>div:first-child img {
                margin-top: 10px;
                width: 128px;
            }

            .footer-boarding>div:first-child img:nth-child(3) {
                margin-left: 7px;
            }

            .footer-boarding>div:nth-child(2)>div {
                width: 100%;
            }

            .footer-boarding>div:nth-child(2)>div img {
                margin-top: 12px;
                float: right;
                width: 105px;
            }

            .footer-boarding>div:nth-child(2) p {
                text-align: right;
                margin: 0;
            }

            .content-status-title {
                margin-top: 2%;
                width: 100%;
                min-height: 30px;
                max-height: 50px;
            }

            .content-status-title>p {
                float: left;
                color: #04847d;
                font-size: 16px;
                margin: 0px;
                margin-right: 8px;
                font-weight: 600;
            }

            .content-status-title>div {
                margin-top: 12px;
                float: left;
                height: 2px;
                width: 82%;
                background-color: gainsboro;
            }

            .content-head-eticket {
                padding-left: 63px;
                padding-right: 30px;
                border-bottom: gainsboro 2px solid;
            }

            #wrap-head-content {
                margin-top: 3%;
                margin-bottom: 3%;
                width: 93%;
                color: #454545;
            }

            #wrap-head-content>img {
                max-width: 100%;
                height: 32px;
            }

            #wrap-head-content>p {
                font-size: 14px;
                margin: 0px;
                width: 50%;
            }

            #wrap-head-content>p:nth-child(3) {
                margin-top: -21px;
                margin-right: -40px;
                float: right;
                text-align: right;
            }

            .content-track {
                float: left;
                margin-top: 2%;
                margin-bottom: 2%;
                padding-left: 63px;
                padding-right: 30px;
                width: 100%;
            }

            .date-text {
                color: #454545;
                font-size: 20px;
                margin: 0px;
                font-weight: bold;
            }

            .time-text {
                color: #454545;
                font-size: 16px;
                margin: 0px;
                font-weight: normal;
            }

            .column-track-time {
                float: left;
                position: relative;
                width: 100%;
                margin-top: 2%;
            }

            .content-track-time {
                float: left;
                width: 120px;
            }

            .content-track-time>p {
                color: #454545;
                font-size: 20px;
                margin: 0px;
                font-weight: normal;
            }

            .content-track-time>p:nth-child(2) {
                float: left;
                margin-top: 70px;
            }

            .content-track-time>span {
                float: left;
                color: #ff8080;
                font-size: 16px;
                padding: 2.5px;
                font-weight: 600;
                margin-top: 70px;
            }

            .content-dot-track {
                width: 3%;
                float: left;
                margin-top: 1%;
            }

            .content-dot-track>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #04847d;
            }

            /*
            .content-dot-track>div:nth-child(2) {
                background-color: #04847d;
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 70px;
            }
            */

            .content-dot-track>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-dot-track-last {
                width: 3%;
                float: left;
                padding-top: 1%;
            }

            .content-dot-track-last>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-dot-track-last>div:nth-child(2) {
                background-color: #04847d;
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 70px;
            }

            .content-dot-track-last>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #04847d;
            }

            .content-dot-track-middle {
                width: 3%;
                float: left;
                padding-top: 1%;
            }

            .content-dot-track-middle>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-dot-track-middle>div:nth-child(2) {
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 70px;
                border-left: 2px dashed #dddddd;
            }

            .content-dot-track-middle>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-track-place {
                float: left;
                color: #454545;
                font-size: 16px;
                width: 80%;
            }

            .content-track-place>div {
                float: left;
                margin-bottom: 46px;
                width: 100%;
            }

            .content-track-place>div p:first-child {
                margin: 0px;
                font-weight: 600;
            }

            .content-track-place>div p:nth-child(2) {
                float: left;
                width: 50%;
                font-size: 14px;
                margin: 0px;
            }

            .wrap-transit {
                margin: -50px 0 -20px 125px;
                float: left;
                min-height: 70px;
                max-height: 200px;
                border-left: 2px dashed #dddddd;
                width: 100%;
            }

            .wrap-transit div {
                float: left;
                border-radius: 5px;
                min-height: 20px;
                max-height: 70px;
                min-width: 170px;
                max-width: 300px;
                margin: 47px 57px 47px 23px;
                padding: 3px 5px;
                border: #575757 1.5px solid;
            }

            .wrap-transit img {
                float: left;
                margin-right: 10px;
                width: 20px;
            }

            .wrap-transit p {
                float: left;
                font-size: 12px;
                font-weight: 600;
                margin: 0;
                margin-top: 1px;
                min-width: 170px;
                max-width: 200px;
            }

            table {
                float: left;
                width: 93%;
                text-align: left;
                border-collapse: collapse;
                margin: 0px 35px 30px;
                color: #454545;
                font-size: 14px;
            }

            thead {
                background-color: #f8f8f8;
                font-weight: 600;
            }

            thead th,
            tbody td {
                padding: 10px 0;
            }

            tbody td {
                border-bottom: gainsboro 0.5px solid;
            }

            thead th:first-child,
            th:nth-child(3),
            th:nth-child(4) {
                text-align: center;
            }

            thead th:first-child {
                padding: 0 5px;
            }

            tbody tr td:first-child,
            td:nth-child(3),
            td:nth-child(4) {
                text-align: center;
            }
        </style>
        <div class="container">
            <div>
                <div class="bg-banner-cabeen">
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/banner_cabeen.png">
                    <div class="content-logo-cabeen">
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_header.svg">
                        <p>E-ticket</p>
                    </div>
                </div>
                <div class="bg-booking-code">
                    <p>Airlane Booking Code (PNR)</p>
                    <p>Flapp Booking ID <span><b>{{.BookID}}</b></span></p> 
                    <p>{{.PNR}}</p>
                </div>
                <div class="content-head-eticket">
                    <div class="content-status-title">
                        <p>{{.FlightType}} Flight</p>
                        <div></div>
                    </div>
                    <div id="wrap-head-content">
                        <img src="{{.Airline.Logo}}">
                        <p>Flight Number <b>{{.Airline.FlightNumber}} </b></p>
                        <p><b>{{.Airline.Class}}</b></p>
                    </div>
                </div>
                <div class="content-track">
                    <p class="date-text">{{.Airline.DepartureDate}}</p>
                    <p class="time-text">{{.Airline.DepartureTime}} - {{.Airline.ArrivedTime}} ({{.Airline.DurationTime}}, {{.Airline.FlightDirect}})</p>
                    {{range $index, $elem := .Schedules}}
                            <div class="column-track-time">
                                <div class="content-track-time">
                                    <p>{{$elem.DepartureTime}}</p>
                                    <p>{{$elem.ArrivedTime}}</p>{{if $elem.IsOvernight }}<span>+1d</span>{{end}}
                                </div>
                                {{if eq $index 0}}
                                <div class="content-dot-track">
                                {{else if last $index $.Schedules}}
                                <div class="content-dot-track-last">
                                {{else}}
                                <div class="content-dot-track-middle">
                                {{end}}
                                    <div></div>
                                    <div style="
                                    {{if eq $index 0}}
                                        {{if eq (len $.Schedules) 1}}
                                        background-color: #c4c4c4;
                                        {{else}}
                                        background-color: #04847d;
                                        {{end}}
                                    {{end}}
                                    width: 1px;
                                    margin: 8px 0 8px 5px;
                                    height: 70px;"></div>
                                    <div></div>
                                </div>
                                <div class="content-track-place">
                                    <div>
                                        <p>{{$elem.DepartureAirportCity}}({{$elem.DepartureAirportIATA}})</p>
                                        <p>{{$elem.DepartureAirportName}} - {{$elem.DepartureAirportTerminal}}</p>
                                    </div>
                                    <div>
										<p>{{$elem.ArrivedAirportCity}}({{$elem.ArrivedAirportIATA}})</p>
										<p>{{$elem.ArrivedAirportName}} - {{$elem.ArrivedAirportTerminal}}</p>
                                    </div>
                                </div>
                            </div>
                        {{if and $elem.TransitCityName}}
                            {{if or (not (last $index $.Schedules)) (and (eq (len $.Schedules) 0) (eq $index 0))}}
                            <div class="wrap-transit">
                                <div>
                                    <img src="https://storage.googleapis.com/cabeen-dev/assets/information.svg">
                                    <p>Transit {{$elem.TransitDurationTime}} in {{$elem.TransitCityName}}</p>
                                </div>
                            </div>
                            {{end}}
                        {{end}}
                    {{end}}
                </div>

                {{if gt (len $.Passengers) 0}}
                <table>
                    <thead>
                        <tr>
                            <th>No.</th>
                            <th>Passenger(s)</th>
                            <th><img src="https://storage.googleapis.com/cabeen-dev/assets/user_head.svg"
                                    class="icon-head-table"></th>
                            <th><img src="https://storage.googleapis.com/cabeen-dev/assets/luggage.svg"
                                    class="icon-head-table"></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range $index, $elem := .Passengers}}
                            <tr>
                                <td>{{add $index $.Index}}</td>
                                <td>{{$elem.Name}}</td>
                                <td>{{$elem.Type}}</td>
                                <td>{{$elem.Baggage}} Kg</td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
            {{end}}

            <div class="footer-boarding">
                <div>
                    <p>Download Flapp Apps</p>
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/playstore.png"
                        onclick="location.href='https://play.google.com/store'">
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/appstore.png"
                        onclick="location.href='https://apps.apple.com/'">
                </div>
                <div>
                    <div>
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_footer.svg">
                    </div>
                    <p>{{now.UTC.Year}}, PT LIONMAS DINAMIKAEXPRESS</p>
                </div>
            </div>
        </div>
    </body>
    </html>
    `

	FlightSecondHTML = `
    <html>
    <head></head>
    <body>
        <style type="text/css">
            @import url("https://fonts.googleapis.com/css?family=Poppins&display=swap");

            body {
                margin: 0;
                padding: 0;
                font-family: "Poppins", sans-serif;
            }

            .container {
                margin: 0 auto;
                position: relative;
                width: auto;
            }

            @media screen and (min-width: 1024px) {
                .container {
                    max-width: 960px;
                }
            }

            .bg-banner-cabeen {
                width: 100%;
                height: 170px;
            }

            .bg-banner-cabeen>img {
                width: 100%;
                z-index: -1;
                position: relative;
            }

            .content-logo-cabeen {
                margin-top: -155px;
                padding: 30px;
            }

            .content-logo-cabeen>img {
                width: 225px;
            }

            .content-logo-cabeen>p {
                font-size: 21px;
                font-weight: normal;
                margin-left: 7%;
                margin-top: 0%;
                color: #575757;
            }

            .bg-booking-code {
                background-color: #f8f8f8;
                min-width: 80%;
                max-width: 100%;
                min-height: 58px;
                max-height: 70px;
                padding: 10px 30px;
                margin-top: -4%;
                margin-bottom: 20px;
            }
    
            .bg-booking-code>p {
                float: left;
                font-size: 14px;
                text-align: left;
                color: #575757;
                margin: 0;
            }
    
            .bg-booking-code>p:nth-child(2) {
                float: left;
                margin-left: 50%;
                font-weight: bold;
            }
    
            .bg-booking-code>p:nth-child(2) b {
                font-weight: bold;
                font-size: 16px;
            }
    
            .bg-booking-code>p:nth-child(3) {
                float: left;
                font-size: 20px;
                font-weight: 900;
                color: #04847d;
            }

            .footer-boarding {
                float: left;
                width: 93%;
                background-color: #f0f9f6;
                padding: 17px 30px;
                color: #99b1ab;
            }

            .footer-boarding div {
                float: left;
                width: 50%;
            }

            .footer-boarding>div:first-child p {
                font-size: 12px;
                margin: 0;
                font-weight: 600;
            }

            .footer-boarding>div:first-child img {
                margin-top: 10px;
                width: 128px;
            }

            .footer-boarding>div:first-child img:nth-child(3) {
                margin-left: 7px;
            }

            .footer-boarding>div:nth-child(2)>div {
                width: 100%;
            }

            .footer-boarding>div:nth-child(2)>div img {
                margin-top: 12px;
                float: right;
                width: 105px;
            }

            .footer-boarding>div:nth-child(2) p {
                text-align: right;
                margin: 0;
            }

            .content-status-title {
                margin-top: 2%;
                width: 100%;
                min-height: 30px;
                max-height: 50px;
            }

            .content-status-title>p {
                float: left;
                color: #04847d;
                font-size: 16px;
                margin: 0px;
                margin-right: 8px;
                font-weight: 600;
            }

            .content-status-title>div {
                margin-top: 12px;
                float: left;
                height: 2px;
                width: 82%;
                background-color: gainsboro;
            }

            .content-head-eticket {
                padding-left: 63px;
                padding-right: 30px;
                border-bottom: gainsboro 2px solid;
            }

            #wrap-head-content {
                margin-top: 3%;
                margin-bottom: 3%;
                width: 93%;
                color: #454545;
            }

            #wrap-head-content>img {
                max-width: 100%;
                height: 32px;
            }

            #wrap-head-content>p {
                font-size: 14px;
                margin: 0px;
                width: 50%;
            }

            #wrap-head-content>p:nth-child(3) {
                margin-top: -21px;
                margin-right: -40px;
                float: right;
                text-align: right;
            }

            .content-track {
                float: left;
                margin-top: 2%;
                margin-bottom: 2%;
                padding-left: 63px;
                padding-right: 30px;
                width: 100%;
            }

            .date-text {
                color: #454545;
                font-size: 20px;
                margin: 0px;
                font-weight: bold;
            }

            .time-text {
                color: #454545;
                font-size: 16px;
                margin: 0px;
                font-weight: normal;
            }

            .column-track-time {
                float: left;
                position: relative;
                width: 100%;
                margin-top: 2%;
            }

            .content-track-time {
                float: left;
                width: 70px;
            }

            .content-track-time>p {
                color: #454545;
                font-size: 20px;
                margin: 0px;
                font-weight: normal;
            }

            .content-track-time>p:nth-child(2) {
                margin-top: 70px;
            }

            .content-dot-track {
                width: 3%;
                float: left;
                margin-top: 1%;
            }

            .content-dot-track>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #04847d;
            }

            /*
            .content-dot-track>div:nth-child(2) {
                background-color: #04847d;
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 70px;
            }
            */

            .content-dot-track>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-dot-track-last {
                width: 3%;
                float: left;
                padding-top: 1%;
            }

            .content-dot-track-last>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-dot-track-last>div:nth-child(2) {
                background-color: #04847d;
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 70px;
            }

            .content-dot-track-last>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #04847d;
            }

            .content-dot-track-middle {
                width: 3%;
                float: left;
                padding-top: 1%;
            }

            .content-dot-track-middle>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-dot-track-middle>div:nth-child(2) {
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 70px;
                border-left: 2px dashed #dddddd;
            }

            .content-dot-track-middle>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-track-place {
                float: left;
                color: #454545;
                font-size: 16px;
                width: 80%;
            }

            .content-track-place>div {
                float: left;
                margin-bottom: 46px;
                width: 100%;
            }

            .content-track-place>div p:first-child {
                margin: 0px;
                font-weight: 600;
            }

            .content-track-place>div p:nth-child(2) {
                float: left;
                width: 50%;
                font-size: 14px;
                margin: 0px;
            }

            .wrap-transit {
                margin: -50px 0 -20px 75px;
                float: left;
                min-height: 70px;
                max-height: 200px;
                border-left: 2px dashed #dddddd;
                width: 100%;
            }

            .wrap-transit div {
                float: left;
                border-radius: 5px;
                min-height: 20px;
                max-height: 70px;
                min-width: 170px;
                max-width: 300px;
                margin: 47px 57px 47px 23px;
                padding: 3px 5px;
                border: #575757 1.5px solid;
            }

            .wrap-transit img {
                float: left;
                margin-right: 10px;
                width: 20px;
            }

            .wrap-transit p {
                float: left;
                font-size: 12px;
                font-weight: 600;
                margin: 0;
                margin-top: 1px;
                min-width: 170px;
                max-width: 200px;
            }

            table {
                float: left;
                width: 93%;
                text-align: left;
                border-collapse: collapse;
                margin: 0px 35px 30px;
                color: #454545;
                font-size: 14px;
            }

            thead {
                background-color: #f8f8f8;
                font-weight: 600;
            }

            thead th,
            tbody td {
                padding: 10px 0;
            }

            tbody td {
                border-bottom: gainsboro 0.5px solid;
            }

            thead th:first-child,
            th:nth-child(3),
            th:nth-child(4) {
                text-align: center;
            }

            thead th:first-child {
                padding: 0 5px;
            }

            tbody tr td:first-child,
            td:nth-child(3),
            td:nth-child(4) {
                text-align: center;
            }
        </style>
        <div class="container">
            <div>
                <div class="bg-banner-cabeen">
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/banner_cabeen.png">
                    <div class="content-logo-cabeen">
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_header.svg">
                        <p>E-ticket</p>
                    </div>
                </div>
                <div class="bg-booking-code">
                    <p>Airlane Booking Code (PNR)</p>
                    <p>Flapp Booking ID <span><b>{{.BookID}}</b></span></p> 
                    <p>{{.PNR}}</p>
                </div>

                <table>
                    <thead>
                        <tr>
                            <th>No.</th>
                            <th>Passenger(s)</th>
                            <th><img src="https://storage.googleapis.com/cabeen-dev/assets/user_head.svg"
                                    class="icon-head-table"></th>
                            <th><img src="https://storage.googleapis.com/cabeen-dev/assets/luggage.svg"
                                    class="icon-head-table"></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range $index, $elem := .Passengers}}
                            <tr>
                                <td>{{add $index $.Index}}</td>
                                <td>{{$elem.Name}}</td>
                                <td>{{$elem.Type}}</td>
                                <td>{{$elem.Baggage}} Kg</td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>

            <div class="footer-boarding">
                <div>
                    <p>Download Flapp Apps</p>
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/playstore.png"
                        onclick="location.href='https://play.google.com/store'">
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/appstore.png"
                        onclick="location.href='https://apps.apple.com/'">
                </div>
                <div>
                    <div>
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_footer.svg">
                    </div>
                    <p>{{now.UTC.Year}}, PT LIONMAS DINAMIKAEXPRESS</p>
                </div>
            </div>
        </div>
    </body>
    </html>
    `

	FlightReceiptHTML = `
    <html>
    <head></head>
    <body>
        <style type="text/css">
            @import url("https://fonts.googleapis.com/css?family=Poppins&display=swap");

            body {
                margin: 0;
                padding: 0;
                font-family: "Poppins", sans-serif;
            }

            .container {
                margin: 0 auto;
                position: relative;
                width: auto;
            }

            @media screen and (min-width: 1024px) {
                .container {
                    max-width: 960px;
                }
            }

            .bg-banner-cabeen {
                width: 100%;
                height: 170px;
            }

            .bg-banner-cabeen>img {
                width: 100%;
                z-index: -1;
                position: relative;
            }

            .content-logo-cabeen {
                margin-top: -155px;
                padding: 30px;
            }

            .content-logo-cabeen>img {
                width: 225px;
            }

            .content-logo-cabeen>p {
                font-size: 21px;
                font-weight: normal;
                margin-left: 7%;
                margin-top: 0%;
                color: #575757;
            }

            .bg-booking-code {
                position: relative;
                float: left;
                background-color: #f8f8f8;
                width: 93.8%;
                padding: 10px 30px;
                margin-top: -4%;
            }

            .bg-booking-code>p {
                float: left;
                font-size: 15px;
                text-align: left;
                color: #575757;
                margin: 0;
            }

            .bg-booking-code>p:nth-child(2) {
                margin-left: 50%;
            }

            .footer-boarding {
                float: left;
                width: 93%;
                background-color: #f0f9f6;
                padding: 17px 30px;
                color: #99b1ab;
            }

            .footer-boarding div {
                float: left;
                width: 50%;
            }

            .footer-boarding>div:first-child p {
                font-size: 12px;
                margin: 0;
                font-weight: 600;
            }

            .footer-boarding>div:first-child img {
                margin-top: 10px;
                width: 128px;
            }

            .footer-boarding>div:first-child img:nth-child(3) {
                margin-left: 7px;
            }

            .footer-boarding>div:nth-child(2)>div {
                width: 100%;
            }

            .footer-boarding>div:nth-child(2)>div img {
                margin-top: 12px;
                float: right;
                width: 105px;
            }

            .footer-boarding>div:nth-child(2) p {
                text-align: right;
                margin: 0;
            }

            .content-status-title {
                margin-top: 2%;
                width: 100%;
                min-height: 30px;
                max-height: 50px;
            }

            .content-status-title>p {
                float: left;
                color: #04847d;
                font-size: 16px;
                margin: 0px;
                margin-right: 8px;
                font-weight: 600;
            }

            .content-status-title>div {
                margin-top: 12px;
                float: left;
                height: 2px;
                width: 82%;
                background-color: gainsboro;
            }

            .content-head-eticket {
                /* padding-left: 63px; */
                padding: 0 30px 0 30px;
                /* padding-right: 30px; */
                border-bottom: gainsboro 2px solid;
            }

            .content-head-eticket ul {
                /* background-color: teal; */
                list-style: none;
                padding: 0;
                width: 35%;
                font-size: 15px;
                color: #575757;
                float: left;
            }

            .content-head-eticket>ol {
                font-size: 15px;
                float: left;
                padding: 0;
                color: #575757;
                font-weight: bold;
                margin-left: 13px;
            }

            .content-head-eticket>ol span {
                font-size: 14px;
                font-weight: normal;
            }

            .content-head-eticket ul:nth-child(2) {
                width: 30%;
                margin-left: 33.7%;
            }

            .content-head-eticket ul li:first-child h3 {
                width: 100%;
                float: left;
                margin: 0 0 10px;
                font-size: 16px;
                font-weight: 600;
                color: #04847d;
            }

            .content-head-eticket div {
                float: left;
                width: 100%;
            }

            .content-head-eticket div h3 {
                float: left;
                color: #04847d;
                font-size: 16px;
                font-weight: 600;
                margin: 0;
                /* background-color: tomato; */
            }

            .content-head-eticket div #hr-head {
                height: 2px;
                background: gainsboro;
                float: left;
                width: 80%;
                margin: 12px 0 13px 19px;
            }

            table.tb-profile {
                font-weight: normal;
                color: #454545;
                font-size: 15px;
                border-collapse: collapse;
            }

            #wrap-head-content {
                margin-top: 3%;
                margin-bottom: 3%;
                width: 93%;
                color: #454545;
            }

            #wrap-head-content>img {
                max-width: 100%;
                height: 32px;
            }

            #wrap-head-content>p {
                font-size: 14px;
                margin: 0px;
                width: 50%;
            }

            #wrap-head-content>p:nth-child(3) {
                margin-top: -21px;
                margin-right: -40px;
                float: right;
                text-align: right;
            }

            .content-track {
                margin-top: 2%;
                margin-bottom: 3%;
                padding-left: 63px;
                padding-right: 30px;
            }

            .date-text {
                color: #454545;
                font-size: 20px;
                margin: 0px;
                font-weight: bold;
            }

            .time-text {
                color: #454545;
                font-size: 16px;
                margin: 0px;
                font-weight: normal;
            }

            .column-track-time {
                float: left;
                width: 100%;
                margin-top: 2%;
                margin-bottom: 20px;
            }

            .content-track-time {
                float: left;
                width: 8%;
            }

            .content-track-time>p {
                color: #454545;
                font-size: 20px;
                margin: 0px;
                font-weight: normal;
            }

            .content-track-time>p:nth-child(2) {
                margin-top: 94px;
            }

            .content-dot-track {
                width: 8%;
                float: left;
                padding-top: 1%;
            }

            .content-dot-track>div:first-child {
                height: 12px;
                width: 12px;
                margin-bottom: 5%;
                border-radius: 6px;
                background-color: #04847d;
            }

            .content-dot-track>div:nth-child(2) {
                background-color: #c4c4c4;
                width: 1px;
                margin: 8px 0 8px 5px;
                height: 94px;
            }

            .content-dot-track>div:nth-child(3) {
                height: 12px;
                width: 12px;
                margin-top: 5%;
                border-radius: 6px;
                background-color: #c4c4c4;
            }

            .content-track-place {
                float: left;
                color: #454545;
                font-size: 16px;
            }

            .content-track-place>div:first-child {
                margin-bottom: 70px;
            }

            .content-track-place>div p:first-child {
                margin: 0px;
                font-weight: bold;
            }

            .content-track-place>div p:nth-child(2) {
                font-size: 14px;
                margin: 0px;
            }

            table#tb-list-cust {
                float: left;
                width: 100%;
                text-align: left;
                border-collapse: collapse;
                margin: 0px 0 130px;
                color: #454545;
                font-size: 14px;
            }

            table#tb-list-cust>thead {
                background-color: #f8f8f8;
                font-weight: 600;
            }

            table#tb-list-cust>thead th,
            table#tb-list-cust>tbody td {
                padding: 10px 0;
            }

            table#tb-list-cust>tbody td {
                border-bottom: gainsboro 0.5px solid;
            }

            table#tb-list-cust>thead th:first-child,
            table#tb-list-cust>thead th:nth-child(2),
            table#tb-list-cust>thead th:nth-child(4) {
                text-align: center;
            }

            thead th:first-child {
                padding: 0 5px;
            }

            table#tb-list-cust>tbody>tr>td:first-child,
            table#tb-list-cust>tbody>tr>td:nth-child(2),
            table#tb-list-cust>tbody>tr>td:nth-child(4) {
                text-align: center;
            }

            table#tb-list-cust>tbody tr td.null-td {
                border: none;
            }

            table#tb-list-cust>tbody tr td.tb-total,
            td.tb-payment {
                background-color: #f8f8f8;
            }

            table#tb-list-cust>thead tr th:first-child {
                padding: 0 20px;
            }

            table#tb-list-cust>thead tr th:nth-child(5) {
                padding-left: 15px;
            }

            table#tb-list-cust>tbody tr td:nth-child(5) {
                padding-left: 15px;
            }

            table#tb-list-cust>tbody tr td:nth-child(6) {
                text-align: right;
                padding-right: 50px;
            }

            table#tb-list-cust>thead tr th:nth-child(6) {
                text-align: right;
                padding-right: 50px;
            }

            #wrap-img-paid {
                position: relative;
                float: left;
                width: 100%;
                padding: 0 0 50px 30px;
            }
        </style>
        <div class="container">
            <div>
                <div class="bg-banner-cabeen">
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/banner_cabeen.png">
                    <div class="content-logo-cabeen">
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_header.svg">
                        <p>Receipt</p>
                    </div>
                </div>
                <div class="bg-booking-code">
                    <p><b>Booking ID : {{.BookID}}</b></p>
                    <p><b>Date : {{.CreatedAt}}</b></p>
                </div>

                <div class="content-head-eticket">
                    <ul>
                        <li>
                            <h3>Customer Details</h3>
                        </li>
                        <li><b>Name :</b> {{.Customer.Name}}</li>
                        <li><b>Email :</b> {{.Customer.Email}}</li>
                        <li><b>Contact Number :</b> {{.Customer.Phone}}</li>
                    </ul>
                    <ul>
                        <li>
                            <h3>Payment Details</h3>
                        </li>
                        <li><b>Booking ID :</b> {{.BookID}}</li>
                        <li><b>Method :</b> {{.Payment.Method}}</li>
                        <li><b>Status :</b> {{.Payment.Status}}</li>
                    </ul>
                    <div>
                        <h3>Passenger Details</h3>
                        <div id="hr-head"></div>
                    </div>
                    <ol>
                        {{range $index, $elem := .Passengers}}
                        <li>{{$elem.Name}} <span>({{$elem.Type}})</span></li>
                        {{end}}
                    </ol>
                </div>

                <table id="tb-list-cust">
                    <thead>
                        <tr>
                            <th>No.</th>
                            <th>Type of Item</th>
                            <th>Item Description</th>
                            <th>Qty</th>
                            <th>Price per unit Rp</th>
                            <th>Total Rp</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range $idx, $el := .Items}}
                        <tr>
                            <td>{{add $idx 1}}</td>
                            <td>{{$el.Type}}</td>
                            <td>{{$el.Desc}}</td>
                            <td>{{$el.Qty}}</td>
                            <td>{{$el.Price}}</td>
                            <td>{{$el.TotalPrice}}</td>
                        </tr>
                        {{end}}
                        <tr>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="tb-total"><b>TOTAL</b></td>
                            <td class="tb-total">{{.Total}}</td>
                        </tr>
                        <tr>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td><b>ADMINISTRATION FEE</b></td>
                            <td>{{.AdminFee}}</td>
                        </tr>
                        <tr>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="null-td"></td>
                            <td class="tb-payment"><b>PAYMENT AMOUNT</b></td>
                            <td class="tb-payment">{{.TotalPayment}}</td>
                        </tr>
                    </tbody>
                </table>
            </div>

            <div id="wrap-img-paid">
                <img src="https://storage.googleapis.com/cabeen-dev/paid_flapp.svg">
            </div>
            <div class="footer-boarding">
                <div>
                    <p>Download Flapp Apps</p>
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/playstore.png"
                        onclick="location.href='https://play.google.com/store'">
                    <img src="https://storage.googleapis.com/cabeen-dev/assets/appstore.png"
                        onclick="location.href='https://apps.apple.com/'">
                </div>
                <div>
                    <div>
                        <img src="https://storage.googleapis.com/cabeen-dev/icon_footer.svg">
                    </div>
                    <p>{{now.UTC.Year}}, PT LIONMAS DINAMIKAEXPRESS</p>
                </div>
            </div>
        </div>
    </body>
    </html>
    `
)
