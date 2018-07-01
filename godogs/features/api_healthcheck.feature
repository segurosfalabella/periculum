Feature: Check api healthness
  In order to maintain an application running
  As an product owner
  I need to constant checking the healthness of an api

  Scenario: Check health of an API
    Given A "api_url" 
    When I check artifact's healthness 
    Then notify the status page