Feature: UC-1: CRUD Referral
  Scenario: Add Referral
    Given the Referral Service is initialized
    When a new referral is added
    Then the referral should be added successfully
