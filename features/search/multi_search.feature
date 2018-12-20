@wip
Feature: Multiple flag searches

  Scenario: Searching by growth 13 and the summer season returns blueberry
    When a search for growth 13 and continuous true is performed
    Then the output must only include blueberry
   
  Scenario: Searching by growth 10 the fall season, and continuous false returns yam
    When a search for season fall, continuous false, and growth 10 is performed
    Then the output must only include yam
