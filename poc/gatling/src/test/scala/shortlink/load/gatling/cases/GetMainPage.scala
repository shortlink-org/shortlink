package shortlink.load.gatling.cases

import io.gatling.http.Predef._
import io.gatling.core.Predef._

object GetMainPage {

  val getMainPage = http("Open main page")
    .get("/")
    .check(status is 200)

}
