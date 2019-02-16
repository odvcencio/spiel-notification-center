package tools

import (
	"fmt"
	"log"
	"os"
	"spiel/notification-center/models"
	"strings"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailPromptToWebUser(webUser, spieler models.User) {
	from := mail.NewEmail("Spiel", "info@tryspiel.com")
	subject := fmt.Sprintf("Congratulations %s!", webUser.FirstName)

	firstAndLast := fmt.Sprintf("%s %s", webUser.FirstName, webUser.LastName)

	to := mail.NewEmail(firstAndLast, webUser.ID)
	plainTextContent := `You have taken the first step in joining Spiel, and receiving
    all the benefits we’ve worked hard to provide for you, and
    all our users! ` + spieler.FirstName + " " + spieler.LastName + `, ` + spieler.Title +
		` at ` + spieler.Company + `,  has received
    your question, download our app and sign up so you can
    see ` + spieler.FirstName + `'s personalized video answer specifically
    for you.`
	htmlContent := generateHTML(webUser, spieler)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}

func generateHTML(webUser, spieler models.User) string {
	return `<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office"><head>
	    <!--[if gte mso 9]><xml>
	     <o:OfficeDocumentSettings>
	      <o:AllowPNG/>
	      <o:PixelsPerInch>96</o:PixelsPerInch>
	     </o:OfficeDocumentSettings>
	    </xml><![endif]-->
	    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	    <meta name="viewport" content="width=device-width">
	    <!--[if !mso]><!--><meta http-equiv="X-UA-Compatible" content="IE=edge"><!--<![endif]-->
	    <title></title>
	    <!--[if !mso]><!-- -->
		<link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro" rel="stylesheet" type="text/css">
		<!--<![endif]-->

	    <style type="text/css" id="media-query">
	      body {
	  margin: 0;
	  padding: 0; }

	table, tr, td {
	  vertical-align: top;
	  border-collapse: collapse; }

	.ie-browser table, .mso-container table {
	  table-layout: fixed; }

	* {
	  line-height: inherit; }

	a[x-apple-data-detectors=true] {
	  color: inherit !important;
	  text-decoration: none !important; }

	[owa] .img-container div, [owa] .img-container button {
	  display: block !important; }

	[owa] .fullwidth button {
	  width: 100% !important; }

	[owa] .block-grid .col {
	  display: table-cell;
	  float: none !important;
	  vertical-align: top; }

	.ie-browser .num12, .ie-browser .block-grid, [owa] .num12, [owa] .block-grid {
	  width: 490px !important; }

	.ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div {
	  line-height: 100%; }

	.ie-browser .mixed-two-up .num4, [owa] .mixed-two-up .num4 {
	  width: 160px !important; }

	.ie-browser .mixed-two-up .num8, [owa] .mixed-two-up .num8 {
	  width: 320px !important; }

	.ie-browser .block-grid.two-up .col, [owa] .block-grid.two-up .col {
	  width: 245px !important; }

	.ie-browser .block-grid.three-up .col, [owa] .block-grid.three-up .col {
	  width: 163px !important; }

	.ie-browser .block-grid.four-up .col, [owa] .block-grid.four-up .col {
	  width: 122px !important; }

	.ie-browser .block-grid.five-up .col, [owa] .block-grid.five-up .col {
	  width: 98px !important; }

	.ie-browser .block-grid.six-up .col, [owa] .block-grid.six-up .col {
	  width: 81px !important; }

	.ie-browser .block-grid.seven-up .col, [owa] .block-grid.seven-up .col {
	  width: 70px !important; }

	.ie-browser .block-grid.eight-up .col, [owa] .block-grid.eight-up .col {
	  width: 61px !important; }

	.ie-browser .block-grid.nine-up .col, [owa] .block-grid.nine-up .col {
	  width: 54px !important; }

	.ie-browser .block-grid.ten-up .col, [owa] .block-grid.ten-up .col {
	  width: 49px !important; }

	.ie-browser .block-grid.eleven-up .col, [owa] .block-grid.eleven-up .col {
	  width: 44px !important; }

	.ie-browser .block-grid.twelve-up .col, [owa] .block-grid.twelve-up .col {
	  width: 40px !important; }

	@media only screen and (min-width: 510px) {
	  .block-grid {
	    width: 490px !important; }
	  .block-grid .col {
	    vertical-align: top; }
	    .block-grid .col.num12 {
	      width: 490px !important; }
	  .block-grid.mixed-two-up .col.num4 {
	    width: 160px !important; }
	  .block-grid.mixed-two-up .col.num8 {
	    width: 320px !important; }
	  .block-grid.two-up .col {
	    width: 245px !important; }
	  .block-grid.three-up .col {
	    width: 163px !important; }
	  .block-grid.four-up .col {
	    width: 122px !important; }
	  .block-grid.five-up .col {
	    width: 98px !important; }
	  .block-grid.six-up .col {
	    width: 81px !important; }
	  .block-grid.seven-up .col {
	    width: 70px !important; }
	  .block-grid.eight-up .col {
	    width: 61px !important; }
	  .block-grid.nine-up .col {
	    width: 54px !important; }
	  .block-grid.ten-up .col {
	    width: 49px !important; }
	  .block-grid.eleven-up .col {
	    width: 44px !important; }
	  .block-grid.twelve-up .col {
	    width: 40px !important; } }

	@media (max-width: 510px) {
	  .block-grid, .col {
	    min-width: 320px !important;
	    max-width: 100% !important;
	    display: block !important; }
	  .block-grid {
	    width: calc(100% - 40px) !important; }
	  .col {
	    width: 100% !important; }
	    .col > div {
	      margin: 0 auto; }
	  img.fullwidth, img.fullwidthOnMobile {
	    max-width: 100% !important; }
	  .no-stack .col {
	    min-width: 0 !important;
	    display: table-cell !important; }
	  .no-stack.two-up .col {
	    width: 50% !important; }
	  .no-stack.mixed-two-up .col.num4 {
	    width: 33% !important; }
	  .no-stack.mixed-two-up .col.num8 {
	    width: 66% !important; }
	  .no-stack.three-up .col.num4 {
	    width: 33% !important; }
	  .no-stack.four-up .col.num3 {
	    width: 25% !important; }
	  .mobile_hide {
	    min-height: 0px;
	    max-height: 0px;
	    max-width: 0px;
	    display: none;
	    overflow: hidden;
	    font-size: 0px; } }

	    </style>
	</head>
	<body class="clean-body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #FFFFFF">
	  <style type="text/css" id="media-query-bodytag">
	    @media (max-width: 520px) {
	      .block-grid {
	        min-width: 320px!important;
	        max-width: 100%!important;
	        width: 100%!important;
	        display: block!important;
	      }

	      .col {
	        min-width: 320px!important;
	        max-width: 100%!important;
	        width: 100%!important;
	        display: block!important;
	      }

	        .col > div {
	          margin: 0 auto;
	        }

	      img.fullwidth {
	        max-width: 100%!important;
	      }
				img.fullwidthOnMobile {
	        max-width: 100%!important;
	      }
	      .no-stack .col {
					min-width: 0!important;
					display: table-cell!important;
				}
				.no-stack.two-up .col {
					width: 50%!important;
				}
				.no-stack.mixed-two-up .col.num4 {
					width: 33%!important;
				}
				.no-stack.mixed-two-up .col.num8 {
					width: 66%!important;
				}
				.no-stack.three-up .col.num4 {
					width: 33%!important;
				}
				.no-stack.four-up .col.num3 {
					width: 25%!important;
				}
	      .mobile_hide {
	        min-height: 0px!important;
	        max-height: 0px!important;
	        max-width: 0px!important;
	        display: none!important;
	        overflow: hidden!important;
	        font-size: 0px!important;
	      }
	    }
	  </style>
	  <!--[if IE]><div class="ie-browser"><![endif]-->
	  <!--[if mso]><div class="mso-container"><![endif]-->
	  <table class="nl-container" style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;min-width: 320px;Margin: 0 auto;background-color: #FFFFFF;width: 100%" cellpadding="0" cellspacing="0">
		<tbody>
		<tr style="vertical-align: top">
			<td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top">
	    <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td align="center" style="background-color: #FFFFFF;"><![endif]-->

	    <div style="background-color:transparent;">
	      <div style="Margin: 0 auto;min-width: 320px;max-width: 490px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;" class="block-grid ">
	        <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
	          <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="background-color:transparent;" align="center"><table cellpadding="0" cellspacing="0" border="0" style="width: 490px;"><tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

	              <!--[if (mso)|(IE)]><td align="center" width="490" style=" width:490px; padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><![endif]-->
	            <div class="col num12" style="min-width: 320px;max-width: 490px;display: table-cell;vertical-align: top;">
	              <div style="background-color: transparent; width: 100% !important;">
	              <!--[if (!mso)&(!IE)]><!--><div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><!--<![endif]-->


	                    <div align="center" class="img-container center  autowidth  fullwidth " style="padding-right: 5px;  padding-left: 5px;">
	<!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr style="line-height:0px;line-height:0px;"><td style="padding-right: 5px; padding-left: 5px;" align="center"><![endif]-->
	<div style="line-height:5px;font-size:1px">&#160;</div>  <img class="center  autowidth  fullwidth" align="center" border="0" src="http://www.fastlaneapps.com/wp-content/uploads/2019/01/emailConfirmationQuestionViaWebApp@3x.jpg" alt="Spiel-logo" title="Spiel-logo" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: block !important;border: 0;height: auto;float: none;width: 100%;max-width: 480px" width="480">
	<div style="line-height:5px;font-size:1px">&#160;</div><!--[if mso]></td></tr></table><![endif]-->
	</div>


	              <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
	              </div>
	            </div>
	          <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
	        </div>
	      </div>
	    </div>
	    <div style="background-color:transparent;">
	      <div style="Margin: 0 auto;min-width: 320px;max-width: 490px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;" class="block-grid ">
	        <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
	          <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="background-color:transparent;" align="center"><table cellpadding="0" cellspacing="0" border="0" style="width: 490px;"><tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

	              <!--[if (mso)|(IE)]><td align="center" width="490" style=" width:490px; padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><![endif]-->
	            <div class="col num12" style="min-width: 320px;max-width: 490px;display: table-cell;vertical-align: top;">
	              <div style="background-color: transparent; width: 100% !important;">
	              <!--[if (!mso)&(!IE)]><!--><div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><!--<![endif]-->


	                    <div class="">
		<!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top: 0px; padding-bottom: 0px;"><![endif]-->
		<div style="color:#9b9b9b;font-family:'Source Sans Pro', Tahoma, Verdana, Segoe, sans-serif;line-height:150%; padding-right: 0px; padding-left: 0px; padding-top: 0px; padding-bottom: 0px;">
			<div style="font-size:12px;line-height:18px;color:#9b9b9b;font-family:'Source Sans Pro', Tahoma, Verdana, Segoe, sans-serif;text-align:left;"><div style="font-size:12px; line-height:18px;">
	<p style="margin: 0;font-size: 12px;line-height: 18px;text-align: center">Congratulations, ` + strings.Title(webUser.FirstName) + `!</p>
	<p style="margin: 0;font-size: 12px;line-height: 18px;text-align: center">&#160;</p>
	<p style="margin: 0;font-size: 12px;line-height: 18px;text-align: center">You have taken the first step in joining Spiel, and receiving <br>all the benefits we’ve worked hard to provide for you, and <br>all our users! ` + spieler.FirstName + " " + spieler.LastName + `, ` + spieler.Title + ` at ` + spieler.Company + ` has received <br>your question, download our app and sign up so you can <br>see ` + spieler.FirstName + `'s personalized video answer specifically <br>for you.</p>
	</div></div>
		</div>
		<!--[if mso]></td></tr></table><![endif]-->
	</div>


	                    <div align="center" class="img-container center  autowidth  " style="padding-right: 15px;  padding-left: 15px;">
	<!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr style="line-height:0px;line-height:0px;"><td style="padding-right: 15px; padding-left: 15px;" align="center"><![endif]-->
	<div style="line-height:15px;font-size:1px">&#160;</div>  <a href="https://itunes.apple.com/us/app/spiel-app/id1449212004?mt=8" target="_blank">
	    <img class="center  autowidth " align="center" border="0" src="images/appStore.png" alt="Image" title="Image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: block !important;border: none;height: auto;float: none;width: 100%;max-width: 120px" width="120">
	  </a>
	<div style="line-height:15px;font-size:1px">&#160;</div><!--[if mso]></td></tr></table><![endif]-->
	</div>


	              <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
	              </div>
	            </div>
	          <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
	        </div>
	      </div>
	    </div>
	   <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
			</td>
	  </tr>
	  </tbody>
	  </table>
	  <!--[if (mso)|(IE)]></div><![endif]-->


	</body></html>`
}
