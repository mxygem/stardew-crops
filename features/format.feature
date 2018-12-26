Feature: Formatting options
  Available options:
   * pretty - When outputting data, this will display things in a table or in other nicely formatted ways.
   * json - output formatted JSON
   * raw - (default for now) outputs non-formatted JSON
 
  Scenario: The pretty formatter returns nicely formatted data
    When info for starfruit is requested with the pretty formatter
    Then the output must be starfruit's pretty info

  Scenario: The json-pretty formatter returns json data with whitespace
    When a search is run with the json formatter
    Then the output must be formatted json

  Scenario: The raw formatter returns search json data with no whitespace
    When a search is run with the raw formatter
    Then the output must be raw json
