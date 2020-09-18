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
            background-color: #ffffff;
            margin: 0;
            padding: 0;
        }

        .container {
            flex-grow: 1;
            margin: 0 auto;
            position: relative;
            width: auto;
        }

        @media screen and (min-width: 1024px) {
            .container {
                max-width: 960px;
            }
        }

        /* @media screen and (min-width: 1216px) {
            .container {
                max-width: 1152px;
            }
        }

        @media screen and (min-width: 1408px) {
            .container {
                max-width: 1344px;
            }
        } */

        .title-boarding-pass {
            font-family: "Poppins", sans-serif;
            font-size: 24px;
            color: #ffffff;
            font-weight: 600;
            padding-left: 7%;
            padding-top: 2%;
            max-width: 100%;
        }

        .header {
            width: 100%;
            height: 225px;
            background-color: #04847D;
            color: white;
        }

        .content-card-boarding {
            margin-left: 6%;
            margin-right: 6%;
            margin-top: -20%;
        }

        .main-card-boarding {
            width: 100%;
            background-color: #ffffff;
            min-height: 270px;
            border-radius: 15px;
            border-style: solid;
            border-width: 1.5px;
            border-color: #c4c4c4;
            margin-top: 3%
        }

        .content-header-card-main-boarding {
            display: flex;
            flex-direction: row;
            justify-content: space-between;
            padding-left: 3%;
            padding-right: 3%;
            align-items: center;
            padding-top: 1.5%;
            padding-bottom: 1.5%;
        }

        .content-body-card-main-boarding {
            display: flex;
            flex-direction: row;
            margin-top: 1%;
            margin-bottom: 1%;
            margin-right: 4%;

        }

        .content-body-card-main-boarding:after {
            content: "";
            display: table;
            clear: both;
        }

        .column-barcode {
            float: left;
            min-height: 170px;
            width: 13%;
            display: flex;
            align-items: center;
            justify-content: center;
            flex-direction: column;
        }

        .column-information-boarding {
            float: left;
            width: 90%;
        }
        

        .content-inside-information {
            display: flex;
            flex-direction: row;
            padding-top: 1%;
        }

        .second-content-inside-information {
            display: flex;
            flex-direction: row;
        }

        .column-inside-information {
            float: left;
            min-width: 34.5%;
            max-width: 34.5%;
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

        .icon-lion {
            max-width: 100%;
            height: 26px;
        }

        .icon-barcode {
            max-width: 100%;
            height: 39.5px;
        }

        .icon-body-barcode {
            width: 200px;
            padding-right: 6%;
            padding-bottom: 10%;
            /* margin-left: -35%; */
            -ms-transform: rotate(-90deg);
            /* IE 9 */
            transform: rotate(-90deg);
        }

        .hr-head {
            border-color: #c4c4c4;
            border-style: solid;
            border-width: 0.5px;
            margin: 0%;
        }

        .hr-body {
            margin-top: 2%;
            margin-bottom: 2%;
            margin-left: 0%;
            border-color: #c4c4c4;
            border-style: solid;
            border-width: 0.5px;
            margin-right: 0%;
            width: 100%;
        }

        .title-boarding {
            font-family: "Poppins", sans-serif;
            font-size: 19px;
            color: #454545;
            margin: 0px;
            width: 100%;
            text-align: center;
            font-weight: bold;
        }

        .content-title-boarding {
            margin-top: 2%;
            margin-bottom: 3%;
            display: flex;
            flex-direction: row;
            justify-content: center;
            align-items: center;
        }

        .border-gray-title {
            background-color: #dddddd;
            height: 1px;
            min-width: 17%;
        }

        .column-card-boarding {
            float: left;
            width: 48%;

        }

        .column-card-boarding-four {
            float: left;
            width: 76%;
        }

        .column-card-boarding-five {
            float: left;
            width: 20%;
        }

        .row {
            display: flex;
            margin: 2%;
            margin-left: 4.5%;
            margin-right: 4.5%;
            justify-content: center;
        }

        /* Clear floats after the columns */
        .row:after {
            content: "";
            display: table;
            clear: both;
        }

        .row-second {
            display: flex;
            margin-left: 4.5%;
            margin-right: 4.5%;
            justify-content: center;
        }

        /* Clear floats after the columns */
        .row-second:after {
            content: "";
            display: table;
            clear: both;
        }

        .gap-column {
            margin-left: 1%;
            margin-right: 1%;
        }

        .gap-inside {
            margin-top: 2%;
            margin-bottom: 2%;
        }

        .card-boarding-one {
            width: 100%;
            min-height: 28%;
            max-height: 35%;
            background-color: #ffffff;
            border-color: #c4c4c4;
            border-style: solid;
            border-width: 1px;
            border-radius: 15px;
        }

        .card-boarding-two {
            width: 100%;
            height: 60.5%;
            margin-top: 3%;
            background-color: #ffffff;
            border-color: #c4c4c4;
            border-style: solid;
            border-width: 1px;
            border-radius: 15px;
        }

        .card-boarding-three {
            width: 100%;
            max-height: 100%;
            background-color: #ffffff;
            border-color: #c4c4c4;
            border-style: solid;
            border-width: 1px;
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
            height: 36px;
            margin-top: -1%;
            background-size: contain;
            background-repeat: no-repeat;
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
        .content-card-travel {
            display: flex;
            flex-direction: row;
        }

        .row-content-travel {
            margin: 6%;
            margin-top: 2%;
            margin-left: 3%;
            margin-bottom: 5%;
            display: flex;
            max-width: 100%;
            flex-direction: row;
        }

        .icon-boarding {
            width: 75px;
            height: 75px;
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

        .desc-travel {
            font-family: "Poppins", sans-serif;
            font-size: 10px;
            margin: 3px 0px;
            color: #454545;
            font-weight: 600;

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
            width: 75%;
        }

        .second-column-desc-boarding {
            float: left;
            width: 75%;
        }

        
        .third-column-desc-boarding {
            float: left;
            width: 85%;
        }


        @media screen and (max-width: 600px) {
            .column-desc-boarding {
                width: 100%;
                display: block;
                margin-bottom: 20px;
            }
        }

        /* Clear floats after the columns */
        .understand-text {
            font-family: "Poppins", sans-serif;
            font-size: 10px;
            color: #04847d;
            margin-top: -6%;
            width: 80%;
            margin-bottom: 3%;
            margin-left: 4%;
            font-weight: 600;
        }

        .content-backpack {
            display: flex;
            flex-direction: row;
            justify-content: center;
            margin-top: 9.5%;
            margin-bottom: 8%;
        }

        .group-backpack {
            width: 276px;
            height: 99px;
            object-fit: contain;
        }

        .bg-thanks {
            background-color: #ddefe9;
            width: 100%;
            height: 55px;
            margin-top: 2%;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .thanks-choosing {
            font-family: "Poppins", sans-serif;
            font-size: 18px;
            padding-top: 2%;
            padding-bottom: 3%;
            color: #04847d;
            font-weight: 600;
            text-align: center;
        }

        .footer-boarding {
            width: 100%;
            height:97.2px;
            background-color: #f0f9f6;
        }

        .icon-cabeen-boarding {
            width: 105px;
            height: 25.5px;
            padding-top: 5%;
            object-fit: contain;
        }

        .content-footer-boarding {
            display: flex;
            flex-direction: row;
            justify-content: space-between;
            margin-top: -1%;
            margin-bottom: 0%;
            margin-left: 7%;
            margin-right: 7%;
        }

        .download-cabeen-text {
            font-family: "Poppins", sans-serif;
            font-size: 12px;
            margin-top: 7px;
            margin-bottom: 0px;
            color: #99b1ab;
            font-weight: 600;
        }

        .content-download-footer {
            display: flex;
            flex-direction: column;
            margin-top: 1%;
        }

        .content-store {
            display: flex;
            flex-direction: row;
        }

        .appstore {
            width: 125px;
            height: 52.5px;
            cursor: pointer;
            object-fit: contain;

        }

        .playstore {
            width: 125px;
            height: 52.5px;
            margin-right: 3%;
            object-fit: contain;
            cursor: pointer;

        }
    </style>
    <div class="container">
        <div class="header">
        <!-- <div style="width: 100%;height: 225px;background-color: #04847d;color: #fff;"> -->
            <p class="title-boarding-pass">Your Boarding Pass</p>
        </div>
        <div class="content-card-boarding">
            <div class="main-card-boarding">
                <div class="content-header-card-main-boarding">
                    <img src="https://storage.googleapis.com/cabeen-dev/icon_lion.svg" class="icon-lion">
                    <img src="https://storage.googleapis.com/cabeen-dev/barcode.svg" class="icon-barcode">
                </div>
                <hr class="hr-head" />
                <div class="content-body-card-main-boarding">
                    <div class="column-barcode">
                        <img src="https://storage.googleapis.com/cabeen-dev/barcode.svg" class="icon-body-barcode">
                    </div>
                    <div class="column-information-boarding">
                        <div class="content-inside-information">
                            <div class="column-inside-information">
                                <p class="title-main-boarding">Name</p>
                                <p class="subtitle-main-boarding">{{.FullName}}</p>
                            </div>
                            <div class="column-inside-information">
                                <p class="title-main-boarding">Frequent Flyer No.</p>
                                <p class="subtitle-main-boarding">{{.FFNumber}}</p>
                            </div>
                            <div class="column-inside-information">
                                <p class="title-main-boarding">Security No.</p>
                                <p class="subtitle-main-boarding">{{.SecurityNumber}}</p>
                            </div>
                        </div>
                        <hr class="hr-body" />
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
            </div>
        </div>


        <!-- TITLE PRESENT BOARDING PASS -->
        <div class="content-title-boarding">
            <div class="border-gray-title"></div>
            <p class="title-boarding">Please present this boarding pass to the airport for flight</p>
            <div class="border-gray-title"></div>
        </div>

        <!-- LINE ONE -->
        <div class="row">
            <div class="column-card-boarding">
                <div class="card-boarding-one">
                    <div class="rectangle-one">
                        <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;width: 22%;">
                        <p class="title-boarding-card">Check-in with luggage</p>
                    </div>
                    <div class="content-card-travel">
                        <div class="row-content-travel">
                            <img src="https://storage.googleapis.com/cabeen-dev/travel_one.svg" class="icon-boarding">
                            <div class="gap-desc-travel-one"></div>
                            <div class="column-desc-boarding">
                                <p class="desc-travel">Please head to the check-in counter no later than 60 minutes
                                    before
                                    departure</p>
                            </div>
                        </div>
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
                            <div class="gap-desc-travel-one"></div>
                            <div class="column-desc-boarding">
                                <div class="content-point">
                                    <div class="ellipse-point">
                                    </div>
                                    <p class="desc-travel">Make sure your items always stay with you </p>
                                </div>
                                <div class="content-point">
                                    <div class="ellipse-point "></div>
                                    <p class="desc-travel">Make sure you know the contents of your luggage</p>
                                </div>
                                <div class="content-point">
                                    <div class="ellipse-point">
                                    </div>
                                    <p class="desc-travel">Pay attention to items that are not allowed to be carried
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <br>
                    <p class="understand-text">Make sure you understand, know and agree to the applicable
                        security rules</p>
                </div>
            </div>
            <div class="gap-column"></div>
            <div class="column-card-boarding">
                <div class="card-boarding-three">
                    <div class="rectangle-three">
                        <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;width: 22%;">
                        <p class="title-boarding-card">Carry-on bag only?</p>
                    </div>
                    <div class="content-card-travel">
                        <div class="row-content-travel">
                            <img src="https://storage.googleapis.com/cabeen-dev/backpack.png" class="icon-boarding">
                            <div class="gap-desc-travel-one"></div>
                            <div class="second-column-desc-boarding">
                                <div class="content-point">
                                    <div class="ellipse-point">
                                    </div>
                                    <p class="desc-travel">Domestic destination - please head to the departure gate
                                    </p>
                                </div>
                                <div class="content-point">
                                    <div class="ellipse-point">
                                    </div>
                                    <p class="desc-travel">International destinations - please go to the check-in
                                        counter for verification of travel documents</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="rectangle-four">
                        <img src="https://storage.googleapis.com/cabeen-dev/rectangle_one.svg" style="position: absolute;width: 22%;">
                        <p class="title-boarding-card">Size of carry-on bag</p>
                    </div>
                    <div class="content-backpack">
                        <img src="https://storage.googleapis.com/cabeen-dev/group_backpack.png" class="group-backpack">
                    </div>
                </div>
            </div>
        </div>

        <!-- LINE TWO -->
        <div class="row-second">
            <div class="column-card-boarding-four">
                <div class="card-boarding-four">
                    <div class="rectangle-five">
                        <img src="https://storage.googleapis.com/cabeen-dev/rectangle_five.svg" style="position: absolute;width: 22%;">
                        <p class="title-boarding-card-last">Important Information</p>
                    </div>
                    <div class="content-card-travel">
                        <div class="row-content-travel">
                            <div class="third-column-desc-boarding">
                                <div class="content-point">
                                    <div class="ellipse-point-second">
                                    </div>
                                    <p class="desc-travel">Make sure your travel documents are still valid</p>
                                </div>
                                <div class="content-point">
                                    <div class="ellipse-point-second">
                                    </div>
                                    <p class="desc-travel">Allowance for baggage as stipulated in your E-Ticket
                                        regulations</p>
                                </div>
                                <div class="content-point">
                                    <div class="ellipse-point-second">
                                    </div>
                                    <p class="desc-travel">Receiving boarding passes is at the discretion of the
                                        airport
                                    </p>
                                </div>
                                <div class="content-point">
                                    <div class="ellipse-point-second">
                                    </div>
                                    <p class="desc-travel">For operational, safety or security reasons we can change
                                        our
                                        seats at any time even after you board the aircraft
                                    </p>
                                </div>
                                <div class="content-point">
                                    <img src="https://storage.googleapis.com/cabeen-dev/ellipse_point.svg" class="ellipse-point-second">
                                    <p class="desc-travel">You are expected to be at the Boarding Gate no later than
                                        40
                                        minutes before departure time to avoid being offroad from the aircraft
                                        and
                                        from
                                        being charged a no-show fee
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="gap-inside"></div>
                </div>
            </div>
            <div class="gap-column"></div>
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
            <p class="thanks-choosing">Thank you for choosing to fly with Lion Air</p>
        </div>
        
        <div class="footer-boarding">
            <div class="content-footer-boarding">
                <img src="https://storage.googleapis.com/cabeen-dev/icon_cabeen.png" class="icon-cabeen-boarding">
                <div class="content-download-footer">
                    <p class="download-cabeen-text">Download Cabeen Apps</p>
                    <div class="content-store">
                        <img src="https://storage.googleapis.com/cabeen-dev/playstore.png" onclick="location.href='https://play.google.com/store'"
                            class="playstore">
                        <img src="https://storage.googleapis.com/cabeen-dev/appstore.png" onclick="location.href='https://apps.apple.com/'"
                            class="appstore">
                    </div>
                </div>
            </div>
        </div>
</body>
</html>
	`
)
