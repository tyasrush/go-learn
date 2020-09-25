package html

var (
	BoardingPassHTML = `
	<html> 
<head>
</head>
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
            min-width: 50px;
            max-width: 70px;
            height: 190px;
        }

        #column-information-boarding {
            display: inline;
            float: left;
        }

        .card-bp-body {
            min-width: 70%;
            max-width: 100%;
            float: left;
            margin: 1% 4% 1% 0;
        }

        .content-inside-information {
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
            min-width: 33.3%;
            max-width: 33.3%;
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
            margin: 0 6% 0 47%;
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
                    <img src="https://storage.googleapis.com/cabeen-dev/barcode.svg" class="icon-barcode">
                </div>
                <div class="card-bp-body">
                    <div class="column-barcode">
                        <img src="https://storage.googleapis.com/cabeen-dev/pdf_test.jpg" class="icon-body-barcode">
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
                                        <p class="second-title-main-boarding">Departure Date</p>
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
                <img src="https://storage.googleapis.com/cabeen-dev/icon_cabeen.png" class="icon-cabeen-boarding">
                <div class="content-download-footer">
                    <p class="download-cabeen-text">Download Cabeen Apps</p>
                    <div class="content-store">
                        <img src="https://storage.googleapis.com/cabeen-dev/playstore.png" onclick="location.href='https://play.google.com/store'">
                        <img src="https://storage.googleapis.com/cabeen-dev/appstore.png" onclick="location.href='https://apps.apple.com/'">
                    </div>
                </div>
        </div>
</body>
</html>
	`
)
