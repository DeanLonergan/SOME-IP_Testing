Feature: Client Subscription to VSOMEIP Service

  Scenario: Client successfully subscribes to the VSOMEIP service
    Given the routingmanagerd is running
    And the service is running
    When the client application starts
    Then the client should successfully subscribe to the service
    And the client should receive a confirmation message "SUBSCRIBE ACK"
