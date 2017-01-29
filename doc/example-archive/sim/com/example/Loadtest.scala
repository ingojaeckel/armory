package com.example

import io.gatling.core.scenario.Simulation
import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class Loadtest extends Simulation {

  val scn = scenario("Example loadtest")
      .exec(http("ping").get("/ping").check(status.is(200)))

  setUp(scn.inject(rampUsers(10) over (1 minute))
    .protocols(http.baseURL("https://load-test-me.appspot.com")
      .acceptEncodingHeader("gzip, deflate")
      .userAgentHeader("Gatling")
      .shareConnections))
}
