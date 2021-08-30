# file: version.feature
Feature: get body
  In order to test openfaas response
  As an API user
  I need to be able to request user payload

  Scenario: allow POST method
    When I send "POST" request to "http://localhost:8001/function/go-app"
    Then the response code should be 200
    And the response should match json:
      """
      {"name":"testUser","age":18}
      """